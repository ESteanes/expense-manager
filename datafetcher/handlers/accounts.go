package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/esteanes/expense-manager/datafetcher/templates"
	"github.com/esteanes/expense-manager/datafetcher/upclient"
)

type AccountHandler struct {
	*BaseHandler
}

func NewAccountHandler(log *log.Logger, upclient *upclient.APIClient, auth context.Context) *AccountHandler {
	handler := &AccountHandler{}
	handler.BaseHandler = &BaseHandler{
		Uri:         "/accounts",
		Log:         log,
		UpClient:    upclient,
		UpAuth:      auth,
		Handler:     handler, // Set the Handler interface to the specific handler
		MaxPageSize: int32(100),
	}
	return handler
}

func (h *AccountHandler) Post(w http.ResponseWriter, r *http.Request) {}
func (h *AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	filterOwnershipType := upclient.OwnershipTypeEnum("INDIVIDUAL")
	accountChannel := make(chan upclient.AccountResource, 10)
	clonedChannels := Clone(accountChannel, 2)
	accountChannel1 := <-clonedChannels
	accountChannel2 := <-clonedChannels
	go h.GetAccounts(accountChannel, filterOwnershipType)

	templ.Handler(templates.Accounts(accountChannel1, accountChannel2), templ.WithStreaming()).ServeHTTP(w, r)
}

func (h *AccountHandler) GetAccounts(accountChannel chan upclient.AccountResource, ownershipType upclient.OwnershipTypeEnum) {
	defer close(accountChannel)
	resp, r2, err := h.UpClient.AccountsAPI.AccountsGet(h.UpAuth).PageSize(h.MaxPageSize).FilterOwnershipType(ownershipType).Execute()
	if err != nil {
		h.Log.Println(fmt.Sprintf("Error when calling `AccountsAPI.AccountsGet`: %s\n", err))
		h.Log.Println(fmt.Sprintf("Full HTTP response: %v\n", r2))
		h.Log.Println("Unable to get account information")
	}

	for _, account := range resp.Data {
		accountChannel <- account
	}
}
func Clone[T any](inCh chan T, size int) <-chan <-chan T {
	// The channel of channels to return at the end of this function call
	ret := make(chan (<-chan T), size)

	// This slice keeps track of all the output channels this function will be creating below.
	outChs := make([]chan T, size)

	// Create channels, keep track of them in the slice and send them on the return channel
	for i := 0; i < size; i++ {
		// The buffer size of the newly created channel is the same as the input channel
		outChs[i] = make(chan T, cap(inCh))
		ret <- outChs[i]
	}

	// Start a goroutine to manage receiving message from the input channels and sending out to the output channels
	// Close the output channels if the input channel has been closed.
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
