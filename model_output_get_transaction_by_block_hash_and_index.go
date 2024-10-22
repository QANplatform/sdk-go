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

// checks if the OutputGetTransactionByBlockHashAndIndex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputGetTransactionByBlockHashAndIndex{}

// OutputGetTransactionByBlockHashAndIndex struct for OutputGetTransactionByBlockHashAndIndex
type OutputGetTransactionByBlockHashAndIndex struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`
	// The transaction response object, or null if no transaction is found
	Transaction ResponseTransaction `json:"Transaction"`
}

type _OutputGetTransactionByBlockHashAndIndex OutputGetTransactionByBlockHashAndIndex

// NewOutputGetTransactionByBlockHashAndIndex instantiates a new OutputGetTransactionByBlockHashAndIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputGetTransactionByBlockHashAndIndex(transaction ResponseTransaction) *OutputGetTransactionByBlockHashAndIndex {
	this := OutputGetTransactionByBlockHashAndIndex{}
	this.Transaction = transaction
	return &this
}

// NewOutputGetTransactionByBlockHashAndIndexWithDefaults instantiates a new OutputGetTransactionByBlockHashAndIndex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputGetTransactionByBlockHashAndIndexWithDefaults() *OutputGetTransactionByBlockHashAndIndex {
	this := OutputGetTransactionByBlockHashAndIndex{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *OutputGetTransactionByBlockHashAndIndex) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputGetTransactionByBlockHashAndIndex) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *OutputGetTransactionByBlockHashAndIndex) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *OutputGetTransactionByBlockHashAndIndex) SetSchema(v string) {
	o.Schema = &v
}

// GetTransaction returns the Transaction field value
func (o *OutputGetTransactionByBlockHashAndIndex) GetTransaction() ResponseTransaction {
	if o == nil {
		var ret ResponseTransaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *OutputGetTransactionByBlockHashAndIndex) GetTransactionOk() (*ResponseTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *OutputGetTransactionByBlockHashAndIndex) SetTransaction(v ResponseTransaction) {
	o.Transaction = v
}

func (o OutputGetTransactionByBlockHashAndIndex) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputGetTransactionByBlockHashAndIndex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["Transaction"] = o.Transaction
	return toSerialize, nil
}

func (o *OutputGetTransactionByBlockHashAndIndex) UnmarshalJSON(data []byte) (err error) {
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

	varOutputGetTransactionByBlockHashAndIndex := _OutputGetTransactionByBlockHashAndIndex{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutputGetTransactionByBlockHashAndIndex)

	if err != nil {
		return err
	}

	*o = OutputGetTransactionByBlockHashAndIndex(varOutputGetTransactionByBlockHashAndIndex)

	return err
}

type NullableOutputGetTransactionByBlockHashAndIndex struct {
	value *OutputGetTransactionByBlockHashAndIndex
	isSet bool
}

func (v NullableOutputGetTransactionByBlockHashAndIndex) Get() *OutputGetTransactionByBlockHashAndIndex {
	return v.value
}

func (v *NullableOutputGetTransactionByBlockHashAndIndex) Set(val *OutputGetTransactionByBlockHashAndIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputGetTransactionByBlockHashAndIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputGetTransactionByBlockHashAndIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputGetTransactionByBlockHashAndIndex(val *OutputGetTransactionByBlockHashAndIndex) *NullableOutputGetTransactionByBlockHashAndIndex {
	return &NullableOutputGetTransactionByBlockHashAndIndex{value: val, isSet: true}
}

func (v NullableOutputGetTransactionByBlockHashAndIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputGetTransactionByBlockHashAndIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


