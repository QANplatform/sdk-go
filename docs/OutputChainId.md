# OutputChainId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**ChainId** | **string** | A decimal value in string format which represents an integer of the chain ID | 

## Methods

### NewOutputChainId

`func NewOutputChainId(chainId string, ) *OutputChainId`

NewOutputChainId instantiates a new OutputChainId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputChainIdWithDefaults

`func NewOutputChainIdWithDefaults() *OutputChainId`

NewOutputChainIdWithDefaults instantiates a new OutputChainId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputChainId) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputChainId) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputChainId) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputChainId) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetChainId

`func (o *OutputChainId) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *OutputChainId) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *OutputChainId) SetChainId(v string)`

SetChainId sets ChainId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


