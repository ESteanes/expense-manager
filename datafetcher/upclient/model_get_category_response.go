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

// checks if the GetCategoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCategoryResponse{}

// GetCategoryResponse Successful response to get a single category and its ancestry. 
type GetCategoryResponse struct {
	// The category returned in this response. 
	Data CategoryResource `json:"data"`
}

type _GetCategoryResponse GetCategoryResponse

// NewGetCategoryResponse instantiates a new GetCategoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCategoryResponse(data CategoryResource) *GetCategoryResponse {
	this := GetCategoryResponse{}
	this.Data = data
	return &this
}

// NewGetCategoryResponseWithDefaults instantiates a new GetCategoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCategoryResponseWithDefaults() *GetCategoryResponse {
	this := GetCategoryResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCategoryResponse) GetData() CategoryResource {
	if o == nil {
		var ret CategoryResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCategoryResponse) GetDataOk() (*CategoryResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCategoryResponse) SetData(v CategoryResource) {
	o.Data = v
}

func (o GetCategoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCategoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCategoryResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetCategoryResponse := _GetCategoryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCategoryResponse)

	if err != nil {
		return err
	}

	*o = GetCategoryResponse(varGetCategoryResponse)

	return err
}

type NullableGetCategoryResponse struct {
	value *GetCategoryResponse
	isSet bool
}

func (v NullableGetCategoryResponse) Get() *GetCategoryResponse {
	return v.value
}

func (v *NullableGetCategoryResponse) Set(val *GetCategoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCategoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCategoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCategoryResponse(val *GetCategoryResponse) *NullableGetCategoryResponse {
	return &NullableGetCategoryResponse{value: val, isSet: true}
}

func (v NullableGetCategoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCategoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


