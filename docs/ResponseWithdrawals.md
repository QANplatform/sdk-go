# ResponseWithdrawals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address to which the withdrawn amount is sent | 
**Amount** | **string** | The amount of value, provided in decimal format, corresponding to a certain value in gwei | 
**Index** | **int32** | The index of the withdrawal to uniquely identify each withdrawal | 
**ValidatorIndex** | **int32** | The index of the validator who initiated the withdrawal | 

## Methods

### NewResponseWithdrawals

`func NewResponseWithdrawals(address string, amount string, index int32, validatorIndex int32, ) *ResponseWithdrawals`

NewResponseWithdrawals instantiates a new ResponseWithdrawals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithdrawalsWithDefaults

`func NewResponseWithdrawalsWithDefaults() *ResponseWithdrawals`

NewResponseWithdrawalsWithDefaults instantiates a new ResponseWithdrawals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ResponseWithdrawals) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ResponseWithdrawals) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ResponseWithdrawals) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *ResponseWithdrawals) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ResponseWithdrawals) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ResponseWithdrawals) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIndex

`func (o *ResponseWithdrawals) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ResponseWithdrawals) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ResponseWithdrawals) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetValidatorIndex

`func (o *ResponseWithdrawals) GetValidatorIndex() int32`

GetValidatorIndex returns the ValidatorIndex field if non-nil, zero value otherwise.

### GetValidatorIndexOk

`func (o *ResponseWithdrawals) GetValidatorIndexOk() (*int32, bool)`

GetValidatorIndexOk returns a tuple with the ValidatorIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorIndex

`func (o *ResponseWithdrawals) SetValidatorIndex(v int32)`

SetValidatorIndex sets ValidatorIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


