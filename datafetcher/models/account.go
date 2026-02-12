package models

import "time"

// Account represents a unified account model across different bank providers
type Account struct {
	ID          string    // Unified ID: "upbank_xxx" or "monzo_xxx"
	ProviderID  string    // Original ID from the provider
	Provider    string    // "upbank" or "monzo"
	DisplayName string    // Account display name
	AccountType string    // TRANSACTIONAL, SAVER, CURRENT, etc.
	Balance     Money     // Current balance
	CreatedAt   time.Time // When the account was created
}

type ProviderStatus struct {
	Name        string
	Description string
	IconLabel   string
	IsUpBank    bool
	Status      string
	IsConnected bool
	ActionURL   string
	ActionLabel string
	ShowAction  bool
}

// Provider constants
const (
	ProviderUpBank = "upbank"
	ProviderMonzo  = "monzo"
)

// Account type constants
const (
	AccountTypeTransactional = "TRANSACTIONAL"
	AccountTypeSaver         = "SAVER"
	AccountTypeCurrent       = "CURRENT"
)

// NewAccount creates a new unified account
func NewAccount(providerID, provider, displayName, accountType string, balance Money, createdAt time.Time) Account {
	return Account{
		ID:          provider + "_" + providerID,
		ProviderID:  providerID,
		Provider:    provider,
		DisplayName: displayName,
		AccountType: accountType,
		Balance:     balance,
		CreatedAt:   createdAt,
	}
}
