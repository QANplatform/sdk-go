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

// checks if the OutputChainId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputChainId{}

// OutputChainId struct for OutputChainId
type OutputChainId struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`
	// A decimal value in string format which represents an integer of the chain ID
	ChainId string `json:"ChainId"`
}

type _OutputChainId OutputChainId

// NewOutputChainId instantiates a new OutputChainId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputChainId(chainId string) *OutputChainId {
	this := OutputChainId{}
	this.ChainId = chainId
	return &this
}

// NewOutputChainIdWithDefaults instantiates a new OutputChainId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputChainIdWithDefaults() *OutputChainId {
	this := OutputChainId{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *OutputChainId) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputChainId) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *OutputChainId) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *OutputChainId) SetSchema(v string) {
	o.Schema = &v
}

// GetChainId returns the ChainId field value
func (o *OutputChainId) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *OutputChainId) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *OutputChainId) SetChainId(v string) {
	o.ChainId = v
}

func (o OutputChainId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputChainId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["ChainId"] = o.ChainId
	return toSerialize, nil
}

func (o *OutputChainId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ChainId",
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

	varOutputChainId := _OutputChainId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutputChainId)

	if err != nil {
		return err
	}

	*o = OutputChainId(varOutputChainId)

	return err
}

type NullableOutputChainId struct {
	value *OutputChainId
	isSet bool
}

func (v NullableOutputChainId) Get() *OutputChainId {
	return v.value
}

func (v *NullableOutputChainId) Set(val *OutputChainId) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputChainId) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputChainId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputChainId(val *OutputChainId) *NullableOutputChainId {
	return &NullableOutputChainId{value: val, isSet: true}
}

func (v NullableOutputChainId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputChainId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

