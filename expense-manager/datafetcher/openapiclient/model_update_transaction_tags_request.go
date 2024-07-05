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

// checks if the UpdateTransactionTagsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTransactionTagsRequest{}

// UpdateTransactionTagsRequest Request to add or remove tags associated with a transaction.
type UpdateTransactionTagsRequest struct {
	// The tags to add to or remove from the transaction.
	Data []TagInputResourceIdentifier `json:"data"`
}

type _UpdateTransactionTagsRequest UpdateTransactionTagsRequest

// NewUpdateTransactionTagsRequest instantiates a new UpdateTransactionTagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTransactionTagsRequest(data []TagInputResourceIdentifier) *UpdateTransactionTagsRequest {
	this := UpdateTransactionTagsRequest{}
	this.Data = data
	return &this
}

// NewUpdateTransactionTagsRequestWithDefaults instantiates a new UpdateTransactionTagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTransactionTagsRequestWithDefaults() *UpdateTransactionTagsRequest {
	this := UpdateTransactionTagsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateTransactionTagsRequest) GetData() []TagInputResourceIdentifier {
	if o == nil {
		var ret []TagInputResourceIdentifier
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateTransactionTagsRequest) GetDataOk() ([]TagInputResourceIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *UpdateTransactionTagsRequest) SetData(v []TagInputResourceIdentifier) {
	o.Data = v
}

func (o UpdateTransactionTagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTransactionTagsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateTransactionTagsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateTransactionTagsRequest := _UpdateTransactionTagsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateTransactionTagsRequest)

	if err != nil {
		return err
	}

	*o = UpdateTransactionTagsRequest(varUpdateTransactionTagsRequest)

	return err
}

type NullableUpdateTransactionTagsRequest struct {
	value *UpdateTransactionTagsRequest
	isSet bool
}

func (v NullableUpdateTransactionTagsRequest) Get() *UpdateTransactionTagsRequest {
	return v.value
}

func (v *NullableUpdateTransactionTagsRequest) Set(val *UpdateTransactionTagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTransactionTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTransactionTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTransactionTagsRequest(val *UpdateTransactionTagsRequest) *NullableUpdateTransactionTagsRequest {
	return &NullableUpdateTransactionTagsRequest{value: val, isSet: true}
}

func (v NullableUpdateTransactionTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTransactionTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
