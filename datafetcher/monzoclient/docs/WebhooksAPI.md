# \WebhooksAPI

All URIs are relative to *https://api.monzo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhook**](WebhooksAPI.md#DeleteWebhook) | **Delete** /webhooks/{webhook_id} | Delete a webhook
[**ListWebhooks**](WebhooksAPI.md#ListWebhooks) | **Get** /webhooks | List webhooks
[**RegisterWebhook**](WebhooksAPI.md#RegisterWebhook) | **Post** /webhooks | Register a webhook



## DeleteWebhook

> map[string]interface{} DeleteWebhook(ctx, webhookId).Execute()

Delete a webhook



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
	webhookId := "webhook_000091yhhOmrXQaVZ1Irsv" // string | The id of the webhook to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.DeleteWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | The id of the webhook to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListWebhooks

> WebhooksResponse ListWebhooks(ctx).AccountId(accountId).Execute()

List webhooks



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
	accountId := "acc_00009237aqC8c5umZmrRdh" // string | The account to list registered webhooks for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ListWebhooks(context.Background()).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhooks`: WebhooksResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The account to list registered webhooks for | 

### Return type

[**WebhooksResponse**](WebhooksResponse.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterWebhook

> RegisterWebhook200Response RegisterWebhook(ctx).AccountId(accountId).Url(url).Execute()

Register a webhook



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
	accountId := "accountId_example" // string | The account to receive notifications for
	url := "url_example" // string | The URL to send notifications to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.RegisterWebhook(context.Background()).AccountId(accountId).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.RegisterWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterWebhook`: RegisterWebhook200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.RegisterWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The account to receive notifications for | 
 **url** | **string** | The URL to send notifications to | 

### Return type

[**RegisterWebhook200Response**](RegisterWebhook200Response.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

