# InputSendRawTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Data** | **string** | The signed transaction (typically signed with a library, using your private key) | 

## Methods

### NewInputSendRawTransaction

`func NewInputSendRawTransaction(data string, ) *InputSendRawTransaction`

NewInputSendRawTransaction instantiates a new InputSendRawTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSendRawTransactionWithDefaults

`func NewInputSendRawTransactionWithDefaults() *InputSendRawTransaction`

NewInputSendRawTransactionWithDefaults instantiates a new InputSendRawTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *InputSendRawTransaction) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *InputSendRawTransaction) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *InputSendRawTransaction) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *InputSendRawTransaction) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetData

`func (o *InputSendRawTransaction) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InputSendRawTransaction) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InputSendRawTransaction) SetData(v string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


