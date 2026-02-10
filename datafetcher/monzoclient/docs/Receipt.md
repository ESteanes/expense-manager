# Receipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The ID of the Transaction to associate the Receipt with | 
**ExternalId** | **string** | A unique identifier used as an idempotency key | 
**Total** | **int64** | The amount in minor units (pennies). Should be positive. | 
**Currency** | **string** | The currency code | 
**Items** | [**[]ReceiptItem**](ReceiptItem.md) |  | 
**Taxes** | Pointer to [**[]ReceiptTax**](ReceiptTax.md) |  | [optional] 
**Payments** | Pointer to [**[]ReceiptPayment**](ReceiptPayment.md) |  | [optional] 
**Merchant** | Pointer to [**ReceiptMerchant**](ReceiptMerchant.md) |  | [optional] 

## Methods

### NewReceipt

`func NewReceipt(transactionId string, externalId string, total int64, currency string, items []ReceiptItem, ) *Receipt`

NewReceipt instantiates a new Receipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptWithDefaults

`func NewReceiptWithDefaults() *Receipt`

NewReceiptWithDefaults instantiates a new Receipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *Receipt) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Receipt) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Receipt) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetExternalId

`func (o *Receipt) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Receipt) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Receipt) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetTotal

`func (o *Receipt) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Receipt) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Receipt) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetCurrency

`func (o *Receipt) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Receipt) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Receipt) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetItems

`func (o *Receipt) GetItems() []ReceiptItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Receipt) GetItemsOk() (*[]ReceiptItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Receipt) SetItems(v []ReceiptItem)`

SetItems sets Items field to given value.


### GetTaxes

`func (o *Receipt) GetTaxes() []ReceiptTax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *Receipt) GetTaxesOk() (*[]ReceiptTax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *Receipt) SetTaxes(v []ReceiptTax)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *Receipt) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetPayments

`func (o *Receipt) GetPayments() []ReceiptPayment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *Receipt) GetPaymentsOk() (*[]ReceiptPayment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *Receipt) SetPayments(v []ReceiptPayment)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *Receipt) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetMerchant

`func (o *Receipt) GetMerchant() ReceiptMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *Receipt) GetMerchantOk() (*ReceiptMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *Receipt) SetMerchant(v ReceiptMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *Receipt) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


