# InputEstimateGas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Object** | Pointer to [**EstimateGasObject**](EstimateGasObject.md) | The state override set with address-to-state mapping where each address maps to an object containing | [optional] 
**Transaction** | [**ParamsTransaction**](ParamsTransaction.md) | The transaction call object | 

## Methods

### NewInputEstimateGas

`func NewInputEstimateGas(transaction ParamsTransaction, ) *InputEstimateGas`

NewInputEstimateGas instantiates a new InputEstimateGas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputEstimateGasWithDefaults

`func NewInputEstimateGasWithDefaults() *InputEstimateGas`

NewInputEstimateGasWithDefaults instantiates a new InputEstimateGas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *InputEstimateGas) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *InputEstimateGas) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *InputEstimateGas) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *InputEstimateGas) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetObject

`func (o *InputEstimateGas) GetObject() EstimateGasObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *InputEstimateGas) GetObjectOk() (*EstimateGasObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *InputEstimateGas) SetObject(v EstimateGasObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *InputEstimateGas) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetTransaction

`func (o *InputEstimateGas) GetTransaction() ParamsTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *InputEstimateGas) GetTransactionOk() (*ParamsTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *InputEstimateGas) SetTransaction(v ParamsTransaction)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


