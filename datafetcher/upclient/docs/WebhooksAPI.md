# \WebhooksAPI

All URIs are relative to *https://api.up.com.au/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhooksGet**](WebhooksAPI.md#WebhooksGet) | **Get** /webhooks | List webhooks
[**WebhooksIdDelete**](WebhooksAPI.md#WebhooksIdDelete) | **Delete** /webhooks/{id} | Delete webhook
[**WebhooksIdGet**](WebhooksAPI.md#WebhooksIdGet) | **Get** /webhooks/{id} | Retrieve webhook
[**WebhooksPost**](WebhooksAPI.md#WebhooksPost) | **Post** /webhooks | Create webhook
[**WebhooksWebhookIdLogsGet**](WebhooksAPI.md#WebhooksWebhookIdLogsGet) | **Get** /webhooks/{webhookId}/logs | List webhook logs
[**WebhooksWebhookIdPingPost**](WebhooksAPI.md#WebhooksWebhookIdPingPost) | **Post** /webhooks/{webhookId}/ping | Ping webhook



## WebhooksGet

> ListWebhooksResponse WebhooksGet(ctx).PageSize(pageSize).Execute()

List webhooks



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
	pageSize := int32(30) // int32 | The number of records to return in each page.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.WebhooksGet(context.Background()).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksGet`: ListWebhooksResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.WebhooksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of records to return in each page.  | 

### Return type

[**ListWebhooksResponse**](ListWebhooksResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksIdDelete

> WebhooksIdDelete(ctx, id).Execute()

Delete webhook



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
	id := "be530308-60e2-489b-9c73-ab0db6fd132d" // string | The unique identifier for the webhook. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the webhook.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksIdGet

> GetWebhookResponse WebhooksIdGet(ctx, id).Execute()

Retrieve webhook



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
	id := "48984142-bb60-4fd9-8db2-ed72cfc2a0ba" // string | The unique identifier for the webhook. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.WebhooksIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksIdGet`: GetWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.WebhooksIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the webhook.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWebhookResponse**](GetWebhookResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksPost

> CreateWebhookResponse WebhooksPost(ctx).CreateWebhookRequest(createWebhookRequest).Execute()

Create webhook



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
	createWebhookRequest := *openapiclient.NewCreateWebhookRequest(*openapiclient.NewWebhookInputResource(*openapiclient.NewWebhookInputResourceAttributes("Url_example"))) // CreateWebhookRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.WebhooksPost(context.Background()).CreateWebhookRequest(createWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksPost`: CreateWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.WebhooksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhookRequest** | [**CreateWebhookRequest**](CreateWebhookRequest.md) |  | 

### Return type

[**CreateWebhookResponse**](CreateWebhookResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksWebhookIdLogsGet

> ListWebhookDeliveryLogsResponse WebhooksWebhookIdLogsGet(ctx, webhookId).PageSize(pageSize).Execute()

List webhook logs



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
	webhookId := "e579c0fe-62e7-47de-b436-4b1dfe110b35" // string | The unique identifier for the webhook. 
	pageSize := int32(30) // int32 | The number of records to return in each page.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.WebhooksWebhookIdLogsGet(context.Background(), webhookId).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksWebhookIdLogsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksWebhookIdLogsGet`: ListWebhookDeliveryLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.WebhooksWebhookIdLogsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | The unique identifier for the webhook.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksWebhookIdLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of records to return in each page.  | 

### Return type

[**ListWebhookDeliveryLogsResponse**](ListWebhookDeliveryLogsResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksWebhookIdPingPost

> WebhookEventCallback WebhooksWebhookIdPingPost(ctx, webhookId).Execute()

Ping webhook



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
	webhookId := "ca59a175-79fa-4467-867f-b0cf582ee6bd" // string | The unique identifier for the webhook. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.WebhooksWebhookIdPingPost(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksWebhookIdPingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksWebhookIdPingPost`: WebhookEventCallback
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.WebhooksWebhookIdPingPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | The unique identifier for the webhook.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksWebhookIdPingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookEventCallback**](WebhookEventCallback.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

