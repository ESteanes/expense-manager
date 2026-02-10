# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The transaction ID | [optional] 
**Created** | Pointer to **time.Time** | When the transaction was created | [optional] 
**Description** | Pointer to **string** | The transaction description | [optional] 
**Amount** | Pointer to **int64** | The amount in minor units (pennies). Negative for debits. | [optional] 
**Currency** | Pointer to **string** | The ISO 4217 currency code | [optional] 
**Merchant** | Pointer to [**TransactionMerchant**](TransactionMerchant.md) |  | [optional] 
**Notes** | Pointer to **string** | User notes | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata key-value pairs | [optional] 
**Category** | Pointer to **string** | Transaction category | [optional] 
**IsLoad** | Pointer to **bool** | Whether this is a top-up transaction | [optional] 
**Settled** | Pointer to **string** | When the transaction settled (empty string if not settled) | [optional] 
**DeclineReason** | Pointer to **string** | Only present on declined transactions | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Transaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *Transaction) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Transaction) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Transaction) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Transaction) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *Transaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *Transaction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transaction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transaction) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Transaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *Transaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Transaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Transaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Transaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMerchant

`func (o *Transaction) GetMerchant() TransactionMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *Transaction) GetMerchantOk() (*TransactionMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *Transaction) SetMerchant(v TransactionMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *Transaction) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetNotes

`func (o *Transaction) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Transaction) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Transaction) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Transaction) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetMetadata

`func (o *Transaction) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Transaction) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Transaction) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Transaction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCategory

`func (o *Transaction) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Transaction) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Transaction) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Transaction) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetIsLoad

`func (o *Transaction) GetIsLoad() bool`

GetIsLoad returns the IsLoad field if non-nil, zero value otherwise.

### GetIsLoadOk

`func (o *Transaction) GetIsLoadOk() (*bool, bool)`

GetIsLoadOk returns a tuple with the IsLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoad

`func (o *Transaction) SetIsLoad(v bool)`

SetIsLoad sets IsLoad field to given value.

### HasIsLoad

`func (o *Transaction) HasIsLoad() bool`

HasIsLoad returns a boolean if a field has been set.

### GetSettled

`func (o *Transaction) GetSettled() string`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *Transaction) GetSettledOk() (*string, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *Transaction) SetSettled(v string)`

SetSettled sets Settled field to given value.

### HasSettled

`func (o *Transaction) HasSettled() bool`

HasSettled returns a boolean if a field has been set.

### GetDeclineReason

`func (o *Transaction) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *Transaction) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *Transaction) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *Transaction) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


