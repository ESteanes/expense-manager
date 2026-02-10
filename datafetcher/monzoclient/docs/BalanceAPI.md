# \BalanceAPI

All URIs are relative to *https://api.monzo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalance**](BalanceAPI.md#GetBalance) | **Get** /balance | Read balance



## GetBalance

> BalanceResponse GetBalance(ctx).AccountId(accountId).Execute()

Read balance



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
	accountId := "acc_00009237aqC8c5umZmrRdh" // string | The id of the account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BalanceAPI.GetBalance(context.Background()).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalanceAPI.GetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalance`: BalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `BalanceAPI.GetBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The id of the account | 

### Return type

[**BalanceResponse**](BalanceResponse.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

