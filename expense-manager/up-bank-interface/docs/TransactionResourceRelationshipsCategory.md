# TransactionResourceRelationshipsCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**NullableCategoryResourceRelationshipsParentData**](CategoryResourceRelationshipsParentData.md) |  | 
**Links** | Pointer to [**TransactionResourceRelationshipsCategoryLinks**](TransactionResourceRelationshipsCategoryLinks.md) |  | [optional] 

## Methods

### NewTransactionResourceRelationshipsCategory

`func NewTransactionResourceRelationshipsCategory(data NullableCategoryResourceRelationshipsParentData, ) *TransactionResourceRelationshipsCategory`

NewTransactionResourceRelationshipsCategory instantiates a new TransactionResourceRelationshipsCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResourceRelationshipsCategoryWithDefaults

`func NewTransactionResourceRelationshipsCategoryWithDefaults() *TransactionResourceRelationshipsCategory`

NewTransactionResourceRelationshipsCategoryWithDefaults instantiates a new TransactionResourceRelationshipsCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransactionResourceRelationshipsCategory) GetData() CategoryResourceRelationshipsParentData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionResourceRelationshipsCategory) GetDataOk() (*CategoryResourceRelationshipsParentData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionResourceRelationshipsCategory) SetData(v CategoryResourceRelationshipsParentData)`

SetData sets Data field to given value.


### SetDataNil

`func (o *TransactionResourceRelationshipsCategory) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *TransactionResourceRelationshipsCategory) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetLinks

`func (o *TransactionResourceRelationshipsCategory) GetLinks() TransactionResourceRelationshipsCategoryLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransactionResourceRelationshipsCategory) GetLinksOk() (*TransactionResourceRelationshipsCategoryLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransactionResourceRelationshipsCategory) SetLinks(v TransactionResourceRelationshipsCategoryLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TransactionResourceRelationshipsCategory) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


