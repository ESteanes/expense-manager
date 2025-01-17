/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package upclient

import (
	"encoding/json"
)

// checks if the ErrorObjectSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorObjectSource{}

// ErrorObjectSource If applicable, location in the request that this error relates to. This may be a parameter in the query string, or a an attribute in the request body.
type ErrorObjectSource struct {
	// If this error relates to a query parameter, the name of the parameter.
	Parameter *string `json:"parameter,omitempty"`
	// If this error relates to an attribute in the request body, a rfc-6901 JSON pointer to the attribute.
	Pointer *string `json:"pointer,omitempty"`
}

// NewErrorObjectSource instantiates a new ErrorObjectSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorObjectSource() *ErrorObjectSource {
	this := ErrorObjectSource{}
	return &this
}

// NewErrorObjectSourceWithDefaults instantiates a new ErrorObjectSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorObjectSourceWithDefaults() *ErrorObjectSource {
	this := ErrorObjectSource{}
	return &this
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *ErrorObjectSource) GetParameter() string {
	if o == nil || IsNil(o.Parameter) {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorObjectSource) GetParameterOk() (*string, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *ErrorObjectSource) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *ErrorObjectSource) SetParameter(v string) {
	o.Parameter = &v
}

// GetPointer returns the Pointer field value if set, zero value otherwise.
func (o *ErrorObjectSource) GetPointer() string {
	if o == nil || IsNil(o.Pointer) {
		var ret string
		return ret
	}
	return *o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorObjectSource) GetPointerOk() (*string, bool) {
	if o == nil || IsNil(o.Pointer) {
		return nil, false
	}
	return o.Pointer, true
}

// HasPointer returns a boolean if a field has been set.
func (o *ErrorObjectSource) HasPointer() bool {
	if o != nil && !IsNil(o.Pointer) {
		return true
	}

	return false
}

// SetPointer gets a reference to the given string and assigns it to the Pointer field.
func (o *ErrorObjectSource) SetPointer(v string) {
	o.Pointer = &v
}

func (o ErrorObjectSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorObjectSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !IsNil(o.Pointer) {
		toSerialize["pointer"] = o.Pointer
	}
	return toSerialize, nil
}

type NullableErrorObjectSource struct {
	value *ErrorObjectSource
	isSet bool
}

func (v NullableErrorObjectSource) Get() *ErrorObjectSource {
	return v.value
}

func (v *NullableErrorObjectSource) Set(val *ErrorObjectSource) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorObjectSource) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorObjectSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorObjectSource(val *ErrorObjectSource) *NullableErrorObjectSource {
	return &NullableErrorObjectSource{value: val, isSet: true}
}

func (v NullableErrorObjectSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorObjectSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
