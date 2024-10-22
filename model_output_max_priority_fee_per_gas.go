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

// checks if the OutputMaxPriorityFeePerGas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputMaxPriorityFeePerGas{}

// OutputMaxPriorityFeePerGas struct for OutputMaxPriorityFeePerGas
type OutputMaxPriorityFeePerGas struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`
	// The hex value of the priority fee needed to be included in a block
	MaxPriorityFeePerGas string `json:"MaxPriorityFeePerGas"`
}

type _OutputMaxPriorityFeePerGas OutputMaxPriorityFeePerGas

// NewOutputMaxPriorityFeePerGas instantiates a new OutputMaxPriorityFeePerGas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputMaxPriorityFeePerGas(maxPriorityFeePerGas string) *OutputMaxPriorityFeePerGas {
	this := OutputMaxPriorityFeePerGas{}
	this.MaxPriorityFeePerGas = maxPriorityFeePerGas
	return &this
}

// NewOutputMaxPriorityFeePerGasWithDefaults instantiates a new OutputMaxPriorityFeePerGas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputMaxPriorityFeePerGasWithDefaults() *OutputMaxPriorityFeePerGas {
	this := OutputMaxPriorityFeePerGas{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *OutputMaxPriorityFeePerGas) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputMaxPriorityFeePerGas) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *OutputMaxPriorityFeePerGas) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *OutputMaxPriorityFeePerGas) SetSchema(v string) {
	o.Schema = &v
}

// GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field value
func (o *OutputMaxPriorityFeePerGas) GetMaxPriorityFeePerGas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxPriorityFeePerGas
}

// GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field value
// and a boolean to check if the value has been set.
func (o *OutputMaxPriorityFeePerGas) GetMaxPriorityFeePerGasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPriorityFeePerGas, true
}

// SetMaxPriorityFeePerGas sets field value
func (o *OutputMaxPriorityFeePerGas) SetMaxPriorityFeePerGas(v string) {
	o.MaxPriorityFeePerGas = v
}

func (o OutputMaxPriorityFeePerGas) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputMaxPriorityFeePerGas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["MaxPriorityFeePerGas"] = o.MaxPriorityFeePerGas
	return toSerialize, nil
}

func (o *OutputMaxPriorityFeePerGas) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"MaxPriorityFeePerGas",
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

	varOutputMaxPriorityFeePerGas := _OutputMaxPriorityFeePerGas{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutputMaxPriorityFeePerGas)

	if err != nil {
		return err
	}

	*o = OutputMaxPriorityFeePerGas(varOutputMaxPriorityFeePerGas)

	return err
}

type NullableOutputMaxPriorityFeePerGas struct {
	value *OutputMaxPriorityFeePerGas
	isSet bool
}

func (v NullableOutputMaxPriorityFeePerGas) Get() *OutputMaxPriorityFeePerGas {
	return v.value
}

func (v *NullableOutputMaxPriorityFeePerGas) Set(val *OutputMaxPriorityFeePerGas) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputMaxPriorityFeePerGas) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputMaxPriorityFeePerGas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputMaxPriorityFeePerGas(val *OutputMaxPriorityFeePerGas) *NullableOutputMaxPriorityFeePerGas {
	return &NullableOutputMaxPriorityFeePerGas{value: val, isSet: true}
}

func (v NullableOutputMaxPriorityFeePerGas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputMaxPriorityFeePerGas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

