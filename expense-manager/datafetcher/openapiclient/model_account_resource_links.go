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

// checks if the AccountResourceLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountResourceLinks{}

// AccountResourceLinks struct for AccountResourceLinks
type AccountResourceLinks struct {
	// The canonical link to this resource within the API.
	Self string `json:"self"`
}

type _AccountResourceLinks AccountResourceLinks

// NewAccountResourceLinks instantiates a new AccountResourceLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountResourceLinks(self string) *AccountResourceLinks {
	this := AccountResourceLinks{}
	this.Self = self
	return &this
}

// NewAccountResourceLinksWithDefaults instantiates a new AccountResourceLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountResourceLinksWithDefaults() *AccountResourceLinks {
	this := AccountResourceLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *AccountResourceLinks) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *AccountResourceLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *AccountResourceLinks) SetSelf(v string) {
	o.Self = v
}

func (o AccountResourceLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountResourceLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	return toSerialize, nil
}

func (o *AccountResourceLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varAccountResourceLinks := _AccountResourceLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountResourceLinks)

	if err != nil {
		return err
	}

	*o = AccountResourceLinks(varAccountResourceLinks)

	return err
}

type NullableAccountResourceLinks struct {
	value *AccountResourceLinks
	isSet bool
}

func (v NullableAccountResourceLinks) Get() *AccountResourceLinks {
	return v.value
}

func (v *NullableAccountResourceLinks) Set(val *AccountResourceLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResourceLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResourceLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResourceLinks(val *AccountResourceLinks) *NullableAccountResourceLinks {
	return &NullableAccountResourceLinks{value: val, isSet: true}
}

func (v NullableAccountResourceLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResourceLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
