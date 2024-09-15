package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/a-h/templ"
	"github.com/esteanes/expense-manager/datafetcher/functions"
	"github.com/esteanes/expense-manager/datafetcher/templates"
	"github.com/esteanes/expense-manager/datafetcher/upclient"
)

type TransactionsHandler struct {
	*BaseHandler
	*AccountHandler
}

func NewTransactionHandler(log *log.Logger, upclient *upclient.APIClient, auth context.Context, accountHandler *AccountHandler) *TransactionsHandler {
	handler := &TransactionsHandler{}
	handler.BaseHandler = &BaseHandler{
		Uri:         "/transactions",
		Log:         log,
		UpClient:    upclient,
		UpAuth:      auth,
		Handler:     handler, // Set the Handler interface to the specific handler
		MaxPageSize: int32(100),
	}
	handler.AccountHandler = accountHandler
	return handler
}

func (h *TransactionsHandler) Post(w http.ResponseWriter, r *http.Request) {}
func (h *TransactionsHandler) Get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	numTransactions, err := strconv.ParseInt(queryParams.Get(functions.TransactionNumQueryParam), 10, 32)
	if err != nil {
		numTransactions = int64(10)
	}
	transactionsChannel := make(chan upclient.TransactionResource, numTransactions)
	accountId := queryParams.Get(functions.AccountIdQueryParam)
	if accountId == "" {
		go h.getTransactionsForAllAccounts(transactionsChannel, int32(numTransactions))
	} else {
		go h.getTransactionsForSpecifiedAccount(transactionsChannel, int32(numTransactions), accountId)
	}
	accountsChannel := make(chan upclient.AccountResource)
	go h.AccountHandler.GetAccounts(accountsChannel, upclient.OwnershipTypeEnum("INDIVIDUAL"))
	templ.Handler(templates.Transactions("Transactions", transactionsChannel, accountsChannel, strconv.Itoa(int(numTransactions))), templ.WithStreaming()).ServeHTTP(w, r)

}
func (h *TransactionsHandler) getTransactionsForAllAccounts(transactionsChannel chan upclient.TransactionResource, numTransactions int32) {
	defer close(transactionsChannel)
	getRequest := h.UpClient.TransactionsAPI.TransactionsGet(h.UpAuth).PageSize(h.MaxPageSize)
	var pageAfter *string
	pageAfter = nil
	countTransactions := int32(0)
	for countTransactions < numTransactions {
		if pageAfter != nil {
			pageKeyParsed, err := ExtractPageAfter(*pageAfter)
			if err != nil {
				h.Log.Panicln(fmt.Sprintf("There was an error parsing the pageKey %s", err))
			}
			getRequest = getRequest.PageAfter(pageKeyParsed)
		}
		h.Log.Println(fmt.Sprintf("request struct is %+v", getRequest))
		resp, r2, err := getRequest.Execute()
		if err != nil {
			h.Log.Println(fmt.Sprintf("Error when calling `TransactionsAPI.TransactionsGet``: %s\n", err))
			if r2 != nil {
				h.Log.Println(fmt.Sprintf("Full HTTP response: %s\n", r2.Body))
			}
			h.Log.Println(r2.Body)
			return
		}
		pageAfter = resp.Links.Next.Get()
		if pageAfter != nil {
			h.Log.Println(fmt.Sprintf("page after link is: %s", *pageAfter))
		}
		for _, transaction := range resp.Data {
			if countTransactions < numTransactions {
				transactionsChannel <- transaction
				countTransactions++
			}
		}
	}
	if pageAfter == nil {
		h.Log.Println("You have reached the end of all transactions")
	}
}

func (h *TransactionsHandler) getTransactionsForSpecifiedAccount(transactionsChannel chan upclient.TransactionResource, numTransactions int32, accountId string) {
	defer close(transactionsChannel)
	getRequest := h.UpClient.TransactionsAPI.AccountsAccountIdTransactionsGet(h.UpAuth, accountId).PageSize(h.MaxPageSize)
	var pageAfter *string
	pageAfter = nil
	countTransactions := int32(0)
	for countTransactions < numTransactions {
		if pageAfter != nil {
			pageKeyParsed, err := ExtractPageAfter(*pageAfter)
			if err != nil {
				h.Log.Panicln(fmt.Sprintf("There was an error parsing the pageKey %s", err))
			}
			getRequest = getRequest.PageAfter(pageKeyParsed)
		}
		h.Log.Println(fmt.Sprintf("request struct is %+v", getRequest))
		resp, r2, err := getRequest.Execute()
		if err != nil {
			h.Log.Println(fmt.Sprintf("Error when calling `TransactionsAPI.TransactionsGet``: %s\n", err))
			if r2 != nil {
				h.Log.Println(fmt.Sprintf("Full HTTP response: %s\n", r2.Body))
			}
			h.Log.Println(r2.Body)
			return
		}
		pageAfter = resp.Links.Next.Get()
		if pageAfter != nil {
			h.Log.Println(fmt.Sprintf("page after link is: %s", *pageAfter))
		}
		for _, transaction := range resp.Data {
			if countTransactions < numTransactions {
				transactionsChannel <- transaction
				countTransactions++
			}
		}
	}
	if pageAfter == nil {
		h.Log.Println("You have reached the end of all transactions")
	}
}

func ExtractPageAfter(inputURL string) (string, error) {
	// Parse the URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	// Extract the query parameters
	queryParams := parsedURL.Query()

	// Get the value of "page[before]"
	pageAfter := queryParams.Get("page[after]")

	return pageAfter, nil
}
