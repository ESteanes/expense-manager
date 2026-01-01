package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/esteanes/up-bank-go/datafetcher/functions"
	"github.com/esteanes/up-bank-go/datafetcher/templates"
	"github.com/esteanes/up-bank-go/datafetcher/upclient"
)

type TransactionsHandler struct {
	*functions.BaseHandler
	*AccountHandler
	*functions.TransactionsLogic
}

func NewTransactionHandler(log *log.Logger, upclient *upclient.APIClient, auth context.Context, accountHandler *AccountHandler) *TransactionsHandler {
	handler := &TransactionsHandler{}
	handler.BaseHandler = &functions.BaseHandler{
		Uri:         "/transactions",
		Log:         log,
		UpClient:    upclient,
		UpAuth:      auth,
		Handler:     handler, // Set the Handler interface to the specific handler
		MaxPageSize: int32(100),
	}
	handler.AccountHandler = accountHandler
	handler.TransactionsLogic = &functions.TransactionsLogic{BaseHandler: handler.BaseHandler}
	return handler
}

func (h *TransactionsHandler) Post(w http.ResponseWriter, r *http.Request) {}

func (h *TransactionsHandler) Get(w http.ResponseWriter, r *http.Request) {
	queryParams := functions.FetchQueryParams(r.URL.Query())
	transactionsChannel := h.TransactionsLogic.FetchAppropriateTransactions(queryParams)
	accountsChannel := make(chan upclient.AccountResource)
	go h.AccountHandler.GetAccounts(accountsChannel, upclient.OwnershipTypeEnum("INDIVIDUAL"))
	templ.Handler(templates.Transactions("Transactions", transactionsChannel, accountsChannel, queryParams), templ.WithStreaming()).ServeHTTP(w, r)

}
