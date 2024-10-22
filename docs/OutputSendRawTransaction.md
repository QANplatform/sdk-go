# OutputSendRawTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**TransactionHash** | **string** | The transaction hash, or the zero hash if the transaction is not yet available | 

## Methods

### NewOutputSendRawTransaction

`func NewOutputSendRawTransaction(transactionHash string, ) *OutputSendRawTransaction`

NewOutputSendRawTransaction instantiates a new OutputSendRawTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSendRawTransactionWithDefaults

`func NewOutputSendRawTransactionWithDefaults() *OutputSendRawTransaction`

NewOutputSendRawTransactionWithDefaults instantiates a new OutputSendRawTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputSendRawTransaction) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputSendRawTransaction) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputSendRawTransaction) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputSendRawTransaction) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTransactionHash

`func (o *OutputSendRawTransaction) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *OutputSendRawTransaction) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *OutputSendRawTransaction) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


