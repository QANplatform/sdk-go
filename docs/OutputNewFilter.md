# OutputNewFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**FilterId** | **string** | A filter id to be used when calling getFilterChanges | 

## Methods

### NewOutputNewFilter

`func NewOutputNewFilter(filterId string, ) *OutputNewFilter`

NewOutputNewFilter instantiates a new OutputNewFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputNewFilterWithDefaults

`func NewOutputNewFilterWithDefaults() *OutputNewFilter`

NewOutputNewFilterWithDefaults instantiates a new OutputNewFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputNewFilter) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputNewFilter) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputNewFilter) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputNewFilter) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetFilterId

`func (o *OutputNewFilter) GetFilterId() string`

GetFilterId returns the FilterId field if non-nil, zero value otherwise.

### GetFilterIdOk

`func (o *OutputNewFilter) GetFilterIdOk() (*string, bool)`

GetFilterIdOk returns a tuple with the FilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterId

`func (o *OutputNewFilter) SetFilterId(v string)`

SetFilterId sets FilterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


