# GetReceipt200ResponseReceipt

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
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewGetReceipt200ResponseReceipt

`func NewGetReceipt200ResponseReceipt(transactionId string, externalId string, total int64, currency string, items []ReceiptItem, ) *GetReceipt200ResponseReceipt`

NewGetReceipt200ResponseReceipt instantiates a new GetReceipt200ResponseReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReceipt200ResponseReceiptWithDefaults

`func NewGetReceipt200ResponseReceiptWithDefaults() *GetReceipt200ResponseReceipt`

NewGetReceipt200ResponseReceiptWithDefaults instantiates a new GetReceipt200ResponseReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *GetReceipt200ResponseReceipt) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetReceipt200ResponseReceipt) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetReceipt200ResponseReceipt) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetExternalId

`func (o *GetReceipt200ResponseReceipt) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetReceipt200ResponseReceipt) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetReceipt200ResponseReceipt) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetTotal

`func (o *GetReceipt200ResponseReceipt) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetReceipt200ResponseReceipt) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetReceipt200ResponseReceipt) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetCurrency

`func (o *GetReceipt200ResponseReceipt) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetReceipt200ResponseReceipt) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetReceipt200ResponseReceipt) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetItems

`func (o *GetReceipt200ResponseReceipt) GetItems() []ReceiptItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetReceipt200ResponseReceipt) GetItemsOk() (*[]ReceiptItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetReceipt200ResponseReceipt) SetItems(v []ReceiptItem)`

SetItems sets Items field to given value.


### GetTaxes

`func (o *GetReceipt200ResponseReceipt) GetTaxes() []ReceiptTax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *GetReceipt200ResponseReceipt) GetTaxesOk() (*[]ReceiptTax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *GetReceipt200ResponseReceipt) SetTaxes(v []ReceiptTax)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *GetReceipt200ResponseReceipt) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetPayments

`func (o *GetReceipt200ResponseReceipt) GetPayments() []ReceiptPayment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *GetReceipt200ResponseReceipt) GetPaymentsOk() (*[]ReceiptPayment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *GetReceipt200ResponseReceipt) SetPayments(v []ReceiptPayment)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *GetReceipt200ResponseReceipt) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetMerchant

`func (o *GetReceipt200ResponseReceipt) GetMerchant() ReceiptMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *GetReceipt200ResponseReceipt) GetMerchantOk() (*ReceiptMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *GetReceipt200ResponseReceipt) SetMerchant(v ReceiptMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *GetReceipt200ResponseReceipt) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetId

`func (o *GetReceipt200ResponseReceipt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReceipt200ResponseReceipt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReceipt200ResponseReceipt) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetReceipt200ResponseReceipt) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


