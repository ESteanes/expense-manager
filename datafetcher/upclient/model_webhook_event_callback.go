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

// checks if the WebhookEventCallback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEventCallback{}

// WebhookEventCallback Asynchronous callback request used for webhook event delivery. 
type WebhookEventCallback struct {
	// The webhook event data sent to the subscribed webhook. 
	Data WebhookEventResource `json:"data"`
}

type _WebhookEventCallback WebhookEventCallback

// NewWebhookEventCallback instantiates a new WebhookEventCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEventCallback(data WebhookEventResource) *WebhookEventCallback {
	this := WebhookEventCallback{}
	this.Data = data
	return &this
}

// NewWebhookEventCallbackWithDefaults instantiates a new WebhookEventCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEventCallbackWithDefaults() *WebhookEventCallback {
	this := WebhookEventCallback{}
	return &this
}

// GetData returns the Data field value
func (o *WebhookEventCallback) GetData() WebhookEventResource {
	if o == nil {
		var ret WebhookEventResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WebhookEventCallback) GetDataOk() (*WebhookEventResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WebhookEventCallback) SetData(v WebhookEventResource) {
	o.Data = v
}

func (o WebhookEventCallback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEventCallback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *WebhookEventCallback) UnmarshalJSON(data []byte) (err error) {
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

	varWebhookEventCallback := _WebhookEventCallback{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookEventCallback)

	if err != nil {
		return err
	}

	*o = WebhookEventCallback(varWebhookEventCallback)

	return err
}

type NullableWebhookEventCallback struct {
	value *WebhookEventCallback
	isSet bool
}

func (v NullableWebhookEventCallback) Get() *WebhookEventCallback {
	return v.value
}

func (v *NullableWebhookEventCallback) Set(val *WebhookEventCallback) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventCallback(val *WebhookEventCallback) *NullableWebhookEventCallback {
	return &NullableWebhookEventCallback{value: val, isSet: true}
}

func (v NullableWebhookEventCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


