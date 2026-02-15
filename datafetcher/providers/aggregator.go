package providers

import (
	"context"
	"errors"
	"fmt"
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
// Providers are checked for IsEnabled() at fetch time, not creation time,
// so providers that become enabled later (e.g., after OAuth) will still work.
func NewAggregator(log *log.Logger, providers ...BankProvider) *Aggregator {
	// Filter out nil providers but keep all non-nil ones
	// IsEnabled() is checked at fetch time to support late authorization
	validProviders := make([]BankProvider, 0)
	for _, p := range providers {
		if p != nil {
			validProviders = append(validProviders, p)
		}
	}
	return &Aggregator{
		providers: validProviders,
		log:       log,
	}
}

// GetAllAccounts fetches accounts from all providers concurrently
// Returns a channel of accounts and a channel of alerts for provider errors
func (a *Aggregator) GetAllAccounts(ctx context.Context) (<-chan models.Account, <-chan models.ProviderAlert) {
	accountChan := make(chan models.Account, 100)
	alertChan := make(chan models.ProviderAlert, 10)

	go func() {
		defer close(accountChan)
		defer close(alertChan)

		var wg sync.WaitGroup
		for _, provider := range a.providers {
			// Check if provider is enabled at fetch time
			if !provider.IsEnabled() {
				a.log.Printf("Skipping disabled provider: %s", provider.ProviderType())
				continue
			}

			wg.Add(1)
			go func(p BankProvider) {
				defer wg.Done()
				// Create a channel for this provider
				providerChan := make(chan models.Account, 50)

				// Start fetching in a goroutine
				errCh := make(chan error, 1)
				go func() {
					errCh <- p.GetAccounts(ctx, providerChan)
				}()

				// Forward accounts to the main channel
				for account := range providerChan {
					select {
					case accountChan <- account:
					case <-ctx.Done():
						return
					}
				}

				// Check for errors after the provider channel is drained
				if err := <-errCh; err != nil {
					a.log.Printf("Error fetching accounts from %s: %v", p.ProviderType(), err)
					a.sendAlert(alertChan, p.ProviderType(), err)
				}
			}(provider)
		}
		wg.Wait()
	}()

	return accountChan, alertChan
}

// GetAllTransactions fetches transactions from all providers concurrently
// Returns a channel of transactions and a channel of alerts for provider errors
func (a *Aggregator) GetAllTransactions(ctx context.Context, params *QueryParams) (<-chan models.Transaction, <-chan models.ProviderAlert) {
	bufSize := 1000
	if params != nil && params.NumTransactions != nil {
		bufSize = int(*params.NumTransactions)
	}
	txChan := make(chan models.Transaction, bufSize)
	alertChan := make(chan models.ProviderAlert, 10)

	go func() {
		defer close(txChan)
		defer close(alertChan)

		var wg sync.WaitGroup
		for _, provider := range a.providers {
			// Check if provider is enabled at fetch time
			if !provider.IsEnabled() {
				a.log.Printf("Skipping disabled provider: %s", provider.ProviderType())
				continue
			}

			wg.Add(1)
			go func(p BankProvider) {
				defer wg.Done()
				// Create a channel for this provider
				providerChan := make(chan models.Transaction, bufSize)

				// Start fetching in a goroutine
				errCh := make(chan error, 1)
				go func() {
					errCh <- p.GetTransactions(ctx, providerChan, params)
				}()

				// Forward transactions to the main channel
				for tx := range providerChan {
					select {
					case txChan <- tx:
					case <-ctx.Done():
						return
					}
				}

				// Check for errors after the provider channel is drained
				if err := <-errCh; err != nil {
					a.log.Printf("Error fetching transactions from %s: %v", p.ProviderType(), err)
					a.sendAlert(alertChan, p.ProviderType(), err)
				}
			}(provider)
		}
		wg.Wait()
	}()

	return txChan, alertChan
}

// sendAlert sends a ProviderAlert based on the error type
func (a *Aggregator) sendAlert(alertChan chan<- models.ProviderAlert, provider string, err error) {
	var scaErr *SCARequiredError
	var msg string
	if errors.As(err, &scaErr) {
		msg = "Strong Customer Authentication required. Please approve access in the Monzo app, then refresh this page."
	} else {
		msg = fmt.Sprintf("Error fetching data from %s: %v", provider, err)
	}
	alertChan <- models.ProviderAlert{
		Provider: provider,
		Message:  msg,
	}
}

// HasProviders returns true if at least one provider is enabled
func (a *Aggregator) HasProviders() bool {
	for _, p := range a.providers {
		if p.IsEnabled() {
			return true
		}
	}
	return false
}

// ProviderCount returns the number of currently enabled providers
func (a *Aggregator) ProviderCount() int {
	count := 0
	for _, p := range a.providers {
		if p.IsEnabled() {
			count++
		}
	}
	return count
}
