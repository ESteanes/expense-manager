/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TransactionsAPIService TransactionsAPI service
type TransactionsAPIService service

type ApiAccountsAccountIdTransactionsGetRequest struct {
	ctx            context.Context
	ApiService     *TransactionsAPIService
	accountId      string
	pageSize       *int32
	filterStatus   *TransactionStatusEnum
	filterSince    *time.Time
	filterUntil    *time.Time
	filterCategory *string
	filterTag      *string
}

// The number of records to return in each page.
func (r ApiAccountsAccountIdTransactionsGetRequest) PageSize(pageSize int32) ApiAccountsAccountIdTransactionsGetRequest {
	r.pageSize = &pageSize
	return r
}

// The transaction status for which to return records. This can be used to filter &#x60;HELD&#x60; transactions from those that are &#x60;SETTLED&#x60;.
func (r ApiAccountsAccountIdTransactionsGetRequest) FilterStatus(filterStatus TransactionStatusEnum) ApiAccountsAccountIdTransactionsGetRequest {
	r.filterStatus = &filterStatus
	return r
}

// The start date-time from which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.
func (r ApiAccountsAccountIdTransactionsGetRequest) FilterSince(filterSince time.Time) ApiAccountsAccountIdTransactionsGetRequest {
	r.filterSince = &filterSince
	return r
}

// The end date-time up to which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.
func (r ApiAccountsAccountIdTransactionsGetRequest) FilterUntil(filterUntil time.Time) ApiAccountsAccountIdTransactionsGetRequest {
	r.filterUntil = &filterUntil
	return r
}

// The category identifier for which to filter transactions. Both parent and child categories can be filtered through this parameter. Providing an invalid category identifier results in a &#x60;404&#x60; response.
func (r ApiAccountsAccountIdTransactionsGetRequest) FilterCategory(filterCategory string) ApiAccountsAccountIdTransactionsGetRequest {
	r.filterCategory = &filterCategory
	return r
}

// A transaction tag to filter for which to return records. If the tag does not exist, zero records are returned and a success response is given.
func (r ApiAccountsAccountIdTransactionsGetRequest) FilterTag(filterTag string) ApiAccountsAccountIdTransactionsGetRequest {
	r.filterTag = &filterTag
	return r
}

func (r ApiAccountsAccountIdTransactionsGetRequest) Execute() (*ListTransactionsResponse, *http.Response, error) {
	return r.ApiService.AccountsAccountIdTransactionsGetExecute(r)
}

/*
AccountsAccountIdTransactionsGet List transactions by account

Retrieve a list of all transactions for a specific account. The returned
list is [paginated](#pagination) and can be scrolled by following the
`next` and `prev` links where present. To narrow the results to a
specific date range pass one or both of `filter[since]` and
`filter[until]` in the query string. These filter parameters
**should not** be used for pagination. Results are ordered newest first
to oldest last.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId The unique identifier for the account.
	@return ApiAccountsAccountIdTransactionsGetRequest
*/
func (a *TransactionsAPIService) AccountsAccountIdTransactionsGet(ctx context.Context, accountId string) ApiAccountsAccountIdTransactionsGetRequest {
	return ApiAccountsAccountIdTransactionsGetRequest{
		ApiService: a,
		ctx:        ctx,
		accountId:  accountId,
	}
}

// Execute executes the request
//
//	@return ListTransactionsResponse
func (a *TransactionsAPIService) AccountsAccountIdTransactionsGetExecute(r ApiAccountsAccountIdTransactionsGetRequest) (*ListTransactionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListTransactionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.AccountsAccountIdTransactionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{accountId}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[size]", r.pageSize, "")
	}
	if r.filterStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[status]", r.filterStatus, "")
	}
	if r.filterSince != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[since]", r.filterSince, "")
	}
	if r.filterUntil != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[until]", r.filterUntil, "")
	}
	if r.filterCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[category]", r.filterCategory, "")
	}
	if r.filterTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[tag]", r.filterTag, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTransactionsGetRequest struct {
	ctx            context.Context
	ApiService     *TransactionsAPIService
	pageSize       *int32
	filterStatus   *TransactionStatusEnum
	filterSince    *time.Time
	filterUntil    *time.Time
	filterCategory *string
	filterTag      *string
}

