# InputFeeHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**BlockCount** | **int32** | The number of blocks in the requested range. Between 1 and 1024 blocks can be requested in a single query. It will return less than the requested range if not all blocks are available | [default to 2]
**NewestBlock** | **string** | The highest number block of the requested range in hexadecimal format or tags. The supported tag values include earliest for the earliest/genesis block, latest for the latest mined block, pending for the pending state/transactions. | [default to "latest"]
**RewardPercentiles** | **[]int32** | A list of percentile values with a monotonic increase in value. The transactions will be ranked by effective tip per gas for each block in the requested range, and the corresponding effective tip for the percentile will be calculated while taking gas consumption into consideration | 

## Methods

### NewInputFeeHistory

`func NewInputFeeHistory(blockCount int32, newestBlock string, rewardPercentiles []int32, ) *InputFeeHistory`

NewInputFeeHistory instantiates a new InputFeeHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputFeeHistoryWithDefaults

`func NewInputFeeHistoryWithDefaults() *InputFeeHistory`

NewInputFeeHistoryWithDefaults instantiates a new InputFeeHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *InputFeeHistory) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *InputFeeHistory) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *InputFeeHistory) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *InputFeeHistory) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetBlockCount

`func (o *InputFeeHistory) GetBlockCount() int32`

GetBlockCount returns the BlockCount field if non-nil, zero value otherwise.

### GetBlockCountOk

`func (o *InputFeeHistory) GetBlockCountOk() (*int32, bool)`

GetBlockCountOk returns a tuple with the BlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCount

`func (o *InputFeeHistory) SetBlockCount(v int32)`

SetBlockCount sets BlockCount field to given value.


### GetNewestBlock

`func (o *InputFeeHistory) GetNewestBlock() string`

GetNewestBlock returns the NewestBlock field if non-nil, zero value otherwise.

### GetNewestBlockOk

`func (o *InputFeeHistory) GetNewestBlockOk() (*string, bool)`

GetNewestBlockOk returns a tuple with the NewestBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewestBlock

`func (o *InputFeeHistory) SetNewestBlock(v string)`

SetNewestBlock sets NewestBlock field to given value.


### GetRewardPercentiles

`func (o *InputFeeHistory) GetRewardPercentiles() []int32`

GetRewardPercentiles returns the RewardPercentiles field if non-nil, zero value otherwise.

### GetRewardPercentilesOk

`func (o *InputFeeHistory) GetRewardPercentilesOk() (*[]int32, bool)`

GetRewardPercentilesOk returns a tuple with the RewardPercentiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardPercentiles

`func (o *InputFeeHistory) SetRewardPercentiles(v []int32)`

SetRewardPercentiles sets RewardPercentiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


