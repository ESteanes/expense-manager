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

// checks if the PingResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PingResponseMeta{}

// PingResponseMeta struct for PingResponseMeta
type PingResponseMeta struct {
	// The unique identifier of the authenticated customer.
	Id string `json:"id"`
	// A cute emoji that represents the response status.
	StatusEmoji string `json:"statusEmoji"`
}

type _PingResponseMeta PingResponseMeta

// NewPingResponseMeta instantiates a new PingResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingResponseMeta(id string, statusEmoji string) *PingResponseMeta {
	this := PingResponseMeta{}
	this.Id = id
	this.StatusEmoji = statusEmoji
	return &this
}

// NewPingResponseMetaWithDefaults instantiates a new PingResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingResponseMetaWithDefaults() *PingResponseMeta {
	this := PingResponseMeta{}
	return &this
}

// GetId returns the Id field value
func (o *PingResponseMeta) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PingResponseMeta) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PingResponseMeta) SetId(v string) {
	o.Id = v
}

// GetStatusEmoji returns the StatusEmoji field value
func (o *PingResponseMeta) GetStatusEmoji() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusEmoji
}

// GetStatusEmojiOk returns a tuple with the StatusEmoji field value
// and a boolean to check if the value has been set.
func (o *PingResponseMeta) GetStatusEmojiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusEmoji, true
}

// SetStatusEmoji sets field value
func (o *PingResponseMeta) SetStatusEmoji(v string) {
	o.StatusEmoji = v
}

func (o PingResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PingResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["statusEmoji"] = o.StatusEmoji
	return toSerialize, nil
}

func (o *PingResponseMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"statusEmoji",
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

	varPingResponseMeta := _PingResponseMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPingResponseMeta)

	if err != nil {
		return err
	}

	*o = PingResponseMeta(varPingResponseMeta)

	return err
}

type NullablePingResponseMeta struct {
	value *PingResponseMeta
	isSet bool
}

func (v NullablePingResponseMeta) Get() *PingResponseMeta {
	return v.value
}

func (v *NullablePingResponseMeta) Set(val *PingResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullablePingResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullablePingResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingResponseMeta(val *PingResponseMeta) *NullablePingResponseMeta {
	return &NullablePingResponseMeta{value: val, isSet: true}
}

func (v NullablePingResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
