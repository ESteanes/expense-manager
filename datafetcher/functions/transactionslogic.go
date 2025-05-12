package functions

import (
	"net/url"

	"github.com/esteanes/up-bank-go/datafetcher/upclient"
)

type TransactionsLogic struct {
	BaseHandler *BaseHandler
}

func (h *TransactionsLogic) FetchAppropriateTransactions(queryParams *QueryParams) chan upclient.TransactionResource {
	transactionsChannel := make(chan upclient.TransactionResource, *queryParams.NumTransactions)
	if queryParams.AccountID == nil {
		go h.getTransactionsForAllAccounts(transactionsChannel, queryParams)
	} else {
		go h.getTransactionsForSpecifiedAccount(transactionsChannel, queryParams)
	}
	return transactionsChannel
}

func (h *TransactionsLogic) getTransactionsForSpecifiedAccount(transactionsChannel chan upclient.TransactionResource, queryParams *QueryParams) {
	defer close(transactionsChannel)
	getRequest := h.BaseHandler.UpClient.TransactionsAPI.AccountsAccountIdTransactionsGet(h.BaseHandler.UpAuth, *queryParams.AccountID).PageSize(h.BaseHandler.MaxPageSize)

	if queryParams.StartDate != nil {
		h.BaseHandler.Log.Printf("Setting Filter Since to: %s", *queryParams.StartDate)
		getRequest = getRequest.FilterSince(*queryParams.StartDate)
	}
	if queryParams.EndDate != nil {
		h.BaseHandler.Log.Printf("Setting Filter Until to: %s", *queryParams.EndDate)
		getRequest = getRequest.FilterUntil(*queryParams.EndDate)
	}

	var pageAfter *string
	pageAfter = nil
	countTransactions := int32(0)
	for countTransactions < *queryParams.NumTransactions {
		if pageAfter != nil {
			pageKeyParsed, err := ExtractPageAfter(*pageAfter)
			if err != nil {
				h.BaseHandler.Log.Panicf("There was an error parsing the pageKey %s", err)
			}
			getRequest = getRequest.PageAfter(pageKeyParsed)
		}
		h.BaseHandler.Log.Printf("request struct is %+v", getRequest)
		resp, r2, err := getRequest.Execute()
		if err != nil {
			h.BaseHandler.Log.Printf("Error when calling `TransactionsAPI.TransactionsGet``: %s\n", err)
			if r2 != nil {
				h.BaseHandler.Log.Printf("Full HTTP response: %s\n", r2.Body)
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
		h.BaseHandler.Log.Printf("page after link is: %s", *pageAfter)
	}
	if pageAfter == nil {
		h.BaseHandler.Log.Println("You have reached the end of all transactions")
	}
}

func (h *TransactionsLogic) getTransactionsForAllAccounts(transactionsChannel chan upclient.TransactionResource, queryParams *QueryParams) {
	defer close(transactionsChannel)
	getRequest := h.BaseHandler.UpClient.TransactionsAPI.TransactionsGet(h.BaseHandler.UpAuth).PageSize(h.BaseHandler.MaxPageSize)

	if queryParams.StartDate != nil {
		getRequest = getRequest.FilterSince(*queryParams.StartDate)
	}
	if queryParams.EndDate != nil {
		getRequest = getRequest.FilterUntil(*queryParams.EndDate)
	}

	var pageAfter *string
	pageAfter = nil
	countTransactions := int32(0)
	for countTransactions < *queryParams.NumTransactions {
		if pageAfter != nil {
			pageKeyParsed, err := ExtractPageAfter(*pageAfter)
			if err != nil {
				h.BaseHandler.Log.Panicf("There was an error parsing the pageKey %s", err)
			}
			getRequest = getRequest.PageAfter(pageKeyParsed)
		}
		h.BaseHandler.Log.Printf("request struct is %+v", getRequest)
		resp, r2, err := getRequest.Execute()
		if err != nil {
			h.BaseHandler.Log.Printf("Error when calling `TransactionsAPI.TransactionsGet``: %s\n", err)
			if r2 != nil {
				h.BaseHandler.Log.Printf("Full HTTP response: %s\n", r2.Body)
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
		h.BaseHandler.Log.Printf("page after link is: %s", *pageAfter)
	}
	if pageAfter == nil {
		h.BaseHandler.Log.Println("You have reached the end of all transactions")
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
