package functions

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"
)

const (
	AccountIdQueryParam      = "accountId"
	TransactionNumQueryParam = "numTransactions"
	StartDateParam           = "startDate"
	EndDateParam             = "endDate"
	TransactionTypeParam     = "transactionTypes"
)

// QueryParams struct to hold the query parameters
type QueryParams struct {
	AccountID        *string
	TransactionTypes []string
	NumTransactions  *int32
	StartDate        *time.Time
	EndDate          *time.Time
}

func (q *QueryParams) String() string {
	return fmt.Sprintf("QueryParams{AccountID: %s, NumTransactions: %d, StartDate: %s, EndDate: %s}", *q.AccountID, *q.NumTransactions, *q.StartDate, *q.EndDate)
}

// FetchQueryParams extracts query parameters from the url.Values map
func FetchQueryParams(queryParams url.Values) *QueryParams {
	params := &QueryParams{}
	defaultTransactions := int32(20)
	params.NumTransactions = &defaultTransactions
	logger := log.New(os.Stdout, "query params:", log.Default().Flags())

	// Fetch "accountId"
	if accountID := queryParams.Get(AccountIdQueryParam); accountID != "" {
		params.AccountID = &accountID
	}

	// Fetch "transactionTypes"
	params.TransactionTypes = queryParams["transactionTypes"]

	// Fetch "numTransactions"
	if numTransactions := queryParams.Get(TransactionNumQueryParam); numTransactions != "" {
		if num, err := strconv.Atoi(numTransactions); err == nil {
			myInt := int32(num)
			params.NumTransactions = &myInt
		}
	}

	// Fetch "startDate"
	if startDate := queryParams.Get(StartDateParam); startDate != "" {
		log.Println(fmt.Sprintf("Start date is %s", startDate))
		timeStartDate, err := time.Parse(time.RFC3339, startDate)
		if err == nil {
			params.StartDate = &timeStartDate
		}
	}

	// Fetch "endDate"
	if endDate := queryParams.Get(EndDateParam); endDate != "" {
		timeEndDate, err := time.Parse(time.RFC3339, endDate)
		if err == nil {
			params.EndDate = &timeEndDate
		}
	}

	if params.StartDate != nil && params.EndDate != nil {
		maxTransactions := int32(10000)
		params.NumTransactions = &maxTransactions
	}

	logger.Println(fmt.Sprintf("query params are: %v", params))

	return params
}
