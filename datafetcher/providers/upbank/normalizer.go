package upbank

import (
	"github.com/esteanes/up-bank-go/datafetcher/models"
	"github.com/esteanes/up-bank-go/datafetcher/upclient"
)

// NormalizeAccount converts an Up Bank account to a unified Account model
func NormalizeAccount(account upclient.AccountResource) models.Account {
	balance := models.NewMoney(
		account.Attributes.Balance.CurrencyCode,
		account.Attributes.Balance.Value,
		int64(account.Attributes.Balance.ValueInBaseUnits),
	)

	// Map Up Bank account types to unified types
	accountType := string(account.Attributes.AccountType)

	return models.NewAccount(
		account.Id,
		models.ProviderUpBank,
		account.Attributes.DisplayName,
		accountType,
		balance,
		account.Attributes.CreatedAt,
	)
}

// NormalizeTransaction converts an Up Bank transaction to a unified Transaction model
func NormalizeTransaction(tx upclient.TransactionResource, accountID string) models.Transaction {
	amount := models.NewMoney(
		tx.Attributes.Amount.CurrencyCode,
		tx.Attributes.Amount.Value,
		int64(tx.Attributes.Amount.ValueInBaseUnits),
	)

	// Extract optional fields
	var rawText *string
	if rt := tx.Attributes.RawText.Get(); rt != nil {
		rawText = rt
	}

	var message *string
	if msg := tx.Attributes.Message.Get(); msg != nil {
		message = msg
	}

	var category *string
	if cat := tx.Relationships.Category.Data.Get(); cat != nil {
		category = &cat.Id
	}

	var settledAt = tx.Attributes.SettledAt.Get()

	status := models.TransactionStatusSettled
	if tx.Attributes.Status == upclient.HELD {
		status = models.TransactionStatusHeld
	}

	// Use the account ID from the transaction's relationship if available
	txAccountID := accountID
	if tx.Relationships.Account.Data.Id != "" {
		txAccountID = tx.Relationships.Account.Data.Id
	}

	return models.NewTransaction(
		tx.Id,
		models.ProviderUpBank,
		txAccountID,
		status,
		tx.Attributes.Description,
		rawText,
		message,
		amount,
		category,
		tx.Attributes.CreatedAt,
		settledAt,
	)
}
