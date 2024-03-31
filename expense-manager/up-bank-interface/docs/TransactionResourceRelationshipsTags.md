# TransactionResourceRelationshipsTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TransactionResourceRelationshipsTagsDataInner**](TransactionResourceRelationshipsTagsDataInner.md) |  | 
**Links** | Pointer to [**TransactionResourceRelationshipsTagsLinks**](TransactionResourceRelationshipsTagsLinks.md) |  | [optional] 

## Methods

### NewTransactionResourceRelationshipsTags

`func NewTransactionResourceRelationshipsTags(data []TransactionResourceRelationshipsTagsDataInner, ) *TransactionResourceRelationshipsTags`

NewTransactionResourceRelationshipsTags instantiates a new TransactionResourceRelationshipsTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResourceRelationshipsTagsWithDefaults

`func NewTransactionResourceRelationshipsTagsWithDefaults() *TransactionResourceRelationshipsTags`

NewTransactionResourceRelationshipsTagsWithDefaults instantiates a new TransactionResourceRelationshipsTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransactionResourceRelationshipsTags) GetData() []TransactionResourceRelationshipsTagsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionResourceRelationshipsTags) GetDataOk() (*[]TransactionResourceRelationshipsTagsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionResourceRelationshipsTags) SetData(v []TransactionResourceRelationshipsTagsDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *TransactionResourceRelationshipsTags) GetLinks() TransactionResourceRelationshipsTagsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransactionResourceRelationshipsTags) GetLinksOk() (*TransactionResourceRelationshipsTagsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransactionResourceRelationshipsTags) SetLinks(v TransactionResourceRelationshipsTagsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TransactionResourceRelationshipsTags) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


