package monzo

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/esteanes/up-bank-go/datafetcher/models"
	"github.com/esteanes/up-bank-go/datafetcher/monzoclient"
	"github.com/esteanes/up-bank-go/datafetcher/providers"
)

const (
	defaultMaxTransactions = int32(100)
	pageSize               = int32(100)
)

// Provider implements the BankProvider interface for Monzo
type Provider struct {
	client       *monzoclient.APIClient
	oauthManager *OAuthManager
	log          *log.Logger
	enabled      bool
}

// NewProvider creates a new Monzo provider with OAuth support
func NewProvider(oauthManager *OAuthManager, log *log.Logger) *Provider {
	if oauthManager == nil {
		return &Provider{enabled: false, log: log}
	}

	configuration := monzoclient.NewConfiguration()
	apiClient := monzoclient.NewAPIClient(configuration)

	return &Provider{
		client:       apiClient,
		oauthManager: oauthManager,
		log:          log,
		enabled:      true,
	}
}

// ProviderType returns the provider identifier
func (p *Provider) ProviderType() string {
	return models.ProviderMonzo
}

// IsEnabled returns whether this provider is configured and authorized
func (p *Provider) IsEnabled() bool {
	return p.enabled && p.oauthManager != nil && p.oauthManager.IsAuthorized()
}

// GetOAuthManager returns the OAuth manager for this provider
func (p *Provider) GetOAuthManager() *OAuthManager {
	return p.oauthManager
}

// getAuthContext returns a context with the current access token
func (p *Provider) getAuthContext(ctx context.Context) (context.Context, error) {
	token, err := p.oauthManager.GetValidToken(ctx)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, monzoclient.ContextAccessToken, token), nil
}

