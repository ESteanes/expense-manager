# TransactionResourceRelationshipsAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TransactionResourceRelationshipsAccountData**](TransactionResourceRelationshipsAccountData.md) |  | 
**Links** | Pointer to [**AccountResourceRelationshipsTransactionsLinks**](AccountResourceRelationshipsTransactionsLinks.md) |  | [optional] 

## Methods

### NewTransactionResourceRelationshipsAccount

`func NewTransactionResourceRelationshipsAccount(data TransactionResourceRelationshipsAccountData, ) *TransactionResourceRelationshipsAccount`

NewTransactionResourceRelationshipsAccount instantiates a new TransactionResourceRelationshipsAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResourceRelationshipsAccountWithDefaults

`func NewTransactionResourceRelationshipsAccountWithDefaults() *TransactionResourceRelationshipsAccount`

NewTransactionResourceRelationshipsAccountWithDefaults instantiates a new TransactionResourceRelationshipsAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransactionResourceRelationshipsAccount) GetData() TransactionResourceRelationshipsAccountData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionResourceRelationshipsAccount) GetDataOk() (*TransactionResourceRelationshipsAccountData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionResourceRelationshipsAccount) SetData(v TransactionResourceRelationshipsAccountData)`

SetData sets Data field to given value.


### GetLinks

`func (o *TransactionResourceRelationshipsAccount) GetLinks() AccountResourceRelationshipsTransactionsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransactionResourceRelationshipsAccount) GetLinksOk() (*AccountResourceRelationshipsTransactionsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransactionResourceRelationshipsAccount) SetLinks(v AccountResourceRelationshipsTransactionsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TransactionResourceRelationshipsAccount) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


