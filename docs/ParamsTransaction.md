# ParamsTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | The hash of the method signature and encoded parameters. For more information, see the Contract ABI description in the Solidity documentation. | [optional] 
**From** | Pointer to **string** | The address from which the transaction is sent | [optional] 
**Gas** | Pointer to **string** | The integer of gas provided for the transaction execution | [optional] 
**GasPrice** | Pointer to **string** | The integer of gasPrice used for each paid gas encoded as hexadecimal | [optional] 
**To** | **string** | The address to which the transaction is addressed | 
**Value** | Pointer to **string** | The integer of value sent with this transaction encoded as hexadecimal | [optional] 

## Methods

### NewParamsTransaction

`func NewParamsTransaction(to string, ) *ParamsTransaction`

NewParamsTransaction instantiates a new ParamsTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParamsTransactionWithDefaults

`func NewParamsTransactionWithDefaults() *ParamsTransaction`

NewParamsTransactionWithDefaults instantiates a new ParamsTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ParamsTransaction) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ParamsTransaction) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ParamsTransaction) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ParamsTransaction) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFrom

`func (o *ParamsTransaction) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ParamsTransaction) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ParamsTransaction) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ParamsTransaction) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGas

`func (o *ParamsTransaction) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *ParamsTransaction) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *ParamsTransaction) SetGas(v string)`

SetGas sets Gas field to given value.

### HasGas

`func (o *ParamsTransaction) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *ParamsTransaction) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ParamsTransaction) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ParamsTransaction) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *ParamsTransaction) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetTo

`func (o *ParamsTransaction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ParamsTransaction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ParamsTransaction) SetTo(v string)`

SetTo sets To field to given value.


### GetValue

`func (o *ParamsTransaction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ParamsTransaction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ParamsTransaction) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ParamsTransaction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


