# InputCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**BlockNumber** | **string** |  | 
**Transaction** | [**ParamsTransaction**](ParamsTransaction.md) | The transaction call object | 

## Methods

### NewInputCall

`func NewInputCall(blockNumber string, transaction ParamsTransaction, ) *InputCall`

NewInputCall instantiates a new InputCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputCallWithDefaults

`func NewInputCallWithDefaults() *InputCall`

NewInputCallWithDefaults instantiates a new InputCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *InputCall) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *InputCall) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *InputCall) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *InputCall) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetBlockNumber

`func (o *InputCall) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *InputCall) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *InputCall) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.


### GetTransaction

`func (o *InputCall) GetTransaction() ParamsTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *InputCall) GetTransactionOk() (*ParamsTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *InputCall) SetTransaction(v ParamsTransaction)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


