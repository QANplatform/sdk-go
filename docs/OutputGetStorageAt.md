# OutputGetStorageAt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Value** | **string** | The value from a storage position at a given address | 

## Methods

### NewOutputGetStorageAt

`func NewOutputGetStorageAt(value string, ) *OutputGetStorageAt`

NewOutputGetStorageAt instantiates a new OutputGetStorageAt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetStorageAtWithDefaults

`func NewOutputGetStorageAtWithDefaults() *OutputGetStorageAt`

NewOutputGetStorageAtWithDefaults instantiates a new OutputGetStorageAt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetStorageAt) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetStorageAt) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetStorageAt) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetStorageAt) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetValue

`func (o *OutputGetStorageAt) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OutputGetStorageAt) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OutputGetStorageAt) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


