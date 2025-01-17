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

// checks if the TagInputResourceIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagInputResourceIdentifier{}

// TagInputResourceIdentifier Uniquely identifies a single tag in the API.
type TagInputResourceIdentifier struct {
	// The type of this resource: `tags`
	Type string `json:"type"`
	// The label of the tag, which also acts as the tag’s unique identifier.
	Id string `json:"id"`
}

type _TagInputResourceIdentifier TagInputResourceIdentifier

// NewTagInputResourceIdentifier instantiates a new TagInputResourceIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagInputResourceIdentifier(type_ string, id string) *TagInputResourceIdentifier {
	this := TagInputResourceIdentifier{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewTagInputResourceIdentifierWithDefaults instantiates a new TagInputResourceIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagInputResourceIdentifierWithDefaults() *TagInputResourceIdentifier {
	this := TagInputResourceIdentifier{}
	return &this
}

// GetType returns the Type field value
func (o *TagInputResourceIdentifier) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TagInputResourceIdentifier) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TagInputResourceIdentifier) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TagInputResourceIdentifier) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagInputResourceIdentifier) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagInputResourceIdentifier) SetId(v string) {
	o.Id = v
}

func (o TagInputResourceIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagInputResourceIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *TagInputResourceIdentifier) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTagInputResourceIdentifier := _TagInputResourceIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagInputResourceIdentifier)

	if err != nil {
		return err
	}

	*o = TagInputResourceIdentifier(varTagInputResourceIdentifier)

	return err
}

type NullableTagInputResourceIdentifier struct {
	value *TagInputResourceIdentifier
	isSet bool
}

func (v NullableTagInputResourceIdentifier) Get() *TagInputResourceIdentifier {
	return v.value
}

func (v *NullableTagInputResourceIdentifier) Set(val *TagInputResourceIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableTagInputResourceIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableTagInputResourceIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagInputResourceIdentifier(val *TagInputResourceIdentifier) *NullableTagInputResourceIdentifier {
	return &NullableTagInputResourceIdentifier{value: val, isSet: true}
}

func (v NullableTagInputResourceIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagInputResourceIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
