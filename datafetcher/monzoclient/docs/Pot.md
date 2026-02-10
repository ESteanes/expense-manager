# Pot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The pot ID | [optional] 
**Name** | Pointer to **string** | The pot name | [optional] 
**Style** | Pointer to **string** | The pot background style | [optional] 
**Balance** | Pointer to **int64** | The pot balance in minor units (pennies) | [optional] 
**Currency** | Pointer to **string** | The pot currency | [optional] 
**Created** | Pointer to **time.Time** | When the pot was created | [optional] 
**Updated** | Pointer to **time.Time** | When the pot was last updated | [optional] 
**Deleted** | Pointer to **bool** | Whether the pot is deleted | [optional] 

## Methods

### NewPot

`func NewPot() *Pot`

NewPot instantiates a new Pot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPotWithDefaults

`func NewPotWithDefaults() *Pot`

NewPotWithDefaults instantiates a new Pot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Pot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Pot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Pot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Pot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStyle

`func (o *Pot) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *Pot) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *Pot) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *Pot) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetBalance

`func (o *Pot) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Pot) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Pot) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *Pot) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *Pot) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Pot) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Pot) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Pot) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCreated

`func (o *Pot) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Pot) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Pot) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Pot) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Pot) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Pot) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Pot) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Pot) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetDeleted

`func (o *Pot) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Pot) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Pot) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Pot) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


