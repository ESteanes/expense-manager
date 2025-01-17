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

// checks if the GetWebhookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWebhookResponse{}

// GetWebhookResponse Successful response to get a single webhook.
type GetWebhookResponse struct {
	// The webhook returned in this response.
	Data WebhookResource `json:"data"`
}

type _GetWebhookResponse GetWebhookResponse

// NewGetWebhookResponse instantiates a new GetWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhookResponse(data WebhookResource) *GetWebhookResponse {
	this := GetWebhookResponse{}
	this.Data = data
	return &this
}

// NewGetWebhookResponseWithDefaults instantiates a new GetWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhookResponseWithDefaults() *GetWebhookResponse {
	this := GetWebhookResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetWebhookResponse) GetData() WebhookResource {
	if o == nil {
		var ret WebhookResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetWebhookResponse) GetDataOk() (*WebhookResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetWebhookResponse) SetData(v WebhookResource) {
	o.Data = v
}

func (o GetWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWebhookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetWebhookResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetWebhookResponse := _GetWebhookResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetWebhookResponse)

	if err != nil {
		return err
	}

	*o = GetWebhookResponse(varGetWebhookResponse)

	return err
}

type NullableGetWebhookResponse struct {
	value *GetWebhookResponse
	isSet bool
}

func (v NullableGetWebhookResponse) Get() *GetWebhookResponse {
	return v.value
}

func (v *NullableGetWebhookResponse) Set(val *GetWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhookResponse(val *GetWebhookResponse) *NullableGetWebhookResponse {
	return &NullableGetWebhookResponse{value: val, isSet: true}
}

func (v NullableGetWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
