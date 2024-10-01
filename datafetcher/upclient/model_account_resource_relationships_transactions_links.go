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

// checks if the AccountResourceRelationshipsTransactionsLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountResourceRelationshipsTransactionsLinks{}

// AccountResourceRelationshipsTransactionsLinks struct for AccountResourceRelationshipsTransactionsLinks
type AccountResourceRelationshipsTransactionsLinks struct {
	// The link to retrieve the related resource(s) in this relationship. 
	Related string `json:"related"`
}

type _AccountResourceRelationshipsTransactionsLinks AccountResourceRelationshipsTransactionsLinks

// NewAccountResourceRelationshipsTransactionsLinks instantiates a new AccountResourceRelationshipsTransactionsLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountResourceRelationshipsTransactionsLinks(related string) *AccountResourceRelationshipsTransactionsLinks {
	this := AccountResourceRelationshipsTransactionsLinks{}
	this.Related = related
	return &this
}

// NewAccountResourceRelationshipsTransactionsLinksWithDefaults instantiates a new AccountResourceRelationshipsTransactionsLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountResourceRelationshipsTransactionsLinksWithDefaults() *AccountResourceRelationshipsTransactionsLinks {
	this := AccountResourceRelationshipsTransactionsLinks{}
	return &this
}

// GetRelated returns the Related field value
func (o *AccountResourceRelationshipsTransactionsLinks) GetRelated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *AccountResourceRelationshipsTransactionsLinks) GetRelatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *AccountResourceRelationshipsTransactionsLinks) SetRelated(v string) {
	o.Related = v
}

func (o AccountResourceRelationshipsTransactionsLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountResourceRelationshipsTransactionsLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["related"] = o.Related
	return toSerialize, nil
}

func (o *AccountResourceRelationshipsTransactionsLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"related",
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

	varAccountResourceRelationshipsTransactionsLinks := _AccountResourceRelationshipsTransactionsLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountResourceRelationshipsTransactionsLinks)

	if err != nil {
		return err
	}

	*o = AccountResourceRelationshipsTransactionsLinks(varAccountResourceRelationshipsTransactionsLinks)

	return err
}

type NullableAccountResourceRelationshipsTransactionsLinks struct {
	value *AccountResourceRelationshipsTransactionsLinks
	isSet bool
}

func (v NullableAccountResourceRelationshipsTransactionsLinks) Get() *AccountResourceRelationshipsTransactionsLinks {
	return v.value
}

func (v *NullableAccountResourceRelationshipsTransactionsLinks) Set(val *AccountResourceRelationshipsTransactionsLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResourceRelationshipsTransactionsLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResourceRelationshipsTransactionsLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResourceRelationshipsTransactionsLinks(val *AccountResourceRelationshipsTransactionsLinks) *NullableAccountResourceRelationshipsTransactionsLinks {
	return &NullableAccountResourceRelationshipsTransactionsLinks{value: val, isSet: true}
}

func (v NullableAccountResourceRelationshipsTransactionsLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResourceRelationshipsTransactionsLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


