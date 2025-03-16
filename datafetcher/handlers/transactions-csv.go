package handlers

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/esteanes/expense-manager/datafetcher/functions"
	"github.com/esteanes/expense-manager/datafetcher/upclient"
)

type TransactionsCsvHandler struct {
	*BaseHandler
	*TransactionsHandler
}

func NewTransactionCsvHandler(log *log.Logger, upclient *upclient.APIClient, auth context.Context, transactionsHandler *TransactionsHandler) *TransactionsCsvHandler {
	handler := &TransactionsCsvHandler{}
	handler.BaseHandler = &BaseHandler{
		Uri:      "/api/v1/transactions/csv",
		Log:      log,
		UpClient: upclient,
		UpAuth:   auth,
		Handler:  handler}
	handler.TransactionsHandler = transactionsHandler
	return handler
}

func (h *TransactionsCsvHandler) Post(w http.ResponseWriter, r *http.Request) {}
func (h *TransactionsCsvHandler) Get(w http.ResponseWriter, r *http.Request) {
	queryParams := functions.FetchQueryParams(r.URL.Query())
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=transactions.csv")
	csvWriter := csv.NewWriter(w)
	transactionsChannel := h.fetchAppropriateTransactions(queryParams)

	header := []string{"Category", "Cost", "empty", "Empty", "rawText", "description", "empty_1", "empty_2", "createdAt", "transactionId"}

	h.Log.Printf("trying to generate a CSV")
	// Print the JSON-formatted response
	if err := csvWriter.Write(header); err != nil {
		http.Error(w, "Error writing CSV header", http.StatusInternalServerError)
		return
	}

	for transaction := range transactionsChannel {
		h.Log.Printf(fmt.Sprintf("transaction is: %v", transaction))
		if queryParams.TransactionTypes != nil {
			transactionType := transaction.Attributes.TransactionType.Get()
			if transactionType == nil {
				continue
			}
			if !slices.Contains(queryParams.TransactionTypes, *transactionType) {
				h.Log.Println("Skipping transaction as it is not part of requested transactionTypes: " + strings.Join(queryParams.TransactionTypes, ", "))
				continue
			}
		}
		rawText := ""
		if transaction.Attributes.RawText.Get() != nil {
			rawText = *transaction.Attributes.RawText.Get()
		}
		relationshipCat := ""
		if relationship, valid := transaction.Relationships.GetCategoryOk(); valid {
			if relationship.Data.Get() != nil {
				relationshipCat = relationship.Data.Get().Id
			}
		}
		record := []string{
			relationshipCat,
			transaction.Attributes.Amount.Value,
			"", "",
			rawText,
			transaction.Attributes.Description,
			"", "",
			transaction.Attributes.CreatedAt.Format("2006-01-02"),
			transaction.Id,
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
