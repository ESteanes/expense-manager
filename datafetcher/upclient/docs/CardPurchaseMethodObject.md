# CardPurchaseMethodObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**CardPurchaseMethodEnum**](CardPurchaseMethodEnum.md) | The type of card purchase.  | 
**CardNumberSuffix** | **NullableString** | The last four digits of the card used for the purchase, if applicable.  | 

## Methods

### NewCardPurchaseMethodObject

`func NewCardPurchaseMethodObject(method CardPurchaseMethodEnum, cardNumberSuffix NullableString, ) *CardPurchaseMethodObject`

NewCardPurchaseMethodObject instantiates a new CardPurchaseMethodObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardPurchaseMethodObjectWithDefaults

`func NewCardPurchaseMethodObjectWithDefaults() *CardPurchaseMethodObject`

NewCardPurchaseMethodObjectWithDefaults instantiates a new CardPurchaseMethodObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *CardPurchaseMethodObject) GetMethod() CardPurchaseMethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CardPurchaseMethodObject) GetMethodOk() (*CardPurchaseMethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CardPurchaseMethodObject) SetMethod(v CardPurchaseMethodEnum)`

SetMethod sets Method field to given value.


### GetCardNumberSuffix

`func (o *CardPurchaseMethodObject) GetCardNumberSuffix() string`

GetCardNumberSuffix returns the CardNumberSuffix field if non-nil, zero value otherwise.

### GetCardNumberSuffixOk

`func (o *CardPurchaseMethodObject) GetCardNumberSuffixOk() (*string, bool)`

GetCardNumberSuffixOk returns a tuple with the CardNumberSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumberSuffix

`func (o *CardPurchaseMethodObject) SetCardNumberSuffix(v string)`

SetCardNumberSuffix sets CardNumberSuffix field to given value.


### SetCardNumberSuffixNil

`func (o *CardPurchaseMethodObject) SetCardNumberSuffixNil(b bool)`

 SetCardNumberSuffixNil sets the value for CardNumberSuffix to be an explicit nil

### UnsetCardNumberSuffix
`func (o *CardPurchaseMethodObject) UnsetCardNumberSuffix()`

UnsetCardNumberSuffix ensures that no value is present for CardNumberSuffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


