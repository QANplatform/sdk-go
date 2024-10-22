# OutputSyncing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**SyncStatus** | [**SyncStatus**](SyncStatus.md) | Returns null when the node is fully synchronized, otherwise returns the sync status | 

## Methods

### NewOutputSyncing

`func NewOutputSyncing(syncStatus SyncStatus, ) *OutputSyncing`

NewOutputSyncing instantiates a new OutputSyncing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSyncingWithDefaults

`func NewOutputSyncingWithDefaults() *OutputSyncing`

NewOutputSyncingWithDefaults instantiates a new OutputSyncing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputSyncing) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputSyncing) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputSyncing) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputSyncing) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSyncStatus

`func (o *OutputSyncing) GetSyncStatus() SyncStatus`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *OutputSyncing) GetSyncStatusOk() (*SyncStatus, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *OutputSyncing) SetSyncStatus(v SyncStatus)`

SetSyncStatus sets SyncStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


