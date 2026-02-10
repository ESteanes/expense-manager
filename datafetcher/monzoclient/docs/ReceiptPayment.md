# ReceiptPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The payment type | 
**Amount** | **int64** | Amount paid in minor units (pennies) | 
**Currency** | **string** | The currency code | 
**Bin** | Pointer to **string** | Bank identification number (for card payments) | [optional] 
**LastFour** | Pointer to **string** | Last four digits of card (for card payments) | [optional] 
**AuthCode** | Pointer to **string** | Authorization code (for card payments) | [optional] 
**Aid** | Pointer to **string** | Application identifier (for card payments) | [optional] 
**Mid** | Pointer to **string** | Merchant identifier (for card payments) | [optional] 
**Tid** | Pointer to **string** | Terminal identifier (for card payments) | [optional] 
**GiftCardType** | Pointer to **string** | Gift card type description (for gift_card payments) | [optional] 

## Methods

### NewReceiptPayment

`func NewReceiptPayment(type_ string, amount int64, currency string, ) *ReceiptPayment`

NewReceiptPayment instantiates a new ReceiptPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptPaymentWithDefaults

`func NewReceiptPaymentWithDefaults() *ReceiptPayment`

NewReceiptPaymentWithDefaults instantiates a new ReceiptPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReceiptPayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReceiptPayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReceiptPayment) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *ReceiptPayment) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ReceiptPayment) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ReceiptPayment) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *ReceiptPayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReceiptPayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReceiptPayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBin

`func (o *ReceiptPayment) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *ReceiptPayment) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *ReceiptPayment) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *ReceiptPayment) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetLastFour

`func (o *ReceiptPayment) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *ReceiptPayment) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *ReceiptPayment) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *ReceiptPayment) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### GetAuthCode

`func (o *ReceiptPayment) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ReceiptPayment) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ReceiptPayment) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *ReceiptPayment) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetAid

`func (o *ReceiptPayment) GetAid() string`

GetAid returns the Aid field if non-nil, zero value otherwise.

### GetAidOk

`func (o *ReceiptPayment) GetAidOk() (*string, bool)`

GetAidOk returns a tuple with the Aid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAid

`func (o *ReceiptPayment) SetAid(v string)`

SetAid sets Aid field to given value.

### HasAid

`func (o *ReceiptPayment) HasAid() bool`

HasAid returns a boolean if a field has been set.

### GetMid

`func (o *ReceiptPayment) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *ReceiptPayment) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *ReceiptPayment) SetMid(v string)`

SetMid sets Mid field to given value.

### HasMid

`func (o *ReceiptPayment) HasMid() bool`

HasMid returns a boolean if a field has been set.

### GetTid

`func (o *ReceiptPayment) GetTid() string`

GetTid returns the Tid field if non-nil, zero value otherwise.

### GetTidOk

`func (o *ReceiptPayment) GetTidOk() (*string, bool)`

GetTidOk returns a tuple with the Tid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTid

`func (o *ReceiptPayment) SetTid(v string)`

SetTid sets Tid field to given value.

### HasTid

`func (o *ReceiptPayment) HasTid() bool`

HasTid returns a boolean if a field has been set.

### GetGiftCardType

`func (o *ReceiptPayment) GetGiftCardType() string`

GetGiftCardType returns the GiftCardType field if non-nil, zero value otherwise.

### GetGiftCardTypeOk

`func (o *ReceiptPayment) GetGiftCardTypeOk() (*string, bool)`

GetGiftCardTypeOk returns a tuple with the GiftCardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardType

`func (o *ReceiptPayment) SetGiftCardType(v string)`

SetGiftCardType sets GiftCardType field to given value.

### HasGiftCardType

`func (o *ReceiptPayment) HasGiftCardType() bool`

HasGiftCardType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


