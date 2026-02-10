# ReceiptTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The tax description | 
**Amount** | **int64** | Total amount of the tax in minor units (pennies) | 
**Currency** | **string** | The currency code | 
**TaxNumber** | Pointer to **string** | The tax number | [optional] 

## Methods

### NewReceiptTax

`func NewReceiptTax(description string, amount int64, currency string, ) *ReceiptTax`

NewReceiptTax instantiates a new ReceiptTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptTaxWithDefaults

`func NewReceiptTaxWithDefaults() *ReceiptTax`

NewReceiptTaxWithDefaults instantiates a new ReceiptTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ReceiptTax) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReceiptTax) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReceiptTax) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmount

`func (o *ReceiptTax) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ReceiptTax) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ReceiptTax) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *ReceiptTax) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReceiptTax) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReceiptTax) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTaxNumber

`func (o *ReceiptTax) GetTaxNumber() string`

GetTaxNumber returns the TaxNumber field if non-nil, zero value otherwise.

### GetTaxNumberOk

`func (o *ReceiptTax) GetTaxNumberOk() (*string, bool)`

GetTaxNumberOk returns a tuple with the TaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNumber

`func (o *ReceiptTax) SetTaxNumber(v string)`

SetTaxNumber sets TaxNumber field to given value.

### HasTaxNumber

`func (o *ReceiptTax) HasTaxNumber() bool`

HasTaxNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


