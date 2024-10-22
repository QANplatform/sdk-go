# FilterObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The contract address or a list of addresses from which logs should originate | 
**FromBlock** | **string** |  | 
**ToBlock** | **string** |  | 
**Topics** | **[]string** |  | 

## Methods

### NewFilterObject

`func NewFilterObject(address string, fromBlock string, toBlock string, topics []string, ) *FilterObject`

NewFilterObject instantiates a new FilterObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterObjectWithDefaults

`func NewFilterObjectWithDefaults() *FilterObject`

NewFilterObjectWithDefaults instantiates a new FilterObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FilterObject) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FilterObject) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FilterObject) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetFromBlock

`func (o *FilterObject) GetFromBlock() string`

GetFromBlock returns the FromBlock field if non-nil, zero value otherwise.

### GetFromBlockOk

`func (o *FilterObject) GetFromBlockOk() (*string, bool)`

GetFromBlockOk returns a tuple with the FromBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromBlock

`func (o *FilterObject) SetFromBlock(v string)`

SetFromBlock sets FromBlock field to given value.


### GetToBlock

`func (o *FilterObject) GetToBlock() string`

GetToBlock returns the ToBlock field if non-nil, zero value otherwise.

### GetToBlockOk

`func (o *FilterObject) GetToBlockOk() (*string, bool)`

GetToBlockOk returns a tuple with the ToBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBlock

`func (o *FilterObject) SetToBlock(v string)`

SetToBlock sets ToBlock field to given value.


### GetTopics

`func (o *FilterObject) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *FilterObject) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *FilterObject) SetTopics(v []string)`

SetTopics sets Topics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


