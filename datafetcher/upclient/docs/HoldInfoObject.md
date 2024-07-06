# HoldInfoObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**MoneyObject**](MoneyObject.md) | The amount of this transaction while in the &#x60;HELD&#x60; status, in Australian dollars.  | 
**ForeignAmount** | [**NullableMoneyObject**](MoneyObject.md) | The foreign currency amount of this transaction while in the &#x60;HELD&#x60; status. This field will be &#x60;null&#x60; for domestic transactions. The amount was converted to the AUD amount reflected in the &#x60;amount&#x60; field.  | 

## Methods

### NewHoldInfoObject

`func NewHoldInfoObject(amount MoneyObject, foreignAmount NullableMoneyObject, ) *HoldInfoObject`

NewHoldInfoObject instantiates a new HoldInfoObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldInfoObjectWithDefaults

`func NewHoldInfoObjectWithDefaults() *HoldInfoObject`

NewHoldInfoObjectWithDefaults instantiates a new HoldInfoObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *HoldInfoObject) GetAmount() MoneyObject`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *HoldInfoObject) GetAmountOk() (*MoneyObject, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *HoldInfoObject) SetAmount(v MoneyObject)`

SetAmount sets Amount field to given value.


### GetForeignAmount

`func (o *HoldInfoObject) GetForeignAmount() MoneyObject`

GetForeignAmount returns the ForeignAmount field if non-nil, zero value otherwise.

### GetForeignAmountOk

`func (o *HoldInfoObject) GetForeignAmountOk() (*MoneyObject, bool)`

GetForeignAmountOk returns a tuple with the ForeignAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmount

`func (o *HoldInfoObject) SetForeignAmount(v MoneyObject)`

SetForeignAmount sets ForeignAmount field to given value.


### SetForeignAmountNil

`func (o *HoldInfoObject) SetForeignAmountNil(b bool)`

 SetForeignAmountNil sets the value for ForeignAmount to be an explicit nil

### UnsetForeignAmount
`func (o *HoldInfoObject) UnsetForeignAmount()`

UnsetForeignAmount ensures that no value is present for ForeignAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


