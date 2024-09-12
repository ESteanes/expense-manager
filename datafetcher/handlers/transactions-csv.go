package handlers

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

	"github.com/esteanes/expense-manager/datafetcher/upclient"
)

type TransactionsCsvHandler struct {
	*BaseHandler
}

func NewTransactionCsvHandler(log *log.Logger, upclient *upclient.APIClient, auth context.Context) *TransactionsCsvHandler {
	handler := &TransactionsCsvHandler{}
	handler.BaseHandler = &BaseHandler{
		Uri:      "/transactions-csv",
		Log:      log,
		UpClient: upclient,
		UpAuth:   auth,
		Handler:  handler}
	return handler
}

func (h *TransactionsCsvHandler) Post(w http.ResponseWriter, r *http.Request) {}
func (h *TransactionsCsvHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=transactions.csv")
	csvWriter := csv.NewWriter(w)
	pageSize := int32(30) // int32 | The number of records to return in each page.  (optional)
	resp2, _, _ := h.UpClient.TransactionsAPI.TransactionsGet(h.UpAuth).PageSize(pageSize).Execute()

	header := []string{"Transaction Description", "Amount", "Date"}

	h.Log.Printf("trying to generate a CSV")
	// Print the JSON-formatted response
	if err := csvWriter.Write(header); err != nil {
		http.Error(w, "Error writing CSV header", http.StatusInternalServerError)
		return
	}
	for _, transaction := range resp2.Data {
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
