# OutputGetBlockByNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Block** | [**ResponseBlock**](ResponseBlock.md) | A block object, or null when no block was found | 

## Methods

### NewOutputGetBlockByNumber

`func NewOutputGetBlockByNumber(block ResponseBlock, ) *OutputGetBlockByNumber`

NewOutputGetBlockByNumber instantiates a new OutputGetBlockByNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetBlockByNumberWithDefaults

`func NewOutputGetBlockByNumberWithDefaults() *OutputGetBlockByNumber`

NewOutputGetBlockByNumberWithDefaults instantiates a new OutputGetBlockByNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetBlockByNumber) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetBlockByNumber) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetBlockByNumber) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetBlockByNumber) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetBlock

`func (o *OutputGetBlockByNumber) GetBlock() ResponseBlock`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *OutputGetBlockByNumber) GetBlockOk() (*ResponseBlock, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *OutputGetBlockByNumber) SetBlock(v ResponseBlock)`

SetBlock sets Block field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


