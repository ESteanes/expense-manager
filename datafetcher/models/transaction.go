package models

import "time"

// Transaction represents a unified transaction model across different bank providers
type Transaction struct {
	ID          string     // Unified ID: "upbank_xxx" or "monzo_xxx"
	ProviderID  string     // Original ID from the provider
	Provider    string     // "upbank" or "monzo"
	AccountID   string     // Unified account ID this transaction belongs to
	Status      string     // HELD or SETTLED
	Description string     // Transaction description (merchant name, etc.)
	RawText     *string    // Original unprocessed text
	Message     *string    // Attached message or note
	Amount      Money      // Transaction amount
	Category    *string    // Category of the transaction
	CreatedAt   time.Time  // When the transaction was first encountered
	SettledAt   *time.Time // When the transaction settled (nil if HELD)
}

// Transaction status constants
const (
	TransactionStatusHeld    = "HELD"
	TransactionStatusSettled = "SETTLED"
)

// NewTransaction creates a new unified transaction
func NewTransaction(
	providerID, provider, accountID, status, description string,
	rawText, message *string,
	amount Money,
	category *string,
	createdAt time.Time,
	settledAt *time.Time,
) Transaction {
	return Transaction{
		ID:          provider + "_" + providerID,
		ProviderID:  providerID,
		Provider:    provider,
		AccountID:   accountID,
		Status:      status,
		Description: description,
		RawText:     rawText,
		Message:     message,
		Amount:      amount,
		Category:    category,
		CreatedAt:   createdAt,
		SettledAt:   settledAt,
	}
}
