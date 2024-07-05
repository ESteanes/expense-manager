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

// checks if the WebhookDeliveryLogResourceAttributesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookDeliveryLogResourceAttributesResponse{}

// WebhookDeliveryLogResourceAttributesResponse Information about the response that was received from the webhook URL.
type WebhookDeliveryLogResourceAttributesResponse struct {
	// The HTTP status code received in the response.
	StatusCode int32 `json:"statusCode"`
	// The payload that was received in the response body.
	Body string `json:"body"`
}

type _WebhookDeliveryLogResourceAttributesResponse WebhookDeliveryLogResourceAttributesResponse

// NewWebhookDeliveryLogResourceAttributesResponse instantiates a new WebhookDeliveryLogResourceAttributesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookDeliveryLogResourceAttributesResponse(statusCode int32, body string) *WebhookDeliveryLogResourceAttributesResponse {
	this := WebhookDeliveryLogResourceAttributesResponse{}
	this.StatusCode = statusCode
	this.Body = body
	return &this
}

// NewWebhookDeliveryLogResourceAttributesResponseWithDefaults instantiates a new WebhookDeliveryLogResourceAttributesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookDeliveryLogResourceAttributesResponseWithDefaults() *WebhookDeliveryLogResourceAttributesResponse {
	this := WebhookDeliveryLogResourceAttributesResponse{}
	return &this
}

// GetStatusCode returns the StatusCode field value
func (o *WebhookDeliveryLogResourceAttributesResponse) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *WebhookDeliveryLogResourceAttributesResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *WebhookDeliveryLogResourceAttributesResponse) SetStatusCode(v int32) {
	o.StatusCode = v
}

// GetBody returns the Body field value
func (o *WebhookDeliveryLogResourceAttributesResponse) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *WebhookDeliveryLogResourceAttributesResponse) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *WebhookDeliveryLogResourceAttributesResponse) SetBody(v string) {
	o.Body = v
}

func (o WebhookDeliveryLogResourceAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookDeliveryLogResourceAttributesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statusCode"] = o.StatusCode
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

func (o *WebhookDeliveryLogResourceAttributesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statusCode",
		"body",
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

	varWebhookDeliveryLogResourceAttributesResponse := _WebhookDeliveryLogResourceAttributesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookDeliveryLogResourceAttributesResponse)

	if err != nil {
		return err
	}

	*o = WebhookDeliveryLogResourceAttributesResponse(varWebhookDeliveryLogResourceAttributesResponse)

	return err
}

type NullableWebhookDeliveryLogResourceAttributesResponse struct {
	value *WebhookDeliveryLogResourceAttributesResponse
	isSet bool
}

func (v NullableWebhookDeliveryLogResourceAttributesResponse) Get() *WebhookDeliveryLogResourceAttributesResponse {
	return v.value
}

func (v *NullableWebhookDeliveryLogResourceAttributesResponse) Set(val *WebhookDeliveryLogResourceAttributesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDeliveryLogResourceAttributesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDeliveryLogResourceAttributesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDeliveryLogResourceAttributesResponse(val *WebhookDeliveryLogResourceAttributesResponse) *NullableWebhookDeliveryLogResourceAttributesResponse {
	return &NullableWebhookDeliveryLogResourceAttributesResponse{value: val, isSet: true}
}

func (v NullableWebhookDeliveryLogResourceAttributesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDeliveryLogResourceAttributesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
