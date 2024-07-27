package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/esteanes/expense-manager/datafetcher/upclient"
)

type TransactionsHandler struct {
	*BaseHandler
}

func NewTransactionHandler(log *log.Logger, upclient *upclient.APIClient, auth context.Context) *TransactionsHandler {
	handler := &TransactionsHandler{}
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
		fmt.Fprintf(w, "Error when calling `TransactionsAPI.TransactionsGet``: %v\n", err)
		fmt.Fprintf(w, "Full HTTP response: %v\n", r2.Body)
		h.Log.Println(r2.Body)
	}
	// response from `TransactionsGet`: ListTransactionsResponse
	fmt.Fprintf(w, "Response from `TransactionsAPI.TransactionsGet`: %v\n", resp2)
	jsonString := `{"account":{"data":{"type":"accounts","id":"a90b55ad-1bcb-4e75-b407-0e0e1e5c8a6d"},"links":{"related":"https://api.up.com.au/api/v1/accounts/a90b55ad-1bcb-4e75-b407-0e0e1e5c8a6d"}},"transferAccount":{"data":null},"category":{"data":{"type":"categories","id":"public-transport"},"links":{"self":"https://api.up.com.au/api/v1/transactions/e33e57c4-efa9-496d-83bd-e1f8c0b9651c/relationships/category","related":"https://api.up.com.au/api/v1/categories/public-transport"}},"parentCategory":{"data":{"type":"categories","id":"transport"},"links":{"related":"https://api.up.com.au/api/v1/categories/transport"}},"tags":{"data":[],"links":{"self":"https://api.up.com.au/api/v1/transactions/e33e57c4-efa9-496d-83bd-e1f8c0b9651c/relationships/tags"}},"attachment":{"data":null}}`
	var manualResponse upclient.TransactionResourceRelationships
	err = json.Unmarshal([]byte(jsonString), &manualResponse)
	fmt.Fprintf(w, "Manual json unmarshalling: %v\n", manualResponse)
	if err != nil {
		fmt.Fprintf(w, "error unmarshalling: %s\n", err)
	}
}
