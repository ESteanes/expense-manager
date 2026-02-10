# WhoamiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticated** | Pointer to **bool** | Whether the request is authenticated | [optional] 
**ClientId** | Pointer to **string** | Your client ID | [optional] 
**UserId** | Pointer to **string** | The user ID | [optional] 

## Methods

### NewWhoamiResponse

`func NewWhoamiResponse() *WhoamiResponse`

NewWhoamiResponse instantiates a new WhoamiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoamiResponseWithDefaults

`func NewWhoamiResponseWithDefaults() *WhoamiResponse`

NewWhoamiResponseWithDefaults instantiates a new WhoamiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticated

`func (o *WhoamiResponse) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *WhoamiResponse) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *WhoamiResponse) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *WhoamiResponse) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### GetClientId

`func (o *WhoamiResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *WhoamiResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *WhoamiResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *WhoamiResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetUserId

`func (o *WhoamiResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WhoamiResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WhoamiResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WhoamiResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


