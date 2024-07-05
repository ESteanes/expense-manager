/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// WebhookDeliveryStatusEnum Specifies the nature of the success or failure of a webhook delivery attempt to the subscribed webhook URL. The currently returned values are described below:  - **`DELIVERED`**: The event was delivered to the webhook URL   successfully and a `200` response was received. - **`UNDELIVERABLE`**: The webhook URL was not reachable, or timed out. - **`BAD_RESPONSE_CODE`**: The event was delivered to the webhook URL   but a non-`200` response was received.
type WebhookDeliveryStatusEnum string

// List of WebhookDeliveryStatusEnum
const (
	DELIVERED         WebhookDeliveryStatusEnum = "DELIVERED"
	UNDELIVERABLE     WebhookDeliveryStatusEnum = "UNDELIVERABLE"
	BAD_RESPONSE_CODE WebhookDeliveryStatusEnum = "BAD_RESPONSE_CODE"
)

// All allowed values of WebhookDeliveryStatusEnum enum
var AllowedWebhookDeliveryStatusEnumEnumValues = []WebhookDeliveryStatusEnum{
	"DELIVERED",
	"UNDELIVERABLE",
	"BAD_RESPONSE_CODE",
}

func (v *WebhookDeliveryStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookDeliveryStatusEnum(value)
	for _, existing := range AllowedWebhookDeliveryStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookDeliveryStatusEnum", value)
}

// NewWebhookDeliveryStatusEnumFromValue returns a pointer to a valid WebhookDeliveryStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookDeliveryStatusEnumFromValue(v string) (*WebhookDeliveryStatusEnum, error) {
	ev := WebhookDeliveryStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookDeliveryStatusEnum: valid values are %v", v, AllowedWebhookDeliveryStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookDeliveryStatusEnum) IsValid() bool {
	for _, existing := range AllowedWebhookDeliveryStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookDeliveryStatusEnum value
func (v WebhookDeliveryStatusEnum) Ptr() *WebhookDeliveryStatusEnum {
	return &v
}

type NullableWebhookDeliveryStatusEnum struct {
	value *WebhookDeliveryStatusEnum
	isSet bool
}

func (v NullableWebhookDeliveryStatusEnum) Get() *WebhookDeliveryStatusEnum {
	return v.value
}

func (v *NullableWebhookDeliveryStatusEnum) Set(val *WebhookDeliveryStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDeliveryStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDeliveryStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDeliveryStatusEnum(val *WebhookDeliveryStatusEnum) *NullableWebhookDeliveryStatusEnum {
	return &NullableWebhookDeliveryStatusEnum{value: val, isSet: true}
}

func (v NullableWebhookDeliveryStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDeliveryStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
