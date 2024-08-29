# \AttachmentsAPI

All URIs are relative to *https://api.up.com.au/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachmentsGet**](AttachmentsAPI.md#AttachmentsGet) | **Get** /attachments | List attachments
[**AttachmentsIdGet**](AttachmentsAPI.md#AttachmentsIdGet) | **Get** /attachments/{id} | Retrieve attachment



## AttachmentsGet

> ListAttachmentsResponse AttachmentsGet(ctx).Execute()

List attachments



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.AttachmentsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentsGet`: ListAttachmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentsGetRequest struct via the builder pattern


### Return type

[**ListAttachmentsResponse**](ListAttachmentsResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentsIdGet

> GetAttachmentResponse AttachmentsIdGet(ctx, id).Execute()

Retrieve attachment



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
	id := "3672d7fb-e56d-4c4a-b546-7c11ddb5e5e7" // string | The unique identifier for the attachment. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.AttachmentsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentsIdGet`: GetAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the attachment.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAttachmentResponse**](GetAttachmentResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

