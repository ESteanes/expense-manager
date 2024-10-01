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

// checks if the CategoryInputResourceIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryInputResourceIdentifier{}

// CategoryInputResourceIdentifier Uniquely identifies a category in the API. 
type CategoryInputResourceIdentifier struct {
	// The type of this resource: `categories`
	Type string `json:"type"`
	// The unique identifier of the category, as returned by the `/categories` endpoint. 
	Id string `json:"id"`
}

type _CategoryInputResourceIdentifier CategoryInputResourceIdentifier

// NewCategoryInputResourceIdentifier instantiates a new CategoryInputResourceIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryInputResourceIdentifier(type_ string, id string) *CategoryInputResourceIdentifier {
	this := CategoryInputResourceIdentifier{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCategoryInputResourceIdentifierWithDefaults instantiates a new CategoryInputResourceIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryInputResourceIdentifierWithDefaults() *CategoryInputResourceIdentifier {
	this := CategoryInputResourceIdentifier{}
	return &this
}

// GetType returns the Type field value
func (o *CategoryInputResourceIdentifier) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CategoryInputResourceIdentifier) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CategoryInputResourceIdentifier) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CategoryInputResourceIdentifier) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CategoryInputResourceIdentifier) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CategoryInputResourceIdentifier) SetId(v string) {
	o.Id = v
}

func (o CategoryInputResourceIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryInputResourceIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *CategoryInputResourceIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varCategoryInputResourceIdentifier := _CategoryInputResourceIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategoryInputResourceIdentifier)

	if err != nil {
		return err
	}

	*o = CategoryInputResourceIdentifier(varCategoryInputResourceIdentifier)

	return err
}

type NullableCategoryInputResourceIdentifier struct {
	value *CategoryInputResourceIdentifier
	isSet bool
}

func (v NullableCategoryInputResourceIdentifier) Get() *CategoryInputResourceIdentifier {
	return v.value
}

func (v *NullableCategoryInputResourceIdentifier) Set(val *CategoryInputResourceIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryInputResourceIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryInputResourceIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryInputResourceIdentifier(val *CategoryInputResourceIdentifier) *NullableCategoryInputResourceIdentifier {
	return &NullableCategoryInputResourceIdentifier{value: val, isSet: true}
}

func (v NullableCategoryInputResourceIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryInputResourceIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


