# OutputGetBlockReceipts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**TransactionReceipts** | [**[]ResponseTransactionReceipt**](ResponseTransactionReceipt.md) | An array of transaction receipt objects | 

## Methods

### NewOutputGetBlockReceipts

`func NewOutputGetBlockReceipts(transactionReceipts []ResponseTransactionReceipt, ) *OutputGetBlockReceipts`

NewOutputGetBlockReceipts instantiates a new OutputGetBlockReceipts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetBlockReceiptsWithDefaults

`func NewOutputGetBlockReceiptsWithDefaults() *OutputGetBlockReceipts`

NewOutputGetBlockReceiptsWithDefaults instantiates a new OutputGetBlockReceipts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetBlockReceipts) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetBlockReceipts) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetBlockReceipts) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetBlockReceipts) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTransactionReceipts

`func (o *OutputGetBlockReceipts) GetTransactionReceipts() []ResponseTransactionReceipt`

GetTransactionReceipts returns the TransactionReceipts field if non-nil, zero value otherwise.

### GetTransactionReceiptsOk

`func (o *OutputGetBlockReceipts) GetTransactionReceiptsOk() (*[]ResponseTransactionReceipt, bool)`

GetTransactionReceiptsOk returns a tuple with the TransactionReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReceipts

`func (o *OutputGetBlockReceipts) SetTransactionReceipts(v []ResponseTransactionReceipt)`

SetTransactionReceipts sets TransactionReceipts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


