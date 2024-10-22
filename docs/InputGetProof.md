# InputGetProof

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Address** | **string** |  | 
**BlockNumber** | Pointer to **string** |  | [optional] [default to "latest"]
**StorageKeys** | **[]string** | An array of storage-keys that should be proofed and included | 

## Methods

### NewInputGetProof

`func NewInputGetProof(address string, storageKeys []string, ) *InputGetProof`

NewInputGetProof instantiates a new InputGetProof object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGetProofWithDefaults

`func NewInputGetProofWithDefaults() *InputGetProof`

NewInputGetProofWithDefaults instantiates a new InputGetProof object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *InputGetProof) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *InputGetProof) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *InputGetProof) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *InputGetProof) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAddress

`func (o *InputGetProof) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InputGetProof) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InputGetProof) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBlockNumber

`func (o *InputGetProof) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *InputGetProof) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *InputGetProof) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *InputGetProof) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetStorageKeys

`func (o *InputGetProof) GetStorageKeys() []string`

GetStorageKeys returns the StorageKeys field if non-nil, zero value otherwise.

### GetStorageKeysOk

`func (o *InputGetProof) GetStorageKeysOk() (*[]string, bool)`

GetStorageKeysOk returns a tuple with the StorageKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKeys

`func (o *InputGetProof) SetStorageKeys(v []string)`

SetStorageKeys sets StorageKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


