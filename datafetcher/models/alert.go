package models

// ProviderAlert represents a warning or error from a bank provider
// that should be displayed to the user.
type ProviderAlert struct {
	Provider string // "monzo", "upbank"
	Message  string
}
