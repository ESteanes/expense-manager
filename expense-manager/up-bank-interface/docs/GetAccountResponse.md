# GetAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AccountResource**](AccountResource.md) | The account returned in this response.  | 

## Methods

### NewGetAccountResponse

`func NewGetAccountResponse(data AccountResource, ) *GetAccountResponse`

NewGetAccountResponse instantiates a new GetAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountResponseWithDefaults

`func NewGetAccountResponseWithDefaults() *GetAccountResponse`

NewGetAccountResponseWithDefaults instantiates a new GetAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAccountResponse) GetData() AccountResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAccountResponse) GetDataOk() (*AccountResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAccountResponse) SetData(v AccountResource)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


