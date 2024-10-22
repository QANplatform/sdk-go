# OutputGasPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**GasPrice** | **string** | The decimal value of the current gas price in wei | 

## Methods

### NewOutputGasPrice

`func NewOutputGasPrice(gasPrice string, ) *OutputGasPrice`

NewOutputGasPrice instantiates a new OutputGasPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGasPriceWithDefaults

`func NewOutputGasPriceWithDefaults() *OutputGasPrice`

NewOutputGasPriceWithDefaults instantiates a new OutputGasPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGasPrice) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGasPrice) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGasPrice) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGasPrice) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetGasPrice

`func (o *OutputGasPrice) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *OutputGasPrice) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *OutputGasPrice) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


