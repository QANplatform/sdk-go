# OutputGetBlockTransactionCountByNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**TransactionCount** | **string** | The number of transactions associated with a specific block, in decimal value | 

## Methods

### NewOutputGetBlockTransactionCountByNumber

`func NewOutputGetBlockTransactionCountByNumber(transactionCount string, ) *OutputGetBlockTransactionCountByNumber`

NewOutputGetBlockTransactionCountByNumber instantiates a new OutputGetBlockTransactionCountByNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetBlockTransactionCountByNumberWithDefaults

`func NewOutputGetBlockTransactionCountByNumberWithDefaults() *OutputGetBlockTransactionCountByNumber`

NewOutputGetBlockTransactionCountByNumberWithDefaults instantiates a new OutputGetBlockTransactionCountByNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetBlockTransactionCountByNumber) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetBlockTransactionCountByNumber) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetBlockTransactionCountByNumber) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetBlockTransactionCountByNumber) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTransactionCount

`func (o *OutputGetBlockTransactionCountByNumber) GetTransactionCount() string`

GetTransactionCount returns the TransactionCount field if non-nil, zero value otherwise.

### GetTransactionCountOk

`func (o *OutputGetBlockTransactionCountByNumber) GetTransactionCountOk() (*string, bool)`

GetTransactionCountOk returns a tuple with the TransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCount

`func (o *OutputGetBlockTransactionCountByNumber) SetTransactionCount(v string)`

SetTransactionCount sets TransactionCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


