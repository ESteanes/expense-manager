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

// checks if the AttachmentResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentResource{}

// AttachmentResource struct for AttachmentResource
type AttachmentResource struct {
	// The type of this resource: `attachments`
	Type string `json:"type"`
	// The unique identifier for this attachment.
	Id            string                          `json:"id"`
	Attributes    AttachmentResourceAttributes    `json:"attributes"`
	Relationships AttachmentResourceRelationships `json:"relationships"`
	Links         *AccountResourceLinks           `json:"links,omitempty"`
}

type _AttachmentResource AttachmentResource

// NewAttachmentResource instantiates a new AttachmentResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentResource(type_ string, id string, attributes AttachmentResourceAttributes, relationships AttachmentResourceRelationships) *AttachmentResource {
	this := AttachmentResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAttachmentResourceWithDefaults instantiates a new AttachmentResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentResourceWithDefaults() *AttachmentResource {
	this := AttachmentResource{}
	return &this
}

// GetType returns the Type field value
func (o *AttachmentResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttachmentResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AttachmentResource) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AttachmentResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachmentResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttachmentResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *AttachmentResource) GetAttributes() AttachmentResourceAttributes {
	if o == nil {
		var ret AttachmentResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AttachmentResource) GetAttributesOk() (*AttachmentResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AttachmentResource) SetAttributes(v AttachmentResourceAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AttachmentResource) GetRelationships() AttachmentResourceRelationships {
	if o == nil {
		var ret AttachmentResourceRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AttachmentResource) GetRelationshipsOk() (*AttachmentResourceRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AttachmentResource) SetRelationships(v AttachmentResourceRelationships) {
	o.Relationships = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AttachmentResource) GetLinks() AccountResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentResource) GetLinksOk() (*AccountResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AttachmentResource) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountResourceLinks and assigns it to the Links field.
func (o *AttachmentResource) SetLinks(v AccountResourceLinks) {
	o.Links = &v
}

func (o AttachmentResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentResource) ToMap() (map[string]interface{}, error) {
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

func (o *AttachmentResource) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAttachmentResource := _AttachmentResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttachmentResource)

	if err != nil {
		return err
	}

	*o = AttachmentResource(varAttachmentResource)

	return err
}

type NullableAttachmentResource struct {
	value *AttachmentResource
	isSet bool
}

func (v NullableAttachmentResource) Get() *AttachmentResource {
	return v.value
}

func (v *NullableAttachmentResource) Set(val *AttachmentResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentResource(val *AttachmentResource) *NullableAttachmentResource {
	return &NullableAttachmentResource{value: val, isSet: true}
}

func (v NullableAttachmentResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
