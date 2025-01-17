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

// checks if the TransactionResourceRelationshipsTransferAccountLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionResourceRelationshipsTransferAccountLinks{}

// TransactionResourceRelationshipsTransferAccountLinks struct for TransactionResourceRelationshipsTransferAccountLinks
type TransactionResourceRelationshipsTransferAccountLinks struct {
	// The link to retrieve the related resource(s) in this relationship.
	Related string `json:"related"`
}

type _TransactionResourceRelationshipsTransferAccountLinks TransactionResourceRelationshipsTransferAccountLinks

// NewTransactionResourceRelationshipsTransferAccountLinks instantiates a new TransactionResourceRelationshipsTransferAccountLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionResourceRelationshipsTransferAccountLinks(related string) *TransactionResourceRelationshipsTransferAccountLinks {
	this := TransactionResourceRelationshipsTransferAccountLinks{}
	this.Related = related
	return &this
}

// NewTransactionResourceRelationshipsTransferAccountLinksWithDefaults instantiates a new TransactionResourceRelationshipsTransferAccountLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionResourceRelationshipsTransferAccountLinksWithDefaults() *TransactionResourceRelationshipsTransferAccountLinks {
	this := TransactionResourceRelationshipsTransferAccountLinks{}
	return &this
}

// GetRelated returns the Related field value
func (o *TransactionResourceRelationshipsTransferAccountLinks) GetRelated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *TransactionResourceRelationshipsTransferAccountLinks) GetRelatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *TransactionResourceRelationshipsTransferAccountLinks) SetRelated(v string) {
	o.Related = v
}

func (o TransactionResourceRelationshipsTransferAccountLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionResourceRelationshipsTransferAccountLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["related"] = o.Related
	return toSerialize, nil
}

func (o *TransactionResourceRelationshipsTransferAccountLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"related",
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

	varTransactionResourceRelationshipsTransferAccountLinks := _TransactionResourceRelationshipsTransferAccountLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionResourceRelationshipsTransferAccountLinks)

	if err != nil {
		return err
	}

	*o = TransactionResourceRelationshipsTransferAccountLinks(varTransactionResourceRelationshipsTransferAccountLinks)

	return err
}

type NullableTransactionResourceRelationshipsTransferAccountLinks struct {
	value *TransactionResourceRelationshipsTransferAccountLinks
	isSet bool
}

func (v NullableTransactionResourceRelationshipsTransferAccountLinks) Get() *TransactionResourceRelationshipsTransferAccountLinks {
	return v.value
}

func (v *NullableTransactionResourceRelationshipsTransferAccountLinks) Set(val *TransactionResourceRelationshipsTransferAccountLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionResourceRelationshipsTransferAccountLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionResourceRelationshipsTransferAccountLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionResourceRelationshipsTransferAccountLinks(val *TransactionResourceRelationshipsTransferAccountLinks) *NullableTransactionResourceRelationshipsTransferAccountLinks {
	return &NullableTransactionResourceRelationshipsTransferAccountLinks{value: val, isSet: true}
}

func (v NullableTransactionResourceRelationshipsTransferAccountLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionResourceRelationshipsTransferAccountLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
