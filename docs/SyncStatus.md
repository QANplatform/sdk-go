# SyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentBlock** | **string** | The current block, same as getFilterChangesblockNumber encoded as decimal | 
**HighestBlock** | **string** | The estimated highest block encoded as decimal | 
**StartingBlock** | **string** | The block at which the import started encoded as decimal | 

## Methods

### NewSyncStatus

`func NewSyncStatus(currentBlock string, highestBlock string, startingBlock string, ) *SyncStatus`

NewSyncStatus instantiates a new SyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncStatusWithDefaults

`func NewSyncStatusWithDefaults() *SyncStatus`

NewSyncStatusWithDefaults instantiates a new SyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentBlock

`func (o *SyncStatus) GetCurrentBlock() string`

GetCurrentBlock returns the CurrentBlock field if non-nil, zero value otherwise.

### GetCurrentBlockOk

`func (o *SyncStatus) GetCurrentBlockOk() (*string, bool)`

GetCurrentBlockOk returns a tuple with the CurrentBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBlock

`func (o *SyncStatus) SetCurrentBlock(v string)`

SetCurrentBlock sets CurrentBlock field to given value.


### GetHighestBlock

`func (o *SyncStatus) GetHighestBlock() string`

GetHighestBlock returns the HighestBlock field if non-nil, zero value otherwise.

### GetHighestBlockOk

`func (o *SyncStatus) GetHighestBlockOk() (*string, bool)`

GetHighestBlockOk returns a tuple with the HighestBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestBlock

`func (o *SyncStatus) SetHighestBlock(v string)`

SetHighestBlock sets HighestBlock field to given value.


### GetStartingBlock

`func (o *SyncStatus) GetStartingBlock() string`

GetStartingBlock returns the StartingBlock field if non-nil, zero value otherwise.

### GetStartingBlockOk

`func (o *SyncStatus) GetStartingBlockOk() (*string, bool)`

GetStartingBlockOk returns a tuple with the StartingBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingBlock

`func (o *SyncStatus) SetStartingBlock(v string)`

SetStartingBlock sets StartingBlock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


