package handlers

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

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
		Uri:      "/transactions-csv",
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
	transactionsChannel := make(chan upclient.TransactionResource, *queryParams.NumTransactions)
	if queryParams.AccountID == nil {
		go h.getTransactionsForAllAccounts(transactionsChannel, queryParams)
	} else {
		go h.getTransactionsForSpecifiedAccount(transactionsChannel, queryParams)

		header := []string{"Transaction Description", "Amount", "Date"}

		h.Log.Printf("trying to generate a CSV")
		// Print the JSON-formatted response
		if err := csvWriter.Write(header); err != nil {
			http.Error(w, "Error writing CSV header", http.StatusInternalServerError)
			return
		}
		for transaction := range transactionsChannel {
			record := []string{
				transaction.Attributes.Description,
				transaction.Attributes.Amount.Value,
				transaction.Attributes.CreatedAt.String(),
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
}
