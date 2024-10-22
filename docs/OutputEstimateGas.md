# OutputEstimateGas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**EstimatedGas** | **string** | The estimated amount of gas used | 

## Methods

### NewOutputEstimateGas

`func NewOutputEstimateGas(estimatedGas string, ) *OutputEstimateGas`

NewOutputEstimateGas instantiates a new OutputEstimateGas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputEstimateGasWithDefaults

`func NewOutputEstimateGasWithDefaults() *OutputEstimateGas`

NewOutputEstimateGasWithDefaults instantiates a new OutputEstimateGas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputEstimateGas) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputEstimateGas) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputEstimateGas) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputEstimateGas) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetEstimatedGas

`func (o *OutputEstimateGas) GetEstimatedGas() string`

GetEstimatedGas returns the EstimatedGas field if non-nil, zero value otherwise.

### GetEstimatedGasOk

`func (o *OutputEstimateGas) GetEstimatedGasOk() (*string, bool)`

GetEstimatedGasOk returns a tuple with the EstimatedGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedGas

`func (o *OutputEstimateGas) SetEstimatedGas(v string)`

SetEstimatedGas sets EstimatedGas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


