# TransactionResourceAttributesRoundUp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**MoneyObject**](MoneyObject.md) | The total amount of this Round Up, including any boosts, represented as a negative value.  | 
**BoostPortion** | [**NullableRoundUpObjectBoostPortion**](RoundUpObjectBoostPortion.md) |  | 

## Methods

### NewTransactionResourceAttributesRoundUp

`func NewTransactionResourceAttributesRoundUp(amount MoneyObject, boostPortion NullableRoundUpObjectBoostPortion, ) *TransactionResourceAttributesRoundUp`

NewTransactionResourceAttributesRoundUp instantiates a new TransactionResourceAttributesRoundUp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResourceAttributesRoundUpWithDefaults

`func NewTransactionResourceAttributesRoundUpWithDefaults() *TransactionResourceAttributesRoundUp`

NewTransactionResourceAttributesRoundUpWithDefaults instantiates a new TransactionResourceAttributesRoundUp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionResourceAttributesRoundUp) GetAmount() MoneyObject`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionResourceAttributesRoundUp) GetAmountOk() (*MoneyObject, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionResourceAttributesRoundUp) SetAmount(v MoneyObject)`

SetAmount sets Amount field to given value.


### GetBoostPortion

`func (o *TransactionResourceAttributesRoundUp) GetBoostPortion() RoundUpObjectBoostPortion`

GetBoostPortion returns the BoostPortion field if non-nil, zero value otherwise.

### GetBoostPortionOk

`func (o *TransactionResourceAttributesRoundUp) GetBoostPortionOk() (*RoundUpObjectBoostPortion, bool)`

GetBoostPortionOk returns a tuple with the BoostPortion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostPortion

`func (o *TransactionResourceAttributesRoundUp) SetBoostPortion(v RoundUpObjectBoostPortion)`

SetBoostPortion sets BoostPortion field to given value.


### SetBoostPortionNil

`func (o *TransactionResourceAttributesRoundUp) SetBoostPortionNil(b bool)`

 SetBoostPortionNil sets the value for BoostPortion to be an explicit nil

### UnsetBoostPortion
`func (o *TransactionResourceAttributesRoundUp) UnsetBoostPortion()`

UnsetBoostPortion ensures that no value is present for BoostPortion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


