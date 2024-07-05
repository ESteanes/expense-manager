/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
	"fmt"
)

// TransactionStatusEnum Specifies which stage of processing a transaction is currently at. Currently returned values are `HELD` and `SETTLED`. When a transaction is held, its account’s `availableBalance` is affected. When a transaction is settled, its account’s `currentBalance` is affected.
type TransactionStatusEnum string

// List of TransactionStatusEnum
const (
	HELD    TransactionStatusEnum = "HELD"
	SETTLED TransactionStatusEnum = "SETTLED"
)

// All allowed values of TransactionStatusEnum enum
var AllowedTransactionStatusEnumEnumValues = []TransactionStatusEnum{
	"HELD",
	"SETTLED",
}

func (v *TransactionStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionStatusEnum(value)
	for _, existing := range AllowedTransactionStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionStatusEnum", value)
}

// NewTransactionStatusEnumFromValue returns a pointer to a valid TransactionStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionStatusEnumFromValue(v string) (*TransactionStatusEnum, error) {
	ev := TransactionStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionStatusEnum: valid values are %v", v, AllowedTransactionStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionStatusEnum) IsValid() bool {
	for _, existing := range AllowedTransactionStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionStatusEnum value
func (v TransactionStatusEnum) Ptr() *TransactionStatusEnum {
	return &v
}

type NullableTransactionStatusEnum struct {
	value *TransactionStatusEnum
	isSet bool
}

func (v NullableTransactionStatusEnum) Get() *TransactionStatusEnum {
	return v.value
}

func (v *NullableTransactionStatusEnum) Set(val *TransactionStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStatusEnum(val *TransactionStatusEnum) *NullableTransactionStatusEnum {
	return &NullableTransactionStatusEnum{value: val, isSet: true}
}

func (v NullableTransactionStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
