package unified

import (
	"context"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/esteanes/up-bank-go/datafetcher/functions"
	"github.com/esteanes/up-bank-go/datafetcher/models"
	"github.com/esteanes/up-bank-go/datafetcher/providers"
	"github.com/esteanes/up-bank-go/datafetcher/templates"
)

// TransactionsHandler handles transaction-related requests using the aggregator
type TransactionsHandler struct {
	Uri            string
	Log            *log.Logger
	Aggregator     *providers.Aggregator
	AccountHandler *AccountHandler
}

// NewTransactionsHandler creates a new unified transactions handler
func NewTransactionsHandler(log *log.Logger, aggregator *providers.Aggregator, accountHandler *AccountHandler) *TransactionsHandler {
	return &TransactionsHandler{
		Uri:            "/transactions",
		Log:            log,
		Aggregator:     aggregator,
		AccountHandler: accountHandler,
	}
}

// ServeHTTP implements http.Handler
func (h *TransactionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

// Get handles GET requests for transactions
func (h *TransactionsHandler) Get(w http.ResponseWriter, r *http.Request) {
	queryParams := functions.FetchQueryParams(r.URL.Query())
	transactionsChannel := h.FetchTransactions(r.Context(), queryParams)
	accountsChannel := h.Aggregator.GetAllAccounts(r.Context())
	templ.Handler(templates.Transactions("Transactions", transactionsChannel, accountsChannel, queryParams), templ.WithStreaming()).ServeHTTP(w, r)
}

// Post handles POST requests (placeholder)
func (h *TransactionsHandler) Post(w http.ResponseWriter, r *http.Request) {}

// FetchTransactions returns a channel of transactions based on query params
func (h *TransactionsHandler) FetchTransactions(ctx context.Context, queryParams *functions.QueryParams) <-chan models.Transaction {
	params := &providers.QueryParams{
		AccountID:       queryParams.AccountID,
		NumTransactions: queryParams.NumTransactions,
		StartDate:       queryParams.StartDate,
		EndDate:         queryParams.EndDate,
	}
	return h.Aggregator.GetAllTransactions(ctx, params)
}
