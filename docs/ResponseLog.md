# ResponseLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | An address from which this log originated | 
**BlockHash** | **NullableString** | The hash of the block where this log was in. null when its a pending log | 
**BlockNumber** | **NullableString** | The block number where this log was in. null when its a pending log | 
**Data** | **string** | It contains one or more 32 Bytes non-indexed arguments of the log | 
**LogIndex** | **NullableString** | The integer of the log index position in the block. null when its a pending log | 
**Removed** | **bool** | It is true when the log was removed due to a chain reorganization, and false if it&#39;s a valid log | 
**Topics** | **[]string** | An array of zero to four 32 Bytes DATA of indexed log arguments. In Solidity, the first topic is the hash of the signature of the event (e.g. Deposit(address, bytes32, uint256)), except you declare the event with the anonymous specifier | 
**TransactionHash** | **NullableString** | The hash of the transactions this log was created from. null when its a pending log | 
**TransactionIndex** | **NullableString** | The integer of the transaction&#39;s index position that the log was created from. null when it&#39;s a pending log | 

## Methods

### NewResponseLog

`func NewResponseLog(address string, blockHash NullableString, blockNumber NullableString, data string, logIndex NullableString, removed bool, topics []string, transactionHash NullableString, transactionIndex NullableString, ) *ResponseLog`

NewResponseLog instantiates a new ResponseLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseLogWithDefaults

`func NewResponseLogWithDefaults() *ResponseLog`

NewResponseLogWithDefaults instantiates a new ResponseLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ResponseLog) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ResponseLog) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ResponseLog) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBlockHash

`func (o *ResponseLog) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ResponseLog) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ResponseLog) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### SetBlockHashNil

`func (o *ResponseLog) SetBlockHashNil(b bool)`

 SetBlockHashNil sets the value for BlockHash to be an explicit nil

### UnsetBlockHash
`func (o *ResponseLog) UnsetBlockHash()`

UnsetBlockHash ensures that no value is present for BlockHash, not even an explicit nil
### GetBlockNumber

`func (o *ResponseLog) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *ResponseLog) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *ResponseLog) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.


### SetBlockNumberNil

`func (o *ResponseLog) SetBlockNumberNil(b bool)`

 SetBlockNumberNil sets the value for BlockNumber to be an explicit nil

### UnsetBlockNumber
`func (o *ResponseLog) UnsetBlockNumber()`

UnsetBlockNumber ensures that no value is present for BlockNumber, not even an explicit nil
### GetData

`func (o *ResponseLog) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseLog) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseLog) SetData(v string)`

SetData sets Data field to given value.


### GetLogIndex

`func (o *ResponseLog) GetLogIndex() string`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *ResponseLog) GetLogIndexOk() (*string, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *ResponseLog) SetLogIndex(v string)`

SetLogIndex sets LogIndex field to given value.


### SetLogIndexNil

`func (o *ResponseLog) SetLogIndexNil(b bool)`

 SetLogIndexNil sets the value for LogIndex to be an explicit nil

### UnsetLogIndex
`func (o *ResponseLog) UnsetLogIndex()`

UnsetLogIndex ensures that no value is present for LogIndex, not even an explicit nil
### GetRemoved

`func (o *ResponseLog) GetRemoved() bool`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *ResponseLog) GetRemovedOk() (*bool, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *ResponseLog) SetRemoved(v bool)`

SetRemoved sets Removed field to given value.


### GetTopics

`func (o *ResponseLog) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *ResponseLog) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *ResponseLog) SetTopics(v []string)`

SetTopics sets Topics field to given value.


### GetTransactionHash

`func (o *ResponseLog) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ResponseLog) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ResponseLog) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### SetTransactionHashNil

`func (o *ResponseLog) SetTransactionHashNil(b bool)`

 SetTransactionHashNil sets the value for TransactionHash to be an explicit nil

### UnsetTransactionHash
`func (o *ResponseLog) UnsetTransactionHash()`

UnsetTransactionHash ensures that no value is present for TransactionHash, not even an explicit nil
### GetTransactionIndex

`func (o *ResponseLog) GetTransactionIndex() string`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *ResponseLog) GetTransactionIndexOk() (*string, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *ResponseLog) SetTransactionIndex(v string)`

SetTransactionIndex sets TransactionIndex field to given value.


### SetTransactionIndexNil

`func (o *ResponseLog) SetTransactionIndexNil(b bool)`

 SetTransactionIndexNil sets the value for TransactionIndex to be an explicit nil

### UnsetTransactionIndex
`func (o *ResponseLog) UnsetTransactionIndex()`

UnsetTransactionIndex ensures that no value is present for TransactionIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


