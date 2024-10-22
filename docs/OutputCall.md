# OutputCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Data** | **string** | The return value of the executed contract method | 

## Methods

### NewOutputCall

`func NewOutputCall(data string, ) *OutputCall`

NewOutputCall instantiates a new OutputCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputCallWithDefaults

`func NewOutputCallWithDefaults() *OutputCall`

NewOutputCallWithDefaults instantiates a new OutputCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputCall) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputCall) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputCall) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputCall) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetData

`func (o *OutputCall) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OutputCall) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OutputCall) SetData(v string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


