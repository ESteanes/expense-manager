# \UtilityEndpointsAPI

All URIs are relative to *https://api.up.com.au/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UtilPingGet**](UtilityEndpointsAPI.md#UtilPingGet) | **Get** /util/ping | Ping



## UtilPingGet

> PingResponse UtilPingGet(ctx).Execute()

Ping



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ESteanes/expense-manager"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityEndpointsAPI.UtilPingGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityEndpointsAPI.UtilPingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UtilPingGet`: PingResponse
	fmt.Fprintf(os.Stdout, "Response from `UtilityEndpointsAPI.UtilPingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUtilPingGetRequest struct via the builder pattern


### Return type

[**PingResponse**](PingResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

