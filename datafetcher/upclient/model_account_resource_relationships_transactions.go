/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package upclient

import (
	"encoding/json"
)

// checks if the AccountResourceRelationshipsTransactions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountResourceRelationshipsTransactions{}

// AccountResourceRelationshipsTransactions struct for AccountResourceRelationshipsTransactions
type AccountResourceRelationshipsTransactions struct {
	Links *AccountResourceRelationshipsTransactionsLinks `json:"links,omitempty"`
}

// NewAccountResourceRelationshipsTransactions instantiates a new AccountResourceRelationshipsTransactions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountResourceRelationshipsTransactions() *AccountResourceRelationshipsTransactions {
	this := AccountResourceRelationshipsTransactions{}
	return &this
}

// NewAccountResourceRelationshipsTransactionsWithDefaults instantiates a new AccountResourceRelationshipsTransactions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountResourceRelationshipsTransactionsWithDefaults() *AccountResourceRelationshipsTransactions {
	this := AccountResourceRelationshipsTransactions{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AccountResourceRelationshipsTransactions) GetLinks() AccountResourceRelationshipsTransactionsLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountResourceRelationshipsTransactionsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResourceRelationshipsTransactions) GetLinksOk() (*AccountResourceRelationshipsTransactionsLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AccountResourceRelationshipsTransactions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceRelationshipsTransactionsLinks and assigns it to the Links field.
func (o *AccountResourceRelationshipsTransactions) SetLinks(v AccountResourceRelationshipsTransactionsLinks) {
	o.Links = &v
}

func (o AccountResourceRelationshipsTransactions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountResourceRelationshipsTransactions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableAccountResourceRelationshipsTransactions struct {
	value *AccountResourceRelationshipsTransactions
	isSet bool
}

func (v NullableAccountResourceRelationshipsTransactions) Get() *AccountResourceRelationshipsTransactions {
	return v.value
}

func (v *NullableAccountResourceRelationshipsTransactions) Set(val *AccountResourceRelationshipsTransactions) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResourceRelationshipsTransactions) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResourceRelationshipsTransactions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResourceRelationshipsTransactions(val *AccountResourceRelationshipsTransactions) *NullableAccountResourceRelationshipsTransactions {
	return &NullableAccountResourceRelationshipsTransactions{value: val, isSet: true}
}

func (v NullableAccountResourceRelationshipsTransactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResourceRelationshipsTransactions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
