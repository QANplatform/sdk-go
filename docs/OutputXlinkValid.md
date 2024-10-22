# OutputXlinkValid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Validity** | **string** | The xlink validity time of the specified address. | 

## Methods

### NewOutputXlinkValid

`func NewOutputXlinkValid(validity string, ) *OutputXlinkValid`

NewOutputXlinkValid instantiates a new OutputXlinkValid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputXlinkValidWithDefaults

`func NewOutputXlinkValidWithDefaults() *OutputXlinkValid`

NewOutputXlinkValidWithDefaults instantiates a new OutputXlinkValid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputXlinkValid) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputXlinkValid) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputXlinkValid) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputXlinkValid) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetValidity

`func (o *OutputXlinkValid) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *OutputXlinkValid) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *OutputXlinkValid) SetValidity(v string)`

SetValidity sets Validity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


