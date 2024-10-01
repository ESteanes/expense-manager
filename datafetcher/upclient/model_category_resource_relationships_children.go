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

// checks if the CategoryResourceRelationshipsChildren type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryResourceRelationshipsChildren{}

// CategoryResourceRelationshipsChildren struct for CategoryResourceRelationshipsChildren
type CategoryResourceRelationshipsChildren struct {
	Data []CategoryResourceRelationshipsChildrenDataInner `json:"data"`
	Links *AccountResourceRelationshipsTransactionsLinks `json:"links,omitempty"`
}

type _CategoryResourceRelationshipsChildren CategoryResourceRelationshipsChildren

// NewCategoryResourceRelationshipsChildren instantiates a new CategoryResourceRelationshipsChildren object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryResourceRelationshipsChildren(data []CategoryResourceRelationshipsChildrenDataInner) *CategoryResourceRelationshipsChildren {
	this := CategoryResourceRelationshipsChildren{}
	this.Data = data
	return &this
}

// NewCategoryResourceRelationshipsChildrenWithDefaults instantiates a new CategoryResourceRelationshipsChildren object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryResourceRelationshipsChildrenWithDefaults() *CategoryResourceRelationshipsChildren {
	this := CategoryResourceRelationshipsChildren{}
	return &this
}

// GetData returns the Data field value
func (o *CategoryResourceRelationshipsChildren) GetData() []CategoryResourceRelationshipsChildrenDataInner {
	if o == nil {
		var ret []CategoryResourceRelationshipsChildrenDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CategoryResourceRelationshipsChildren) GetDataOk() ([]CategoryResourceRelationshipsChildrenDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CategoryResourceRelationshipsChildren) SetData(v []CategoryResourceRelationshipsChildrenDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CategoryResourceRelationshipsChildren) GetLinks() AccountResourceRelationshipsTransactionsLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountResourceRelationshipsTransactionsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryResourceRelationshipsChildren) GetLinksOk() (*AccountResourceRelationshipsTransactionsLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CategoryResourceRelationshipsChildren) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceRelationshipsTransactionsLinks and assigns it to the Links field.
func (o *CategoryResourceRelationshipsChildren) SetLinks(v AccountResourceRelationshipsTransactionsLinks) {
	o.Links = &v
}

func (o CategoryResourceRelationshipsChildren) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryResourceRelationshipsChildren) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *CategoryResourceRelationshipsChildren) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varCategoryResourceRelationshipsChildren := _CategoryResourceRelationshipsChildren{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategoryResourceRelationshipsChildren)

	if err != nil {
		return err
	}

	*o = CategoryResourceRelationshipsChildren(varCategoryResourceRelationshipsChildren)

	return err
}

type NullableCategoryResourceRelationshipsChildren struct {
	value *CategoryResourceRelationshipsChildren
	isSet bool
}

func (v NullableCategoryResourceRelationshipsChildren) Get() *CategoryResourceRelationshipsChildren {
	return v.value
}

func (v *NullableCategoryResourceRelationshipsChildren) Set(val *CategoryResourceRelationshipsChildren) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryResourceRelationshipsChildren) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryResourceRelationshipsChildren) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryResourceRelationshipsChildren(val *CategoryResourceRelationshipsChildren) *NullableCategoryResourceRelationshipsChildren {
	return &NullableCategoryResourceRelationshipsChildren{value: val, isSet: true}
}

func (v NullableCategoryResourceRelationshipsChildren) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryResourceRelationshipsChildren) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


