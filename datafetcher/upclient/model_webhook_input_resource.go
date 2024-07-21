/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package upclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the WebhookInputResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookInputResource{}

// WebhookInputResource Represents a webhook specified as request input.
type WebhookInputResource struct {
	Attributes WebhookInputResourceAttributes `json:"attributes"`
}

type _WebhookInputResource WebhookInputResource

// NewWebhookInputResource instantiates a new WebhookInputResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookInputResource(attributes WebhookInputResourceAttributes) *WebhookInputResource {
	this := WebhookInputResource{}
	this.Attributes = attributes
	return &this
}

// NewWebhookInputResourceWithDefaults instantiates a new WebhookInputResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookInputResourceWithDefaults() *WebhookInputResource {
	this := WebhookInputResource{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *WebhookInputResource) GetAttributes() WebhookInputResourceAttributes {
	if o == nil {
		var ret WebhookInputResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WebhookInputResource) GetAttributesOk() (*WebhookInputResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *WebhookInputResource) SetAttributes(v WebhookInputResourceAttributes) {
	o.Attributes = v
}

func (o WebhookInputResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookInputResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *WebhookInputResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebhookInputResource := _WebhookInputResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookInputResource)

	if err != nil {
		return err
	}

	*o = WebhookInputResource(varWebhookInputResource)

	return err
}

type NullableWebhookInputResource struct {
	value *WebhookInputResource
	isSet bool
}

func (v NullableWebhookInputResource) Get() *WebhookInputResource {
	return v.value
}

func (v *NullableWebhookInputResource) Set(val *WebhookInputResource) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookInputResource) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookInputResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookInputResource(val *WebhookInputResource) *NullableWebhookInputResource {
	return &NullableWebhookInputResource{value: val, isSet: true}
}

func (v NullableWebhookInputResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookInputResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}