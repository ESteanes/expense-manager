# \TagsAPI

All URIs are relative to *https://api.up.com.au/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagsGet**](TagsAPI.md#TagsGet) | **Get** /tags | List tags
[**TransactionsTransactionIdRelationshipsTagsDelete**](TagsAPI.md#TransactionsTransactionIdRelationshipsTagsDelete) | **Delete** /transactions/{transactionId}/relationships/tags | Remove tags from transaction
[**TransactionsTransactionIdRelationshipsTagsPost**](TagsAPI.md#TransactionsTransactionIdRelationshipsTagsPost) | **Post** /transactions/{transactionId}/relationships/tags | Add tags to transaction



## TagsGet

> ListTagsResponse TagsGet(ctx).PageSize(pageSize).Execute()

List tags



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/esteanes/expense-manager/datafetcher/upclient"
)

func main() {
	pageSize := int32(50) // int32 | The number of records to return in each page.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.TagsGet(context.Background()).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.TagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagsGet`: ListTagsResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.TagsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of records to return in each page.  | 

### Return type

[**ListTagsResponse**](ListTagsResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsTransactionIdRelationshipsTagsDelete

> TransactionsTransactionIdRelationshipsTagsDelete(ctx, transactionId).UpdateTransactionTagsRequest(updateTransactionTagsRequest).Execute()

Remove tags from transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/esteanes/expense-manager/datafetcher/upclient"
)

func main() {
	transactionId := "c3feb4ba-829c-4482-b882-1b9bd23da82d" // string | The unique identifier for the transaction. 
	updateTransactionTagsRequest := *openapiclient.NewUpdateTransactionTagsRequest([]openapiclient.TagInputResourceIdentifier{*openapiclient.NewTagInputResourceIdentifier("Type_example", "Id_example")}) // UpdateTransactionTagsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.TransactionsTransactionIdRelationshipsTagsDelete(context.Background(), transactionId).UpdateTransactionTagsRequest(updateTransactionTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.TransactionsTransactionIdRelationshipsTagsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The unique identifier for the transaction.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsTransactionIdRelationshipsTagsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTransactionTagsRequest** | [**UpdateTransactionTagsRequest**](UpdateTransactionTagsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsTransactionIdRelationshipsTagsPost

> TransactionsTransactionIdRelationshipsTagsPost(ctx, transactionId).UpdateTransactionTagsRequest(updateTransactionTagsRequest).Execute()

Add tags to transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/esteanes/expense-manager/datafetcher/upclient"
)

func main() {
	transactionId := "acde4631-db56-49a6-aea3-4e2311ef1d6a" // string | The unique identifier for the transaction. 
	updateTransactionTagsRequest := *openapiclient.NewUpdateTransactionTagsRequest([]openapiclient.TagInputResourceIdentifier{*openapiclient.NewTagInputResourceIdentifier("Type_example", "Id_example")}) // UpdateTransactionTagsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.TransactionsTransactionIdRelationshipsTagsPost(context.Background(), transactionId).UpdateTransactionTagsRequest(updateTransactionTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.TransactionsTransactionIdRelationshipsTagsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The unique identifier for the transaction.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsTransactionIdRelationshipsTagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTransactionTagsRequest** | [**UpdateTransactionTagsRequest**](UpdateTransactionTagsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

