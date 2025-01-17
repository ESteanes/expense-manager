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

// checks if the CategoryResourceRelationshipsParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryResourceRelationshipsParent{}

// CategoryResourceRelationshipsParent struct for CategoryResourceRelationshipsParent
type CategoryResourceRelationshipsParent struct {
	Data  NullableCategoryResourceRelationshipsParentData `json:"data"`
	Links *AccountResourceRelationshipsTransactionsLinks  `json:"links,omitempty"`
}

type _CategoryResourceRelationshipsParent CategoryResourceRelationshipsParent

// NewCategoryResourceRelationshipsParent instantiates a new CategoryResourceRelationshipsParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryResourceRelationshipsParent(data NullableCategoryResourceRelationshipsParentData) *CategoryResourceRelationshipsParent {
	this := CategoryResourceRelationshipsParent{}
	this.Data = data
	return &this
}

// NewCategoryResourceRelationshipsParentWithDefaults instantiates a new CategoryResourceRelationshipsParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryResourceRelationshipsParentWithDefaults() *CategoryResourceRelationshipsParent {
	this := CategoryResourceRelationshipsParent{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for CategoryResourceRelationshipsParentData will be returned
func (o *CategoryResourceRelationshipsParent) GetData() CategoryResourceRelationshipsParentData {
	if o == nil || o.Data.Get() == nil {
		var ret CategoryResourceRelationshipsParentData
		return ret
	}

	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CategoryResourceRelationshipsParent) GetDataOk() (*CategoryResourceRelationshipsParentData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value
func (o *CategoryResourceRelationshipsParent) SetData(v CategoryResourceRelationshipsParentData) {
	o.Data.Set(&v)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CategoryResourceRelationshipsParent) GetLinks() AccountResourceRelationshipsTransactionsLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountResourceRelationshipsTransactionsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryResourceRelationshipsParent) GetLinksOk() (*AccountResourceRelationshipsTransactionsLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CategoryResourceRelationshipsParent) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceRelationshipsTransactionsLinks and assigns it to the Links field.
func (o *CategoryResourceRelationshipsParent) SetLinks(v AccountResourceRelationshipsTransactionsLinks) {
	o.Links = &v
}

func (o CategoryResourceRelationshipsParent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryResourceRelationshipsParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data.Get()
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *CategoryResourceRelationshipsParent) UnmarshalJSON(data []byte) (err error) {
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

	varCategoryResourceRelationshipsParent := _CategoryResourceRelationshipsParent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategoryResourceRelationshipsParent)

	if err != nil {
		return err
	}

	*o = CategoryResourceRelationshipsParent(varCategoryResourceRelationshipsParent)

	return err
}

type NullableCategoryResourceRelationshipsParent struct {
	value *CategoryResourceRelationshipsParent
	isSet bool
}

func (v NullableCategoryResourceRelationshipsParent) Get() *CategoryResourceRelationshipsParent {
	return v.value
}

func (v *NullableCategoryResourceRelationshipsParent) Set(val *CategoryResourceRelationshipsParent) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryResourceRelationshipsParent) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryResourceRelationshipsParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryResourceRelationshipsParent(val *CategoryResourceRelationshipsParent) *NullableCategoryResourceRelationshipsParent {
	return &NullableCategoryResourceRelationshipsParent{value: val, isSet: true}
}

func (v NullableCategoryResourceRelationshipsParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryResourceRelationshipsParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
