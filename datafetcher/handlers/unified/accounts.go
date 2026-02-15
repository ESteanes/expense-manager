package unified

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/esteanes/up-bank-go/datafetcher/models"
	"github.com/esteanes/up-bank-go/datafetcher/providers"
	"github.com/esteanes/up-bank-go/datafetcher/templates"
)

// AccountHandler handles account-related requests using the aggregator
type AccountHandler struct {
	Uri        string
	Log        *log.Logger
	Aggregator *providers.Aggregator
}

// NewAccountHandler creates a new unified account handler
func NewAccountHandler(log *log.Logger, aggregator *providers.Aggregator) *AccountHandler {
	return &AccountHandler{
		Uri:        "/accounts",
		Log:        log,
		Aggregator: aggregator,
	}
}

// ServeHTTP implements http.Handler
func (h *AccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Log.Printf("%s %s params: %s", r.Method, r.RequestURI, r.URL.Query())
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	case http.MethodPost:
		h.Post(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Get handles GET requests for accounts
func (h *AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	accountChannel, alertChannel := h.Aggregator.GetAllAccounts(r.Context())
	templ.Handler(templates.Accounts("Account Information", accountChannel, alertChannel, true), templ.WithStreaming()).ServeHTTP(w, r)
}

// Post handles POST requests (placeholder)
func (h *AccountHandler) Post(w http.ResponseWriter, r *http.Request) {}

// GetAccounts returns a channel of accounts (for use by other handlers)
func (h *AccountHandler) GetAccounts() <-chan models.Account {
	accounts, _ := h.Aggregator.GetAllAccounts(nil)
	return accounts
}
