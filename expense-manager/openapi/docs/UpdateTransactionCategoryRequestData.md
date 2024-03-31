# UpdateTransactionCategoryRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this resource: &#x60;categories&#x60; | 
**Id** | **string** | The unique identifier of the category, as returned by the &#x60;/categories&#x60; endpoint.  | 

## Methods

### NewUpdateTransactionCategoryRequestData

`func NewUpdateTransactionCategoryRequestData(type_ string, id string, ) *UpdateTransactionCategoryRequestData`

NewUpdateTransactionCategoryRequestData instantiates a new UpdateTransactionCategoryRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransactionCategoryRequestDataWithDefaults

`func NewUpdateTransactionCategoryRequestDataWithDefaults() *UpdateTransactionCategoryRequestData`

NewUpdateTransactionCategoryRequestDataWithDefaults instantiates a new UpdateTransactionCategoryRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateTransactionCategoryRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateTransactionCategoryRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateTransactionCategoryRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *UpdateTransactionCategoryRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTransactionCategoryRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTransactionCategoryRequestData) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


