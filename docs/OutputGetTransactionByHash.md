# OutputGetTransactionByHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Transaction** | [**ResponseTransaction**](ResponseTransaction.md) | The transaction response object, or null if no transaction is found | 

## Methods

### NewOutputGetTransactionByHash

`func NewOutputGetTransactionByHash(transaction ResponseTransaction, ) *OutputGetTransactionByHash`

NewOutputGetTransactionByHash instantiates a new OutputGetTransactionByHash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetTransactionByHashWithDefaults

`func NewOutputGetTransactionByHashWithDefaults() *OutputGetTransactionByHash`

NewOutputGetTransactionByHashWithDefaults instantiates a new OutputGetTransactionByHash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetTransactionByHash) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetTransactionByHash) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetTransactionByHash) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetTransactionByHash) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTransaction

`func (o *OutputGetTransactionByHash) GetTransaction() ResponseTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *OutputGetTransactionByHash) GetTransactionOk() (*ResponseTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *OutputGetTransactionByHash) SetTransaction(v ResponseTransaction)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


