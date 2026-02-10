# \TransactionsAPI

All URIs are relative to *https://api.monzo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnnotateTransaction**](TransactionsAPI.md#AnnotateTransaction) | **Patch** /transactions/{transaction_id} | Annotate transaction
[**GetTransaction**](TransactionsAPI.md#GetTransaction) | **Get** /transactions/{transaction_id} | Retrieve transaction
[**ListTransactions**](TransactionsAPI.md#ListTransactions) | **Get** /transactions | List transactions



## AnnotateTransaction

> GetTransaction200Response AnnotateTransaction(ctx, transactionId).Execute()

Annotate transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/esteanes/up-bank-go/datafetcher/monzoclient"
)

func main() {
	transactionId := "tx_00008zIcpb1TB4yeIFXMzx" // string | The id of the transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.AnnotateTransaction(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.AnnotateTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnnotateTransaction`: GetTransaction200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.AnnotateTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The id of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnnotateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTransaction200Response**](GetTransaction200Response.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransaction

> GetTransaction200Response GetTransaction(ctx, transactionId).Expand(expand).Execute()

Retrieve transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/esteanes/up-bank-go/datafetcher/monzoclient"
)

func main() {
	transactionId := "tx_00008zIcpb1TB4yeIFXMzx" // string | The id of the transaction
	expand := "expand_example" // string | Expand merchant inline (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransaction(context.Background(), transactionId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransaction`: GetTransaction200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The id of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Expand merchant inline | 

### Return type

[**GetTransaction200Response**](GetTransaction200Response.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> TransactionsResponse ListTransactions(ctx).AccountId(accountId).Since(since).Before(before).Limit(limit).Expand(expand).Execute()

List transactions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/esteanes/up-bank-go/datafetcher/monzoclient"
)

func main() {
	accountId := "acc_00009237aqC8c5umZmrRdh" // string | The account to retrieve transactions from
	since := "2009-11-10T23:00:00Z" // string | Start time as RFC3339 encoded timestamp or object id (optional)
	before := "2009-11-10T23:00:00Z" // string | End time as RFC3339 encoded timestamp (optional)
	limit := int32(56) // int32 | Limits the number of results per-page (optional) (default to 30)
	expand := "expand_example" // string | Expand merchant inline (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.ListTransactions(context.Background()).AccountId(accountId).Since(since).Before(before).Limit(limit).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransactions`: TransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The account to retrieve transactions from | 
 **since** | **string** | Start time as RFC3339 encoded timestamp or object id | 
 **before** | **string** | End time as RFC3339 encoded timestamp | 
 **limit** | **int32** | Limits the number of results per-page | [default to 30]
 **expand** | **string** | Expand merchant inline | 

### Return type

[**TransactionsResponse**](TransactionsResponse.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

