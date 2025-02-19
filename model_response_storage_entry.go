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

// checks if the ResponseStorageEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseStorageEntry{}

// ResponseStorageEntry struct for ResponseStorageEntry
type ResponseStorageEntry struct {
	// The requested storage key
	Key string `json:"Key"`
	// An array of rlp-serialized MerkleTree-Nodes which starts with the stateRoot-Node and follows the path of the SHA3 address as key
	Proof string `json:"Proof"`
	// The storage value
	Value string `json:"Value"`
}

type _ResponseStorageEntry ResponseStorageEntry

// NewResponseStorageEntry instantiates a new ResponseStorageEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseStorageEntry(key string, proof string, value string) *ResponseStorageEntry {
	this := ResponseStorageEntry{}
	this.Key = key
	this.Proof = proof
	this.Value = value
	return &this
}

// NewResponseStorageEntryWithDefaults instantiates a new ResponseStorageEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseStorageEntryWithDefaults() *ResponseStorageEntry {
	this := ResponseStorageEntry{}
	return &this
}

// GetKey returns the Key field value
func (o *ResponseStorageEntry) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ResponseStorageEntry) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ResponseStorageEntry) SetKey(v string) {
	o.Key = v
}

// GetProof returns the Proof field value
func (o *ResponseStorageEntry) GetProof() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proof
}

// GetProofOk returns a tuple with the Proof field value
// and a boolean to check if the value has been set.
func (o *ResponseStorageEntry) GetProofOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proof, true
}

// SetProof sets field value
func (o *ResponseStorageEntry) SetProof(v string) {
	o.Proof = v
}

// GetValue returns the Value field value
func (o *ResponseStorageEntry) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ResponseStorageEntry) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ResponseStorageEntry) SetValue(v string) {
	o.Value = v
}

func (o ResponseStorageEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseStorageEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Key"] = o.Key
	toSerialize["Proof"] = o.Proof
	toSerialize["Value"] = o.Value
	return toSerialize, nil
}

func (o *ResponseStorageEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Key",
		"Proof",
		"Value",
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

	varResponseStorageEntry := _ResponseStorageEntry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseStorageEntry)

	if err != nil {
		return err
	}

	*o = ResponseStorageEntry(varResponseStorageEntry)

	return err
}

type NullableResponseStorageEntry struct {
	value *ResponseStorageEntry
	isSet bool
}

func (v NullableResponseStorageEntry) Get() *ResponseStorageEntry {
	return v.value
}

func (v *NullableResponseStorageEntry) Set(val *ResponseStorageEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseStorageEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseStorageEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseStorageEntry(val *ResponseStorageEntry) *NullableResponseStorageEntry {
	return &NullableResponseStorageEntry{value: val, isSet: true}
}

func (v NullableResponseStorageEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseStorageEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


