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

// checks if the CustomerObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerObject{}

// CustomerObject Provides information about the customer who initiated a transaction 
type CustomerObject struct {
	// The Upname or preferred name of the customer 
	DisplayName string `json:"displayName"`
}

type _CustomerObject CustomerObject

// NewCustomerObject instantiates a new CustomerObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerObject(displayName string) *CustomerObject {
	this := CustomerObject{}
	this.DisplayName = displayName
	return &this
}

// NewCustomerObjectWithDefaults instantiates a new CustomerObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerObjectWithDefaults() *CustomerObject {
	this := CustomerObject{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *CustomerObject) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CustomerObject) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CustomerObject) SetDisplayName(v string) {
	o.DisplayName = v
}

func (o CustomerObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	return toSerialize, nil
}

func (o *CustomerObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
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

	varCustomerObject := _CustomerObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerObject)

	if err != nil {
		return err
	}

	*o = CustomerObject(varCustomerObject)

	return err
}

type NullableCustomerObject struct {
	value *CustomerObject
	isSet bool
}

func (v NullableCustomerObject) Get() *CustomerObject {
	return v.value
}

func (v *NullableCustomerObject) Set(val *CustomerObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerObject(val *CustomerObject) *NullableCustomerObject {
	return &NullableCustomerObject{value: val, isSet: true}
}

func (v NullableCustomerObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


