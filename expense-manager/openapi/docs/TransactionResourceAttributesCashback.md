# TransactionResourceAttributesCashback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of why this cashback was paid.  | 
**Amount** | [**MoneyObject**](MoneyObject.md) | The total amount of cashback paid, represented as a positive value.  | 

## Methods

### NewTransactionResourceAttributesCashback

`func NewTransactionResourceAttributesCashback(description string, amount MoneyObject, ) *TransactionResourceAttributesCashback`

NewTransactionResourceAttributesCashback instantiates a new TransactionResourceAttributesCashback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResourceAttributesCashbackWithDefaults

`func NewTransactionResourceAttributesCashbackWithDefaults() *TransactionResourceAttributesCashback`

NewTransactionResourceAttributesCashbackWithDefaults instantiates a new TransactionResourceAttributesCashback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TransactionResourceAttributesCashback) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionResourceAttributesCashback) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionResourceAttributesCashback) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmount

`func (o *TransactionResourceAttributesCashback) GetAmount() MoneyObject`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionResourceAttributesCashback) GetAmountOk() (*MoneyObject, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionResourceAttributesCashback) SetAmount(v MoneyObject)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


