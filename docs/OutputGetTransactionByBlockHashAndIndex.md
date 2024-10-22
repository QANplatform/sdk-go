# OutputGetTransactionByBlockHashAndIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Transaction** | [**ResponseTransaction**](ResponseTransaction.md) | The transaction response object, or null if no transaction is found | 

## Methods

### NewOutputGetTransactionByBlockHashAndIndex

`func NewOutputGetTransactionByBlockHashAndIndex(transaction ResponseTransaction, ) *OutputGetTransactionByBlockHashAndIndex`

NewOutputGetTransactionByBlockHashAndIndex instantiates a new OutputGetTransactionByBlockHashAndIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetTransactionByBlockHashAndIndexWithDefaults

`func NewOutputGetTransactionByBlockHashAndIndexWithDefaults() *OutputGetTransactionByBlockHashAndIndex`

NewOutputGetTransactionByBlockHashAndIndexWithDefaults instantiates a new OutputGetTransactionByBlockHashAndIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetTransactionByBlockHashAndIndex) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetTransactionByBlockHashAndIndex) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetTransactionByBlockHashAndIndex) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetTransactionByBlockHashAndIndex) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTransaction

`func (o *OutputGetTransactionByBlockHashAndIndex) GetTransaction() ResponseTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *OutputGetTransactionByBlockHashAndIndex) GetTransactionOk() (*ResponseTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *OutputGetTransactionByBlockHashAndIndex) SetTransaction(v ResponseTransaction)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


