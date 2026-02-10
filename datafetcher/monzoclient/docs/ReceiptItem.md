# ReceiptItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The product description | 
**Amount** | **int64** | The amount paid for the item in minor units (pennies) | 
**Currency** | **string** | The currency code | 
**Quantity** | Pointer to **float64** | How many of the product were bought | [optional] 
**Unit** | Pointer to **string** | The unit the quantity is measured in | [optional] 
**Tax** | Pointer to **int64** | The tax in minor units (pennies) | [optional] 
**SubItems** | Pointer to [**[]ReceiptItemSubItemsInner**](ReceiptItemSubItemsInner.md) |  | [optional] 

## Methods

### NewReceiptItem

`func NewReceiptItem(description string, amount int64, currency string, ) *ReceiptItem`

NewReceiptItem instantiates a new ReceiptItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptItemWithDefaults

`func NewReceiptItemWithDefaults() *ReceiptItem`

NewReceiptItemWithDefaults instantiates a new ReceiptItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ReceiptItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReceiptItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReceiptItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmount

`func (o *ReceiptItem) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ReceiptItem) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ReceiptItem) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *ReceiptItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReceiptItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReceiptItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetQuantity

`func (o *ReceiptItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReceiptItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReceiptItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReceiptItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnit

`func (o *ReceiptItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ReceiptItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ReceiptItem) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ReceiptItem) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetTax

`func (o *ReceiptItem) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *ReceiptItem) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *ReceiptItem) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *ReceiptItem) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetSubItems

`func (o *ReceiptItem) GetSubItems() []ReceiptItemSubItemsInner`

GetSubItems returns the SubItems field if non-nil, zero value otherwise.

### GetSubItemsOk

`func (o *ReceiptItem) GetSubItemsOk() (*[]ReceiptItemSubItemsInner, bool)`

GetSubItemsOk returns a tuple with the SubItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubItems

`func (o *ReceiptItem) SetSubItems(v []ReceiptItemSubItemsInner)`

SetSubItems sets SubItems field to given value.

### HasSubItems

`func (o *ReceiptItem) HasSubItems() bool`

HasSubItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


