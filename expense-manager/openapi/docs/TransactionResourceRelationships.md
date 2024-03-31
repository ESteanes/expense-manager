# TransactionResourceRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**TransactionResourceRelationshipsAccount**](TransactionResourceRelationshipsAccount.md) |  | 
**TransferAccount** | [**TransactionResourceRelationshipsTransferAccount**](TransactionResourceRelationshipsTransferAccount.md) |  | 
**Category** | [**TransactionResourceRelationshipsCategory**](TransactionResourceRelationshipsCategory.md) |  | 
**ParentCategory** | [**CategoryResourceRelationshipsParent**](CategoryResourceRelationshipsParent.md) |  | 
**Tags** | [**TransactionResourceRelationshipsTags**](TransactionResourceRelationshipsTags.md) |  | 

## Methods

### NewTransactionResourceRelationships

`func NewTransactionResourceRelationships(account TransactionResourceRelationshipsAccount, transferAccount TransactionResourceRelationshipsTransferAccount, category TransactionResourceRelationshipsCategory, parentCategory CategoryResourceRelationshipsParent, tags TransactionResourceRelationshipsTags, ) *TransactionResourceRelationships`

NewTransactionResourceRelationships instantiates a new TransactionResourceRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResourceRelationshipsWithDefaults

`func NewTransactionResourceRelationshipsWithDefaults() *TransactionResourceRelationships`

NewTransactionResourceRelationshipsWithDefaults instantiates a new TransactionResourceRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TransactionResourceRelationships) GetAccount() TransactionResourceRelationshipsAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TransactionResourceRelationships) GetAccountOk() (*TransactionResourceRelationshipsAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TransactionResourceRelationships) SetAccount(v TransactionResourceRelationshipsAccount)`

SetAccount sets Account field to given value.


### GetTransferAccount

`func (o *TransactionResourceRelationships) GetTransferAccount() TransactionResourceRelationshipsTransferAccount`

GetTransferAccount returns the TransferAccount field if non-nil, zero value otherwise.

### GetTransferAccountOk

`func (o *TransactionResourceRelationships) GetTransferAccountOk() (*TransactionResourceRelationshipsTransferAccount, bool)`

GetTransferAccountOk returns a tuple with the TransferAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAccount

`func (o *TransactionResourceRelationships) SetTransferAccount(v TransactionResourceRelationshipsTransferAccount)`

SetTransferAccount sets TransferAccount field to given value.


### GetCategory

`func (o *TransactionResourceRelationships) GetCategory() TransactionResourceRelationshipsCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TransactionResourceRelationships) GetCategoryOk() (*TransactionResourceRelationshipsCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TransactionResourceRelationships) SetCategory(v TransactionResourceRelationshipsCategory)`

SetCategory sets Category field to given value.


### GetParentCategory

`func (o *TransactionResourceRelationships) GetParentCategory() CategoryResourceRelationshipsParent`

GetParentCategory returns the ParentCategory field if non-nil, zero value otherwise.

### GetParentCategoryOk

`func (o *TransactionResourceRelationships) GetParentCategoryOk() (*CategoryResourceRelationshipsParent, bool)`

GetParentCategoryOk returns a tuple with the ParentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCategory

`func (o *TransactionResourceRelationships) SetParentCategory(v CategoryResourceRelationshipsParent)`

SetParentCategory sets ParentCategory field to given value.


### GetTags

`func (o *TransactionResourceRelationships) GetTags() TransactionResourceRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TransactionResourceRelationships) GetTagsOk() (*TransactionResourceRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TransactionResourceRelationships) SetTags(v TransactionResourceRelationshipsTags)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


