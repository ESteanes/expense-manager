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

// checks if the CreateWebhookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebhookResponse{}

// CreateWebhookResponse Successful response after creating a webhook.
type CreateWebhookResponse struct {
	// The webhook that was created.
	Data WebhookResource `json:"data"`
}

type _CreateWebhookResponse CreateWebhookResponse

// NewCreateWebhookResponse instantiates a new CreateWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebhookResponse(data WebhookResource) *CreateWebhookResponse {
	this := CreateWebhookResponse{}
	this.Data = data
	return &this
}

// NewCreateWebhookResponseWithDefaults instantiates a new CreateWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebhookResponseWithDefaults() *CreateWebhookResponse {
	this := CreateWebhookResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateWebhookResponse) GetData() WebhookResource {
	if o == nil {
		var ret WebhookResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateWebhookResponse) GetDataOk() (*WebhookResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateWebhookResponse) SetData(v WebhookResource) {
	o.Data = v
}

func (o CreateWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebhookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateWebhookResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateWebhookResponse := _CreateWebhookResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWebhookResponse)

	if err != nil {
		return err
	}

	*o = CreateWebhookResponse(varCreateWebhookResponse)

	return err
}

type NullableCreateWebhookResponse struct {
	value *CreateWebhookResponse
	isSet bool
}

func (v NullableCreateWebhookResponse) Get() *CreateWebhookResponse {
	return v.value
}

func (v *NullableCreateWebhookResponse) Set(val *CreateWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebhookResponse(val *CreateWebhookResponse) *NullableCreateWebhookResponse {
	return &NullableCreateWebhookResponse{value: val, isSet: true}
}

func (v NullableCreateWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
