package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/esteanes/expense-manager/datafetcher/upclient"
)

type TransactionsHandler struct {
	*BaseHandler
}

func NewTransactionHandler(log *log.Logger, upclient *upclient.APIClient, auth context.Context) *AccountHandler {
	handler := &AccountHandler{}
	handler.BaseHandler = &BaseHandler{
		Uri:      "/transactions",
		Log:      log,
		UpClient: upclient,
		UpAuth:   auth,
		Handler:  handler, // Set the Handler interface to the specific handler
	}
	return handler
}

func (h *TransactionsHandler) Post(w http.ResponseWriter, r *http.Request) {}
func (h *TransactionsHandler) Get(w http.ResponseWriter, r *http.Request) {
	pageSize := int32(30) // int32 | The number of records to return in each page.  (optional)
	resp2, r2, err := h.UpClient.TransactionsAPI.TransactionsGet(h.UpAuth).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.TransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r2.Body)
	}
	// response from `TransactionsGet`: ListTransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.TransactionsGet`: %v\n", resp2)
}
