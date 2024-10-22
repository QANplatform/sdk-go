# OutputGetFilterLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Logs** | [**[]ResponseLog**](ResponseLog.md) | An array of log objects, or an empty array if nothing has changed since last poll | 

## Methods

### NewOutputGetFilterLogs

`func NewOutputGetFilterLogs(logs []ResponseLog, ) *OutputGetFilterLogs`

NewOutputGetFilterLogs instantiates a new OutputGetFilterLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetFilterLogsWithDefaults

`func NewOutputGetFilterLogsWithDefaults() *OutputGetFilterLogs`

NewOutputGetFilterLogsWithDefaults instantiates a new OutputGetFilterLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetFilterLogs) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetFilterLogs) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetFilterLogs) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetFilterLogs) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetLogs

`func (o *OutputGetFilterLogs) GetLogs() []ResponseLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *OutputGetFilterLogs) GetLogsOk() (*[]ResponseLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *OutputGetFilterLogs) SetLogs(v []ResponseLog)`

SetLogs sets Logs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


