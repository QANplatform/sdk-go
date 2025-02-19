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

// checks if the OutputGetTransactionByBlockNumberAndIndex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputGetTransactionByBlockNumberAndIndex{}

// OutputGetTransactionByBlockNumberAndIndex struct for OutputGetTransactionByBlockNumberAndIndex
type OutputGetTransactionByBlockNumberAndIndex struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`
	// The transaction response object, or null if no transaction is found
	Transaction ResponseTransaction `json:"Transaction"`
}

type _OutputGetTransactionByBlockNumberAndIndex OutputGetTransactionByBlockNumberAndIndex

// NewOutputGetTransactionByBlockNumberAndIndex instantiates a new OutputGetTransactionByBlockNumberAndIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputGetTransactionByBlockNumberAndIndex(transaction ResponseTransaction) *OutputGetTransactionByBlockNumberAndIndex {
	this := OutputGetTransactionByBlockNumberAndIndex{}
	this.Transaction = transaction
	return &this
}

// NewOutputGetTransactionByBlockNumberAndIndexWithDefaults instantiates a new OutputGetTransactionByBlockNumberAndIndex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputGetTransactionByBlockNumberAndIndexWithDefaults() *OutputGetTransactionByBlockNumberAndIndex {
	this := OutputGetTransactionByBlockNumberAndIndex{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *OutputGetTransactionByBlockNumberAndIndex) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputGetTransactionByBlockNumberAndIndex) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *OutputGetTransactionByBlockNumberAndIndex) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *OutputGetTransactionByBlockNumberAndIndex) SetSchema(v string) {
	o.Schema = &v
}

// GetTransaction returns the Transaction field value
func (o *OutputGetTransactionByBlockNumberAndIndex) GetTransaction() ResponseTransaction {
	if o == nil {
		var ret ResponseTransaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *OutputGetTransactionByBlockNumberAndIndex) GetTransactionOk() (*ResponseTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *OutputGetTransactionByBlockNumberAndIndex) SetTransaction(v ResponseTransaction) {
	o.Transaction = v
}

func (o OutputGetTransactionByBlockNumberAndIndex) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputGetTransactionByBlockNumberAndIndex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["Transaction"] = o.Transaction
	return toSerialize, nil
}

func (o *OutputGetTransactionByBlockNumberAndIndex) UnmarshalJSON(data []byte) (err error) {
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

	varOutputGetTransactionByBlockNumberAndIndex := _OutputGetTransactionByBlockNumberAndIndex{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutputGetTransactionByBlockNumberAndIndex)

	if err != nil {
		return err
	}

	*o = OutputGetTransactionByBlockNumberAndIndex(varOutputGetTransactionByBlockNumberAndIndex)

	return err
}

type NullableOutputGetTransactionByBlockNumberAndIndex struct {
	value *OutputGetTransactionByBlockNumberAndIndex
	isSet bool
}

func (v NullableOutputGetTransactionByBlockNumberAndIndex) Get() *OutputGetTransactionByBlockNumberAndIndex {
	return v.value
}

func (v *NullableOutputGetTransactionByBlockNumberAndIndex) Set(val *OutputGetTransactionByBlockNumberAndIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputGetTransactionByBlockNumberAndIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputGetTransactionByBlockNumberAndIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputGetTransactionByBlockNumberAndIndex(val *OutputGetTransactionByBlockNumberAndIndex) *NullableOutputGetTransactionByBlockNumberAndIndex {
	return &NullableOutputGetTransactionByBlockNumberAndIndex{value: val, isSet: true}
}

func (v NullableOutputGetTransactionByBlockNumberAndIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputGetTransactionByBlockNumberAndIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


