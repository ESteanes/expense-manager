# TransactionResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**TransactionStatusEnum**](TransactionStatusEnum.md) | The current processing status of this transaction, according to whether or not this transaction has settled or is still held.  | 
**RawText** | **NullableString** | The original, unprocessed text of the transaction. This is often not a perfect indicator of the actual merchant, but it is useful for reconciliation purposes in some cases.  | 
**Description** | **string** | A short description for this transaction. Usually the merchant name for purchases.  | 
**Message** | **NullableString** | Attached message for this transaction, such as a payment message, or a transfer note.  | 
**IsCategorizable** | **bool** | Boolean flag set to true on transactions that support the use of categories.  | 
**HoldInfo** | [**NullableHoldInfoObject**](HoldInfoObject.md) | If this transaction is currently in the &#x60;HELD&#x60; status, or was ever in the &#x60;HELD&#x60; status, the &#x60;amount&#x60; and &#x60;foreignAmount&#x60; of the transaction while &#x60;HELD&#x60;.  | 
**RoundUp** | [**NullableRoundUpObject**](RoundUpObject.md) | Details of how this transaction was rounded-up. If no Round Up was applied this field will be &#x60;null&#x60;.  | 
**Cashback** | [**NullableCashbackObject**](CashbackObject.md) | If all or part of this transaction was instantly reimbursed in the form of cashback, details of the reimbursement.  | 
**Amount** | [**MoneyObject**](MoneyObject.md) | The amount of this transaction in Australian dollars. For transactions that were once &#x60;HELD&#x60; but are now &#x60;SETTLED&#x60;, refer to the &#x60;holdInfo&#x60; field for the original &#x60;amount&#x60; the transaction was &#x60;HELD&#x60; at.  | 
**ForeignAmount** | [**NullableMoneyObject**](MoneyObject.md) | The foreign currency amount of this transaction. This field will be &#x60;null&#x60; for domestic transactions. The amount was converted to the AUD amount reflected in the &#x60;amount&#x60; of this transaction. Refer to the &#x60;holdInfo&#x60; field for the original &#x60;foreignAmount&#x60; the transaction was &#x60;HELD&#x60; at.  | 
**SettledAt** | **NullableTime** | The date-time at which this transaction settled. This field will be &#x60;null&#x60; for transactions that are currently in the &#x60;HELD&#x60; status.  | 
**CreatedAt** | **time.Time** | The date-time at which this transaction was first encountered.  | 

## Methods

### NewTransactionResourceAttributes

`func NewTransactionResourceAttributes(status TransactionStatusEnum, rawText NullableString, description string, message NullableString, isCategorizable bool, holdInfo NullableHoldInfoObject, roundUp NullableRoundUpObject, cashback NullableCashbackObject, amount MoneyObject, foreignAmount NullableMoneyObject, settledAt NullableTime, createdAt time.Time, ) *TransactionResourceAttributes`

NewTransactionResourceAttributes instantiates a new TransactionResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResourceAttributesWithDefaults

`func NewTransactionResourceAttributesWithDefaults() *TransactionResourceAttributes`

