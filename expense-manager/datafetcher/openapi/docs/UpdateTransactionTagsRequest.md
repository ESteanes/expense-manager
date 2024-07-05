# UpdateTransactionTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TagInputResourceIdentifier**](TagInputResourceIdentifier.md) | The tags to add to or remove from the transaction.  | 

## Methods

### NewUpdateTransactionTagsRequest

`func NewUpdateTransactionTagsRequest(data []TagInputResourceIdentifier, ) *UpdateTransactionTagsRequest`

NewUpdateTransactionTagsRequest instantiates a new UpdateTransactionTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransactionTagsRequestWithDefaults

`func NewUpdateTransactionTagsRequestWithDefaults() *UpdateTransactionTagsRequest`

NewUpdateTransactionTagsRequestWithDefaults instantiates a new UpdateTransactionTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UpdateTransactionTagsRequest) GetData() []TagInputResourceIdentifier`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateTransactionTagsRequest) GetDataOk() (*[]TagInputResourceIdentifier, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateTransactionTagsRequest) SetData(v []TagInputResourceIdentifier)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


