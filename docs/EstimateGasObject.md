# EstimateGasObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **string** | Fake balance to set for the account before executing the call | [optional] 
**Code** | Pointer to **int32** | Fake EVM bytecode to inject into the account before executing the call | [optional] 
**Nonce** | Pointer to **string** | Fake nonce to set for the account before executing the call | [optional] 
**State** | Pointer to **string** | Fake key-value mapping to override all slots in the account storage before executing the call | [optional] 
**StateDiff** | Pointer to **string** | Fake key-value mapping to override individual slots in the account storage before executing the call | [optional] 

## Methods

### NewEstimateGasObject

`func NewEstimateGasObject() *EstimateGasObject`

NewEstimateGasObject instantiates a new EstimateGasObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateGasObjectWithDefaults

`func NewEstimateGasObjectWithDefaults() *EstimateGasObject`

NewEstimateGasObjectWithDefaults instantiates a new EstimateGasObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *EstimateGasObject) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *EstimateGasObject) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *EstimateGasObject) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *EstimateGasObject) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCode

`func (o *EstimateGasObject) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EstimateGasObject) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EstimateGasObject) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *EstimateGasObject) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetNonce

`func (o *EstimateGasObject) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *EstimateGasObject) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *EstimateGasObject) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *EstimateGasObject) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetState

`func (o *EstimateGasObject) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EstimateGasObject) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EstimateGasObject) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EstimateGasObject) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateDiff

`func (o *EstimateGasObject) GetStateDiff() string`

GetStateDiff returns the StateDiff field if non-nil, zero value otherwise.

### GetStateDiffOk

`func (o *EstimateGasObject) GetStateDiffOk() (*string, bool)`

GetStateDiffOk returns a tuple with the StateDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDiff

`func (o *EstimateGasObject) SetStateDiff(v string)`

SetStateDiff sets StateDiff field to given value.

### HasStateDiff

`func (o *EstimateGasObject) HasStateDiff() bool`

HasStateDiff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


