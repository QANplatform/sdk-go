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

// checks if the OutputGetTransactionByHash type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputGetTransactionByHash{}

// OutputGetTransactionByHash struct for OutputGetTransactionByHash
type OutputGetTransactionByHash struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`
	// The transaction response object, or null if no transaction is found
	Transaction ResponseTransaction `json:"Transaction"`
}

type _OutputGetTransactionByHash OutputGetTransactionByHash

// NewOutputGetTransactionByHash instantiates a new OutputGetTransactionByHash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputGetTransactionByHash(transaction ResponseTransaction) *OutputGetTransactionByHash {
	this := OutputGetTransactionByHash{}
	this.Transaction = transaction
	return &this
}

// NewOutputGetTransactionByHashWithDefaults instantiates a new OutputGetTransactionByHash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputGetTransactionByHashWithDefaults() *OutputGetTransactionByHash {
	this := OutputGetTransactionByHash{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *OutputGetTransactionByHash) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputGetTransactionByHash) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *OutputGetTransactionByHash) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *OutputGetTransactionByHash) SetSchema(v string) {
	o.Schema = &v
}

// GetTransaction returns the Transaction field value
func (o *OutputGetTransactionByHash) GetTransaction() ResponseTransaction {
	if o == nil {
		var ret ResponseTransaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *OutputGetTransactionByHash) GetTransactionOk() (*ResponseTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *OutputGetTransactionByHash) SetTransaction(v ResponseTransaction) {
	o.Transaction = v
}

func (o OutputGetTransactionByHash) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputGetTransactionByHash) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["Transaction"] = o.Transaction
	return toSerialize, nil
}

func (o *OutputGetTransactionByHash) UnmarshalJSON(data []byte) (err error) {
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

	varOutputGetTransactionByHash := _OutputGetTransactionByHash{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutputGetTransactionByHash)

	if err != nil {
		return err
	}

	*o = OutputGetTransactionByHash(varOutputGetTransactionByHash)

	return err
}

type NullableOutputGetTransactionByHash struct {
	value *OutputGetTransactionByHash
	isSet bool
}

func (v NullableOutputGetTransactionByHash) Get() *OutputGetTransactionByHash {
	return v.value
}

func (v *NullableOutputGetTransactionByHash) Set(val *OutputGetTransactionByHash) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputGetTransactionByHash) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputGetTransactionByHash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputGetTransactionByHash(val *OutputGetTransactionByHash) *NullableOutputGetTransactionByHash {
	return &NullableOutputGetTransactionByHash{value: val, isSet: true}
}

func (v NullableOutputGetTransactionByHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputGetTransactionByHash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


