# OutputGetTransactionReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**TransactionReceipt** | [**ResponseTransactionReceipt**](ResponseTransactionReceipt.md) | A transaction receipt object, or null when the transaction is not available | 

## Methods

### NewOutputGetTransactionReceipt

`func NewOutputGetTransactionReceipt(transactionReceipt ResponseTransactionReceipt, ) *OutputGetTransactionReceipt`

NewOutputGetTransactionReceipt instantiates a new OutputGetTransactionReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetTransactionReceiptWithDefaults

`func NewOutputGetTransactionReceiptWithDefaults() *OutputGetTransactionReceipt`

NewOutputGetTransactionReceiptWithDefaults instantiates a new OutputGetTransactionReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetTransactionReceipt) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetTransactionReceipt) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetTransactionReceipt) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetTransactionReceipt) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTransactionReceipt

`func (o *OutputGetTransactionReceipt) GetTransactionReceipt() ResponseTransactionReceipt`

GetTransactionReceipt returns the TransactionReceipt field if non-nil, zero value otherwise.

### GetTransactionReceiptOk

`func (o *OutputGetTransactionReceipt) GetTransactionReceiptOk() (*ResponseTransactionReceipt, bool)`

GetTransactionReceiptOk returns a tuple with the TransactionReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReceipt

`func (o *OutputGetTransactionReceipt) SetTransactionReceipt(v ResponseTransactionReceipt)`

SetTransactionReceipt sets TransactionReceipt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


