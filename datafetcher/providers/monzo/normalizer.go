package monzo

import (
	"time"

	"github.com/esteanes/up-bank-go/datafetcher/models"
	"github.com/esteanes/up-bank-go/datafetcher/monzoclient"
)

// NormalizeAccount converts a Monzo account and balance to a unified Account model
func NormalizeAccount(account monzoclient.Account, balance *monzoclient.BalanceResponse) models.Account {
	id := ""
	if account.Id != nil {
		id = *account.Id
	}

	displayName := ""
	if account.Description != nil {
		displayName = *account.Description
	}

	createdAt := time.Time{}
	if account.Created != nil {
		createdAt = *account.Created
	}

	// Convert balance from pennies to Money
	var money models.Money
	if balance != nil {
		currency := "GBP"
		if balance.Currency != nil {
			currency = *balance.Currency
		}
		balanceValue := int64(0)
		if balance.Balance != nil {
			balanceValue = *balance.Balance
		}
		money = models.NewMoneyFromBaseUnits(currency, balanceValue)
	} else {
		money = models.NewMoney("GBP", "0.00", 0)
	}

	// Monzo accounts are typically "CURRENT" type (uk_retail)
	accountType := models.AccountTypeCurrent

	return models.NewAccount(
		id,
		models.ProviderMonzo,
		displayName,
		accountType,
		money,
		createdAt,
	)
}

// NormalizeTransaction converts a Monzo transaction to a unified Transaction model
func NormalizeTransaction(tx monzoclient.Transaction, accountID string) models.Transaction {
	id := ""
	if tx.Id != nil {
		id = *tx.Id
	}

	description := ""
	if tx.Description != nil {
		description = *tx.Description
	}

	// Convert amount from pennies to Money
	currency := "GBP"
	if tx.Currency != nil {
		currency = *tx.Currency
	}
	amount := int64(0)
	if tx.Amount != nil {
		amount = *tx.Amount
	}
	money := models.NewMoneyFromBaseUnits(currency, amount)

	// Determine status: empty settled string = HELD, otherwise SETTLED
	status := models.TransactionStatusHeld
	var settledAt *time.Time
	if tx.Settled != nil && *tx.Settled != "" {
		status = models.TransactionStatusSettled
		// Parse settled time if available
		settledTime, err := time.Parse(time.RFC3339, *tx.Settled)
		if err == nil {
			settledAt = &settledTime
		}
	}

	createdAt := time.Time{}
	if tx.Created != nil {
		createdAt = *tx.Created
	}

	// Map notes to message
	var message *string
	if tx.Notes != nil && *tx.Notes != "" {
		message = tx.Notes
	}

	// Category
	var category *string
	if tx.Category != nil && *tx.Category != "" {
		category = tx.Category
	}

	return models.NewTransaction(
		id,
		models.ProviderMonzo,
		accountID,
		status,
		description,
		nil, // rawText - Monzo doesn't have this
		message,
		money,
		category,
		createdAt,
		settledAt,
	)
}
