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
)

// TagsAPIService TagsAPI service
type TagsAPIService service

type ApiTagsGetRequest struct {
	ctx        context.Context
	ApiService *TagsAPIService
	pageSize   *int32
}

// The number of records to return in each page.
func (r ApiTagsGetRequest) PageSize(pageSize int32) ApiTagsGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiTagsGetRequest) Execute() (*ListTagsResponse, *http.Response, error) {
	return r.ApiService.TagsGetExecute(r)
}

/*
TagsGet List tags

Retrieve a list of all tags currently in use. The returned list is
[paginated](#pagination) and can be scrolled by following the `next`
and `prev` links where present. Results are ordered lexicographically.
The `transactions` relationship for each tag exposes a link
to get the transactions with the given tag.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiTagsGetRequest
*/
func (a *TagsAPIService) TagsGet(ctx context.Context) ApiTagsGetRequest {
	return ApiTagsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListTagsResponse
func (a *TagsAPIService) TagsGetExecute(r ApiTagsGetRequest) (*ListTagsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListTagsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.TagsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[size]", r.pageSize, "")
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

type ApiTransactionsTransactionIdRelationshipsTagsDeleteRequest struct {
	ctx                          context.Context
	ApiService                   *TagsAPIService
	transactionId                string
	updateTransactionTagsRequest *UpdateTransactionTagsRequest
}

func (r ApiTransactionsTransactionIdRelationshipsTagsDeleteRequest) UpdateTransactionTagsRequest(updateTransactionTagsRequest UpdateTransactionTagsRequest) ApiTransactionsTransactionIdRelationshipsTagsDeleteRequest {
	r.updateTransactionTagsRequest = &updateTransactionTagsRequest
	return r
}

func (r ApiTransactionsTransactionIdRelationshipsTagsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.TransactionsTransactionIdRelationshipsTagsDeleteExecute(r)
}

/*
TransactionsTransactionIdRelationshipsTagsDelete Remove tags from transaction

Disassociates one or more tags from a specific transaction. Tags that are
not associated are silently ignored. An HTTP `204` is returned on
success. The associated tags, along with this request URL, are also
exposed via the `tags` relationship on the transaction resource returned
from `/transactions/{id}`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param transactionId The unique identifier for the transaction.
	@return ApiTransactionsTransactionIdRelationshipsTagsDeleteRequest
*/
func (a *TagsAPIService) TransactionsTransactionIdRelationshipsTagsDelete(ctx context.Context, transactionId string) ApiTransactionsTransactionIdRelationshipsTagsDeleteRequest {
	return ApiTransactionsTransactionIdRelationshipsTagsDeleteRequest{
		ApiService:    a,
		ctx:           ctx,
		transactionId: transactionId,
	}
}

// Execute executes the request
func (a *TagsAPIService) TransactionsTransactionIdRelationshipsTagsDeleteExecute(r ApiTransactionsTransactionIdRelationshipsTagsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.TransactionsTransactionIdRelationshipsTagsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions/{transactionId}/relationships/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", url.PathEscape(parameterValueToString(r.transactionId, "transactionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateTransactionTagsRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiTransactionsTransactionIdRelationshipsTagsPostRequest struct {
	ctx                          context.Context
	ApiService                   *TagsAPIService
	transactionId                string
	updateTransactionTagsRequest *UpdateTransactionTagsRequest
}

func (r ApiTransactionsTransactionIdRelationshipsTagsPostRequest) UpdateTransactionTagsRequest(updateTransactionTagsRequest UpdateTransactionTagsRequest) ApiTransactionsTransactionIdRelationshipsTagsPostRequest {
	r.updateTransactionTagsRequest = &updateTransactionTagsRequest
	return r
}

func (r ApiTransactionsTransactionIdRelationshipsTagsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.TransactionsTransactionIdRelationshipsTagsPostExecute(r)
}

/*
TransactionsTransactionIdRelationshipsTagsPost Add tags to transaction

Associates one or more tags with a specific transaction. No more than 6
tags may be present on any single transaction. Duplicate tags are
silently ignored. An HTTP `204` is returned on success. The associated
tags, along with this request URL, are also exposed via the `tags`
relationship on the transaction resource returned from
`/transactions/{id}`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param transactionId The unique identifier for the transaction.
	@return ApiTransactionsTransactionIdRelationshipsTagsPostRequest
*/
func (a *TagsAPIService) TransactionsTransactionIdRelationshipsTagsPost(ctx context.Context, transactionId string) ApiTransactionsTransactionIdRelationshipsTagsPostRequest {
	return ApiTransactionsTransactionIdRelationshipsTagsPostRequest{
		ApiService:    a,
		ctx:           ctx,
		transactionId: transactionId,
	}
}

// Execute executes the request
func (a *TagsAPIService) TransactionsTransactionIdRelationshipsTagsPostExecute(r ApiTransactionsTransactionIdRelationshipsTagsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.TransactionsTransactionIdRelationshipsTagsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions/{transactionId}/relationships/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", url.PathEscape(parameterValueToString(r.transactionId, "transactionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateTransactionTagsRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
