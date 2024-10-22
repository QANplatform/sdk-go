# OutputFeeHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**BaseFeePerGas** | **[]string** | An array of block base fees per gas. This includes the next block after the newest of the returned range, because this value can be derived from the newest block. Zeroes are returned for pre-EIP-1559 blocks | 
**GasUsedRatio** | **[]float64** | An array of block gas used ratios. These are calculated as the ratio of gasUsed and gasLimit | 
**OldestBlock** | **string** | The lowest number block of the returned range encoded in decimal format | 
**Reward** | **[][]string** | An array of effective priority fees per gas data points from a single block. All zeroes are returned if the block is empty | 

## Methods

### NewOutputFeeHistory

`func NewOutputFeeHistory(baseFeePerGas []string, gasUsedRatio []float64, oldestBlock string, reward [][]string, ) *OutputFeeHistory`

NewOutputFeeHistory instantiates a new OutputFeeHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputFeeHistoryWithDefaults

`func NewOutputFeeHistoryWithDefaults() *OutputFeeHistory`

NewOutputFeeHistoryWithDefaults instantiates a new OutputFeeHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputFeeHistory) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputFeeHistory) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputFeeHistory) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputFeeHistory) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetBaseFeePerGas

`func (o *OutputFeeHistory) GetBaseFeePerGas() []string`

GetBaseFeePerGas returns the BaseFeePerGas field if non-nil, zero value otherwise.

### GetBaseFeePerGasOk

`func (o *OutputFeeHistory) GetBaseFeePerGasOk() (*[]string, bool)`

GetBaseFeePerGasOk returns a tuple with the BaseFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFeePerGas

`func (o *OutputFeeHistory) SetBaseFeePerGas(v []string)`

SetBaseFeePerGas sets BaseFeePerGas field to given value.


### GetGasUsedRatio

`func (o *OutputFeeHistory) GetGasUsedRatio() []float64`

GetGasUsedRatio returns the GasUsedRatio field if non-nil, zero value otherwise.

### GetGasUsedRatioOk

`func (o *OutputFeeHistory) GetGasUsedRatioOk() (*[]float64, bool)`

GetGasUsedRatioOk returns a tuple with the GasUsedRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsedRatio

`func (o *OutputFeeHistory) SetGasUsedRatio(v []float64)`

SetGasUsedRatio sets GasUsedRatio field to given value.


### GetOldestBlock

`func (o *OutputFeeHistory) GetOldestBlock() string`

GetOldestBlock returns the OldestBlock field if non-nil, zero value otherwise.

### GetOldestBlockOk

`func (o *OutputFeeHistory) GetOldestBlockOk() (*string, bool)`

GetOldestBlockOk returns a tuple with the OldestBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestBlock

`func (o *OutputFeeHistory) SetOldestBlock(v string)`

SetOldestBlock sets OldestBlock field to given value.


### GetReward

`func (o *OutputFeeHistory) GetReward() [][]string`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *OutputFeeHistory) GetRewardOk() (*[][]string, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *OutputFeeHistory) SetReward(v [][]string)`

SetReward sets Reward field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


