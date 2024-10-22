# InputGetLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Address** | Pointer to **string** | The contract address or a list of addresses from which logs should originate | [optional] 
**BlockHash** | Pointer to **string** | With the addition of EIP-234, blockHash is a new filter option that restricts the logs returned to the block number referenced in the blockHash. Using the blockHash field is equivalent to setting the fromBlock and toBlock to the block number the blockHash references. If blockHash is present in the filter criteria, neither fromBlock nor toBlock is allowed | [optional] 
**FromBlock** | Pointer to **string** | The block number as a string in decimal format or tags. The supported tag values include earliest for the earliest/genesis block, latest for the latest mined block, pending for the pending state/transactions. | [optional] 
**ToBlock** | Pointer to **string** | The block number as a string in decimal format or tags. The supported tag values include earliest for the earliest/genesis block, latest for the latest mined block, pending for the pending state/transactions. | [optional] 
**Topics** | Pointer to **[]string** | An array of DATA topics and also, the topics are order-dependent. Visit this official page to learn more about topics | [optional] 

## Methods

### NewInputGetLogs

`func NewInputGetLogs() *InputGetLogs`

NewInputGetLogs instantiates a new InputGetLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGetLogsWithDefaults

`func NewInputGetLogsWithDefaults() *InputGetLogs`

NewInputGetLogsWithDefaults instantiates a new InputGetLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *InputGetLogs) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *InputGetLogs) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *InputGetLogs) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *InputGetLogs) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAddress

`func (o *InputGetLogs) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InputGetLogs) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InputGetLogs) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InputGetLogs) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBlockHash

`func (o *InputGetLogs) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *InputGetLogs) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *InputGetLogs) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *InputGetLogs) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetFromBlock

`func (o *InputGetLogs) GetFromBlock() string`

GetFromBlock returns the FromBlock field if non-nil, zero value otherwise.

### GetFromBlockOk

`func (o *InputGetLogs) GetFromBlockOk() (*string, bool)`

GetFromBlockOk returns a tuple with the FromBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromBlock

`func (o *InputGetLogs) SetFromBlock(v string)`

SetFromBlock sets FromBlock field to given value.

### HasFromBlock

`func (o *InputGetLogs) HasFromBlock() bool`

HasFromBlock returns a boolean if a field has been set.

### GetToBlock

`func (o *InputGetLogs) GetToBlock() string`

GetToBlock returns the ToBlock field if non-nil, zero value otherwise.

### GetToBlockOk

`func (o *InputGetLogs) GetToBlockOk() (*string, bool)`

GetToBlockOk returns a tuple with the ToBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBlock

`func (o *InputGetLogs) SetToBlock(v string)`

SetToBlock sets ToBlock field to given value.

### HasToBlock

`func (o *InputGetLogs) HasToBlock() bool`

HasToBlock returns a boolean if a field has been set.

### GetTopics

`func (o *InputGetLogs) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *InputGetLogs) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *InputGetLogs) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *InputGetLogs) HasTopics() bool`

HasTopics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


