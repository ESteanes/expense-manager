/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ErrorObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorObject{}

// ErrorObject Provides information about an error processing a request. 
type ErrorObject struct {
	// The HTTP status code associated with this error. This can also be obtained from the response headers. The status indicates the broad type of error according to HTTP semantics. 
	Status string `json:"status"`
	// A short description of this error. This should be stable across multiple occurrences of this type of error and typically expands on the reason for the status code. 
	Title string `json:"title"`
	// A detailed description of this error. This should be considered unique to individual occurrences of an error and subject to change. It is useful for debugging purposes. 
	Detail string `json:"detail"`
	Source *ErrorObjectSource `json:"source,omitempty"`
}

type _ErrorObject ErrorObject

// NewErrorObject instantiates a new ErrorObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorObject(status string, title string, detail string) *ErrorObject {
	this := ErrorObject{}
	this.Status = status
	this.Title = title
	this.Detail = detail
	return &this
}

// NewErrorObjectWithDefaults instantiates a new ErrorObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorObjectWithDefaults() *ErrorObject {
	this := ErrorObject{}
	return &this
}

// GetStatus returns the Status field value
func (o *ErrorObject) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ErrorObject) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ErrorObject) SetStatus(v string) {
	o.Status = v
}

// GetTitle returns the Title field value
func (o *ErrorObject) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ErrorObject) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ErrorObject) SetTitle(v string) {
	o.Title = v
}

// GetDetail returns the Detail field value
func (o *ErrorObject) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *ErrorObject) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *ErrorObject) SetDetail(v string) {
	o.Detail = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ErrorObject) GetSource() ErrorObjectSource {
	if o == nil || IsNil(o.Source) {
		var ret ErrorObjectSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorObject) GetSourceOk() (*ErrorObjectSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ErrorObject) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given ErrorObjectSource and assigns it to the Source field.
func (o *ErrorObject) SetSource(v ErrorObjectSource) {
	o.Source = &v
}

func (o ErrorObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["detail"] = o.Detail
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

func (o *ErrorObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"title",
		"detail",
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

	varErrorObject := _ErrorObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorObject)

	if err != nil {
		return err
	}

	*o = ErrorObject(varErrorObject)

	return err
}

type NullableErrorObject struct {
	value *ErrorObject
	isSet bool
}

func (v NullableErrorObject) Get() *ErrorObject {
	return v.value
}

func (v *NullableErrorObject) Set(val *ErrorObject) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorObject) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorObject(val *ErrorObject) *NullableErrorObject {
	return &NullableErrorObject{value: val, isSet: true}
}

func (v NullableErrorObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


