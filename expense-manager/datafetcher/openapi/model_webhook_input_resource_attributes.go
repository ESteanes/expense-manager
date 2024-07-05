/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the WebhookInputResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookInputResourceAttributes{}

// WebhookInputResourceAttributes struct for WebhookInputResourceAttributes
type WebhookInputResourceAttributes struct {
	// The URL that this webhook should post events to. This must be a valid HTTP or HTTPS URL that does not exceed 300 characters in length.
	Url string `json:"url"`
	// An optional description for this webhook, up to 64 characters in length.
	Description NullableString `json:"description,omitempty"`
}

type _WebhookInputResourceAttributes WebhookInputResourceAttributes

// NewWebhookInputResourceAttributes instantiates a new WebhookInputResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookInputResourceAttributes(url string) *WebhookInputResourceAttributes {
	this := WebhookInputResourceAttributes{}
	this.Url = url
	return &this
}

// NewWebhookInputResourceAttributesWithDefaults instantiates a new WebhookInputResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookInputResourceAttributesWithDefaults() *WebhookInputResourceAttributes {
	this := WebhookInputResourceAttributes{}
	return &this
}

// GetUrl returns the Url field value
func (o *WebhookInputResourceAttributes) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookInputResourceAttributes) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookInputResourceAttributes) SetUrl(v string) {
	o.Url = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookInputResourceAttributes) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookInputResourceAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *WebhookInputResourceAttributes) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *WebhookInputResourceAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *WebhookInputResourceAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *WebhookInputResourceAttributes) UnsetDescription() {
	o.Description.Unset()
}

func (o WebhookInputResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookInputResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

func (o *WebhookInputResourceAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varWebhookInputResourceAttributes := _WebhookInputResourceAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookInputResourceAttributes)

	if err != nil {
		return err
	}

	*o = WebhookInputResourceAttributes(varWebhookInputResourceAttributes)

	return err
}

type NullableWebhookInputResourceAttributes struct {
	value *WebhookInputResourceAttributes
	isSet bool
}

func (v NullableWebhookInputResourceAttributes) Get() *WebhookInputResourceAttributes {
	return v.value
}

func (v *NullableWebhookInputResourceAttributes) Set(val *WebhookInputResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookInputResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookInputResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookInputResourceAttributes(val *WebhookInputResourceAttributes) *NullableWebhookInputResourceAttributes {
	return &NullableWebhookInputResourceAttributes{value: val, isSet: true}
}

func (v NullableWebhookInputResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookInputResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
