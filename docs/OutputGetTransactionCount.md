# OutputGetTransactionCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**TransactionCount** | **string** | The integer of the number of transactions sent from an address encoded as decimal | 

## Methods

### NewOutputGetTransactionCount

`func NewOutputGetTransactionCount(transactionCount string, ) *OutputGetTransactionCount`

NewOutputGetTransactionCount instantiates a new OutputGetTransactionCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetTransactionCountWithDefaults

`func NewOutputGetTransactionCountWithDefaults() *OutputGetTransactionCount`

NewOutputGetTransactionCountWithDefaults instantiates a new OutputGetTransactionCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetTransactionCount) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetTransactionCount) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetTransactionCount) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetTransactionCount) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTransactionCount

`func (o *OutputGetTransactionCount) GetTransactionCount() string`

GetTransactionCount returns the TransactionCount field if non-nil, zero value otherwise.

### GetTransactionCountOk

`func (o *OutputGetTransactionCount) GetTransactionCountOk() (*string, bool)`

GetTransactionCountOk returns a tuple with the TransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCount

`func (o *OutputGetTransactionCount) SetTransactionCount(v string)`

SetTransactionCount sets TransactionCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


