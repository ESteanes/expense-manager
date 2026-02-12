package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/esteanes/up-bank-go/datafetcher/functions"
	"github.com/esteanes/up-bank-go/datafetcher/upclient"
)

// TransactionsHandler is the legacy handler for Up Bank transactions
// Deprecated: Use unified.TransactionsHandler instead
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
		Handler:     handler,
		MaxPageSize: int32(100),
	}
	handler.AccountHandler = accountHandler
	handler.TransactionsLogic = &functions.TransactionsLogic{BaseHandler: handler.BaseHandler}
	return handler
}

func (h *TransactionsHandler) Post(w http.ResponseWriter, r *http.Request) {}

func (h *TransactionsHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Legacy handler - functionality moved to unified handlers
	http.Error(w, "Use unified handlers", http.StatusNotImplemented)
}
