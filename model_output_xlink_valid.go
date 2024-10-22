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

// checks if the OutputXlinkValid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputXlinkValid{}

// OutputXlinkValid struct for OutputXlinkValid
type OutputXlinkValid struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`
	// The xlink validity time of the specified address.
	Validity string `json:"Validity"`
}

type _OutputXlinkValid OutputXlinkValid

// NewOutputXlinkValid instantiates a new OutputXlinkValid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputXlinkValid(validity string) *OutputXlinkValid {
	this := OutputXlinkValid{}
	this.Validity = validity
	return &this
}

// NewOutputXlinkValidWithDefaults instantiates a new OutputXlinkValid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputXlinkValidWithDefaults() *OutputXlinkValid {
	this := OutputXlinkValid{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *OutputXlinkValid) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputXlinkValid) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *OutputXlinkValid) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *OutputXlinkValid) SetSchema(v string) {
	o.Schema = &v
}

// GetValidity returns the Validity field value
func (o *OutputXlinkValid) GetValidity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Validity
}

// GetValidityOk returns a tuple with the Validity field value
// and a boolean to check if the value has been set.
func (o *OutputXlinkValid) GetValidityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Validity, true
}

// SetValidity sets field value
func (o *OutputXlinkValid) SetValidity(v string) {
	o.Validity = v
}

func (o OutputXlinkValid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputXlinkValid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["Validity"] = o.Validity
	return toSerialize, nil
}

func (o *OutputXlinkValid) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Validity",
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

	varOutputXlinkValid := _OutputXlinkValid{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutputXlinkValid)

	if err != nil {
		return err
	}

	*o = OutputXlinkValid(varOutputXlinkValid)

	return err
}

type NullableOutputXlinkValid struct {
	value *OutputXlinkValid
	isSet bool
}

func (v NullableOutputXlinkValid) Get() *OutputXlinkValid {
	return v.value
}

func (v *NullableOutputXlinkValid) Set(val *OutputXlinkValid) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputXlinkValid) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputXlinkValid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputXlinkValid(val *OutputXlinkValid) *NullableOutputXlinkValid {
	return &NullableOutputXlinkValid{value: val, isSet: true}
}

func (v NullableOutputXlinkValid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputXlinkValid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


