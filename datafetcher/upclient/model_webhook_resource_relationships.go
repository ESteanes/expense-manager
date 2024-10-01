/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package upclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WebhookResourceRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookResourceRelationships{}

// WebhookResourceRelationships struct for WebhookResourceRelationships
type WebhookResourceRelationships struct {
	Logs AccountResourceRelationshipsTransactions `json:"logs"`
}

type _WebhookResourceRelationships WebhookResourceRelationships

// NewWebhookResourceRelationships instantiates a new WebhookResourceRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResourceRelationships(logs AccountResourceRelationshipsTransactions) *WebhookResourceRelationships {
	this := WebhookResourceRelationships{}
	this.Logs = logs
	return &this
}

// NewWebhookResourceRelationshipsWithDefaults instantiates a new WebhookResourceRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResourceRelationshipsWithDefaults() *WebhookResourceRelationships {
	this := WebhookResourceRelationships{}
	return &this
}

// GetLogs returns the Logs field value
func (o *WebhookResourceRelationships) GetLogs() AccountResourceRelationshipsTransactions {
	if o == nil {
		var ret AccountResourceRelationshipsTransactions
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *WebhookResourceRelationships) GetLogsOk() (*AccountResourceRelationshipsTransactions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logs, true
}

// SetLogs sets field value
func (o *WebhookResourceRelationships) SetLogs(v AccountResourceRelationshipsTransactions) {
	o.Logs = v
}

func (o WebhookResourceRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookResourceRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logs"] = o.Logs
	return toSerialize, nil
}

func (o *WebhookResourceRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logs",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebhookResourceRelationships := _WebhookResourceRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookResourceRelationships)

	if err != nil {
		return err
	}

	*o = WebhookResourceRelationships(varWebhookResourceRelationships)

	return err
}

type NullableWebhookResourceRelationships struct {
	value *WebhookResourceRelationships
	isSet bool
}

func (v NullableWebhookResourceRelationships) Get() *WebhookResourceRelationships {
	return v.value
}

func (v *NullableWebhookResourceRelationships) Set(val *WebhookResourceRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResourceRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResourceRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResourceRelationships(val *WebhookResourceRelationships) *NullableWebhookResourceRelationships {
	return &NullableWebhookResourceRelationships{value: val, isSet: true}
}

func (v NullableWebhookResourceRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResourceRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


