# QAN Golang SDK

This repository is guaranteed up-to-date with the upstream QAN API definitions, and leverages OpenAPI technology to stay consistent.

Versioning is based on SEMVER, meaning:

- Stable releases guarantee backwards compatibility for the same major versions.
- Minor releases will not contain breaking changes.
- Patch releases only focus on fixing issues.

## Documentation for API Endpoints

All URIs are relative to *https://rpc-testnet.qanplatform.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QanBlockNumber**](docs/QANAPI.md#QanBlockNumber) | **Get** /blockNumber/ | Returns the latest block number of the blockchain.
[**QanCall**](docs/QANAPI.md#QanCall) | **Post** /call/ | Executes a new message call immediately without creating a transaction on the block chain.
[**QanChainId**](docs/QANAPI.md#QanChainId) | **Get** /chainId/ | Returns the current network/chain ID, used to sign replay-protected transaction introduced in EIP-155.
[**QanEstimateGas**](docs/QANAPI.md#QanEstimateGas) | **Post** /estimateGas/ | Returns an estimation of gas for a given transaction.
[**QanFeeHistory**](docs/QANAPI.md#QanFeeHistory) | **Post** /feeHistory/ | Returns the collection of historical gas information.
[**QanGasPrice**](docs/QANAPI.md#QanGasPrice) | **Get** /gasPrice/ | Returns the current gas price on the network in wei.
[**QanGetBalance**](docs/QANAPI.md#QanGetBalance) | **Get** /getBalance/{Address}/ | Returns the balance of the account of given address.
[**QanGetBlockByHash**](docs/QANAPI.md#QanGetBlockByHash) | **Get** /getBlockByHash/{Hash}/{TransactionDetailFlag}/ | Returns information of the block matching the given block hash.
[**QanGetBlockByNumber**](docs/QANAPI.md#QanGetBlockByNumber) | **Get** /getBlockByNumber/{BlockNumber}/{TransactionDetailFlag}/ | Returns information of the block matching the given block number.
[**QanGetBlockReceipts**](docs/QANAPI.md#QanGetBlockReceipts) | **Get** /getBlockReceipts/{BlockNumber}/ | Returns all transaction receipts for a given block.
[**QanGetBlockTransactionCountByHash**](docs/QANAPI.md#QanGetBlockTransactionCountByHash) | **Get** /getBlockTransactionCountByHash/{Hash}/ | Returns the number of transactions for the block matching the given block hash.
[**QanGetBlockTransactionCountByNumber**](docs/QANAPI.md#QanGetBlockTransactionCountByNumber) | **Get** /getBlockTransactionCountByNumber/{BlockNumber}/ | Returns the number of transactions for the block matching the given block number.
[**QanGetCode**](docs/QANAPI.md#QanGetCode) | **Get** /getCode/{Address}/ | Returns the compiled bytecode of a smart contract.
[**QanGetFilterChanges**](docs/QANAPI.md#QanGetFilterChanges) | **Get** /getFilterChanges/{FilterId}/ | Polling method for a filter, which returns an array of events that have occurred since the last poll.
[**QanGetFilterLogs**](docs/QANAPI.md#QanGetFilterLogs) | **Get** /getFilterLogs/{Id}/ | Returns an array of all logs matching filter with given id.
[**QanGetLogs**](docs/QANAPI.md#QanGetLogs) | **Post** /getLogs/ | Returns an array of all logs matching a given filter object.
[**QanGetProof**](docs/QANAPI.md#QanGetProof) | **Post** /getProof/ | Returns the account and storage values of the specified account including the Merkle-proof.
[**QanGetStorageAt**](docs/QANAPI.md#QanGetStorageAt) | **Post** /getStorageAt/ | Returns the value from a storage position at a given address.
[**QanGetTransactionByBlockHashAndIndex**](docs/QANAPI.md#QanGetTransactionByBlockHashAndIndex) | **Get** /getTransactionByBlockHashAndIndex/{blockHash}/{index}/ | Returns information about a transaction given a blockhash and transaction index position.
[**QanGetTransactionByBlockNumberAndIndex**](docs/QANAPI.md#QanGetTransactionByBlockNumberAndIndex) | **Get** /getTransactionByBlockNumberAndIndex/{blockNumber}/{index}/ | Returns information about a transaction given a block number and transaction index position.
[**QanGetTransactionByHash**](docs/QANAPI.md#QanGetTransactionByHash) | **Get** /getTransactionByHash/{hash}/ | Returns the information about a transaction from a transaction hash.
[**QanGetTransactionCount**](docs/QANAPI.md#QanGetTransactionCount) | **Get** /getTransactionCount/{Address}/{BlockNumber}/ | Returns the number of transactions sent from an address.
[**QanGetTransactionReceipt**](docs/QANAPI.md#QanGetTransactionReceipt) | **Get** /getTransactionReceipt/{Hash}/ | Returns the receipt of a transaction by transaction hash.
[**QanMaxPriorityFeePerGas**](docs/QANAPI.md#QanMaxPriorityFeePerGas) | **Get** /maxPriorityFeePerGas/ | Get the priority fee needed to be included in a block.
[**QanNewBlockFilter**](docs/QANAPI.md#QanNewBlockFilter) | **Get** /newBlockFilter/ | Creates a filter in the node, to notify when a new block arrives.
[**QanNewFilter**](docs/QANAPI.md#QanNewFilter) | **Post** /newFilter/ | Creates a filter object, based on filter options, to notify when the state changes (logs).
[**QanNewPendingTransactionFilter**](docs/QANAPI.md#QanNewPendingTransactionFilter) | **Get** /newPendingTransactionFilter/ | Creates a filter in the node to notify when new pending transactions arrive.
[**QanSendRawTransaction**](docs/QANAPI.md#QanSendRawTransaction) | **Post** /sendRawTransaction/ | Creates new message call transaction or a contract creation for signed transactions.
[**QanSyncing**](docs/QANAPI.md#QanSyncing) | **Get** /syncing/ | Returns an object with the sync status of the node if the node is out-of-sync and is syncing. Returns null when the node is already in sync.
[**QanUninstallFilter**](docs/QANAPI.md#QanUninstallFilter) | **Get** /uninstallFilter/{FilterId}/ | Uninstalls a filter with the given filter id.
[**QanXlinkValid**](docs/QANAPI.md#QanXlinkValid) | **Get** /xlinkValid/{Address}/ | Returns the xlink validity time of the account of given address.


## Documentation For Models

 - [ErrorDetail](docs/ErrorDetail.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [EstimateGasObject](docs/EstimateGasObject.md)
 - [FilterObject](docs/FilterObject.md)
 - [InputCall](docs/InputCall.md)
 - [InputEstimateGas](docs/InputEstimateGas.md)
 - [InputFeeHistory](docs/InputFeeHistory.md)
 - [InputGetLogs](docs/InputGetLogs.md)
 - [InputGetProof](docs/InputGetProof.md)
 - [InputGetStorageAt](docs/InputGetStorageAt.md)
 - [InputNewFilter](docs/InputNewFilter.md)
 - [InputSendRawTransaction](docs/InputSendRawTransaction.md)
 - [InputSubscribe](docs/InputSubscribe.md)
 - [OutputAccounts](docs/OutputAccounts.md)
 - [OutputBlobBaseFee](docs/OutputBlobBaseFee.md)
 - [OutputBlockNumber](docs/OutputBlockNumber.md)
 - [OutputCall](docs/OutputCall.md)
 - [OutputChainId](docs/OutputChainId.md)
 - [OutputEstimateGas](docs/OutputEstimateGas.md)
 - [OutputFeeHistory](docs/OutputFeeHistory.md)
 - [OutputGasPrice](docs/OutputGasPrice.md)
 - [OutputGetAccount](docs/OutputGetAccount.md)
 - [OutputGetBalance](docs/OutputGetBalance.md)
 - [OutputGetBlockByHash](docs/OutputGetBlockByHash.md)
 - [OutputGetBlockByNumber](docs/OutputGetBlockByNumber.md)
 - [OutputGetBlockReceipts](docs/OutputGetBlockReceipts.md)
 - [OutputGetBlockTransactionCountByHash](docs/OutputGetBlockTransactionCountByHash.md)
 - [OutputGetBlockTransactionCountByNumber](docs/OutputGetBlockTransactionCountByNumber.md)
 - [OutputGetCode](docs/OutputGetCode.md)
 - [OutputGetFilterChanges](docs/OutputGetFilterChanges.md)
 - [OutputGetFilterLogs](docs/OutputGetFilterLogs.md)
 - [OutputGetLogs](docs/OutputGetLogs.md)
 - [OutputGetProof](docs/OutputGetProof.md)
 - [OutputGetStorageAt](docs/OutputGetStorageAt.md)
 - [OutputGetTransactionByBlockHashAndIndex](docs/OutputGetTransactionByBlockHashAndIndex.md)
 - [OutputGetTransactionByBlockNumberAndIndex](docs/OutputGetTransactionByBlockNumberAndIndex.md)
 - [OutputGetTransactionByHash](docs/OutputGetTransactionByHash.md)
 - [OutputGetTransactionCount](docs/OutputGetTransactionCount.md)
 - [OutputGetTransactionReceipt](docs/OutputGetTransactionReceipt.md)
 - [OutputGetUncleCountByBlockHash](docs/OutputGetUncleCountByBlockHash.md)
 - [OutputGetUncleCountByBlockNumber](docs/OutputGetUncleCountByBlockNumber.md)
 - [OutputMaxPriorityFeePerGas](docs/OutputMaxPriorityFeePerGas.md)
 - [OutputNewBlockFilter](docs/OutputNewBlockFilter.md)
 - [OutputNewFilter](docs/OutputNewFilter.md)
 - [OutputNewPendingTransactionFilter](docs/OutputNewPendingTransactionFilter.md)
 - [OutputSendRawTransaction](docs/OutputSendRawTransaction.md)
 - [OutputSubscribe](docs/OutputSubscribe.md)
 - [OutputSyncing](docs/OutputSyncing.md)
 - [OutputUninstallFilter](docs/OutputUninstallFilter.md)
 - [OutputUnsubscribe](docs/OutputUnsubscribe.md)
 - [ParamsTransaction](docs/ParamsTransaction.md)
 - [QanXlinkValidResponse](docs/QanXlinkValidResponse.md)
 - [ResponseBlock](docs/ResponseBlock.md)
 - [ResponseLog](docs/ResponseLog.md)
 - [ResponseStorageEntry](docs/ResponseStorageEntry.md)
 - [ResponseTransaction](docs/ResponseTransaction.md)
 - [ResponseTransactionReceipt](docs/ResponseTransactionReceipt.md)
 - [ResponseWithdrawals](docs/ResponseWithdrawals.md)
 - [SyncStatus](docs/SyncStatus.md)

## Acknowledgements

We would like to thank Smartbear and OpenAPITools tech for making building declarative APIs possible.
A huge benefit for the whole industry!
