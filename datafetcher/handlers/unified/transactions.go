package unified

import (
	"context"
	"log"
	"net/http"
	"sync"

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
	transactionsChannel, txAlerts := h.FetchTransactions(r.Context(), queryParams)
	accountsChannel, accAlerts := h.Aggregator.GetAllAccounts(r.Context())
	alerts := mergeAlertChannels(txAlerts, accAlerts)
	templ.Handler(templates.Transactions("Transactions", transactionsChannel, accountsChannel, alerts, queryParams), templ.WithStreaming()).ServeHTTP(w, r)
}

// Post handles POST requests (placeholder)
func (h *TransactionsHandler) Post(w http.ResponseWriter, r *http.Request) {}

// FetchTransactions returns a channel of transactions and alerts based on query params
func (h *TransactionsHandler) FetchTransactions(ctx context.Context, queryParams *functions.QueryParams) (<-chan models.Transaction, <-chan models.ProviderAlert) {
	params := &providers.QueryParams{
		AccountID:       queryParams.AccountID,
		NumTransactions: queryParams.NumTransactions,
		StartDate:       queryParams.StartDate,
		EndDate:         queryParams.EndDate,
	}
	return h.Aggregator.GetAllTransactions(ctx, params)
}

// mergeAlertChannels merges multiple alert channels into one
func mergeAlertChannels(channels ...<-chan models.ProviderAlert) <-chan models.ProviderAlert {
	merged := make(chan models.ProviderAlert, 10)
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan models.ProviderAlert) {
			defer wg.Done()
			for alert := range c {
				merged <- alert
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(merged)
	}()
	return merged
}
