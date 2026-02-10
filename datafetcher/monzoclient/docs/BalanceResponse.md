# BalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **int64** | The currently available balance in minor units (pennies) | [optional] 
**TotalBalance** | Pointer to **int64** | The sum of available balance and all pots | [optional] 
**Currency** | Pointer to **string** | The ISO 4217 currency code | [optional] 
**SpendToday** | Pointer to **int64** | The amount spent today (from ~4am) as a negative value | [optional] 

## Methods

### NewBalanceResponse

`func NewBalanceResponse() *BalanceResponse`

NewBalanceResponse instantiates a new BalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceResponseWithDefaults

`func NewBalanceResponseWithDefaults() *BalanceResponse`

NewBalanceResponseWithDefaults instantiates a new BalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *BalanceResponse) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalanceResponse) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalanceResponse) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BalanceResponse) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetTotalBalance

`func (o *BalanceResponse) GetTotalBalance() int64`

GetTotalBalance returns the TotalBalance field if non-nil, zero value otherwise.

### GetTotalBalanceOk

`func (o *BalanceResponse) GetTotalBalanceOk() (*int64, bool)`

GetTotalBalanceOk returns a tuple with the TotalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalance

`func (o *BalanceResponse) SetTotalBalance(v int64)`

SetTotalBalance sets TotalBalance field to given value.

### HasTotalBalance

`func (o *BalanceResponse) HasTotalBalance() bool`

HasTotalBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *BalanceResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BalanceResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BalanceResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BalanceResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSpendToday

`func (o *BalanceResponse) GetSpendToday() int64`

GetSpendToday returns the SpendToday field if non-nil, zero value otherwise.

### GetSpendTodayOk

`func (o *BalanceResponse) GetSpendTodayOk() (*int64, bool)`

GetSpendTodayOk returns a tuple with the SpendToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendToday

`func (o *BalanceResponse) SetSpendToday(v int64)`

SetSpendToday sets SpendToday field to given value.

### HasSpendToday

`func (o *BalanceResponse) HasSpendToday() bool`

HasSpendToday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


