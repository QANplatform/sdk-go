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

// checks if the InputEstimateGas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputEstimateGas{}

// InputEstimateGas struct for InputEstimateGas
type InputEstimateGas struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`
	// The state override set with address-to-state mapping where each address maps to an object containing
	Object *EstimateGasObject `json:"Object,omitempty"`
	// The transaction call object
	Transaction ParamsTransaction `json:"Transaction"`
}

type _InputEstimateGas InputEstimateGas

// NewInputEstimateGas instantiates a new InputEstimateGas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputEstimateGas(transaction ParamsTransaction) *InputEstimateGas {
	this := InputEstimateGas{}
	this.Transaction = transaction
	return &this
}

// NewInputEstimateGasWithDefaults instantiates a new InputEstimateGas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputEstimateGasWithDefaults() *InputEstimateGas {
	this := InputEstimateGas{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *InputEstimateGas) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEstimateGas) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *InputEstimateGas) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *InputEstimateGas) SetSchema(v string) {
	o.Schema = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *InputEstimateGas) GetObject() EstimateGasObject {
	if o == nil || IsNil(o.Object) {
		var ret EstimateGasObject
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputEstimateGas) GetObjectOk() (*EstimateGasObject, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *InputEstimateGas) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given EstimateGasObject and assigns it to the Object field.
func (o *InputEstimateGas) SetObject(v EstimateGasObject) {
	o.Object = &v
}

// GetTransaction returns the Transaction field value
func (o *InputEstimateGas) GetTransaction() ParamsTransaction {
	if o == nil {
		var ret ParamsTransaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *InputEstimateGas) GetTransactionOk() (*ParamsTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *InputEstimateGas) SetTransaction(v ParamsTransaction) {
	o.Transaction = v
}

func (o InputEstimateGas) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputEstimateGas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.Object) {
		toSerialize["Object"] = o.Object
	}
	toSerialize["Transaction"] = o.Transaction
	return toSerialize, nil
}

func (o *InputEstimateGas) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Transaction",
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

	varInputEstimateGas := _InputEstimateGas{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputEstimateGas)

	if err != nil {
		return err
	}

	*o = InputEstimateGas(varInputEstimateGas)

	return err
}

type NullableInputEstimateGas struct {
	value *InputEstimateGas
	isSet bool
}

func (v NullableInputEstimateGas) Get() *InputEstimateGas {
	return v.value
}

func (v *NullableInputEstimateGas) Set(val *InputEstimateGas) {
	v.value = val
	v.isSet = true
}

func (v NullableInputEstimateGas) IsSet() bool {
	return v.isSet
}

func (v *NullableInputEstimateGas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputEstimateGas(val *InputEstimateGas) *NullableInputEstimateGas {
	return &NullableInputEstimateGas{value: val, isSet: true}
}

func (v NullableInputEstimateGas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputEstimateGas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


