# TransactionResourceAttributesHoldInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**MoneyObject**](MoneyObject.md) | The amount of this transaction while in the &#x60;HELD&#x60; status, in Australian dollars.  | 
**ForeignAmount** | [**NullableHoldInfoObjectForeignAmount**](HoldInfoObjectForeignAmount.md) |  | 

## Methods

### NewTransactionResourceAttributesHoldInfo

`func NewTransactionResourceAttributesHoldInfo(amount MoneyObject, foreignAmount NullableHoldInfoObjectForeignAmount, ) *TransactionResourceAttributesHoldInfo`

NewTransactionResourceAttributesHoldInfo instantiates a new TransactionResourceAttributesHoldInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResourceAttributesHoldInfoWithDefaults

`func NewTransactionResourceAttributesHoldInfoWithDefaults() *TransactionResourceAttributesHoldInfo`

NewTransactionResourceAttributesHoldInfoWithDefaults instantiates a new TransactionResourceAttributesHoldInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionResourceAttributesHoldInfo) GetAmount() MoneyObject`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionResourceAttributesHoldInfo) GetAmountOk() (*MoneyObject, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionResourceAttributesHoldInfo) SetAmount(v MoneyObject)`

SetAmount sets Amount field to given value.


### GetForeignAmount

`func (o *TransactionResourceAttributesHoldInfo) GetForeignAmount() HoldInfoObjectForeignAmount`

GetForeignAmount returns the ForeignAmount field if non-nil, zero value otherwise.

### GetForeignAmountOk

`func (o *TransactionResourceAttributesHoldInfo) GetForeignAmountOk() (*HoldInfoObjectForeignAmount, bool)`

GetForeignAmountOk returns a tuple with the ForeignAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmount

`func (o *TransactionResourceAttributesHoldInfo) SetForeignAmount(v HoldInfoObjectForeignAmount)`

SetForeignAmount sets ForeignAmount field to given value.


### SetForeignAmountNil

`func (o *TransactionResourceAttributesHoldInfo) SetForeignAmountNil(b bool)`

 SetForeignAmountNil sets the value for ForeignAmount to be an explicit nil

### UnsetForeignAmount
`func (o *TransactionResourceAttributesHoldInfo) UnsetForeignAmount()`

UnsetForeignAmount ensures that no value is present for ForeignAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


