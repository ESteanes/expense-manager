# TransactionResourceRelationshipsTransferAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**NullableTransactionResourceRelationshipsTransferAccountData**](TransactionResourceRelationshipsTransferAccountData.md) |  | 
**Links** | Pointer to [**AccountResourceRelationshipsTransactionsLinks**](AccountResourceRelationshipsTransactionsLinks.md) |  | [optional] 

## Methods

### NewTransactionResourceRelationshipsTransferAccount

`func NewTransactionResourceRelationshipsTransferAccount(data NullableTransactionResourceRelationshipsTransferAccountData, ) *TransactionResourceRelationshipsTransferAccount`

NewTransactionResourceRelationshipsTransferAccount instantiates a new TransactionResourceRelationshipsTransferAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResourceRelationshipsTransferAccountWithDefaults

`func NewTransactionResourceRelationshipsTransferAccountWithDefaults() *TransactionResourceRelationshipsTransferAccount`

NewTransactionResourceRelationshipsTransferAccountWithDefaults instantiates a new TransactionResourceRelationshipsTransferAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransactionResourceRelationshipsTransferAccount) GetData() TransactionResourceRelationshipsTransferAccountData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionResourceRelationshipsTransferAccount) GetDataOk() (*TransactionResourceRelationshipsTransferAccountData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionResourceRelationshipsTransferAccount) SetData(v TransactionResourceRelationshipsTransferAccountData)`

SetData sets Data field to given value.


### SetDataNil

`func (o *TransactionResourceRelationshipsTransferAccount) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *TransactionResourceRelationshipsTransferAccount) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetLinks

`func (o *TransactionResourceRelationshipsTransferAccount) GetLinks() AccountResourceRelationshipsTransactionsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransactionResourceRelationshipsTransferAccount) GetLinksOk() (*AccountResourceRelationshipsTransactionsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransactionResourceRelationshipsTransferAccount) SetLinks(v AccountResourceRelationshipsTransactionsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TransactionResourceRelationshipsTransferAccount) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


