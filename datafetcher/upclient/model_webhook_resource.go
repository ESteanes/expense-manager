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

// checks if the WebhookResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookResource{}

// WebhookResource Provides information about a webhook. 
type WebhookResource struct {
	// The type of this resource: `webhooks`
	Type string `json:"type"`
	// The unique identifier for this webhook. 
	Id string `json:"id"`
	Attributes WebhookResourceAttributes `json:"attributes"`
	Relationships WebhookResourceRelationships `json:"relationships"`
	Links *AccountResourceLinks `json:"links,omitempty"`
}

type _WebhookResource WebhookResource

// NewWebhookResource instantiates a new WebhookResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResource(type_ string, id string, attributes WebhookResourceAttributes, relationships WebhookResourceRelationships) *WebhookResource {
	this := WebhookResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewWebhookResourceWithDefaults instantiates a new WebhookResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResourceWithDefaults() *WebhookResource {
	this := WebhookResource{}
	return &this
}

// GetType returns the Type field value
func (o *WebhookResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhookResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebhookResource) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *WebhookResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *WebhookResource) GetAttributes() WebhookResourceAttributes {
	if o == nil {
		var ret WebhookResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WebhookResource) GetAttributesOk() (*WebhookResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *WebhookResource) SetAttributes(v WebhookResourceAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *WebhookResource) GetRelationships() WebhookResourceRelationships {
	if o == nil {
		var ret WebhookResourceRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *WebhookResource) GetRelationshipsOk() (*WebhookResourceRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *WebhookResource) SetRelationships(v WebhookResourceRelationships) {
	o.Relationships = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WebhookResource) GetLinks() AccountResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResource) GetLinksOk() (*AccountResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WebhookResource) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceLinks and assigns it to the Links field.
func (o *WebhookResource) SetLinks(v AccountResourceLinks) {
	o.Links = &v
}

func (o WebhookResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *WebhookResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"attributes",
		"relationships",
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

	varWebhookResource := _WebhookResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookResource)

	if err != nil {
		return err
	}

	*o = WebhookResource(varWebhookResource)

	return err
}

type NullableWebhookResource struct {
	value *WebhookResource
	isSet bool
}

func (v NullableWebhookResource) Get() *WebhookResource {
	return v.value
}

func (v *NullableWebhookResource) Set(val *WebhookResource) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResource) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResource(val *WebhookResource) *NullableWebhookResource {
	return &NullableWebhookResource{value: val, isSet: true}
}

func (v NullableWebhookResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


