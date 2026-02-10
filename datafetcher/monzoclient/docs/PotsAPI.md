# \PotsAPI

All URIs are relative to *https://api.monzo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DepositIntoPot**](PotsAPI.md#DepositIntoPot) | **Put** /pots/{pot_id}/deposit | Deposit into a pot
[**ListPots**](PotsAPI.md#ListPots) | **Get** /pots | List pots
[**WithdrawFromPot**](PotsAPI.md#WithdrawFromPot) | **Put** /pots/{pot_id}/withdraw | Withdraw from a pot



## DepositIntoPot

> Pot DepositIntoPot(ctx, potId).SourceAccountId(sourceAccountId).Amount(amount).DedupeId(dedupeId).Execute()

Deposit into a pot



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
	potId := "pot_0000778xxfgh4iu8z83nWb" // string | The id of the pot
	sourceAccountId := "sourceAccountId_example" // string | The id of the account to withdraw from
	amount := int64(789) // int64 | The amount to deposit in minor units (pennies)
	dedupeId := "dedupeId_example" // string | A unique string used to de-duplicate deposits

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PotsAPI.DepositIntoPot(context.Background(), potId).SourceAccountId(sourceAccountId).Amount(amount).DedupeId(dedupeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PotsAPI.DepositIntoPot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DepositIntoPot`: Pot
	fmt.Fprintf(os.Stdout, "Response from `PotsAPI.DepositIntoPot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**potId** | **string** | The id of the pot | 

### Other Parameters

Other parameters are passed through a pointer to a apiDepositIntoPotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceAccountId** | **string** | The id of the account to withdraw from | 
 **amount** | **int64** | The amount to deposit in minor units (pennies) | 
 **dedupeId** | **string** | A unique string used to de-duplicate deposits | 

### Return type

[**Pot**](Pot.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPots

> PotsResponse ListPots(ctx).CurrentAccountId(currentAccountId).Execute()

List pots



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
	currentAccountId := "acc_00009237aqC8c5umZmrRdh" // string | The account id to retrieve pots for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PotsAPI.ListPots(context.Background()).CurrentAccountId(currentAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PotsAPI.ListPots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPots`: PotsResponse
	fmt.Fprintf(os.Stdout, "Response from `PotsAPI.ListPots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currentAccountId** | **string** | The account id to retrieve pots for | 

### Return type

[**PotsResponse**](PotsResponse.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WithdrawFromPot

> Pot WithdrawFromPot(ctx, potId).DestinationAccountId(destinationAccountId).Amount(amount).DedupeId(dedupeId).Execute()

Withdraw from a pot



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
	potId := "pot_0000778xxfgh4iu8z83nWb" // string | The id of the pot
	destinationAccountId := "destinationAccountId_example" // string | The id of the account to deposit into
	amount := int64(789) // int64 | The amount to withdraw in minor units (pennies)
	dedupeId := "dedupeId_example" // string | A unique string used to de-duplicate withdrawals

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PotsAPI.WithdrawFromPot(context.Background(), potId).DestinationAccountId(destinationAccountId).Amount(amount).DedupeId(dedupeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PotsAPI.WithdrawFromPot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WithdrawFromPot`: Pot
	fmt.Fprintf(os.Stdout, "Response from `PotsAPI.WithdrawFromPot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**potId** | **string** | The id of the pot | 

### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawFromPotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationAccountId** | **string** | The id of the account to deposit into | 
 **amount** | **int64** | The amount to withdraw in minor units (pennies) | 
 **dedupeId** | **string** | A unique string used to de-duplicate withdrawals | 

### Return type

[**Pot**](Pot.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

