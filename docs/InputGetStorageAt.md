# InputGetStorageAt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Address** | **string** | The address to check for storage | 
**BlockNumber** | **string** |  | 
**Position** | **string** | The integer of the position in storage | 

## Methods

### NewInputGetStorageAt

`func NewInputGetStorageAt(address string, blockNumber string, position string, ) *InputGetStorageAt`

NewInputGetStorageAt instantiates a new InputGetStorageAt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputGetStorageAtWithDefaults

`func NewInputGetStorageAtWithDefaults() *InputGetStorageAt`

NewInputGetStorageAtWithDefaults instantiates a new InputGetStorageAt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *InputGetStorageAt) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *InputGetStorageAt) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *InputGetStorageAt) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *InputGetStorageAt) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAddress

`func (o *InputGetStorageAt) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InputGetStorageAt) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InputGetStorageAt) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBlockNumber

`func (o *InputGetStorageAt) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *InputGetStorageAt) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *InputGetStorageAt) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.


### GetPosition

`func (o *InputGetStorageAt) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *InputGetStorageAt) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *InputGetStorageAt) SetPosition(v string)`

SetPosition sets Position field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


