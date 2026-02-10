# \ReceiptsAPI

All URIs are relative to *https://api.monzo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReceipt**](ReceiptsAPI.md#CreateReceipt) | **Put** /transaction-receipts | Create receipt
[**DeleteReceipt**](ReceiptsAPI.md#DeleteReceipt) | **Delete** /transaction-receipts | Delete receipt
[**GetReceipt**](ReceiptsAPI.md#GetReceipt) | **Get** /transaction-receipts | Retrieve receipt



## CreateReceipt

> CreateReceipt200Response CreateReceipt(ctx).Receipt(receipt).Execute()

Create receipt



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
	receipt := *openapiclient.NewReceipt("tx_00008zIcpb1TB4yeIFXMzx", "Order-12345678", int64(1299), "GBP", []openapiclient.ReceiptItem{*openapiclient.NewReceiptItem("Burger", int64(539), "GBP")}) // Receipt | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceiptsAPI.CreateReceipt(context.Background()).Receipt(receipt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsAPI.CreateReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReceipt`: CreateReceipt200Response
	fmt.Fprintf(os.Stdout, "Response from `ReceiptsAPI.CreateReceipt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **receipt** | [**Receipt**](Receipt.md) |  | 

### Return type

[**CreateReceipt200Response**](CreateReceipt200Response.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReceipt

> map[string]interface{} DeleteReceipt(ctx).ExternalId(externalId).Execute()

Delete receipt



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
	externalId := "test-receipt-1" // string | The external ID of the receipt

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceiptsAPI.DeleteReceipt(context.Background()).ExternalId(externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsAPI.DeleteReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteReceipt`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReceiptsAPI.DeleteReceipt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalId** | **string** | The external ID of the receipt | 

### Return type

**map[string]interface{}**

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceipt

> GetReceipt200Response GetReceipt(ctx).ExternalId(externalId).Execute()

Retrieve receipt



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
	externalId := "test-receipt-1" // string | The external ID of the receipt

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceiptsAPI.GetReceipt(context.Background()).ExternalId(externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsAPI.GetReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReceipt`: GetReceipt200Response
	fmt.Fprintf(os.Stdout, "Response from `ReceiptsAPI.GetReceipt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalId** | **string** | The external ID of the receipt | 

### Return type

[**GetReceipt200Response**](GetReceipt200Response.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

