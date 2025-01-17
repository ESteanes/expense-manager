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

// checks if the CardPurchaseMethodObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardPurchaseMethodObject{}

// CardPurchaseMethodObject Provides information about the card used for a transaction.
type CardPurchaseMethodObject struct {
	// The type of card purchase.
	Method CardPurchaseMethodEnum `json:"method"`
	// The last four digits of the card used for the purchase, if applicable.
	CardNumberSuffix NullableString `json:"cardNumberSuffix"`
}

type _CardPurchaseMethodObject CardPurchaseMethodObject

// NewCardPurchaseMethodObject instantiates a new CardPurchaseMethodObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardPurchaseMethodObject(method CardPurchaseMethodEnum, cardNumberSuffix NullableString) *CardPurchaseMethodObject {
	this := CardPurchaseMethodObject{}
	this.Method = method
	this.CardNumberSuffix = cardNumberSuffix
	return &this
}

// NewCardPurchaseMethodObjectWithDefaults instantiates a new CardPurchaseMethodObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardPurchaseMethodObjectWithDefaults() *CardPurchaseMethodObject {
	this := CardPurchaseMethodObject{}
	return &this
}

// GetMethod returns the Method field value
func (o *CardPurchaseMethodObject) GetMethod() CardPurchaseMethodEnum {
	if o == nil {
		var ret CardPurchaseMethodEnum
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *CardPurchaseMethodObject) GetMethodOk() (*CardPurchaseMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *CardPurchaseMethodObject) SetMethod(v CardPurchaseMethodEnum) {
	o.Method = v
}

// GetCardNumberSuffix returns the CardNumberSuffix field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CardPurchaseMethodObject) GetCardNumberSuffix() string {
	if o == nil || o.CardNumberSuffix.Get() == nil {
		var ret string
		return ret
	}

	return *o.CardNumberSuffix.Get()
}

// GetCardNumberSuffixOk returns a tuple with the CardNumberSuffix field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardPurchaseMethodObject) GetCardNumberSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardNumberSuffix.Get(), o.CardNumberSuffix.IsSet()
}

// SetCardNumberSuffix sets field value
func (o *CardPurchaseMethodObject) SetCardNumberSuffix(v string) {
	o.CardNumberSuffix.Set(&v)
}

func (o CardPurchaseMethodObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardPurchaseMethodObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["method"] = o.Method
	toSerialize["cardNumberSuffix"] = o.CardNumberSuffix.Get()
	return toSerialize, nil
}

func (o *CardPurchaseMethodObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"cardNumberSuffix",
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

	varCardPurchaseMethodObject := _CardPurchaseMethodObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCardPurchaseMethodObject)

	if err != nil {
		return err
	}

	*o = CardPurchaseMethodObject(varCardPurchaseMethodObject)

	return err
}

type NullableCardPurchaseMethodObject struct {
	value *CardPurchaseMethodObject
	isSet bool
}

func (v NullableCardPurchaseMethodObject) Get() *CardPurchaseMethodObject {
	return v.value
}

func (v *NullableCardPurchaseMethodObject) Set(val *CardPurchaseMethodObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCardPurchaseMethodObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCardPurchaseMethodObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardPurchaseMethodObject(val *CardPurchaseMethodObject) *NullableCardPurchaseMethodObject {
	return &NullableCardPurchaseMethodObject{value: val, isSet: true}
}

func (v NullableCardPurchaseMethodObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardPurchaseMethodObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
