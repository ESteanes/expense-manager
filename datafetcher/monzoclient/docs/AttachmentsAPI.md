# \AttachmentsAPI

All URIs are relative to *https://api.monzo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeregisterAttachment**](AttachmentsAPI.md#DeregisterAttachment) | **Post** /attachment/deregister | Deregister attachment
[**RegisterAttachment**](AttachmentsAPI.md#RegisterAttachment) | **Post** /attachment/register | Register attachment
[**UploadAttachment**](AttachmentsAPI.md#UploadAttachment) | **Post** /attachment/upload | Upload attachment



## DeregisterAttachment

> map[string]interface{} DeregisterAttachment(ctx).Id(id).Execute()

Deregister attachment



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
	id := "id_example" // string | The id of the attachment to deregister

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.DeregisterAttachment(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.DeregisterAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeregisterAttachment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.DeregisterAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the attachment to deregister | 

### Return type

**map[string]interface{}**

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAttachment

> RegisterAttachment200Response RegisterAttachment(ctx).ExternalId(externalId).FileUrl(fileUrl).FileType(fileType).Execute()

Register attachment



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
	externalId := "externalId_example" // string | The id of the transaction to associate the attachment with
	fileUrl := "fileUrl_example" // string | The URL of the uploaded attachment
	fileType := "fileType_example" // string | The content type of the attachment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.RegisterAttachment(context.Background()).ExternalId(externalId).FileUrl(fileUrl).FileType(fileType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.RegisterAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAttachment`: RegisterAttachment200Response
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.RegisterAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalId** | **string** | The id of the transaction to associate the attachment with | 
 **fileUrl** | **string** | The URL of the uploaded attachment | 
 **fileType** | **string** | The content type of the attachment | 

### Return type

[**RegisterAttachment200Response**](RegisterAttachment200Response.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadAttachment

> AttachmentUploadResponse UploadAttachment(ctx).FileName(fileName).FileType(fileType).ContentLength(contentLength).Execute()

Upload attachment



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
	fileName := "fileName_example" // string | The name of the file to be uploaded
	fileType := "fileType_example" // string | The content type of the file
	contentLength := int32(56) // int32 | The HTTP Content-Length of the upload request body, in bytes

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.UploadAttachment(context.Background()).FileName(fileName).FileType(fileType).ContentLength(contentLength).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.UploadAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadAttachment`: AttachmentUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.UploadAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileName** | **string** | The name of the file to be uploaded | 
 **fileType** | **string** | The content type of the file | 
 **contentLength** | **int32** | The HTTP Content-Length of the upload request body, in bytes | 

### Return type

[**AttachmentUploadResponse**](AttachmentUploadResponse.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

