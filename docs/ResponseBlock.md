# ResponseBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseFeePerGas** | **NullableString** | A string of the base fee encoded in decimal format. Please note that this response field will not be included in a block requested before the EIP-1559 upgrade | 
**Difficulty** | **string** | The integer of the difficulty for this block encoded as a decimal | 
**ExtraData** | **string** | The “extra data” field of this block | 
**GasLimit** | **string** | The maximum gas allowed in this block encoded as a decimal | 
**GasUsed** | **string** | The total used gas by all transactions in this block encoded as a decimal | 
**Hash** | **NullableString** | The block hash of the requested block. null if pending | 
**LogsBloom** | **NullableString** | The bloom filter for the logs of the block. null if pending | 
**Miner** | **string** | The address of the beneficiary to whom the mining rewards were given | 
**MixHash** | **string** | A string of a 256-bit hash encoded as a decimal | 
**Nonce** | **NullableString** | The hash of the generated proof-of-work. null if pending | 
**Number** | **NullableString** | The block number of the requested block encoded as a decimal. null if pending | 
**ParentHash** | **string** | The hash of the parent block | 
**ReceiptsRoot** | **string** | The root of the receipts trie of the bloc | 
**Sha3Uncles** | **string** | The SHA3 of the uncles data in the block | 
**Size** | **string** | The size of this block in bytes as an Integer value encoded as decimal | 
**StateRoot** | **string** | The root of the final state trie of the block | 
**Timestamp** | **string** | The unix timestamp for when the block was collated | 
**TotalDifficulty** | **string** | The integer of the total difficulty of the chain until this block encoded as a decimal | 
**Transactions** | [**[]ResponseTransaction**](ResponseTransaction.md) | An array of transaction objects - please see getTransactionByHash for exact shape | 
**TransactionsRoot** | **string** | The root of the transaction trie of the block | 
**Uncles** | **[]string** | An array of uncle hashes | 
**Withdrawals** | [**ResponseWithdrawals**](ResponseWithdrawals.md) | A withdrawals object contains information about withdrawals made by validators. Please note that this field will not be included in a block requested before the EIP-4895 upgrade | 
**WithdrawalsRoot** | **NullableString** | The Merkle root of withdrawal data. Also, please note that this field will not be included in a block requested before the EIP-4895 upgrade | 

## Methods

### NewResponseBlock

`func NewResponseBlock(baseFeePerGas NullableString, difficulty string, extraData string, gasLimit string, gasUsed string, hash NullableString, logsBloom NullableString, miner string, mixHash string, nonce NullableString, number NullableString, parentHash string, receiptsRoot string, sha3Uncles string, size string, stateRoot string, timestamp string, totalDifficulty string, transactions []ResponseTransaction, transactionsRoot string, uncles []string, withdrawals ResponseWithdrawals, withdrawalsRoot NullableString, ) *ResponseBlock`

NewResponseBlock instantiates a new ResponseBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBlockWithDefaults

`func NewResponseBlockWithDefaults() *ResponseBlock`

NewResponseBlockWithDefaults instantiates a new ResponseBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseFeePerGas

`func (o *ResponseBlock) GetBaseFeePerGas() string`

GetBaseFeePerGas returns the BaseFeePerGas field if non-nil, zero value otherwise.

### GetBaseFeePerGasOk

`func (o *ResponseBlock) GetBaseFeePerGasOk() (*string, bool)`

GetBaseFeePerGasOk returns a tuple with the BaseFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFeePerGas

`func (o *ResponseBlock) SetBaseFeePerGas(v string)`

SetBaseFeePerGas sets BaseFeePerGas field to given value.


### SetBaseFeePerGasNil

`func (o *ResponseBlock) SetBaseFeePerGasNil(b bool)`

 SetBaseFeePerGasNil sets the value for BaseFeePerGas to be an explicit nil

### UnsetBaseFeePerGas
`func (o *ResponseBlock) UnsetBaseFeePerGas()`

UnsetBaseFeePerGas ensures that no value is present for BaseFeePerGas, not even an explicit nil
### GetDifficulty

