# ResponseTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessList** | Pointer to **string** | A list of addresses and storage keys that the transaction plans to access | [optional] 
**BlockHash** | Pointer to **string** | The hash of the block where this transaction was in. Null when it&#39;s a pending transaction | [optional] 
**BlockNumber** | Pointer to **string** | The block number where this transaction was in. Null when it&#39;s a pending transaction | [optional] 
**ChainId** | Pointer to **string** | The chain id of the transaction, if any | [optional] 
**From** | Pointer to **string** | The address of the sender | [optional] 
**Gas** | Pointer to **string** | The gas provided by the sender, encoded as decimal | [optional] 
**GasPrice** | Pointer to **string** | The gas price provided by the sender in wei encoded as decimal | [optional] 
**Hash** | Pointer to **string** | The hash of the transaction | [optional] 
**Input** | Pointer to **string** | The data sent along with the transaction | [optional] 
**MaxFeePerGas** | Pointer to **string** | The maximum fee per gas set in the transaction | [optional] 
**MaxPriorityFeePerGas** | Pointer to **string** | The maximum priority gas fee set in the transaction | [optional] 
**Nonce** | Pointer to **string** | The number of transactions made by the sender prior to this one encoded as decimal | [optional] 
**R** | Pointer to **string** | The R field of the signature | [optional] 
**S** | Pointer to **string** | The S field of the signature | [optional] 
**To** | Pointer to **string** | The address of the receiver. Null when its a contract creation transaction | [optional] 
**TransactionIndex** | Pointer to **string** | The integer of the transaction&#39;s index position that the log was created from. Null when it&#39;s a pending log | [optional] 
**Type** | Pointer to **string** | The transaction type | [optional] 
**V** | Pointer to **string** | The standardized V field of the signature | [optional] 
**Value** | Pointer to **string** | The value transferred in wei encoded as decimal | [optional] 

## Methods

### NewResponseTransaction

`func NewResponseTransaction() *ResponseTransaction`

NewResponseTransaction instantiates a new ResponseTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTransactionWithDefaults

`func NewResponseTransactionWithDefaults() *ResponseTransaction`

NewResponseTransactionWithDefaults instantiates a new ResponseTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessList

`func (o *ResponseTransaction) GetAccessList() string`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *ResponseTransaction) GetAccessListOk() (*string, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *ResponseTransaction) SetAccessList(v string)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *ResponseTransaction) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetBlockHash

`func (o *ResponseTransaction) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ResponseTransaction) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ResponseTransaction) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *ResponseTransaction) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBlockNumber

`func (o *ResponseTransaction) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *ResponseTransaction) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *ResponseTransaction) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *ResponseTransaction) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetChainId

`func (o *ResponseTransaction) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ResponseTransaction) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ResponseTransaction) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ResponseTransaction) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetFrom

`func (o *ResponseTransaction) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ResponseTransaction) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ResponseTransaction) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ResponseTransaction) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGas

`func (o *ResponseTransaction) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *ResponseTransaction) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *ResponseTransaction) SetGas(v string)`

SetGas sets Gas field to given value.

### HasGas

`func (o *ResponseTransaction) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *ResponseTransaction) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ResponseTransaction) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ResponseTransaction) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *ResponseTransaction) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetHash

`func (o *ResponseTransaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ResponseTransaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ResponseTransaction) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ResponseTransaction) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetInput

`func (o *ResponseTransaction) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ResponseTransaction) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ResponseTransaction) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *ResponseTransaction) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *ResponseTransaction) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *ResponseTransaction) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *ResponseTransaction) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *ResponseTransaction) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### GetMaxPriorityFeePerGas

`func (o *ResponseTransaction) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *ResponseTransaction) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *ResponseTransaction) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *ResponseTransaction) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.

### GetNonce

`func (o *ResponseTransaction) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ResponseTransaction) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ResponseTransaction) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *ResponseTransaction) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetR

`func (o *ResponseTransaction) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *ResponseTransaction) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *ResponseTransaction) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *ResponseTransaction) HasR() bool`

HasR returns a boolean if a field has been set.

### GetS

`func (o *ResponseTransaction) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *ResponseTransaction) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *ResponseTransaction) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *ResponseTransaction) HasS() bool`

HasS returns a boolean if a field has been set.

### GetTo

`func (o *ResponseTransaction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ResponseTransaction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ResponseTransaction) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ResponseTransaction) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTransactionIndex

`func (o *ResponseTransaction) GetTransactionIndex() string`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *ResponseTransaction) GetTransactionIndexOk() (*string, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *ResponseTransaction) SetTransactionIndex(v string)`

SetTransactionIndex sets TransactionIndex field to given value.

### HasTransactionIndex

`func (o *ResponseTransaction) HasTransactionIndex() bool`

HasTransactionIndex returns a boolean if a field has been set.

### GetType

`func (o *ResponseTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseTransaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetV

`func (o *ResponseTransaction) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ResponseTransaction) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ResponseTransaction) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *ResponseTransaction) HasV() bool`

HasV returns a boolean if a field has been set.

### GetValue

`func (o *ResponseTransaction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponseTransaction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponseTransaction) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResponseTransaction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


