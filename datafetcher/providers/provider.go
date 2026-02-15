package providers

import (
	"context"
	"fmt"
	"time"

	"github.com/esteanes/up-bank-go/datafetcher/models"
)

// SCARequiredError is returned when a provider requires Strong Customer Authentication
type SCARequiredError struct {
	Provider string
}

func (e *SCARequiredError) Error() string {
	return fmt.Sprintf("%s: strong customer authentication required", e.Provider)
}

// QueryParams holds the parameters for fetching transactions
type QueryParams struct {
	AccountID       *string    // Filter by specific account (provider ID, not unified ID)
	NumTransactions *int32     // Maximum number of transactions to fetch
	StartDate       *time.Time // Filter transactions after this date
	EndDate         *time.Time // Filter transactions before this date
}

// BankProvider defines the interface that all bank providers must implement
type BankProvider interface {
	// GetAccounts fetches all accounts and sends them to the channel
	// The provider is responsible for closing the channel when done
	GetAccounts(ctx context.Context, accountChan chan<- models.Account) error

	// GetTransactions fetches transactions based on query parameters
	// The provider is responsible for closing the channel when done
	GetTransactions(ctx context.Context, txChan chan<- models.Transaction, params *QueryParams) error

	// ProviderType returns the provider identifier (e.g., "upbank", "monzo")
	ProviderType() string

	// IsEnabled returns whether this provider is configured and available
	IsEnabled() bool
}
