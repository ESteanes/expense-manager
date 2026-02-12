package providers

import (
	"context"
	"log"
	"sync"

	"github.com/esteanes/up-bank-go/datafetcher/models"
)

// Aggregator combines multiple bank providers and fetches data from all of them
type Aggregator struct {
	providers []BankProvider
	log       *log.Logger
}

// NewAggregator creates a new Aggregator with the given providers
func NewAggregator(log *log.Logger, providers ...BankProvider) *Aggregator {
	// Filter to only enabled providers
	enabled := make([]BankProvider, 0)
	for _, p := range providers {
		if p != nil && p.IsEnabled() {
			enabled = append(enabled, p)
		}
	}
	return &Aggregator{
		providers: enabled,
		log:       log,
	}
}

// GetAllAccounts fetches accounts from all providers concurrently
// Returns a channel that receives accounts from all enabled providers
func (a *Aggregator) GetAllAccounts(ctx context.Context) <-chan models.Account {
	accountChan := make(chan models.Account, 100)

	go func() {
		defer close(accountChan)

		var wg sync.WaitGroup
		for _, provider := range a.providers {
			wg.Add(1)
			go func(p BankProvider) {
				defer wg.Done()
				// Create a channel for this provider
				providerChan := make(chan models.Account, 50)

				// Start fetching in a goroutine
				go func() {
					err := p.GetAccounts(ctx, providerChan)
					if err != nil {
						a.log.Printf("Error fetching accounts from %s: %v", p.ProviderType(), err)
					}
				}()

				// Forward accounts to the main channel
				for account := range providerChan {
					select {
					case accountChan <- account:
					case <-ctx.Done():
						return
					}
				}
			}(provider)
		}
		wg.Wait()
	}()

	return accountChan
}

// GetAllTransactions fetches transactions from all providers concurrently
// Returns a channel that receives transactions from all enabled providers
func (a *Aggregator) GetAllTransactions(ctx context.Context, params *QueryParams) <-chan models.Transaction {
	txChan := make(chan models.Transaction, 1000)

	go func() {
		defer close(txChan)

		var wg sync.WaitGroup
		for _, provider := range a.providers {
			wg.Add(1)
			go func(p BankProvider) {
				defer wg.Done()
				// Create a channel for this provider
				providerChan := make(chan models.Transaction, 500)

				// Start fetching in a goroutine
				go func() {
					err := p.GetTransactions(ctx, providerChan, params)
					if err != nil {
						a.log.Printf("Error fetching transactions from %s: %v", p.ProviderType(), err)
					}
				}()

				// Forward transactions to the main channel
				for tx := range providerChan {
					select {
					case txChan <- tx:
					case <-ctx.Done():
						return
					}
				}
			}(provider)
		}
		wg.Wait()
	}()

	return txChan
}

// HasProviders returns true if at least one provider is configured
func (a *Aggregator) HasProviders() bool {
	return len(a.providers) > 0
}

// ProviderCount returns the number of enabled providers
func (a *Aggregator) ProviderCount() int {
	return len(a.providers)
}
