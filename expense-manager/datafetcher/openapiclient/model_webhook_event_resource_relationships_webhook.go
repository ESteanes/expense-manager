/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the WebhookEventResourceRelationshipsWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEventResourceRelationshipsWebhook{}

// WebhookEventResourceRelationshipsWebhook struct for WebhookEventResourceRelationshipsWebhook
type WebhookEventResourceRelationshipsWebhook struct {
	Data  WebhookEventResourceRelationshipsWebhookData   `json:"data"`
	Links *AccountResourceRelationshipsTransactionsLinks `json:"links,omitempty"`
}

type _WebhookEventResourceRelationshipsWebhook WebhookEventResourceRelationshipsWebhook

// NewWebhookEventResourceRelationshipsWebhook instantiates a new WebhookEventResourceRelationshipsWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEventResourceRelationshipsWebhook(data WebhookEventResourceRelationshipsWebhookData) *WebhookEventResourceRelationshipsWebhook {
	this := WebhookEventResourceRelationshipsWebhook{}
	this.Data = data
	return &this
}

// NewWebhookEventResourceRelationshipsWebhookWithDefaults instantiates a new WebhookEventResourceRelationshipsWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEventResourceRelationshipsWebhookWithDefaults() *WebhookEventResourceRelationshipsWebhook {
	this := WebhookEventResourceRelationshipsWebhook{}
	return &this
}

// GetData returns the Data field value
func (o *WebhookEventResourceRelationshipsWebhook) GetData() WebhookEventResourceRelationshipsWebhookData {
	if o == nil {
		var ret WebhookEventResourceRelationshipsWebhookData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WebhookEventResourceRelationshipsWebhook) GetDataOk() (*WebhookEventResourceRelationshipsWebhookData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WebhookEventResourceRelationshipsWebhook) SetData(v WebhookEventResourceRelationshipsWebhookData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WebhookEventResourceRelationshipsWebhook) GetLinks() AccountResourceRelationshipsTransactionsLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountResourceRelationshipsTransactionsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookEventResourceRelationshipsWebhook) GetLinksOk() (*AccountResourceRelationshipsTransactionsLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WebhookEventResourceRelationshipsWebhook) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceRelationshipsTransactionsLinks and assigns it to the Links field.
func (o *WebhookEventResourceRelationshipsWebhook) SetLinks(v AccountResourceRelationshipsTransactionsLinks) {
	o.Links = &v
}

func (o WebhookEventResourceRelationshipsWebhook) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEventResourceRelationshipsWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *WebhookEventResourceRelationshipsWebhook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varWebhookEventResourceRelationshipsWebhook := _WebhookEventResourceRelationshipsWebhook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookEventResourceRelationshipsWebhook)

	if err != nil {
		return err
	}

	*o = WebhookEventResourceRelationshipsWebhook(varWebhookEventResourceRelationshipsWebhook)

	return err
}

type NullableWebhookEventResourceRelationshipsWebhook struct {
	value *WebhookEventResourceRelationshipsWebhook
	isSet bool
}

func (v NullableWebhookEventResourceRelationshipsWebhook) Get() *WebhookEventResourceRelationshipsWebhook {
	return v.value
}

func (v *NullableWebhookEventResourceRelationshipsWebhook) Set(val *WebhookEventResourceRelationshipsWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventResourceRelationshipsWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventResourceRelationshipsWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventResourceRelationshipsWebhook(val *WebhookEventResourceRelationshipsWebhook) *NullableWebhookEventResourceRelationshipsWebhook {
	return &NullableWebhookEventResourceRelationshipsWebhook{value: val, isSet: true}
}

func (v NullableWebhookEventResourceRelationshipsWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventResourceRelationshipsWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
