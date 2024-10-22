# OutputBlockNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**BlockNumber** | **string** | An integer value of the latest block number encoded as decimal | 

## Methods

### NewOutputBlockNumber

`func NewOutputBlockNumber(blockNumber string, ) *OutputBlockNumber`

NewOutputBlockNumber instantiates a new OutputBlockNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputBlockNumberWithDefaults

`func NewOutputBlockNumberWithDefaults() *OutputBlockNumber`

NewOutputBlockNumberWithDefaults instantiates a new OutputBlockNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputBlockNumber) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputBlockNumber) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputBlockNumber) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputBlockNumber) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetBlockNumber

`func (o *OutputBlockNumber) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *OutputBlockNumber) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *OutputBlockNumber) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


