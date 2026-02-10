# ReceiptItemSubItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Amount** | **int64** |  | 
**Currency** | **string** |  | 
**Quantity** | Pointer to **float64** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Tax** | Pointer to **int64** |  | [optional] 

## Methods

### NewReceiptItemSubItemsInner

`func NewReceiptItemSubItemsInner(description string, amount int64, currency string, ) *ReceiptItemSubItemsInner`

NewReceiptItemSubItemsInner instantiates a new ReceiptItemSubItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptItemSubItemsInnerWithDefaults

`func NewReceiptItemSubItemsInnerWithDefaults() *ReceiptItemSubItemsInner`

NewReceiptItemSubItemsInnerWithDefaults instantiates a new ReceiptItemSubItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ReceiptItemSubItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReceiptItemSubItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReceiptItemSubItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmount

`func (o *ReceiptItemSubItemsInner) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ReceiptItemSubItemsInner) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ReceiptItemSubItemsInner) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *ReceiptItemSubItemsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReceiptItemSubItemsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReceiptItemSubItemsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetQuantity

`func (o *ReceiptItemSubItemsInner) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReceiptItemSubItemsInner) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReceiptItemSubItemsInner) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReceiptItemSubItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnit

`func (o *ReceiptItemSubItemsInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ReceiptItemSubItemsInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ReceiptItemSubItemsInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ReceiptItemSubItemsInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetTax

`func (o *ReceiptItemSubItemsInner) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *ReceiptItemSubItemsInner) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *ReceiptItemSubItemsInner) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *ReceiptItemSubItemsInner) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


