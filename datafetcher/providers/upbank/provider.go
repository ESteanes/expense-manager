package upbank

import (
	"context"
	"log"
	"net/url"

	"github.com/esteanes/up-bank-go/datafetcher/models"
	"github.com/esteanes/up-bank-go/datafetcher/providers"
	"github.com/esteanes/up-bank-go/datafetcher/upclient"
)

const maxPageSize = int32(100)

// Provider implements the BankProvider interface for Up Bank
type Provider struct {
	client  *upclient.APIClient
	auth    context.Context
	log     *log.Logger
	enabled bool
}

// NewProvider creates a new Up Bank provider
func NewProvider(token string, log *log.Logger) *Provider {
	if token == "" {
		return &Provider{enabled: false, log: log}
	}

	auth := context.WithValue(context.Background(), upclient.ContextAccessToken, token)
	configuration := upclient.NewConfiguration()
	apiClient := upclient.NewAPIClient(configuration)

	return &Provider{
		client:  apiClient,
		auth:    auth,
		log:     log,
		enabled: true,
	}
}

// ProviderType returns the provider identifier
func (p *Provider) ProviderType() string {
	return models.ProviderUpBank
}

// IsEnabled returns whether this provider is configured
func (p *Provider) IsEnabled() bool {
	return p.enabled
}

// GetAccounts fetches all accounts from Up Bank
func (p *Provider) GetAccounts(ctx context.Context, accountChan chan<- models.Account) error {
	defer close(accountChan)

	if !p.enabled {
		return nil
	}

	// Merge the provided context with auth context
	authCtx := context.WithValue(ctx, upclient.ContextAccessToken, p.auth.Value(upclient.ContextAccessToken))

	resp, r, err := p.client.AccountsAPI.AccountsGet(authCtx).
		PageSize(maxPageSize).
		FilterOwnershipType(upclient.INDIVIDUAL).
		Execute()
	if err != nil {
		p.log.Printf("Error fetching Up Bank accounts: %v", err)
		if r != nil {
			p.log.Printf("HTTP response: %v", r)
		}
		return err
	}

	for _, account := range resp.Data {
		select {
		case accountChan <- NormalizeAccount(account):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}

// GetTransactions fetches transactions from Up Bank
func (p *Provider) GetTransactions(ctx context.Context, txChan chan<- models.Transaction, params *providers.QueryParams) error {
	defer close(txChan)

	if !p.enabled {
		return nil
	}

	// Merge the provided context with auth context
	authCtx := context.WithValue(ctx, upclient.ContextAccessToken, p.auth.Value(upclient.ContextAccessToken))

	// Determine the maximum number of transactions to fetch
	maxTransactions := int32(100)
	if params != nil && params.NumTransactions != nil {
		maxTransactions = *params.NumTransactions
	}

	// Choose appropriate API call based on whether account is specified
	if params != nil && params.AccountID != nil {
		return p.getTransactionsForAccount(authCtx, txChan, params, maxTransactions)
	}
	return p.getTransactionsForAllAccounts(authCtx, txChan, params, maxTransactions)
}

func (p *Provider) getTransactionsForAccount(ctx context.Context, txChan chan<- models.Transaction, params *providers.QueryParams, maxTransactions int32) error {
	getRequest := p.client.TransactionsAPI.AccountsAccountIdTransactionsGet(ctx, *params.AccountID).PageSize(maxPageSize)

	if params.StartDate != nil {
		getRequest = getRequest.FilterSince(*params.StartDate)
	}
	if params.EndDate != nil {
		getRequest = getRequest.FilterUntil(*params.EndDate)
	}

	return p.fetchTransactionPages(ctx, txChan, func(pageAfter *string) ([]upclient.TransactionResource, *string, error) {
		if pageAfter != nil {
			pageKey, err := extractPageAfter(*pageAfter)
			if err != nil {
				return nil, nil, err
			}
			getRequest = getRequest.PageAfter(pageKey)
		}

		resp, r, err := getRequest.Execute()
		if err != nil {
			p.log.Printf("Error fetching Up Bank transactions: %v", err)
			if r != nil {
				p.log.Printf("HTTP response: %v", r)
			}
			return nil, nil, err
		}

		return resp.Data, resp.Links.Next.Get(), nil
	}, *params.AccountID, maxTransactions)
}

func (p *Provider) getTransactionsForAllAccounts(ctx context.Context, txChan chan<- models.Transaction, params *providers.QueryParams, maxTransactions int32) error {
	getRequest := p.client.TransactionsAPI.TransactionsGet(ctx).PageSize(maxPageSize)

	if params != nil && params.StartDate != nil {
		getRequest = getRequest.FilterSince(*params.StartDate)
	}
	if params != nil && params.EndDate != nil {
		getRequest = getRequest.FilterUntil(*params.EndDate)
	}

	return p.fetchTransactionPages(ctx, txChan, func(pageAfter *string) ([]upclient.TransactionResource, *string, error) {
		if pageAfter != nil {
			pageKey, err := extractPageAfter(*pageAfter)
			if err != nil {
				return nil, nil, err
			}
			getRequest = getRequest.PageAfter(pageKey)
		}

		resp, r, err := getRequest.Execute()
		if err != nil {
			p.log.Printf("Error fetching Up Bank transactions: %v", err)
			if r != nil {
				p.log.Printf("HTTP response: %v", r)
			}
			return nil, nil, err
		}

		return resp.Data, resp.Links.Next.Get(), nil
	}, "", maxTransactions)
}

type fetchPageFunc func(pageAfter *string) ([]upclient.TransactionResource, *string, error)

func (p *Provider) fetchTransactionPages(ctx context.Context, txChan chan<- models.Transaction, fetchPage fetchPageFunc, accountID string, maxTransactions int32) error {
	var pageAfter *string
	count := int32(0)

	for count < maxTransactions {
		transactions, nextPage, err := fetchPage(pageAfter)
		if err != nil {
			return err
		}

		for _, tx := range transactions {
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

		pageAfter = nextPage
		if pageAfter == nil {
			break
		}
	}

	return nil
}

func extractPageAfter(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	return parsedURL.Query().Get("page[after]"), nil
}
