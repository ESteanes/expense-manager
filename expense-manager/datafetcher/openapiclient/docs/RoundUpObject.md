# RoundUpObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**MoneyObject**](MoneyObject.md) | The total amount of this Round Up, including any boosts, represented as a negative value.  | 
**BoostPortion** | [**NullableRoundUpObjectBoostPortion**](RoundUpObjectBoostPortion.md) |  | 

## Methods

### NewRoundUpObject

`func NewRoundUpObject(amount MoneyObject, boostPortion NullableRoundUpObjectBoostPortion, ) *RoundUpObject`

NewRoundUpObject instantiates a new RoundUpObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoundUpObjectWithDefaults

`func NewRoundUpObjectWithDefaults() *RoundUpObject`

NewRoundUpObjectWithDefaults instantiates a new RoundUpObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RoundUpObject) GetAmount() MoneyObject`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RoundUpObject) GetAmountOk() (*MoneyObject, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RoundUpObject) SetAmount(v MoneyObject)`

SetAmount sets Amount field to given value.


### GetBoostPortion

`func (o *RoundUpObject) GetBoostPortion() RoundUpObjectBoostPortion`

GetBoostPortion returns the BoostPortion field if non-nil, zero value otherwise.

### GetBoostPortionOk

`func (o *RoundUpObject) GetBoostPortionOk() (*RoundUpObjectBoostPortion, bool)`

GetBoostPortionOk returns a tuple with the BoostPortion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostPortion

`func (o *RoundUpObject) SetBoostPortion(v RoundUpObjectBoostPortion)`

SetBoostPortion sets BoostPortion field to given value.


### SetBoostPortionNil

`func (o *RoundUpObject) SetBoostPortionNil(b bool)`

 SetBoostPortionNil sets the value for BoostPortion to be an explicit nil

### UnsetBoostPortion
`func (o *RoundUpObject) UnsetBoostPortion()`

UnsetBoostPortion ensures that no value is present for BoostPortion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


