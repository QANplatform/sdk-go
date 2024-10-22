# OutputGetBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Balance** | **string** | The balance of the specified address in decimal value, measured in wei | 

## Methods

### NewOutputGetBalance

`func NewOutputGetBalance(balance string, ) *OutputGetBalance`

NewOutputGetBalance instantiates a new OutputGetBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetBalanceWithDefaults

`func NewOutputGetBalanceWithDefaults() *OutputGetBalance`

NewOutputGetBalanceWithDefaults instantiates a new OutputGetBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetBalance) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetBalance) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetBalance) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetBalance) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetBalance

`func (o *OutputGetBalance) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *OutputGetBalance) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *OutputGetBalance) SetBalance(v string)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


