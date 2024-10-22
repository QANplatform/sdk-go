# OutputGetBlockByHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Block** | [**ResponseBlock**](ResponseBlock.md) | A block object, or null when no block was found | 

## Methods

### NewOutputGetBlockByHash

`func NewOutputGetBlockByHash(block ResponseBlock, ) *OutputGetBlockByHash`

NewOutputGetBlockByHash instantiates a new OutputGetBlockByHash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetBlockByHashWithDefaults

`func NewOutputGetBlockByHashWithDefaults() *OutputGetBlockByHash`

NewOutputGetBlockByHashWithDefaults instantiates a new OutputGetBlockByHash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetBlockByHash) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetBlockByHash) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetBlockByHash) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetBlockByHash) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetBlock

`func (o *OutputGetBlockByHash) GetBlock() ResponseBlock`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *OutputGetBlockByHash) GetBlockOk() (*ResponseBlock, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *OutputGetBlockByHash) SetBlock(v ResponseBlock)`

SetBlock sets Block field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


