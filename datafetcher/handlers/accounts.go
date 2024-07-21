package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/esteanes/expense-manager/datafetcher/upclient"
)

type AccountHandler struct {
	*BaseHandler
}

func NewAccountHandler(log *log.Logger, upclient *upclient.APIClient, auth context.Context) *AccountHandler {
	handler := &AccountHandler{}
	handler.BaseHandler = &BaseHandler{
		Uri:      "/accounts",
		Log:      log,
		UpClient: upclient,
		UpAuth:   auth,
		Handler:  handler, // Set the Handler interface to the specific handler
	}
	return handler
}

func (h *AccountHandler) Post(w http.ResponseWriter, r *http.Request) {}
func (h *AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	pageSize := int32(30)
	filterAccountType := upclient.AccountTypeEnum("SAVER")
	filterOwnershipType := upclient.OwnershipTypeEnum("INDIVIDUAL")
	resp, r2, err := h.UpClient.AccountsAPI.AccountsGet(h.UpAuth).PageSize(pageSize).FilterAccountType(filterAccountType).FilterOwnershipType(filterOwnershipType).Execute()

	if err != nil {
		fmt.Fprintf(w, "Error when calling `AccountsAPI.AccountsGet``: %v\n", err)
		fmt.Fprintf(w, "Full HTTP response: %v\n", r2)
		h.Log.Println("Unable to get account information")
	}

	fmt.Fprintf(w, "Response from `AccountsAPI.AccountsGet`: %v\n", resp)
}
