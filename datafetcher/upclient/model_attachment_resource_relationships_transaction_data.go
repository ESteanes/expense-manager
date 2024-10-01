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

// checks if the AttachmentResourceRelationshipsTransactionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentResourceRelationshipsTransactionData{}

// AttachmentResourceRelationshipsTransactionData struct for AttachmentResourceRelationshipsTransactionData
type AttachmentResourceRelationshipsTransactionData struct {
	// The type of this resource: `transactions`
	Type string `json:"type"`
	// The unique identifier of the resource within its type. 
	Id string `json:"id"`
}

type _AttachmentResourceRelationshipsTransactionData AttachmentResourceRelationshipsTransactionData

// NewAttachmentResourceRelationshipsTransactionData instantiates a new AttachmentResourceRelationshipsTransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentResourceRelationshipsTransactionData(type_ string, id string) *AttachmentResourceRelationshipsTransactionData {
	this := AttachmentResourceRelationshipsTransactionData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAttachmentResourceRelationshipsTransactionDataWithDefaults instantiates a new AttachmentResourceRelationshipsTransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentResourceRelationshipsTransactionDataWithDefaults() *AttachmentResourceRelationshipsTransactionData {
	this := AttachmentResourceRelationshipsTransactionData{}
	return &this
}

// GetType returns the Type field value
func (o *AttachmentResourceRelationshipsTransactionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttachmentResourceRelationshipsTransactionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AttachmentResourceRelationshipsTransactionData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AttachmentResourceRelationshipsTransactionData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachmentResourceRelationshipsTransactionData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttachmentResourceRelationshipsTransactionData) SetId(v string) {
	o.Id = v
}

func (o AttachmentResourceRelationshipsTransactionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentResourceRelationshipsTransactionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AttachmentResourceRelationshipsTransactionData) UnmarshalJSON(data []byte) (err error) {
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

	varAttachmentResourceRelationshipsTransactionData := _AttachmentResourceRelationshipsTransactionData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttachmentResourceRelationshipsTransactionData)

	if err != nil {
		return err
	}

	*o = AttachmentResourceRelationshipsTransactionData(varAttachmentResourceRelationshipsTransactionData)

	return err
}

type NullableAttachmentResourceRelationshipsTransactionData struct {
	value *AttachmentResourceRelationshipsTransactionData
	isSet bool
}

func (v NullableAttachmentResourceRelationshipsTransactionData) Get() *AttachmentResourceRelationshipsTransactionData {
	return v.value
}

func (v *NullableAttachmentResourceRelationshipsTransactionData) Set(val *AttachmentResourceRelationshipsTransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentResourceRelationshipsTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentResourceRelationshipsTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentResourceRelationshipsTransactionData(val *AttachmentResourceRelationshipsTransactionData) *NullableAttachmentResourceRelationshipsTransactionData {
	return &NullableAttachmentResourceRelationshipsTransactionData{value: val, isSet: true}
}

func (v NullableAttachmentResourceRelationshipsTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentResourceRelationshipsTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


