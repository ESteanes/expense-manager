package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/esteanes/up-bank-go/datafetcher/functions"
	"github.com/esteanes/up-bank-go/datafetcher/upclient"
)

// AccountHandler is the legacy handler for Up Bank accounts
// Deprecated: Use unified.AccountHandler instead
type AccountHandler struct {
	*functions.BaseHandler
}

func NewAccountHandler(log *log.Logger, upclient *upclient.APIClient, auth context.Context) *AccountHandler {
	handler := &AccountHandler{}
	handler.BaseHandler = &functions.BaseHandler{
		Uri:         "/accounts",
		Log:         log,
		UpClient:    upclient,
		UpAuth:      auth,
		Handler:     handler,
		MaxPageSize: int32(100),
	}
	return handler
}

func (h *AccountHandler) Post(w http.ResponseWriter, r *http.Request) {}
func (h *AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Legacy handler - functionality moved to unified handlers
	http.Error(w, "Use unified handlers", http.StatusNotImplemented)
}

func (h *AccountHandler) GetAccounts(accountChannel chan upclient.AccountResource, ownershipType upclient.OwnershipTypeEnum) {
	defer close(accountChannel)
	resp, r2, err := h.UpClient.AccountsAPI.AccountsGet(h.UpAuth).PageSize(h.MaxPageSize).FilterOwnershipType(ownershipType).Execute()
	if err != nil {
		h.Log.Printf("Error when calling `AccountsAPI.AccountsGet`: %s\n", err)
		if r2 != nil {
			h.Log.Printf("Full HTTP response: %v\n", r2)
		}
		h.Log.Println("Unable to get account information")
	}

	for _, account := range resp.Data {
		accountChannel <- account
	}
}

func Clone[T any](inCh chan T, size int) <-chan <-chan T {
	ret := make(chan (<-chan T), size)
	outChs := make([]chan T, size)

	for i := 0; i < size; i++ {
		outChs[i] = make(chan T, cap(inCh))
		ret <- outChs[i]
	}

	go func() {
		for {
			msg, more := <-inCh
			if more {
				for _, ch := range outChs {
					ch <- msg
				}
			} else {
				for _, ch := range outChs {
					close(ch)
				}
				return
			}
		}
	}()

	return ret
}