`func (o *ResponseBlock) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *ResponseBlock) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *ResponseBlock) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetExtraData

`func (o *ResponseBlock) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *ResponseBlock) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *ResponseBlock) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.


### GetGasLimit

`func (o *ResponseBlock) GetGasLimit() string`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ResponseBlock) GetGasLimitOk() (*string, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ResponseBlock) SetGasLimit(v string)`

SetGasLimit sets GasLimit field to given value.


### GetGasUsed

`func (o *ResponseBlock) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ResponseBlock) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ResponseBlock) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetHash

`func (o *ResponseBlock) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ResponseBlock) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ResponseBlock) SetHash(v string)`

SetHash sets Hash field to given value.


### SetHashNil

`func (o *ResponseBlock) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *ResponseBlock) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetLogsBloom

`func (o *ResponseBlock) GetLogsBloom() string`

GetLogsBloom returns the LogsBloom field if non-nil, zero value otherwise.

### GetLogsBloomOk

`func (o *ResponseBlock) GetLogsBloomOk() (*string, bool)`

GetLogsBloomOk returns a tuple with the LogsBloom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsBloom

`func (o *ResponseBlock) SetLogsBloom(v string)`

SetLogsBloom sets LogsBloom field to given value.


### SetLogsBloomNil

`func (o *ResponseBlock) SetLogsBloomNil(b bool)`

 SetLogsBloomNil sets the value for LogsBloom to be an explicit nil

### UnsetLogsBloom
`func (o *ResponseBlock) UnsetLogsBloom()`

UnsetLogsBloom ensures that no value is present for LogsBloom, not even an explicit nil
### GetMiner

`func (o *ResponseBlock) GetMiner() string`

GetMiner returns the Miner field if non-nil, zero value otherwise.

### GetMinerOk

`func (o *ResponseBlock) GetMinerOk() (*string, bool)`

GetMinerOk returns a tuple with the Miner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiner

`func (o *ResponseBlock) SetMiner(v string)`

SetMiner sets Miner field to given value.


### GetMixHash

`func (o *ResponseBlock) GetMixHash() string`

GetMixHash returns the MixHash field if non-nil, zero value otherwise.

### GetMixHashOk

`func (o *ResponseBlock) GetMixHashOk() (*string, bool)`

GetMixHashOk returns a tuple with the MixHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixHash

`func (o *ResponseBlock) SetMixHash(v string)`

SetMixHash sets MixHash field to given value.


### GetNonce

`func (o *ResponseBlock) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ResponseBlock) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ResponseBlock) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### SetNonceNil

`func (o *ResponseBlock) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *ResponseBlock) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
### GetNumber

`func (o *ResponseBlock) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ResponseBlock) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ResponseBlock) SetNumber(v string)`

SetNumber sets Number field to given value.


### SetNumberNil

`func (o *ResponseBlock) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *ResponseBlock) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetParentHash

`func (o *ResponseBlock) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *ResponseBlock) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *ResponseBlock) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### GetReceiptsRoot

`func (o *ResponseBlock) GetReceiptsRoot() string`

GetReceiptsRoot returns the ReceiptsRoot field if non-nil, zero value otherwise.

### GetReceiptsRootOk

`func (o *ResponseBlock) GetReceiptsRootOk() (*string, bool)`

GetReceiptsRootOk returns a tuple with the ReceiptsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptsRoot

`func (o *ResponseBlock) SetReceiptsRoot(v string)`

SetReceiptsRoot sets ReceiptsRoot field to given value.


### GetSha3Uncles

`func (o *ResponseBlock) GetSha3Uncles() string`

GetSha3Uncles returns the Sha3Uncles field if non-nil, zero value otherwise.

### GetSha3UnclesOk

`func (o *ResponseBlock) GetSha3UnclesOk() (*string, bool)`

GetSha3UnclesOk returns a tuple with the Sha3Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha3Uncles

`func (o *ResponseBlock) SetSha3Uncles(v string)`

SetSha3Uncles sets Sha3Uncles field to given value.


### GetSize

`func (o *ResponseBlock) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponseBlock) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponseBlock) SetSize(v string)`

SetSize sets Size field to given value.


### GetStateRoot

`func (o *ResponseBlock) GetStateRoot() string`

GetStateRoot returns the StateRoot field if non-nil, zero value otherwise.

### GetStateRootOk

`func (o *ResponseBlock) GetStateRootOk() (*string, bool)`

GetStateRootOk returns a tuple with the StateRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRoot

`func (o *ResponseBlock) SetStateRoot(v string)`

SetStateRoot sets StateRoot field to given value.


### GetTimestamp

`func (o *ResponseBlock) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponseBlock) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponseBlock) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetTotalDifficulty

