package models

import "fmt"

// Money represents a monetary value in a unified format across providers
type Money struct {
	CurrencyCode     string // ISO 4217 currency code (e.g., "AUD", "GBP")
	Value            string // Formatted string value (e.g., "10.56")
	ValueInBaseUnits int64  // Amount in smallest denomination (cents/pennies)
}

// NewMoney creates a Money instance from formatted values
func NewMoney(currencyCode, value string, valueInBaseUnits int64) Money {
	return Money{
		CurrencyCode:     currencyCode,
		Value:            value,
		ValueInBaseUnits: valueInBaseUnits,
	}
}

// NewMoneyFromBaseUnits creates a Money instance from base units (pennies/cents)
// and derives the formatted string value
func NewMoneyFromBaseUnits(currencyCode string, valueInBaseUnits int64) Money {
	// Convert base units to formatted value (e.g., 1056 -> "10.56")
	sign := ""
	absValue := valueInBaseUnits
	if valueInBaseUnits < 0 {
		sign = "-"
		absValue = -valueInBaseUnits
	}
	whole := absValue / 100
	fraction := absValue % 100
	value := fmt.Sprintf("%s%d.%02d", sign, whole, fraction)

	return Money{
		CurrencyCode:     currencyCode,
		Value:            value,
		ValueInBaseUnits: valueInBaseUnits,
	}
}

// String returns a formatted string representation of the money value
func (m Money) String() string {
	return fmt.Sprintf("%s %s", m.CurrencyCode, m.Value)
}
