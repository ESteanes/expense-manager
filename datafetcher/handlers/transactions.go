package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/a-h/templ"
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
	numTransactions, err := strconv.ParseInt(queryParams.Get("numTransactions"), 10, 32)
	if err != nil {
		numTransactions = int64(10)
	}
	accountId := queryParams.Get("accountId")
	transactionsChannel := make(chan upclient.TransactionResource, numTransactions)
	go h.getTransactions(transactionsChannel, int32(numTransactions), accountId)
	accountsChannel := make(chan upclient.AccountResource)
	go h.AccountHandler.GetAccounts(accountsChannel, upclient.OwnershipTypeEnum("INDIVIDUAL"))
	templ.Handler(templates.Transactions(transactionsChannel, accountsChannel), templ.WithStreaming()).ServeHTTP(w, r)

}
func (h *TransactionsHandler) getTransactions(transactionsChannel chan upclient.TransactionResource, numTransactions int32, accountId string) {
	defer close(transactionsChannel)
	getRequest := h.UpClient.TransactionsAPI.AccountsAccountIdTransactionsGet(h.UpAuth, accountId).PageSize(h.MaxPageSize)
	var pageAfter *string
	pageAfter = nil
	countTransactions := int32(0)
	for countTransactions < numTransactions {
		if pageAfter != nil {
			pageKeyParsed, _ := ExtractPageAfter(*pageAfter)
			getRequest = getRequest.PageAfter(pageKeyParsed)
			h.Log.Println(fmt.Sprintf("setting pageAfter to %s", pageKeyParsed))
		}
		h.Log.Println(fmt.Sprintf("request struct is %+v", getRequest))
		resp, r2, err := getRequest.Execute()
		if err != nil {
			h.Log.Println(fmt.Sprintf("Error when calling `TransactionsAPI.TransactionsGet``: %s\n", err))
			h.Log.Println(fmt.Sprintf("Full HTTP response: %s\n", r2.Body))
			h.Log.Println(r2.Body)
			return
		}
		// h.Log.Println(fmt.Sprintf("resp object is %#v", resp))
		pageAfter = resp.Links.Next.Get()
		h.Log.Println(fmt.Sprintf("pageAfter is %p", pageAfter))
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
