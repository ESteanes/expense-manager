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

const maxPageSize = int32(100)

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

	// Fetch accounts
	resp, r, err := p.client.AccountsAPI.ListAccounts(authCtx).Execute()
	if err != nil {
		p.log.Printf("Error fetching Monzo accounts: %v", err)
		if r != nil {
			p.log.Printf("HTTP response: %v", r)
		}
		return err
	}

	if resp.Accounts == nil {
		return nil
	}

	// For each account, fetch its balance
	for _, account := range resp.Accounts {
		if account.Id == nil {
			continue
		}

		// Fetch balance for this account
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

	maxTransactions := int32(100)
	if params != nil && params.NumTransactions != nil {
		maxTransactions = *params.NumTransactions
	}

	// If specific account is requested
	if params != nil && params.AccountID != nil {
		return p.getTransactionsForAccount(authCtx, txChan, *params.AccountID, params, maxTransactions)
	}

	// Otherwise, fetch accounts first then get transactions for each
	return p.getTransactionsForAllAccounts(authCtx, txChan, params, maxTransactions)
}

func (p *Provider) getTransactionsForAccount(ctx context.Context, txChan chan<- models.Transaction, accountID string, params *providers.QueryParams, maxTransactions int32) error {
	req := p.client.TransactionsAPI.ListTransactions(ctx).
		AccountId(accountID).
		Limit(maxPageSize)

	if params != nil && params.StartDate != nil {
		req = req.Since(params.StartDate.Format(time.RFC3339))
	}
	if params != nil && params.EndDate != nil {
		req = req.Before(params.EndDate.Format(time.RFC3339))
	}

	resp, r, err := req.Execute()
	if err != nil {
		p.log.Printf("Error fetching Monzo transactions: %v", err)
		if r != nil {
			p.log.Printf("HTTP response: %v", r)
		}
		return err
	}

	if resp.Transactions == nil {
		return nil
	}

	count := int32(0)
	for _, tx := range resp.Transactions {
		if count >= maxTransactions {
			break
		}
		select {
		case txChan <- NormalizeTransaction(tx, accountID):
			count++
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}

func (p *Provider) getTransactionsForAllAccounts(ctx context.Context, txChan chan<- models.Transaction, params *providers.QueryParams, maxTransactions int32) error {
	// First fetch accounts
	resp, r, err := p.client.AccountsAPI.ListAccounts(ctx).Execute()
	if err != nil {
		p.log.Printf("Error fetching Monzo accounts: %v", err)
		if r != nil {
			p.log.Printf("HTTP response: %v", r)
		}
		return err
	}

	if len(resp.Accounts) == 0 {
		return nil
	}

	// Track total transactions across all accounts
	var mu sync.Mutex
	totalCount := int32(0)

	// Fetch transactions for each account concurrently
	var wg sync.WaitGroup
	for _, account := range resp.Accounts {
		if account.Id == nil {
			continue
		}

		wg.Add(1)
		go func(accountID string) {
			defer wg.Done()

			req := p.client.TransactionsAPI.ListTransactions(ctx).
				AccountId(accountID).
				Limit(maxPageSize)

			if params != nil && params.StartDate != nil {
				req = req.Since(params.StartDate.Format(time.RFC3339))
			}
			if params != nil && params.EndDate != nil {
				req = req.Before(params.EndDate.Format(time.RFC3339))
			}

			txResp, _, txErr := req.Execute()
			if txErr != nil {
				p.log.Printf("Error fetching transactions for account %s: %v", accountID, txErr)
				return
			}

			if txResp.Transactions == nil {
				return
			}

			for _, tx := range txResp.Transactions {
				mu.Lock()
				if totalCount >= maxTransactions {
					mu.Unlock()
					return
				}
				totalCount++
				mu.Unlock()

				select {
				case txChan <- NormalizeTransaction(tx, accountID):
				case <-ctx.Done():
					return
				}
			}
		}(*account.Id)
	}

	wg.Wait()
	return nil
}
