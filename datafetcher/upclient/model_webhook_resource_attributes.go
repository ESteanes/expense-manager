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
	"time"
)

// checks if the WebhookResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookResourceAttributes{}

// WebhookResourceAttributes struct for WebhookResourceAttributes
type WebhookResourceAttributes struct {
	// The URL that this webhook is configured to `POST` events to.
	Url string `json:"url"`
	// An optional description that was provided at the time the webhook was created.
	Description NullableString `json:"description"`
	// A shared secret key used to sign all webhook events sent to the configured webhook URL. This field is returned only once, upon the initial creation of the webhook. If lost, create a new webhook and delete this webhook.  The webhook URL receives a request with a `X-Up-Authenticity-Signature` header, which is the SHA-256 HMAC of the entire raw request body signed using this `secretKey`. It is advised to compute and check this signature to verify the authenticity of requests sent to the webhook URL. See [Handling webhook events](#callback_post_webhookURL) for full details.
	SecretKey *string `json:"secretKey,omitempty"`
	// The date-time at which this webhook was created.
	CreatedAt time.Time `json:"createdAt"`
}

type _WebhookResourceAttributes WebhookResourceAttributes

// NewWebhookResourceAttributes instantiates a new WebhookResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResourceAttributes(url string, description NullableString, createdAt time.Time) *WebhookResourceAttributes {
	this := WebhookResourceAttributes{}
	this.Url = url
	this.Description = description
	this.CreatedAt = createdAt
	return &this
}

// NewWebhookResourceAttributesWithDefaults instantiates a new WebhookResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResourceAttributesWithDefaults() *WebhookResourceAttributes {
	this := WebhookResourceAttributes{}
	return &this
}

// GetUrl returns the Url field value
func (o *WebhookResourceAttributes) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookResourceAttributes) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookResourceAttributes) SetUrl(v string) {
	o.Url = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookResourceAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookResourceAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *WebhookResourceAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *WebhookResourceAttributes) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResourceAttributes) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *WebhookResourceAttributes) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *WebhookResourceAttributes) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WebhookResourceAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookResourceAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WebhookResourceAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o WebhookResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["description"] = o.Description.Get()
	if !IsNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	toSerialize["createdAt"] = o.CreatedAt
	return toSerialize, nil
}

func (o *WebhookResourceAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"description",
		"createdAt",
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

	varWebhookResourceAttributes := _WebhookResourceAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookResourceAttributes)

	if err != nil {
		return err
	}

	*o = WebhookResourceAttributes(varWebhookResourceAttributes)

	return err
}

type NullableWebhookResourceAttributes struct {
	value *WebhookResourceAttributes
	isSet bool
}

func (v NullableWebhookResourceAttributes) Get() *WebhookResourceAttributes {
	return v.value
}

func (v *NullableWebhookResourceAttributes) Set(val *WebhookResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResourceAttributes(val *WebhookResourceAttributes) *NullableWebhookResourceAttributes {
	return &NullableWebhookResourceAttributes{value: val, isSet: true}
}

func (v NullableWebhookResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
