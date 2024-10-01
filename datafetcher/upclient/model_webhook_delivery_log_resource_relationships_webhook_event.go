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

// checks if the WebhookDeliveryLogResourceRelationshipsWebhookEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookDeliveryLogResourceRelationshipsWebhookEvent{}

// WebhookDeliveryLogResourceRelationshipsWebhookEvent struct for WebhookDeliveryLogResourceRelationshipsWebhookEvent
type WebhookDeliveryLogResourceRelationshipsWebhookEvent struct {
	Data WebhookDeliveryLogResourceRelationshipsWebhookEventData `json:"data"`
}

type _WebhookDeliveryLogResourceRelationshipsWebhookEvent WebhookDeliveryLogResourceRelationshipsWebhookEvent

// NewWebhookDeliveryLogResourceRelationshipsWebhookEvent instantiates a new WebhookDeliveryLogResourceRelationshipsWebhookEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookDeliveryLogResourceRelationshipsWebhookEvent(data WebhookDeliveryLogResourceRelationshipsWebhookEventData) *WebhookDeliveryLogResourceRelationshipsWebhookEvent {
	this := WebhookDeliveryLogResourceRelationshipsWebhookEvent{}
	this.Data = data
	return &this
}

// NewWebhookDeliveryLogResourceRelationshipsWebhookEventWithDefaults instantiates a new WebhookDeliveryLogResourceRelationshipsWebhookEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookDeliveryLogResourceRelationshipsWebhookEventWithDefaults() *WebhookDeliveryLogResourceRelationshipsWebhookEvent {
	this := WebhookDeliveryLogResourceRelationshipsWebhookEvent{}
	return &this
}

// GetData returns the Data field value
func (o *WebhookDeliveryLogResourceRelationshipsWebhookEvent) GetData() WebhookDeliveryLogResourceRelationshipsWebhookEventData {
	if o == nil {
		var ret WebhookDeliveryLogResourceRelationshipsWebhookEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WebhookDeliveryLogResourceRelationshipsWebhookEvent) GetDataOk() (*WebhookDeliveryLogResourceRelationshipsWebhookEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WebhookDeliveryLogResourceRelationshipsWebhookEvent) SetData(v WebhookDeliveryLogResourceRelationshipsWebhookEventData) {
	o.Data = v
}

func (o WebhookDeliveryLogResourceRelationshipsWebhookEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookDeliveryLogResourceRelationshipsWebhookEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *WebhookDeliveryLogResourceRelationshipsWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varWebhookDeliveryLogResourceRelationshipsWebhookEvent := _WebhookDeliveryLogResourceRelationshipsWebhookEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookDeliveryLogResourceRelationshipsWebhookEvent)

	if err != nil {
		return err
	}

	*o = WebhookDeliveryLogResourceRelationshipsWebhookEvent(varWebhookDeliveryLogResourceRelationshipsWebhookEvent)

	return err
}

type NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent struct {
	value *WebhookDeliveryLogResourceRelationshipsWebhookEvent
	isSet bool
}

func (v NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) Get() *WebhookDeliveryLogResourceRelationshipsWebhookEvent {
	return v.value
}

func (v *NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) Set(val *WebhookDeliveryLogResourceRelationshipsWebhookEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDeliveryLogResourceRelationshipsWebhookEvent(val *WebhookDeliveryLogResourceRelationshipsWebhookEvent) *NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent {
	return &NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent{value: val, isSet: true}
}

func (v NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDeliveryLogResourceRelationshipsWebhookEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


