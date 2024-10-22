# OutputGetCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Bytecode** | **string** | The bytecode from a given address | 

## Methods

### NewOutputGetCode

`func NewOutputGetCode(bytecode string, ) *OutputGetCode`

NewOutputGetCode instantiates a new OutputGetCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetCodeWithDefaults

`func NewOutputGetCodeWithDefaults() *OutputGetCode`

NewOutputGetCodeWithDefaults instantiates a new OutputGetCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetCode) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetCode) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetCode) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetCode) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetBytecode

`func (o *OutputGetCode) GetBytecode() string`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *OutputGetCode) GetBytecodeOk() (*string, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *OutputGetCode) SetBytecode(v string)`

SetBytecode sets Bytecode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


