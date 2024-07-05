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

// checks if the HoldInfoObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HoldInfoObject{}

// HoldInfoObject Provides information about the amount at which a transaction was in the `HELD` status.
type HoldInfoObject struct {
	// The amount of this transaction while in the `HELD` status, in Australian dollars.
	Amount        MoneyObject                         `json:"amount"`
	ForeignAmount NullableHoldInfoObjectForeignAmount `json:"foreignAmount"`
}

type _HoldInfoObject HoldInfoObject

// NewHoldInfoObject instantiates a new HoldInfoObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoldInfoObject(amount MoneyObject, foreignAmount NullableHoldInfoObjectForeignAmount) *HoldInfoObject {
	this := HoldInfoObject{}
	this.Amount = amount
	this.ForeignAmount = foreignAmount
	return &this
}

// NewHoldInfoObjectWithDefaults instantiates a new HoldInfoObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldInfoObjectWithDefaults() *HoldInfoObject {
	this := HoldInfoObject{}
	return &this
}

// GetAmount returns the Amount field value
func (o *HoldInfoObject) GetAmount() MoneyObject {
	if o == nil {
		var ret MoneyObject
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *HoldInfoObject) GetAmountOk() (*MoneyObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *HoldInfoObject) SetAmount(v MoneyObject) {
	o.Amount = v
}

// GetForeignAmount returns the ForeignAmount field value
// If the value is explicit nil, the zero value for HoldInfoObjectForeignAmount will be returned
func (o *HoldInfoObject) GetForeignAmount() HoldInfoObjectForeignAmount {
	if o == nil || o.ForeignAmount.Get() == nil {
		var ret HoldInfoObjectForeignAmount
		return ret
	}

	return *o.ForeignAmount.Get()
}

// GetForeignAmountOk returns a tuple with the ForeignAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HoldInfoObject) GetForeignAmountOk() (*HoldInfoObjectForeignAmount, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForeignAmount.Get(), o.ForeignAmount.IsSet()
}

// SetForeignAmount sets field value
func (o *HoldInfoObject) SetForeignAmount(v HoldInfoObjectForeignAmount) {
	o.ForeignAmount.Set(&v)
}

func (o HoldInfoObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HoldInfoObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["foreignAmount"] = o.ForeignAmount.Get()
	return toSerialize, nil
}

func (o *HoldInfoObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"foreignAmount",
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

	varHoldInfoObject := _HoldInfoObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHoldInfoObject)

	if err != nil {
		return err
	}

	*o = HoldInfoObject(varHoldInfoObject)

	return err
}

type NullableHoldInfoObject struct {
	value *HoldInfoObject
	isSet bool
}

func (v NullableHoldInfoObject) Get() *HoldInfoObject {
	return v.value
}

func (v *NullableHoldInfoObject) Set(val *HoldInfoObject) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldInfoObject) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldInfoObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldInfoObject(val *HoldInfoObject) *NullableHoldInfoObject {
	return &NullableHoldInfoObject{value: val, isSet: true}
}

func (v NullableHoldInfoObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldInfoObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
