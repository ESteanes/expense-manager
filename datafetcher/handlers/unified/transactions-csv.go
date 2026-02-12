package unified

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

	"github.com/esteanes/up-bank-go/datafetcher/functions"
	"github.com/esteanes/up-bank-go/datafetcher/models"
	"github.com/esteanes/up-bank-go/datafetcher/providers"
)

// TransactionsCsvHandler handles CSV export of transactions
type TransactionsCsvHandler struct {
	Uri        string
	Log        *log.Logger
	Aggregator *providers.Aggregator
}

// NewTransactionsCsvHandler creates a new unified transactions CSV handler
func NewTransactionsCsvHandler(log *log.Logger, aggregator *providers.Aggregator) *TransactionsCsvHandler {
	return &TransactionsCsvHandler{
		Uri:        "/api/v1/transactions/csv",
		Log:        log,
		Aggregator: aggregator,
	}
}

// ServeHTTP implements http.Handler
func (h *TransactionsCsvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

// Get handles GET requests for CSV export
func (h *TransactionsCsvHandler) Get(w http.ResponseWriter, r *http.Request) {
	queryParams := functions.FetchQueryParams(r.URL.Query())
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=transactions.csv")

	csvWriter := csv.NewWriter(w)
	transactionsChannel := h.FetchTransactions(r.Context(), queryParams)

	// Updated header to include provider
	header := []string{"Category", "Cost", "Provider", "Status", "RawText", "Description", "Message", "AccountID", "CreatedAt", "TransactionID", "SettledAt"}

	h.Log.Printf("Generating CSV export")
	if err := csvWriter.Write(header); err != nil {
		http.Error(w, "Error writing CSV header", http.StatusInternalServerError)
		return
	}

	for transaction := range transactionsChannel {
		rawText := ""
		if transaction.RawText != nil {
			rawText = *transaction.RawText
		}

		category := ""
		if transaction.Category != nil {
			category = *transaction.Category
		}

		message := ""
		if transaction.Message != nil {
			message = *transaction.Message
		}

		settledAt := ""
		if transaction.SettledAt != nil {
			settledAt = transaction.SettledAt.Format("2006-01-02")
		}

		record := []string{
			category,
			transaction.Amount.Value,
			transaction.Provider,
			transaction.Status,
			rawText,
			transaction.Description,
			message,
			transaction.AccountID,
			transaction.CreatedAt.Format("2006-01-02"),
			transaction.ID,
			settledAt,
		}

		if err := csvWriter.Write(record); err != nil {
			fmt.Fprintf(w, "Error writing CSV line %v\n", err)
			return
		}
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		fmt.Fprintf(w, "Error flushing CSV writer: %v\n", err)
	}
}

// Post handles POST requests (placeholder)
func (h *TransactionsCsvHandler) Post(w http.ResponseWriter, r *http.Request) {}

// FetchTransactions returns a channel of transactions based on query params
func (h *TransactionsCsvHandler) FetchTransactions(ctx context.Context, queryParams *functions.QueryParams) <-chan models.Transaction {
	params := &providers.QueryParams{
		AccountID:       queryParams.AccountID,
		NumTransactions: queryParams.NumTransactions,
		StartDate:       queryParams.StartDate,
		EndDate:         queryParams.EndDate,
	}
	return h.Aggregator.GetAllTransactions(ctx, params)
}
