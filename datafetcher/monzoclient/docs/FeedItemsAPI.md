# \FeedItemsAPI

All URIs are relative to *https://api.monzo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeedItem**](FeedItemsAPI.md#CreateFeedItem) | **Post** /feed | Create feed item



## CreateFeedItem

> map[string]interface{} CreateFeedItem(ctx).AccountId(accountId).Type_(type_).Url(url).ParamsTitle(paramsTitle).ParamsImageUrl(paramsImageUrl).ParamsBody(paramsBody).ParamsBackgroundColor(paramsBackgroundColor).ParamsTitleColor(paramsTitleColor).ParamsBodyColor(paramsBodyColor).Execute()

Create feed item



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
	accountId := "accountId_example" // string | The account to create a feed item for
	type_ := "type__example" // string | Type of feed item
	url := "url_example" // string | A URL to open when the feed item is tapped (optional)
	paramsTitle := "paramsTitle_example" // string | The title to display (optional)
	paramsImageUrl := "paramsImageUrl_example" // string | URL of the image to display (optional)
	paramsBody := "paramsBody_example" // string | The body text of the feed item (optional)
	paramsBackgroundColor := "paramsBackgroundColor_example" // string | Hex value for the background colour (optional)
	paramsTitleColor := "paramsTitleColor_example" // string | Hex value for the title text colour (optional)
	paramsBodyColor := "paramsBodyColor_example" // string | Hex value for the body text colour (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedItemsAPI.CreateFeedItem(context.Background()).AccountId(accountId).Type_(type_).Url(url).ParamsTitle(paramsTitle).ParamsImageUrl(paramsImageUrl).ParamsBody(paramsBody).ParamsBackgroundColor(paramsBackgroundColor).ParamsTitleColor(paramsTitleColor).ParamsBodyColor(paramsBodyColor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedItemsAPI.CreateFeedItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFeedItem`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FeedItemsAPI.CreateFeedItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The account to create a feed item for | 
 **type_** | **string** | Type of feed item | 
 **url** | **string** | A URL to open when the feed item is tapped | 
 **paramsTitle** | **string** | The title to display | 
 **paramsImageUrl** | **string** | URL of the image to display | 
 **paramsBody** | **string** | The body text of the feed item | 
 **paramsBackgroundColor** | **string** | Hex value for the background colour | 
 **paramsTitleColor** | **string** | Hex value for the title text colour | 
 **paramsBodyColor** | **string** | Hex value for the body text colour | 

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

