/*
QAN AutoApi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qan

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OutputGetBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputGetBalance{}

// OutputGetBalance struct for OutputGetBalance
type OutputGetBalance struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`
	// The balance of the specified address in decimal value, measured in wei
	Balance string `json:"Balance"`
}

type _OutputGetBalance OutputGetBalance

// NewOutputGetBalance instantiates a new OutputGetBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputGetBalance(balance string) *OutputGetBalance {
	this := OutputGetBalance{}
	this.Balance = balance
	return &this
}

// NewOutputGetBalanceWithDefaults instantiates a new OutputGetBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputGetBalanceWithDefaults() *OutputGetBalance {
	this := OutputGetBalance{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *OutputGetBalance) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputGetBalance) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *OutputGetBalance) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *OutputGetBalance) SetSchema(v string) {
	o.Schema = &v
}

// GetBalance returns the Balance field value
func (o *OutputGetBalance) GetBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *OutputGetBalance) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *OutputGetBalance) SetBalance(v string) {
	o.Balance = v
}

func (o OutputGetBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputGetBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["Balance"] = o.Balance
	return toSerialize, nil
}

func (o *OutputGetBalance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Balance",
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

	varOutputGetBalance := _OutputGetBalance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutputGetBalance)

	if err != nil {
		return err
	}

	*o = OutputGetBalance(varOutputGetBalance)

	return err
}

type NullableOutputGetBalance struct {
	value *OutputGetBalance
	isSet bool
}

func (v NullableOutputGetBalance) Get() *OutputGetBalance {
	return v.value
}

func (v *NullableOutputGetBalance) Set(val *OutputGetBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputGetBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputGetBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputGetBalance(val *OutputGetBalance) *NullableOutputGetBalance {
	return &NullableOutputGetBalance{value: val, isSet: true}
}

func (v NullableOutputGetBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputGetBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


