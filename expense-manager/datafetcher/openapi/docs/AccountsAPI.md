# \AccountsAPI

All URIs are relative to *https://api.up.com.au/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsGet**](AccountsAPI.md#AccountsGet) | **Get** /accounts | List accounts
[**AccountsIdGet**](AccountsAPI.md#AccountsIdGet) | **Get** /accounts/{id} | Retrieve account



## AccountsGet

> ListAccountsResponse AccountsGet(ctx).PageSize(pageSize).FilterAccountType(filterAccountType).FilterOwnershipType(filterOwnershipType).Execute()

List accounts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pageSize := int32(30) // int32 | The number of records to return in each page.  (optional)
	filterAccountType := openapiclient.AccountTypeEnum("SAVER") // AccountTypeEnum | The type of account for which to return records. This can be used to filter Savers from spending accounts.  (optional)
	filterOwnershipType := openapiclient.OwnershipTypeEnum("INDIVIDUAL") // OwnershipTypeEnum | The account ownership structure for which to return records. This can be used to filter 2Up accounts from Up accounts.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.AccountsGet(context.Background()).PageSize(pageSize).FilterAccountType(filterAccountType).FilterOwnershipType(filterOwnershipType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsGet`: ListAccountsResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of records to return in each page.  | 
 **filterAccountType** | [**AccountTypeEnum**](AccountTypeEnum.md) | The type of account for which to return records. This can be used to filter Savers from spending accounts.  | 
 **filterOwnershipType** | [**OwnershipTypeEnum**](OwnershipTypeEnum.md) | The account ownership structure for which to return records. This can be used to filter 2Up accounts from Up accounts.  | 

### Return type

[**ListAccountsResponse**](ListAccountsResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsIdGet

> GetAccountResponse AccountsIdGet(ctx, id).Execute()

Retrieve account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "92b41408-6b7b-4fca-982b-3fb1fdd77220" // string | The unique identifier for the account. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.AccountsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsIdGet`: GetAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the account.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAccountResponse**](GetAccountResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

