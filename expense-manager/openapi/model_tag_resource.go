/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TagResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagResource{}

// TagResource Provides information about a tag.
type TagResource struct {
	// The type of this resource: `tags`
	Type string `json:"type"`
	// The label of the tag, which also acts as the tag’s unique identifier.
	Id            string                       `json:"id"`
	Relationships AccountResourceRelationships `json:"relationships"`
}

type _TagResource TagResource

// NewTagResource instantiates a new TagResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagResource(type_ string, id string, relationships AccountResourceRelationships) *TagResource {
	this := TagResource{}
	this.Type = type_
	this.Id = id
	this.Relationships = relationships
	return &this
}

// NewTagResourceWithDefaults instantiates a new TagResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagResourceWithDefaults() *TagResource {
	this := TagResource{}
	return &this
}

// GetType returns the Type field value
func (o *TagResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TagResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TagResource) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TagResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagResource) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value
func (o *TagResource) GetRelationships() AccountResourceRelationships {
	if o == nil {
		var ret AccountResourceRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *TagResource) GetRelationshipsOk() (*AccountResourceRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *TagResource) SetRelationships(v AccountResourceRelationships) {
	o.Relationships = v
}

func (o TagResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *TagResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"relationships",
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

	varTagResource := _TagResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagResource)

	if err != nil {
		return err
	}

	*o = TagResource(varTagResource)

	return err
}

type NullableTagResource struct {
	value *TagResource
	isSet bool
}

func (v NullableTagResource) Get() *TagResource {
	return v.value
}

func (v *NullableTagResource) Set(val *TagResource) {
	v.value = val
	v.isSet = true
}

func (v NullableTagResource) IsSet() bool {
	return v.isSet
}

func (v *NullableTagResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagResource(val *TagResource) *NullableTagResource {
	return &NullableTagResource{value: val, isSet: true}
}

func (v NullableTagResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