// GetAccounts fetches all accounts from Monzo
func (p *Provider) GetAccounts(ctx context.Context, accountChan chan<- models.Account) error {
	defer close(accountChan)

	if !p.IsEnabled() {
		p.log.Println("Monzo provider not enabled or not authorized, skipping account fetch")
		return nil
	}

	authCtx, err := p.getAuthContext(ctx)
	if err != nil {
		p.log.Printf("Failed to get auth context: %v", err)
		return err
	}

	resp, r, err := p.client.AccountsAPI.ListAccounts(authCtx).Execute()
	if err != nil {
		p.log.Printf("Error fetching Monzo accounts: %v", err)
		if r != nil {
			p.log.Printf("HTTP Status: %d", r.StatusCode)
			if r.StatusCode == 403 {
				p.log.Println("403 Forbidden - You may need to approve access in the Monzo app (Strong Customer Authentication)")
			}
		}
		return err
	}

	if resp.Accounts == nil {
		return nil
	}

	for _, account := range resp.Accounts {
		if account.Id == nil {
			continue
		}

		balanceResp, _, balanceErr := p.client.BalanceAPI.GetBalance(authCtx).
			AccountId(*account.Id).
			Execute()
		if balanceErr != nil {
			p.log.Printf("Error fetching balance for Monzo account %s: %v", *account.Id, balanceErr)
			balanceResp = nil
		}

		normalizedAccount := NormalizeAccount(account, balanceResp)
		select {
		case accountChan <- normalizedAccount:
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}

// GetTransactions fetches transactions from Monzo
func (p *Provider) GetTransactions(ctx context.Context, txChan chan<- models.Transaction, params *providers.QueryParams) error {
	defer close(txChan)

	if !p.IsEnabled() {
		p.log.Println("Monzo provider not enabled or not authorized, skipping transaction fetch")
		return nil
	}

	authCtx, err := p.getAuthContext(ctx)
	if err != nil {
		p.log.Printf("Failed to get auth context: %v", err)
		return err
	}

	if params != nil && params.AccountID != nil {
		return p.getTransactionsForAccount(authCtx, txChan, *params.AccountID, params)
	}

	return p.getTransactionsForAllAccounts(authCtx, txChan, params)
}

func (p *Provider) getTransactionsForAccount(ctx context.Context, txChan chan<- models.Transaction, accountID string, params *providers.QueryParams) error {
	maxTx := getMaxTransactions(params)
	count := int32(0)

	// Always use forward pagination with 'since' parameter.
	// The Monzo API returns transactions oldest-first when 'since' is set,
	// and 'since' supports object IDs for reliable cursor-based pagination.
	// 'before' is only used as an upper-bound filter (end date), not as a pagination cursor.
	var sinceCursor, beforeCursor string
	var prevCursor string

	if params != nil && params.StartDate != nil {
		sinceCursor = params.StartDate.Format(time.RFC3339)
	}
	if params != nil && params.EndDate != nil {
		beforeCursor = params.EndDate.Format(time.RFC3339)
	}

	for {
		req := p.client.TransactionsAPI.ListTransactions(ctx).
			AccountId(accountID).
			Limit(pageSize)

		if sinceCursor != "" {
			req = req.Since(sinceCursor)
		}
		if beforeCursor != "" {
			req = req.Before(beforeCursor)
		}

		resp, r, err := req.Execute()
		if err != nil {
			p.log.Printf("Error fetching Monzo transactions: %v", err)
			if r != nil {
				p.log.Printf("HTTP Status: %d", r.StatusCode)
			}
			return err
		}

		if len(resp.Transactions) == 0 {
			break
		}

		for _, tx := range resp.Transactions {
			if count >= maxTx {
				return nil
			}
			select {
			case txChan <- NormalizeTransaction(tx, accountID):
				count++
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		// Advance cursor using the last transaction's ID (oldest-first ordering
		// means the last element is the newest in this batch).
		// 'since' accepts object IDs per the Monzo API spec.
		lastTx := resp.Transactions[len(resp.Transactions)-1]
		if lastTx.Id == nil {
			break
		}
		sinceCursor = *lastTx.Id

		// Safety: if cursor hasn't advanced, stop to prevent infinite loops
		if sinceCursor == prevCursor {
			break
		}
		prevCursor = sinceCursor
	}

	return nil
}

func (p *Provider) getTransactionsForAllAccounts(ctx context.Context, txChan chan<- models.Transaction, params *providers.QueryParams) error {
	resp, r, err := p.client.AccountsAPI.ListAccounts(ctx).Execute()
	if err != nil {
		p.log.Printf("Error fetching Monzo accounts: %v", err)
		if r != nil {
			p.log.Printf("HTTP Status: %d", r.StatusCode)
		}
		return err
	}

	if len(resp.Accounts) == 0 {
		return nil
	}

	maxTx := getMaxTransactions(params)
	var mu sync.Mutex
	totalCount := int32(0)

	var wg sync.WaitGroup
	for _, account := range resp.Accounts {
		if account.Id == nil {
			continue
		}

		wg.Add(1)
		go func(accountID string) {
			defer wg.Done()
			p.fetchTransactionsForAccountPaginated(ctx, txChan, accountID, params, maxTx, &totalCount, &mu)
		}(*account.Id)
	}

	wg.Wait()
	return nil
}

func (p *Provider) fetchTransactionsForAccountPaginated(
	ctx context.Context,
	txChan chan<- models.Transaction,
	accountID string,
	params *providers.QueryParams,
	maxTx int32,
	totalCount *int32,
	mu *sync.Mutex,
) {
	var sinceCursor, beforeCursor string
	var prevCursor string

	if params != nil && params.StartDate != nil {
		sinceCursor = params.StartDate.Format(time.RFC3339)
	}
	if params != nil && params.EndDate != nil {
		beforeCursor = params.EndDate.Format(time.RFC3339)
	}

	for {
		// Check if we've already hit the max
		mu.Lock()
		if *totalCount >= maxTx {
			mu.Unlock()
			return
		}
		mu.Unlock()

		req := p.client.TransactionsAPI.ListTransactions(ctx).
			AccountId(accountID).
			Limit(pageSize)

		if sinceCursor != "" {
			req = req.Since(sinceCursor)
		}
		if beforeCursor != "" {
			req = req.Before(beforeCursor)
		}

		txResp, _, txErr := req.Execute()
		if txErr != nil {
			p.log.Printf("Error fetching transactions for account %s: %v", accountID, txErr)
			return
		}

		if len(txResp.Transactions) == 0 {
			return
		}

		for _, tx := range txResp.Transactions {
			mu.Lock()
			if *totalCount >= maxTx {
				mu.Unlock()
				return
			}
			*totalCount++
			mu.Unlock()

			select {
			case txChan <- NormalizeTransaction(tx, accountID):
			case <-ctx.Done():
				return
			}
		}

		// Advance cursor using the last transaction's ID
		lastTx := txResp.Transactions[len(txResp.Transactions)-1]
		if lastTx.Id == nil {
			return
		}
		sinceCursor = *lastTx.Id

		// Safety: if cursor hasn't advanced, stop to prevent infinite loops
		if sinceCursor == prevCursor {
			return
		}
		prevCursor = sinceCursor
	}
}

// getMaxTransactions returns the max transactions from params or the default
func getMaxTransactions(params *providers.QueryParams) int32 {
	if params != nil && params.NumTransactions != nil {
		return *params.NumTransactions
	}
	return defaultMaxTransactions
}