NewTransactionResourceAttributesWithDefaults instantiates a new TransactionResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TransactionResourceAttributes) GetStatus() TransactionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionResourceAttributes) GetStatusOk() (*TransactionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionResourceAttributes) SetStatus(v TransactionStatusEnum)`

SetStatus sets Status field to given value.


### GetRawText

`func (o *TransactionResourceAttributes) GetRawText() string`

GetRawText returns the RawText field if non-nil, zero value otherwise.

### GetRawTextOk

`func (o *TransactionResourceAttributes) GetRawTextOk() (*string, bool)`

GetRawTextOk returns a tuple with the RawText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawText

`func (o *TransactionResourceAttributes) SetRawText(v string)`

SetRawText sets RawText field to given value.


### SetRawTextNil

`func (o *TransactionResourceAttributes) SetRawTextNil(b bool)`

 SetRawTextNil sets the value for RawText to be an explicit nil

### UnsetRawText
`func (o *TransactionResourceAttributes) UnsetRawText()`

UnsetRawText ensures that no value is present for RawText, not even an explicit nil
### GetDescription

`func (o *TransactionResourceAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionResourceAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionResourceAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMessage

`func (o *TransactionResourceAttributes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionResourceAttributes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionResourceAttributes) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *TransactionResourceAttributes) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TransactionResourceAttributes) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetIsCategorizable

`func (o *TransactionResourceAttributes) GetIsCategorizable() bool`

GetIsCategorizable returns the IsCategorizable field if non-nil, zero value otherwise.

### GetIsCategorizableOk

`func (o *TransactionResourceAttributes) GetIsCategorizableOk() (*bool, bool)`

GetIsCategorizableOk returns a tuple with the IsCategorizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCategorizable

`func (o *TransactionResourceAttributes) SetIsCategorizable(v bool)`

SetIsCategorizable sets IsCategorizable field to given value.


### GetHoldInfo

`func (o *TransactionResourceAttributes) GetHoldInfo() HoldInfoObject`

GetHoldInfo returns the HoldInfo field if non-nil, zero value otherwise.

### GetHoldInfoOk

`func (o *TransactionResourceAttributes) GetHoldInfoOk() (*HoldInfoObject, bool)`

GetHoldInfoOk returns a tuple with the HoldInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldInfo

`func (o *TransactionResourceAttributes) SetHoldInfo(v HoldInfoObject)`

SetHoldInfo sets HoldInfo field to given value.


### SetHoldInfoNil

`func (o *TransactionResourceAttributes) SetHoldInfoNil(b bool)`

 SetHoldInfoNil sets the value for HoldInfo to be an explicit nil

### UnsetHoldInfo
`func (o *TransactionResourceAttributes) UnsetHoldInfo()`

UnsetHoldInfo ensures that no value is present for HoldInfo, not even an explicit nil
### GetRoundUp

`func (o *TransactionResourceAttributes) GetRoundUp() RoundUpObject`

GetRoundUp returns the RoundUp field if non-nil, zero value otherwise.

### GetRoundUpOk

`func (o *TransactionResourceAttributes) GetRoundUpOk() (*RoundUpObject, bool)`

GetRoundUpOk returns a tuple with the RoundUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundUp

`func (o *TransactionResourceAttributes) SetRoundUp(v RoundUpObject)`

SetRoundUp sets RoundUp field to given value.


### SetRoundUpNil

`func (o *TransactionResourceAttributes) SetRoundUpNil(b bool)`

 SetRoundUpNil sets the value for RoundUp to be an explicit nil

### UnsetRoundUp
`func (o *TransactionResourceAttributes) UnsetRoundUp()`

UnsetRoundUp ensures that no value is present for RoundUp, not even an explicit nil
### GetCashback

`func (o *TransactionResourceAttributes) GetCashback() CashbackObject`

GetCashback returns the Cashback field if non-nil, zero value otherwise.

### GetCashbackOk

`func (o *TransactionResourceAttributes) GetCashbackOk() (*CashbackObject, bool)`

GetCashbackOk returns a tuple with the Cashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashback

`func (o *TransactionResourceAttributes) SetCashback(v CashbackObject)`

SetCashback sets Cashback field to given value.


### SetCashbackNil

`func (o *TransactionResourceAttributes) SetCashbackNil(b bool)`

 SetCashbackNil sets the value for Cashback to be an explicit nil

### UnsetCashback
`func (o *TransactionResourceAttributes) UnsetCashback()`

UnsetCashback ensures that no value is present for Cashback, not even an explicit nil
### GetAmount

`func (o *TransactionResourceAttributes) GetAmount() MoneyObject`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionResourceAttributes) GetAmountOk() (*MoneyObject, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionResourceAttributes) SetAmount(v MoneyObject)`

SetAmount sets Amount field to given value.


### GetForeignAmount

`func (o *TransactionResourceAttributes) GetForeignAmount() MoneyObject`

GetForeignAmount returns the ForeignAmount field if non-nil, zero value otherwise.

### GetForeignAmountOk

`func (o *TransactionResourceAttributes) GetForeignAmountOk() (*MoneyObject, bool)`

GetForeignAmountOk returns a tuple with the ForeignAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmount

`func (o *TransactionResourceAttributes) SetForeignAmount(v MoneyObject)`

SetForeignAmount sets ForeignAmount field to given value.


### SetForeignAmountNil

`func (o *TransactionResourceAttributes) SetForeignAmountNil(b bool)`

 SetForeignAmountNil sets the value for ForeignAmount to be an explicit nil

### UnsetForeignAmount
`func (o *TransactionResourceAttributes) UnsetForeignAmount()`

UnsetForeignAmount ensures that no value is present for ForeignAmount, not even an explicit nil
### GetSettledAt

`func (o *TransactionResourceAttributes) GetSettledAt() time.Time`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *TransactionResourceAttributes) GetSettledAtOk() (*time.Time, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *TransactionResourceAttributes) SetSettledAt(v time.Time)`

SetSettledAt sets SettledAt field to given value.


### SetSettledAtNil

`func (o *TransactionResourceAttributes) SetSettledAtNil(b bool)`

 SetSettledAtNil sets the value for SettledAt to be an explicit nil

### UnsetSettledAt
`func (o *TransactionResourceAttributes) UnsetSettledAt()`

UnsetSettledAt ensures that no value is present for SettledAt, not even an explicit nil
### GetCreatedAt

`func (o *TransactionResourceAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionResourceAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionResourceAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


