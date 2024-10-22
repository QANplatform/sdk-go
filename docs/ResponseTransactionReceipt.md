# ResponseTransactionReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHash** | Pointer to **string** | The hash of the block. null when pending | [optional] 
**BlockNumber** | Pointer to **string** |  | [optional] 
**ContractAddress** | Pointer to **string** | The contract address created if the transaction was a contract creation, otherwise null | [optional] 
**CumulativeGasUsed** | Pointer to **string** | The total amount of gas used when this transaction was executed in the block | [optional] 
**EffectiveGasPrice** | Pointer to **string** | The actual value per gas deducted from the sender account | [optional] 
**From** | Pointer to **string** | The address of the sender | [optional] 
**GasUsed** | Pointer to **string** | The amount of gas used by this specific transaction alone | [optional] 
**Logs** | Pointer to [**[]ResponseLog**](ResponseLog.md) | An array of log objects that generated this transaction | [optional] 
**LogsBloom** | Pointer to **string** | The bloom filter for light clients to quickly retrieve related logs | [optional] 
**Status** | Pointer to **string** | It is either 1 (success) or 0 (failure) encoded as a decimal | [optional] 
**To** | Pointer to **string** | The address of the receiver. null when it&#39;s a contract creation transaction | [optional] 
**TransactionHash** | Pointer to **string** | The hash of the transaction | [optional] 
**TransactionIndex** | Pointer to **string** | An index of the transaction in the block | [optional] 
**Type** | Pointer to **string** | The value type | [optional] 

## Methods

### NewResponseTransactionReceipt

`func NewResponseTransactionReceipt() *ResponseTransactionReceipt`

NewResponseTransactionReceipt instantiates a new ResponseTransactionReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTransactionReceiptWithDefaults

`func NewResponseTransactionReceiptWithDefaults() *ResponseTransactionReceipt`

NewResponseTransactionReceiptWithDefaults instantiates a new ResponseTransactionReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHash

`func (o *ResponseTransactionReceipt) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ResponseTransactionReceipt) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ResponseTransactionReceipt) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *ResponseTransactionReceipt) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBlockNumber

`func (o *ResponseTransactionReceipt) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *ResponseTransactionReceipt) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *ResponseTransactionReceipt) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *ResponseTransactionReceipt) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetContractAddress

`func (o *ResponseTransactionReceipt) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ResponseTransactionReceipt) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ResponseTransactionReceipt) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *ResponseTransactionReceipt) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetCumulativeGasUsed

`func (o *ResponseTransactionReceipt) GetCumulativeGasUsed() string`

GetCumulativeGasUsed returns the CumulativeGasUsed field if non-nil, zero value otherwise.

### GetCumulativeGasUsedOk

`func (o *ResponseTransactionReceipt) GetCumulativeGasUsedOk() (*string, bool)`

GetCumulativeGasUsedOk returns a tuple with the CumulativeGasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeGasUsed

`func (o *ResponseTransactionReceipt) SetCumulativeGasUsed(v string)`

SetCumulativeGasUsed sets CumulativeGasUsed field to given value.

### HasCumulativeGasUsed

`func (o *ResponseTransactionReceipt) HasCumulativeGasUsed() bool`

HasCumulativeGasUsed returns a boolean if a field has been set.

### GetEffectiveGasPrice

`func (o *ResponseTransactionReceipt) GetEffectiveGasPrice() string`

GetEffectiveGasPrice returns the EffectiveGasPrice field if non-nil, zero value otherwise.

### GetEffectiveGasPriceOk

`func (o *ResponseTransactionReceipt) GetEffectiveGasPriceOk() (*string, bool)`

GetEffectiveGasPriceOk returns a tuple with the EffectiveGasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveGasPrice

`func (o *ResponseTransactionReceipt) SetEffectiveGasPrice(v string)`

SetEffectiveGasPrice sets EffectiveGasPrice field to given value.

### HasEffectiveGasPrice

`func (o *ResponseTransactionReceipt) HasEffectiveGasPrice() bool`

HasEffectiveGasPrice returns a boolean if a field has been set.

### GetFrom

`func (o *ResponseTransactionReceipt) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ResponseTransactionReceipt) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ResponseTransactionReceipt) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ResponseTransactionReceipt) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGasUsed

`func (o *ResponseTransactionReceipt) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ResponseTransactionReceipt) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ResponseTransactionReceipt) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.

### HasGasUsed

`func (o *ResponseTransactionReceipt) HasGasUsed() bool`

HasGasUsed returns a boolean if a field has been set.

### GetLogs

`func (o *ResponseTransactionReceipt) GetLogs() []ResponseLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ResponseTransactionReceipt) GetLogsOk() (*[]ResponseLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ResponseTransactionReceipt) SetLogs(v []ResponseLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ResponseTransactionReceipt) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetLogsBloom

`func (o *ResponseTransactionReceipt) GetLogsBloom() string`

GetLogsBloom returns the LogsBloom field if non-nil, zero value otherwise.

### GetLogsBloomOk

`func (o *ResponseTransactionReceipt) GetLogsBloomOk() (*string, bool)`

GetLogsBloomOk returns a tuple with the LogsBloom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsBloom

`func (o *ResponseTransactionReceipt) SetLogsBloom(v string)`

SetLogsBloom sets LogsBloom field to given value.

### HasLogsBloom

`func (o *ResponseTransactionReceipt) HasLogsBloom() bool`

HasLogsBloom returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseTransactionReceipt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseTransactionReceipt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseTransactionReceipt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseTransactionReceipt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTo

`func (o *ResponseTransactionReceipt) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ResponseTransactionReceipt) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ResponseTransactionReceipt) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ResponseTransactionReceipt) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTransactionHash

`func (o *ResponseTransactionReceipt) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ResponseTransactionReceipt) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ResponseTransactionReceipt) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *ResponseTransactionReceipt) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetTransactionIndex

`func (o *ResponseTransactionReceipt) GetTransactionIndex() string`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *ResponseTransactionReceipt) GetTransactionIndexOk() (*string, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *ResponseTransactionReceipt) SetTransactionIndex(v string)`

SetTransactionIndex sets TransactionIndex field to given value.

### HasTransactionIndex

`func (o *ResponseTransactionReceipt) HasTransactionIndex() bool`

HasTransactionIndex returns a boolean if a field has been set.

### GetType

`func (o *ResponseTransactionReceipt) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseTransactionReceipt) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseTransactionReceipt) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseTransactionReceipt) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


