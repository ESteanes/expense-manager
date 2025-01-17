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

// checks if the WebhookEventResourceRelationshipsTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEventResourceRelationshipsTransaction{}

// WebhookEventResourceRelationshipsTransaction struct for WebhookEventResourceRelationshipsTransaction
type WebhookEventResourceRelationshipsTransaction struct {
	Data  WebhookEventResourceRelationshipsTransactionData `json:"data"`
	Links *AccountResourceRelationshipsTransactionsLinks   `json:"links,omitempty"`
}

type _WebhookEventResourceRelationshipsTransaction WebhookEventResourceRelationshipsTransaction

// NewWebhookEventResourceRelationshipsTransaction instantiates a new WebhookEventResourceRelationshipsTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEventResourceRelationshipsTransaction(data WebhookEventResourceRelationshipsTransactionData) *WebhookEventResourceRelationshipsTransaction {
	this := WebhookEventResourceRelationshipsTransaction{}
	this.Data = data
	return &this
}

// NewWebhookEventResourceRelationshipsTransactionWithDefaults instantiates a new WebhookEventResourceRelationshipsTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEventResourceRelationshipsTransactionWithDefaults() *WebhookEventResourceRelationshipsTransaction {
	this := WebhookEventResourceRelationshipsTransaction{}
	return &this
}

// GetData returns the Data field value
func (o *WebhookEventResourceRelationshipsTransaction) GetData() WebhookEventResourceRelationshipsTransactionData {
	if o == nil {
		var ret WebhookEventResourceRelationshipsTransactionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WebhookEventResourceRelationshipsTransaction) GetDataOk() (*WebhookEventResourceRelationshipsTransactionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WebhookEventResourceRelationshipsTransaction) SetData(v WebhookEventResourceRelationshipsTransactionData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WebhookEventResourceRelationshipsTransaction) GetLinks() AccountResourceRelationshipsTransactionsLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountResourceRelationshipsTransactionsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookEventResourceRelationshipsTransaction) GetLinksOk() (*AccountResourceRelationshipsTransactionsLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WebhookEventResourceRelationshipsTransaction) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceRelationshipsTransactionsLinks and assigns it to the Links field.
func (o *WebhookEventResourceRelationshipsTransaction) SetLinks(v AccountResourceRelationshipsTransactionsLinks) {
	o.Links = &v
}

func (o WebhookEventResourceRelationshipsTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEventResourceRelationshipsTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *WebhookEventResourceRelationshipsTransaction) UnmarshalJSON(data []byte) (err error) {
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

	varWebhookEventResourceRelationshipsTransaction := _WebhookEventResourceRelationshipsTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookEventResourceRelationshipsTransaction)

	if err != nil {
		return err
	}

	*o = WebhookEventResourceRelationshipsTransaction(varWebhookEventResourceRelationshipsTransaction)

	return err
}

type NullableWebhookEventResourceRelationshipsTransaction struct {
	value *WebhookEventResourceRelationshipsTransaction
	isSet bool
}

func (v NullableWebhookEventResourceRelationshipsTransaction) Get() *WebhookEventResourceRelationshipsTransaction {
	return v.value
}

func (v *NullableWebhookEventResourceRelationshipsTransaction) Set(val *WebhookEventResourceRelationshipsTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventResourceRelationshipsTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventResourceRelationshipsTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventResourceRelationshipsTransaction(val *WebhookEventResourceRelationshipsTransaction) *NullableWebhookEventResourceRelationshipsTransaction {
	return &NullableWebhookEventResourceRelationshipsTransaction{value: val, isSet: true}
}

func (v NullableWebhookEventResourceRelationshipsTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventResourceRelationshipsTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
