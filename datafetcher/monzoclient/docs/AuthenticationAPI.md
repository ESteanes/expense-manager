# \AuthenticationAPI

All URIs are relative to *https://api.monzo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExchangeToken**](AuthenticationAPI.md#ExchangeToken) | **Post** /oauth2/token | Exchange authorization code or refresh token
[**Logout**](AuthenticationAPI.md#Logout) | **Post** /oauth2/logout | Log out and invalidate access token
[**Whoami**](AuthenticationAPI.md#Whoami) | **Get** /ping/whoami | Get information about access token



## ExchangeToken

> TokenResponse ExchangeToken(ctx).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).RedirectUri(redirectUri).Code(code).RefreshToken(refreshToken).Execute()

Exchange authorization code or refresh token



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
	grantType := "grantType_example" // string | Must be set to refresh_token (optional)
	clientId := "clientId_example" // string | Your client ID (optional)
	clientSecret := "clientSecret_example" // string | Your client secret (optional)
	redirectUri := "redirectUri_example" // string | The URL in your app where users were sent after authorisation (optional)
	code := "code_example" // string | The authorization code received when the user was redirected back (optional)
	refreshToken := "refreshToken_example" // string | The refresh token received along with the original access token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.ExchangeToken(context.Background()).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).RedirectUri(redirectUri).Code(code).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ExchangeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExchangeToken`: TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.ExchangeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | Must be set to refresh_token | 
 **clientId** | **string** | Your client ID | 
 **clientSecret** | **string** | Your client secret | 
 **redirectUri** | **string** | The URL in your app where users were sent after authorisation | 
 **code** | **string** | The authorization code received when the user was redirected back | 
 **refreshToken** | **string** | The refresh token received along with the original access token | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> map[string]interface{} Logout(ctx).Execute()

Log out and invalidate access token



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Logout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Logout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Logout`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Logout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## Whoami

> WhoamiResponse Whoami(ctx).Execute()

Get information about access token



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Whoami(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Whoami``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Whoami`: WhoamiResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Whoami`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWhoamiRequest struct via the builder pattern


### Return type

[**WhoamiResponse**](WhoamiResponse.md)

### Authorization

[MonzoOAuth2](../README.md#MonzoOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

