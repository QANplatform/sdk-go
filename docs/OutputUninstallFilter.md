# OutputUninstallFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Result** | **bool** | Returns true if the filter was successfully uninstalled, otherwise false | 

## Methods

### NewOutputUninstallFilter

`func NewOutputUninstallFilter(result bool, ) *OutputUninstallFilter`

NewOutputUninstallFilter instantiates a new OutputUninstallFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputUninstallFilterWithDefaults

`func NewOutputUninstallFilterWithDefaults() *OutputUninstallFilter`

NewOutputUninstallFilterWithDefaults instantiates a new OutputUninstallFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputUninstallFilter) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputUninstallFilter) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputUninstallFilter) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputUninstallFilter) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetResult

`func (o *OutputUninstallFilter) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OutputUninstallFilter) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OutputUninstallFilter) SetResult(v bool)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


