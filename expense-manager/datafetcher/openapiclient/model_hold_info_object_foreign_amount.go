/*
Up API

The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the HoldInfoObjectForeignAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HoldInfoObjectForeignAmount{}

// HoldInfoObjectForeignAmount The foreign currency amount of this transaction while in the `HELD` status. This field will be `null` for domestic transactions. The amount was converted to the AUD amount reflected in the `amount` field.
type HoldInfoObjectForeignAmount struct {
	// The ISO 4217 currency code.
	CurrencyCode string `json:"currencyCode"`
	// The amount of money, formatted as a string in the relevant currency. For example, for an Australian dollar value of $10.56, this field will be `\"10.56\"`. The currency symbol is not included in the string.
	Value string `json:"value"`
	// The amount of money in the smallest denomination for the currency, as a 64-bit integer.  For example, for an Australian dollar value of $10.56, this field will be `1056`.
	ValueInBaseUnits int32 `json:"valueInBaseUnits"`
}

type _HoldInfoObjectForeignAmount HoldInfoObjectForeignAmount

// NewHoldInfoObjectForeignAmount instantiates a new HoldInfoObjectForeignAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoldInfoObjectForeignAmount(currencyCode string, value string, valueInBaseUnits int32) *HoldInfoObjectForeignAmount {
	this := HoldInfoObjectForeignAmount{}
	this.CurrencyCode = currencyCode
	this.Value = value
	this.ValueInBaseUnits = valueInBaseUnits
	return &this
}

// NewHoldInfoObjectForeignAmountWithDefaults instantiates a new HoldInfoObjectForeignAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldInfoObjectForeignAmountWithDefaults() *HoldInfoObjectForeignAmount {
	this := HoldInfoObjectForeignAmount{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *HoldInfoObjectForeignAmount) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *HoldInfoObjectForeignAmount) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *HoldInfoObjectForeignAmount) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetValue returns the Value field value
func (o *HoldInfoObjectForeignAmount) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HoldInfoObjectForeignAmount) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *HoldInfoObjectForeignAmount) SetValue(v string) {
	o.Value = v
}

// GetValueInBaseUnits returns the ValueInBaseUnits field value
func (o *HoldInfoObjectForeignAmount) GetValueInBaseUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValueInBaseUnits
}

// GetValueInBaseUnitsOk returns a tuple with the ValueInBaseUnits field value
// and a boolean to check if the value has been set.
func (o *HoldInfoObjectForeignAmount) GetValueInBaseUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueInBaseUnits, true
}

// SetValueInBaseUnits sets field value
func (o *HoldInfoObjectForeignAmount) SetValueInBaseUnits(v int32) {
	o.ValueInBaseUnits = v
}

func (o HoldInfoObjectForeignAmount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HoldInfoObjectForeignAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currencyCode"] = o.CurrencyCode
	toSerialize["value"] = o.Value
	toSerialize["valueInBaseUnits"] = o.ValueInBaseUnits
	return toSerialize, nil
}

func (o *HoldInfoObjectForeignAmount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currencyCode",
		"value",
		"valueInBaseUnits",
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

	varHoldInfoObjectForeignAmount := _HoldInfoObjectForeignAmount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHoldInfoObjectForeignAmount)

	if err != nil {
		return err
	}

	*o = HoldInfoObjectForeignAmount(varHoldInfoObjectForeignAmount)

	return err
}

type NullableHoldInfoObjectForeignAmount struct {
	value *HoldInfoObjectForeignAmount
	isSet bool
}

func (v NullableHoldInfoObjectForeignAmount) Get() *HoldInfoObjectForeignAmount {
	return v.value
}

func (v *NullableHoldInfoObjectForeignAmount) Set(val *HoldInfoObjectForeignAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldInfoObjectForeignAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldInfoObjectForeignAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldInfoObjectForeignAmount(val *HoldInfoObjectForeignAmount) *NullableHoldInfoObjectForeignAmount {
	return &NullableHoldInfoObjectForeignAmount{value: val, isSet: true}
}

func (v NullableHoldInfoObjectForeignAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldInfoObjectForeignAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
