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

// checks if the CategoryResourceRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryResourceRelationships{}

// CategoryResourceRelationships struct for CategoryResourceRelationships
type CategoryResourceRelationships struct {
	Parent CategoryResourceRelationshipsParent `json:"parent"`
	Children CategoryResourceRelationshipsChildren `json:"children"`
}

type _CategoryResourceRelationships CategoryResourceRelationships

// NewCategoryResourceRelationships instantiates a new CategoryResourceRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryResourceRelationships(parent CategoryResourceRelationshipsParent, children CategoryResourceRelationshipsChildren) *CategoryResourceRelationships {
	this := CategoryResourceRelationships{}
	this.Parent = parent
	this.Children = children
	return &this
}

// NewCategoryResourceRelationshipsWithDefaults instantiates a new CategoryResourceRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryResourceRelationshipsWithDefaults() *CategoryResourceRelationships {
	this := CategoryResourceRelationships{}
	return &this
}

// GetParent returns the Parent field value
func (o *CategoryResourceRelationships) GetParent() CategoryResourceRelationshipsParent {
	if o == nil {
		var ret CategoryResourceRelationshipsParent
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *CategoryResourceRelationships) GetParentOk() (*CategoryResourceRelationshipsParent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *CategoryResourceRelationships) SetParent(v CategoryResourceRelationshipsParent) {
	o.Parent = v
}

// GetChildren returns the Children field value
func (o *CategoryResourceRelationships) GetChildren() CategoryResourceRelationshipsChildren {
	if o == nil {
		var ret CategoryResourceRelationshipsChildren
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *CategoryResourceRelationships) GetChildrenOk() (*CategoryResourceRelationshipsChildren, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Children, true
}

// SetChildren sets field value
func (o *CategoryResourceRelationships) SetChildren(v CategoryResourceRelationshipsChildren) {
	o.Children = v
}

func (o CategoryResourceRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryResourceRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parent"] = o.Parent
	toSerialize["children"] = o.Children
	return toSerialize, nil
}

func (o *CategoryResourceRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parent",
		"children",
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

	varCategoryResourceRelationships := _CategoryResourceRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategoryResourceRelationships)

	if err != nil {
		return err
	}

	*o = CategoryResourceRelationships(varCategoryResourceRelationships)

	return err
}

type NullableCategoryResourceRelationships struct {
	value *CategoryResourceRelationships
	isSet bool
}

func (v NullableCategoryResourceRelationships) Get() *CategoryResourceRelationships {
	return v.value
}

func (v *NullableCategoryResourceRelationships) Set(val *CategoryResourceRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryResourceRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryResourceRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryResourceRelationships(val *CategoryResourceRelationships) *NullableCategoryResourceRelationships {
	return &NullableCategoryResourceRelationships{value: val, isSet: true}
}

func (v NullableCategoryResourceRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryResourceRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


