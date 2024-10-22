# InputNewFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**FilterObject** | [**FilterObject**](FilterObject.md) |  | 

## Methods

### NewInputNewFilter

`func NewInputNewFilter(filterObject FilterObject, ) *InputNewFilter`

NewInputNewFilter instantiates a new InputNewFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputNewFilterWithDefaults

`func NewInputNewFilterWithDefaults() *InputNewFilter`

NewInputNewFilterWithDefaults instantiates a new InputNewFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *InputNewFilter) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *InputNewFilter) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *InputNewFilter) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *InputNewFilter) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetFilterObject

`func (o *InputNewFilter) GetFilterObject() FilterObject`

GetFilterObject returns the FilterObject field if non-nil, zero value otherwise.

### GetFilterObjectOk

`func (o *InputNewFilter) GetFilterObjectOk() (*FilterObject, bool)`

GetFilterObjectOk returns a tuple with the FilterObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterObject

`func (o *InputNewFilter) SetFilterObject(v FilterObject)`

SetFilterObject sets FilterObject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


