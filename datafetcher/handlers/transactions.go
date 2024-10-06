package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

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
	queryParams := functions.FetchQueryParams(r.URL.Query())
	transactionsChannel := make(chan upclient.TransactionResource, *queryParams.NumTransactions)
	if queryParams.AccountID == nil {
		go h.getTransactionsForAllAccounts(transactionsChannel, queryParams)
	} else {
		go h.getTransactionsForSpecifiedAccount(transactionsChannel, queryParams)
	}
	accountsChannel := make(chan upclient.AccountResource)
	go h.AccountHandler.GetAccounts(accountsChannel, upclient.OwnershipTypeEnum("INDIVIDUAL"))
	templ.Handler(templates.Transactions("Transactions", transactionsChannel, accountsChannel, queryParams), templ.WithStreaming()).ServeHTTP(w, r)

}
func (h *TransactionsHandler) getTransactionsForAllAccounts(transactionsChannel chan upclient.TransactionResource, queryParams *functions.QueryParams) {
	defer close(transactionsChannel)
	getRequest := h.UpClient.TransactionsAPI.TransactionsGet(h.UpAuth).PageSize(h.MaxPageSize)

	if queryParams.StartDate != nil {
		getRequest.FilterSince(*queryParams.StartDate)
	}
	if queryParams.EndDate != nil {
		getRequest.FilterUntil(*queryParams.EndDate)
	}

	var pageAfter *string
	pageAfter = nil
	countTransactions := int32(0)
	for countTransactions < *queryParams.NumTransactions {
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
			return
		}
		pageAfter = resp.Links.Next.Get()
		if pageAfter != nil {
			h.Log.Println(fmt.Sprintf("page after link is: %s", *pageAfter))
		}
		for _, transaction := range resp.Data {
			if countTransactions < *queryParams.NumTransactions {
				transactionsChannel <- transaction
				countTransactions++
			}
		}
	}
	if pageAfter == nil {
		h.Log.Println("You have reached the end of all transactions")
	}
}

func (h *TransactionsHandler) getTransactionsForSpecifiedAccount(transactionsChannel chan upclient.TransactionResource, queryParams *functions.QueryParams) {
	defer close(transactionsChannel)
	getRequest := h.UpClient.TransactionsAPI.AccountsAccountIdTransactionsGet(h.UpAuth, *queryParams.AccountID).PageSize(h.MaxPageSize)

	if queryParams.StartDate != nil {
		h.Log.Println(fmt.Sprintf("Setting Filter Since to: %s", *queryParams.StartDate))
		getRequest = getRequest.FilterSince(*queryParams.StartDate)
	}
	if queryParams.EndDate != nil {
		h.Log.Println(fmt.Sprintf("Setting Filter Until to: %s", *queryParams.EndDate))
		getRequest = getRequest.FilterUntil(*queryParams.EndDate)
	}

	var pageAfter *string
	pageAfter = nil
	countTransactions := int32(0)
	for countTransactions < *queryParams.NumTransactions {
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
			return
		}
		for _, transaction := range resp.Data {
			if countTransactions < *queryParams.NumTransactions {
				transactionsChannel <- transaction
				countTransactions++
			}
		}
		pageAfter = resp.Links.Next.Get()
		if pageAfter == nil {
			break
		}
		h.Log.Println(fmt.Sprintf("page after link is: %s", *pageAfter))
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
