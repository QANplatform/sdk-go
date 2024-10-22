/*
QAN AutoApi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qan

import (
	"encoding/json"
)

// checks if the ResponseTransactionReceipt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTransactionReceipt{}

// ResponseTransactionReceipt struct for ResponseTransactionReceipt
type ResponseTransactionReceipt struct {
	// The hash of the block. null when pending
	BlockHash *string `json:"BlockHash,omitempty"`
	BlockNumber *string `json:"BlockNumber,omitempty"`
	// The contract address created if the transaction was a contract creation, otherwise null
	ContractAddress *string `json:"ContractAddress,omitempty"`
	// The total amount of gas used when this transaction was executed in the block
	CumulativeGasUsed *string `json:"CumulativeGasUsed,omitempty"`
	// The actual value per gas deducted from the sender account
	EffectiveGasPrice *string `json:"EffectiveGasPrice,omitempty"`
	// The address of the sender
	From *string `json:"From,omitempty"`
	// The amount of gas used by this specific transaction alone
	GasUsed *string `json:"GasUsed,omitempty"`
	// An array of log objects that generated this transaction
	Logs []ResponseLog `json:"Logs,omitempty"`
	// The bloom filter for light clients to quickly retrieve related logs
	LogsBloom *string `json:"LogsBloom,omitempty"`
	// It is either 1 (success) or 0 (failure) encoded as a decimal
	Status *string `json:"Status,omitempty"`
	// The address of the receiver. null when it's a contract creation transaction
	To *string `json:"To,omitempty"`
	// The hash of the transaction
	TransactionHash *string `json:"TransactionHash,omitempty"`
	// An index of the transaction in the block
	TransactionIndex *string `json:"TransactionIndex,omitempty"`
	// The value type
	Type *string `json:"Type,omitempty"`
}

// NewResponseTransactionReceipt instantiates a new ResponseTransactionReceipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTransactionReceipt() *ResponseTransactionReceipt {
	this := ResponseTransactionReceipt{}
	return &this
}

// NewResponseTransactionReceiptWithDefaults instantiates a new ResponseTransactionReceipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTransactionReceiptWithDefaults() *ResponseTransactionReceipt {
	this := ResponseTransactionReceipt{}
	return &this
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetBlockHash() string {
	if o == nil || IsNil(o.BlockHash) {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetBlockHashOk() (*string, bool) {
	if o == nil || IsNil(o.BlockHash) {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasBlockHash() bool {
	if o != nil && !IsNil(o.BlockHash) {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *ResponseTransactionReceipt) SetBlockHash(v string) {
	o.BlockHash = &v
}

// GetBlockNumber returns the BlockNumber field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetBlockNumber() string {
	if o == nil || IsNil(o.BlockNumber) {
		var ret string
		return ret
	}
	return *o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetBlockNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BlockNumber) {
		return nil, false
	}
	return o.BlockNumber, true
}

// HasBlockNumber returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasBlockNumber() bool {
	if o != nil && !IsNil(o.BlockNumber) {
		return true
	}

	return false
}

// SetBlockNumber gets a reference to the given string and assigns it to the BlockNumber field.
func (o *ResponseTransactionReceipt) SetBlockNumber(v string) {
	o.BlockNumber = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetContractAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasContractAddress() bool {
	if o != nil && !IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *ResponseTransactionReceipt) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetCumulativeGasUsed returns the CumulativeGasUsed field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetCumulativeGasUsed() string {
	if o == nil || IsNil(o.CumulativeGasUsed) {
		var ret string
		return ret
	}
	return *o.CumulativeGasUsed
}

// GetCumulativeGasUsedOk returns a tuple with the CumulativeGasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetCumulativeGasUsedOk() (*string, bool) {
	if o == nil || IsNil(o.CumulativeGasUsed) {
		return nil, false
	}
	return o.CumulativeGasUsed, true
}

// HasCumulativeGasUsed returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasCumulativeGasUsed() bool {
	if o != nil && !IsNil(o.CumulativeGasUsed) {
		return true
	}

	return false
}

// SetCumulativeGasUsed gets a reference to the given string and assigns it to the CumulativeGasUsed field.
func (o *ResponseTransactionReceipt) SetCumulativeGasUsed(v string) {
	o.CumulativeGasUsed = &v
}

// GetEffectiveGasPrice returns the EffectiveGasPrice field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetEffectiveGasPrice() string {
	if o == nil || IsNil(o.EffectiveGasPrice) {
		var ret string
		return ret
	}
	return *o.EffectiveGasPrice
}

// GetEffectiveGasPriceOk returns a tuple with the EffectiveGasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetEffectiveGasPriceOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveGasPrice) {
		return nil, false
	}
	return o.EffectiveGasPrice, true
}

// HasEffectiveGasPrice returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasEffectiveGasPrice() bool {
	if o != nil && !IsNil(o.EffectiveGasPrice) {
		return true
	}

	return false
}

// SetEffectiveGasPrice gets a reference to the given string and assigns it to the EffectiveGasPrice field.
func (o *ResponseTransactionReceipt) SetEffectiveGasPrice(v string) {
	o.EffectiveGasPrice = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *ResponseTransactionReceipt) SetFrom(v string) {
	o.From = &v
}

// GetGasUsed returns the GasUsed field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetGasUsed() string {
	if o == nil || IsNil(o.GasUsed) {
		var ret string
		return ret
	}
	return *o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetGasUsedOk() (*string, bool) {
	if o == nil || IsNil(o.GasUsed) {
		return nil, false
	}
	return o.GasUsed, true
}

// HasGasUsed returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasGasUsed() bool {
	if o != nil && !IsNil(o.GasUsed) {
		return true
	}

	return false
}

// SetGasUsed gets a reference to the given string and assigns it to the GasUsed field.
func (o *ResponseTransactionReceipt) SetGasUsed(v string) {
	o.GasUsed = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetLogs() []ResponseLog {
	if o == nil || IsNil(o.Logs) {
		var ret []ResponseLog
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetLogsOk() ([]ResponseLog, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []ResponseLog and assigns it to the Logs field.
func (o *ResponseTransactionReceipt) SetLogs(v []ResponseLog) {
	o.Logs = v
}

// GetLogsBloom returns the LogsBloom field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetLogsBloom() string {
	if o == nil || IsNil(o.LogsBloom) {
		var ret string
		return ret
	}
	return *o.LogsBloom
}

// GetLogsBloomOk returns a tuple with the LogsBloom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetLogsBloomOk() (*string, bool) {
	if o == nil || IsNil(o.LogsBloom) {
		return nil, false
	}
	return o.LogsBloom, true
}

// HasLogsBloom returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasLogsBloom() bool {
	if o != nil && !IsNil(o.LogsBloom) {
		return true
	}

	return false
}

// SetLogsBloom gets a reference to the given string and assigns it to the LogsBloom field.
func (o *ResponseTransactionReceipt) SetLogsBloom(v string) {
	o.LogsBloom = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ResponseTransactionReceipt) SetStatus(v string) {
	o.Status = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *ResponseTransactionReceipt) SetTo(v string) {
	o.To = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetTransactionHashOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasTransactionHash() bool {
	if o != nil && !IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given string and assigns it to the TransactionHash field.
func (o *ResponseTransactionReceipt) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

// GetTransactionIndex returns the TransactionIndex field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetTransactionIndex() string {
	if o == nil || IsNil(o.TransactionIndex) {
		var ret string
		return ret
	}
	return *o.TransactionIndex
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetTransactionIndexOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionIndex) {
		return nil, false
	}
	return o.TransactionIndex, true
}

// HasTransactionIndex returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasTransactionIndex() bool {
	if o != nil && !IsNil(o.TransactionIndex) {
		return true
	}

	return false
}

// SetTransactionIndex gets a reference to the given string and assigns it to the TransactionIndex field.
func (o *ResponseTransactionReceipt) SetTransactionIndex(v string) {
	o.TransactionIndex = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResponseTransactionReceipt) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionReceipt) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResponseTransactionReceipt) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ResponseTransactionReceipt) SetType(v string) {
	o.Type = &v
}

func (o ResponseTransactionReceipt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTransactionReceipt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockHash) {
		toSerialize["BlockHash"] = o.BlockHash
	}
	if !IsNil(o.BlockNumber) {
		toSerialize["BlockNumber"] = o.BlockNumber
	}
	if !IsNil(o.ContractAddress) {
		toSerialize["ContractAddress"] = o.ContractAddress
	}
	if !IsNil(o.CumulativeGasUsed) {
		toSerialize["CumulativeGasUsed"] = o.CumulativeGasUsed
	}
	if !IsNil(o.EffectiveGasPrice) {
		toSerialize["EffectiveGasPrice"] = o.EffectiveGasPrice
	}
	if !IsNil(o.From) {
		toSerialize["From"] = o.From
	}
	if !IsNil(o.GasUsed) {
		toSerialize["GasUsed"] = o.GasUsed
	}
	if !IsNil(o.Logs) {
		toSerialize["Logs"] = o.Logs
	}
	if !IsNil(o.LogsBloom) {
		toSerialize["LogsBloom"] = o.LogsBloom
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.To) {
		toSerialize["To"] = o.To
	}
	if !IsNil(o.TransactionHash) {
		toSerialize["TransactionHash"] = o.TransactionHash
	}
	if !IsNil(o.TransactionIndex) {
		toSerialize["TransactionIndex"] = o.TransactionIndex
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	return toSerialize, nil
}

type NullableResponseTransactionReceipt struct {
	value *ResponseTransactionReceipt
	isSet bool
}

func (v NullableResponseTransactionReceipt) Get() *ResponseTransactionReceipt {
	return v.value
}

func (v *NullableResponseTransactionReceipt) Set(val *ResponseTransactionReceipt) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTransactionReceipt) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTransactionReceipt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTransactionReceipt(val *ResponseTransactionReceipt) *NullableResponseTransactionReceipt {
	return &NullableResponseTransactionReceipt{value: val, isSet: true}
}

func (v NullableResponseTransactionReceipt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTransactionReceipt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


