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

// checks if the CategoryResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryResourceAttributes{}

// CategoryResourceAttributes struct for CategoryResourceAttributes
type CategoryResourceAttributes struct {
	// The name of this category as seen in the Up application.
	Name string `json:"name"`
}

type _CategoryResourceAttributes CategoryResourceAttributes

// NewCategoryResourceAttributes instantiates a new CategoryResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryResourceAttributes(name string) *CategoryResourceAttributes {
	this := CategoryResourceAttributes{}
	this.Name = name
	return &this
}

// NewCategoryResourceAttributesWithDefaults instantiates a new CategoryResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryResourceAttributesWithDefaults() *CategoryResourceAttributes {
	this := CategoryResourceAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *CategoryResourceAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CategoryResourceAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CategoryResourceAttributes) SetName(v string) {
	o.Name = v
}

func (o CategoryResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *CategoryResourceAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCategoryResourceAttributes := _CategoryResourceAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategoryResourceAttributes)

	if err != nil {
		return err
	}

	*o = CategoryResourceAttributes(varCategoryResourceAttributes)

	return err
}

type NullableCategoryResourceAttributes struct {
	value *CategoryResourceAttributes
	isSet bool
}

func (v NullableCategoryResourceAttributes) Get() *CategoryResourceAttributes {
	return v.value
}

func (v *NullableCategoryResourceAttributes) Set(val *CategoryResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryResourceAttributes(val *CategoryResourceAttributes) *NullableCategoryResourceAttributes {
	return &NullableCategoryResourceAttributes{value: val, isSet: true}
}

func (v NullableCategoryResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