// The number of records to return in each page.
func (r ApiTransactionsGetRequest) PageSize(pageSize int32) ApiTransactionsGetRequest {
	r.pageSize = &pageSize
	return r
}

// The transaction status for which to return records. This can be used to filter &#x60;HELD&#x60; transactions from those that are &#x60;SETTLED&#x60;.
func (r ApiTransactionsGetRequest) FilterStatus(filterStatus TransactionStatusEnum) ApiTransactionsGetRequest {
	r.filterStatus = &filterStatus
	return r
}

// The start date-time from which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.
func (r ApiTransactionsGetRequest) FilterSince(filterSince time.Time) ApiTransactionsGetRequest {
	r.filterSince = &filterSince
	return r
}

// The end date-time up to which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.
func (r ApiTransactionsGetRequest) FilterUntil(filterUntil time.Time) ApiTransactionsGetRequest {
	r.filterUntil = &filterUntil
	return r
}

// The category identifier for which to filter transactions. Both parent and child categories can be filtered through this parameter. Providing an invalid category identifier results in a &#x60;404&#x60; response.
func (r ApiTransactionsGetRequest) FilterCategory(filterCategory string) ApiTransactionsGetRequest {
	r.filterCategory = &filterCategory
	return r
}

// A transaction tag to filter for which to return records. If the tag does not exist, zero records are returned and a success response is given.
func (r ApiTransactionsGetRequest) FilterTag(filterTag string) ApiTransactionsGetRequest {
	r.filterTag = &filterTag
	return r
}

func (r ApiTransactionsGetRequest) Execute() (*ListTransactionsResponse, *http.Response, error) {
	return r.ApiService.TransactionsGetExecute(r)
}

/*
TransactionsGet List transactions

Retrieve a list of all transactions across all accounts for the currently
authenticated user. The returned list is [paginated](#pagination) and can
be scrolled by following the `next` and `prev` links where present. To
narrow the results to a specific date range pass one or both of
`filter[since]` and `filter[until]` in the query string. These filter
parameters **should not** be used for pagination. Results are ordered
newest first to oldest last.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiTransactionsGetRequest
*/
func (a *TransactionsAPIService) TransactionsGet(ctx context.Context) ApiTransactionsGetRequest {
	return ApiTransactionsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListTransactionsResponse
func (a *TransactionsAPIService) TransactionsGetExecute(r ApiTransactionsGetRequest) (*ListTransactionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListTransactionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.TransactionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[size]", r.pageSize, "")
	}
	if r.filterStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[status]", r.filterStatus, "")
	}
	if r.filterSince != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[since]", r.filterSince, "")
	}
	if r.filterUntil != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[until]", r.filterUntil, "")
	}
	if r.filterCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[category]", r.filterCategory, "")
	}
	if r.filterTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[tag]", r.filterTag, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTransactionsIdGetRequest struct {
	ctx        context.Context
	ApiService *TransactionsAPIService
	id         string
}

func (r ApiTransactionsIdGetRequest) Execute() (*GetTransactionResponse, *http.Response, error) {
	return r.ApiService.TransactionsIdGetExecute(r)
}

/*
TransactionsIdGet Retrieve transaction

Retrieve a specific transaction by providing its unique identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The unique identifier for the transaction.
	@return ApiTransactionsIdGetRequest
*/
func (a *TransactionsAPIService) TransactionsIdGet(ctx context.Context, id string) ApiTransactionsIdGetRequest {
	return ApiTransactionsIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return GetTransactionResponse
func (a *TransactionsAPIService) TransactionsIdGetExecute(r ApiTransactionsIdGetRequest) (*GetTransactionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetTransactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.TransactionsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