`func (o *ResponseBlock) GetTotalDifficulty() string`

GetTotalDifficulty returns the TotalDifficulty field if non-nil, zero value otherwise.

### GetTotalDifficultyOk

`func (o *ResponseBlock) GetTotalDifficultyOk() (*string, bool)`

GetTotalDifficultyOk returns a tuple with the TotalDifficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDifficulty

`func (o *ResponseBlock) SetTotalDifficulty(v string)`

SetTotalDifficulty sets TotalDifficulty field to given value.


### GetTransactions

`func (o *ResponseBlock) GetTransactions() []ResponseTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ResponseBlock) GetTransactionsOk() (*[]ResponseTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ResponseBlock) SetTransactions(v []ResponseTransaction)`

SetTransactions sets Transactions field to given value.


### GetTransactionsRoot

`func (o *ResponseBlock) GetTransactionsRoot() string`

GetTransactionsRoot returns the TransactionsRoot field if non-nil, zero value otherwise.

### GetTransactionsRootOk

`func (o *ResponseBlock) GetTransactionsRootOk() (*string, bool)`

GetTransactionsRootOk returns a tuple with the TransactionsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsRoot

`func (o *ResponseBlock) SetTransactionsRoot(v string)`

SetTransactionsRoot sets TransactionsRoot field to given value.


### GetUncles

`func (o *ResponseBlock) GetUncles() []string`

GetUncles returns the Uncles field if non-nil, zero value otherwise.

### GetUnclesOk

`func (o *ResponseBlock) GetUnclesOk() (*[]string, bool)`

GetUnclesOk returns a tuple with the Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncles

`func (o *ResponseBlock) SetUncles(v []string)`

SetUncles sets Uncles field to given value.


### GetWithdrawals

`func (o *ResponseBlock) GetWithdrawals() ResponseWithdrawals`

GetWithdrawals returns the Withdrawals field if non-nil, zero value otherwise.

### GetWithdrawalsOk

`func (o *ResponseBlock) GetWithdrawalsOk() (*ResponseWithdrawals, bool)`

GetWithdrawalsOk returns a tuple with the Withdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawals

`func (o *ResponseBlock) SetWithdrawals(v ResponseWithdrawals)`

SetWithdrawals sets Withdrawals field to given value.


### GetWithdrawalsRoot

`func (o *ResponseBlock) GetWithdrawalsRoot() string`

GetWithdrawalsRoot returns the WithdrawalsRoot field if non-nil, zero value otherwise.

### GetWithdrawalsRootOk

`func (o *ResponseBlock) GetWithdrawalsRootOk() (*string, bool)`

GetWithdrawalsRootOk returns a tuple with the WithdrawalsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalsRoot

`func (o *ResponseBlock) SetWithdrawalsRoot(v string)`

SetWithdrawalsRoot sets WithdrawalsRoot field to given value.


### SetWithdrawalsRootNil

`func (o *ResponseBlock) SetWithdrawalsRootNil(b bool)`

 SetWithdrawalsRootNil sets the value for WithdrawalsRoot to be an explicit nil

### UnsetWithdrawalsRoot
`func (o *ResponseBlock) UnsetWithdrawalsRoot()`

UnsetWithdrawalsRoot ensures that no value is present for WithdrawalsRoot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


