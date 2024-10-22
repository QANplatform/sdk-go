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

// checks if the FilterObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterObject{}

// FilterObject struct for FilterObject
type FilterObject struct {
	// The contract address or a list of addresses from which logs should originate
	Address string `json:"Address"`
	FromBlock string `json:"FromBlock"`
	ToBlock string `json:"ToBlock"`
	Topics []string `json:"Topics"`
}

type _FilterObject FilterObject

// NewFilterObject instantiates a new FilterObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterObject(address string, fromBlock string, toBlock string, topics []string) *FilterObject {
	this := FilterObject{}
	this.Address = address
	this.FromBlock = fromBlock
	this.ToBlock = toBlock
	this.Topics = topics
	return &this
}

// NewFilterObjectWithDefaults instantiates a new FilterObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterObjectWithDefaults() *FilterObject {
	this := FilterObject{}
	return &this
}

// GetAddress returns the Address field value
func (o *FilterObject) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *FilterObject) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *FilterObject) SetAddress(v string) {
	o.Address = v
}

// GetFromBlock returns the FromBlock field value
func (o *FilterObject) GetFromBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromBlock
}

// GetFromBlockOk returns a tuple with the FromBlock field value
// and a boolean to check if the value has been set.
func (o *FilterObject) GetFromBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromBlock, true
}

// SetFromBlock sets field value
func (o *FilterObject) SetFromBlock(v string) {
	o.FromBlock = v
}

// GetToBlock returns the ToBlock field value
func (o *FilterObject) GetToBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToBlock
}

// GetToBlockOk returns a tuple with the ToBlock field value
// and a boolean to check if the value has been set.
func (o *FilterObject) GetToBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToBlock, true
}

// SetToBlock sets field value
func (o *FilterObject) SetToBlock(v string) {
	o.ToBlock = v
}

// GetTopics returns the Topics field value
func (o *FilterObject) GetTopics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value
// and a boolean to check if the value has been set.
func (o *FilterObject) GetTopicsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Topics, true
}

// SetTopics sets field value
func (o *FilterObject) SetTopics(v []string) {
	o.Topics = v
}

func (o FilterObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Address"] = o.Address
	toSerialize["FromBlock"] = o.FromBlock
	toSerialize["ToBlock"] = o.ToBlock
	toSerialize["Topics"] = o.Topics
	return toSerialize, nil
}

func (o *FilterObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Address",
		"FromBlock",
		"ToBlock",
		"Topics",
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

	varFilterObject := _FilterObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFilterObject)

	if err != nil {
		return err
	}

	*o = FilterObject(varFilterObject)

	return err
}

type NullableFilterObject struct {
	value *FilterObject
	isSet bool
}

func (v NullableFilterObject) Get() *FilterObject {
	return v.value
}

func (v *NullableFilterObject) Set(val *FilterObject) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterObject) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterObject(val *FilterObject) *NullableFilterObject {
	return &NullableFilterObject{value: val, isSet: true}
}

func (v NullableFilterObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


