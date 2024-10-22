# qan\QANAPI

All URIs are relative to *https://rpc-testnet.qanplatform.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QanBlockNumber**](QANAPI.md#QanBlockNumber) | **Get** /blockNumber/ | Returns the latest block number of the blockchain.
[**QanCall**](QANAPI.md#QanCall) | **Post** /call/ | Executes a new message call immediately without creating a transaction on the block chain.
[**QanChainId**](QANAPI.md#QanChainId) | **Get** /chainId/ | Returns the current network/chain ID, used to sign replay-protected transaction introduced in EIP-155.
[**QanEstimateGas**](QANAPI.md#QanEstimateGas) | **Post** /estimateGas/ | Returns an estimation of gas for a given transaction.
[**QanFeeHistory**](QANAPI.md#QanFeeHistory) | **Post** /feeHistory/ | Returns the collection of historical gas information.
[**QanGasPrice**](QANAPI.md#QanGasPrice) | **Get** /gasPrice/ | Returns the current gas price on the network in wei.
[**QanGetBalance**](QANAPI.md#QanGetBalance) | **Get** /getBalance/{Address}/ | Returns the balance of the account of given address.
[**QanGetBlockByHash**](QANAPI.md#QanGetBlockByHash) | **Get** /getBlockByHash/{Hash}/{TransactionDetailFlag}/ | Returns information of the block matching the given block hash.
[**QanGetBlockByNumber**](QANAPI.md#QanGetBlockByNumber) | **Get** /getBlockByNumber/{BlockNumber}/{TransactionDetailFlag}/ | Returns information of the block matching the given block number.
[**QanGetBlockReceipts**](QANAPI.md#QanGetBlockReceipts) | **Get** /getBlockReceipts/{BlockNumber}/ | Returns all transaction receipts for a given block.
[**QanGetBlockTransactionCountByHash**](QANAPI.md#QanGetBlockTransactionCountByHash) | **Get** /getBlockTransactionCountByHash/{Hash}/ | Returns the number of transactions for the block matching the given block hash.
[**QanGetBlockTransactionCountByNumber**](QANAPI.md#QanGetBlockTransactionCountByNumber) | **Get** /getBlockTransactionCountByNumber/{BlockNumber}/ | Returns the number of transactions for the block matching the given block number.
[**QanGetCode**](QANAPI.md#QanGetCode) | **Get** /getCode/{Address}/ | Returns the compiled bytecode of a smart contract.
[**QanGetFilterChanges**](QANAPI.md#QanGetFilterChanges) | **Get** /getFilterChanges/{FilterId}/ | Polling method for a filter, which returns an array of events that have occurred since the last poll.
[**QanGetFilterLogs**](QANAPI.md#QanGetFilterLogs) | **Get** /getFilterLogs/{Id}/ | Returns an array of all logs matching filter with given id.
[**QanGetLogs**](QANAPI.md#QanGetLogs) | **Post** /getLogs/ | Returns an array of all logs matching a given filter object.
[**QanGetProof**](QANAPI.md#QanGetProof) | **Post** /getProof/ | Returns the account and storage values of the specified account including the Merkle-proof.
[**QanGetStorageAt**](QANAPI.md#QanGetStorageAt) | **Post** /getStorageAt/ | Returns the value from a storage position at a given address.
[**QanGetTransactionByBlockHashAndIndex**](QANAPI.md#QanGetTransactionByBlockHashAndIndex) | **Get** /getTransactionByBlockHashAndIndex/{blockHash}/{index}/ | Returns information about a transaction given a blockhash and transaction index position.
[**QanGetTransactionByBlockNumberAndIndex**](QANAPI.md#QanGetTransactionByBlockNumberAndIndex) | **Get** /getTransactionByBlockNumberAndIndex/{blockNumber}/{index}/ | Returns information about a transaction given a block number and transaction index position.
[**QanGetTransactionByHash**](QANAPI.md#QanGetTransactionByHash) | **Get** /getTransactionByHash/{hash}/ | Returns the information about a transaction from a transaction hash.
[**QanGetTransactionCount**](QANAPI.md#QanGetTransactionCount) | **Get** /getTransactionCount/{Address}/{BlockNumber}/ | Returns the number of transactions sent from an address.
[**QanGetTransactionReceipt**](QANAPI.md#QanGetTransactionReceipt) | **Get** /getTransactionReceipt/{Hash}/ | Returns the receipt of a transaction by transaction hash.
[**QanMaxPriorityFeePerGas**](QANAPI.md#QanMaxPriorityFeePerGas) | **Get** /maxPriorityFeePerGas/ | Get the priority fee needed to be included in a block.
[**QanNewBlockFilter**](QANAPI.md#QanNewBlockFilter) | **Get** /newBlockFilter/ | Creates a filter in the node, to notify when a new block arrives.
[**QanNewFilter**](QANAPI.md#QanNewFilter) | **Post** /newFilter/ | Creates a filter object, based on filter options, to notify when the state changes (logs).
[**QanNewPendingTransactionFilter**](QANAPI.md#QanNewPendingTransactionFilter) | **Get** /newPendingTransactionFilter/ | Creates a filter in the node to notify when new pending transactions arrive.
[**QanSendRawTransaction**](QANAPI.md#QanSendRawTransaction) | **Post** /sendRawTransaction/ | Creates new message call transaction or a contract creation for signed transactions.
[**QanSyncing**](QANAPI.md#QanSyncing) | **Get** /syncing/ | Returns an object with the sync status of the node if the node is out-of-sync and is syncing. Returns null when the node is already in sync.
[**QanUninstallFilter**](QANAPI.md#QanUninstallFilter) | **Get** /uninstallFilter/{FilterId}/ | Uninstalls a filter with the given filter id.
[**QanXlinkValid**](QANAPI.md#QanXlinkValid) | **Get** /xlinkValid/{Address}/ | Returns the xlink validity time of the account of given address.



## QanBlockNumber

> OutputBlockNumber QanBlockNumber(ctx).Execute()

Returns the latest block number of the blockchain.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanBlockNumber(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanBlockNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanBlockNumber`: OutputBlockNumber
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanBlockNumber`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanBlockNumberRequest struct via the builder pattern


### Return type

[**OutputBlockNumber**](OutputBlockNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanCall

> OutputCall QanCall(ctx).InputCall(inputCall).Execute()

Executes a new message call immediately without creating a transaction on the block chain.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	inputCall := *openapiclient.NewInputCall("BlockNumber_example", *openapiclient.NewParamsTransaction("To_example")) // InputCall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanCall(context.Background()).InputCall(inputCall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanCall`: OutputCall
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputCall** | [**InputCall**](InputCall.md) |  | 

### Return type

[**OutputCall**](OutputCall.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanChainId

> OutputChainId QanChainId(ctx).Execute()

Returns the current network/chain ID, used to sign replay-protected transaction introduced in EIP-155.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanChainId(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanChainId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanChainId`: OutputChainId
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanChainId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanChainIdRequest struct via the builder pattern


### Return type

[**OutputChainId**](OutputChainId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanEstimateGas

> OutputEstimateGas QanEstimateGas(ctx).InputEstimateGas(inputEstimateGas).Execute()

Returns an estimation of gas for a given transaction.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	inputEstimateGas := *openapiclient.NewInputEstimateGas(*openapiclient.NewParamsTransaction("To_example")) // InputEstimateGas | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanEstimateGas(context.Background()).InputEstimateGas(inputEstimateGas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanEstimateGas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanEstimateGas`: OutputEstimateGas
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanEstimateGas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanEstimateGasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputEstimateGas** | [**InputEstimateGas**](InputEstimateGas.md) |  | 

### Return type

[**OutputEstimateGas**](OutputEstimateGas.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanFeeHistory

> OutputFeeHistory QanFeeHistory(ctx).InputFeeHistory(inputFeeHistory).Execute()

Returns the collection of historical gas information.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	inputFeeHistory := *openapiclient.NewInputFeeHistory(int32(123), "NewestBlock_example", []int32{int32(123)}) // InputFeeHistory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanFeeHistory(context.Background()).InputFeeHistory(inputFeeHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanFeeHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanFeeHistory`: OutputFeeHistory
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanFeeHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanFeeHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputFeeHistory** | [**InputFeeHistory**](InputFeeHistory.md) |  | 

### Return type

[**OutputFeeHistory**](OutputFeeHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGasPrice

> OutputGasPrice QanGasPrice(ctx).Execute()

Returns the current gas price on the network in wei.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGasPrice(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGasPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGasPrice`: OutputGasPrice
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGasPrice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanGasPriceRequest struct via the builder pattern


### Return type

[**OutputGasPrice**](OutputGasPrice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBalance

> OutputGetBalance QanGetBalance(ctx, address).BlockNumber(blockNumber).Execute()

Returns the balance of the account of given address.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	address := "0xa1e4380a3b1f749673e270229993ee55f35663b4" // string | A 20 bytes long hexadecimal value representing an address
	blockNumber := "blockNumber_example" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending (optional) (default to "latest")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetBalance(context.Background(), address).BlockNumber(blockNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBalance`: OutputGetBalance
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | A 20 bytes long hexadecimal value representing an address | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending | [default to &quot;latest&quot;]

### Return type

[**OutputGetBalance**](OutputGetBalance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBlockByHash

> OutputGetBlockByHash QanGetBlockByHash(ctx, hash, transactionDetailFlag).Execute()

Returns information of the block matching the given block hash.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	hash := "0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd" // string | The hash (32 bytes) of the block
	transactionDetailFlag := true // bool | The method returns the full transaction objects when this value is true otherwise, it returns only the hashes of the transactions (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetBlockByHash(context.Background(), hash, transactionDetailFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetBlockByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBlockByHash`: OutputGetBlockByHash
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetBlockByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | The hash (32 bytes) of the block | 
**transactionDetailFlag** | **bool** | The method returns the full transaction objects when this value is true otherwise, it returns only the hashes of the transactions | [default to false]

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBlockByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputGetBlockByHash**](OutputGetBlockByHash.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBlockByNumber

> OutputGetBlockByNumber QanGetBlockByNumber(ctx, blockNumber, transactionDetailFlag).Execute()

Returns information of the block matching the given block number.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	blockNumber := "blockNumber_example" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending (default to "latest")
	transactionDetailFlag := true // bool | The method returns the full transaction objects when this value is true otherwise, it returns only the hashes of the transactions (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetBlockByNumber(context.Background(), blockNumber, transactionDetailFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetBlockByNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBlockByNumber`: OutputGetBlockByNumber
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetBlockByNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending | [default to &quot;latest&quot;]
**transactionDetailFlag** | **bool** | The method returns the full transaction objects when this value is true otherwise, it returns only the hashes of the transactions | [default to false]

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBlockByNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputGetBlockByNumber**](OutputGetBlockByNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBlockReceipts

> OutputGetBlockReceipts QanGetBlockReceipts(ctx, blockNumber).Execute()

Returns all transaction receipts for a given block.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	blockNumber := "blockNumber_example" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending (default to "latest")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetBlockReceipts(context.Background(), blockNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetBlockReceipts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBlockReceipts`: OutputGetBlockReceipts
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetBlockReceipts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending | [default to &quot;latest&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBlockReceiptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetBlockReceipts**](OutputGetBlockReceipts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBlockTransactionCountByHash

> OutputGetBlockTransactionCountByHash QanGetBlockTransactionCountByHash(ctx, hash).Execute()

Returns the number of transactions for the block matching the given block hash.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	hash := "0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd" // string | The hash of the block

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetBlockTransactionCountByHash(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetBlockTransactionCountByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBlockTransactionCountByHash`: OutputGetBlockTransactionCountByHash
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetBlockTransactionCountByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | The hash of the block | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBlockTransactionCountByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetBlockTransactionCountByHash**](OutputGetBlockTransactionCountByHash.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetBlockTransactionCountByNumber

> OutputGetBlockTransactionCountByNumber QanGetBlockTransactionCountByNumber(ctx, blockNumber).Execute()

Returns the number of transactions for the block matching the given block number.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	blockNumber := "latest" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetBlockTransactionCountByNumber(context.Background(), blockNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetBlockTransactionCountByNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetBlockTransactionCountByNumber`: OutputGetBlockTransactionCountByNumber
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetBlockTransactionCountByNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetBlockTransactionCountByNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetBlockTransactionCountByNumber**](OutputGetBlockTransactionCountByNumber.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetCode

> OutputGetCode QanGetCode(ctx, address).BlockNumber(blockNumber).Execute()

Returns the compiled bytecode of a smart contract.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	address := "0xa1e4380a3b1f749673e270229993ee55f35663b4" // string | The address of the smart contract from which the bytecode will be obtained
	blockNumber := "latest" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending (optional) (default to "latest")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetCode(context.Background(), address).BlockNumber(blockNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetCode`: OutputGetCode
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | The address of the smart contract from which the bytecode will be obtained | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending | [default to &quot;latest&quot;]

### Return type

[**OutputGetCode**](OutputGetCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetFilterChanges

> OutputGetFilterChanges QanGetFilterChanges(ctx, filterId).Execute()

Polling method for a filter, which returns an array of events that have occurred since the last poll.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	filterId := "filterId_example" // string | The filter id that is returned from getFilterChangesnewFilter, getFilterChangesnewBlockFilter or getFilterChangesnewPendingTransactionFilter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetFilterChanges(context.Background(), filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetFilterChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetFilterChanges`: OutputGetFilterChanges
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetFilterChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **string** | The filter id that is returned from getFilterChangesnewFilter, getFilterChangesnewBlockFilter or getFilterChangesnewPendingTransactionFilter | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetFilterChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetFilterChanges**](OutputGetFilterChanges.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetFilterLogs

> OutputGetFilterLogs QanGetFilterLogs(ctx, id).Execute()

Returns an array of all logs matching filter with given id.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	id := "id_example" // string | The filter ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetFilterLogs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetFilterLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetFilterLogs`: OutputGetFilterLogs
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetFilterLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The filter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetFilterLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetFilterLogs**](OutputGetFilterLogs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetLogs

> OutputGetLogs QanGetLogs(ctx).InputGetLogs(inputGetLogs).Execute()

Returns an array of all logs matching a given filter object.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	inputGetLogs := *openapiclient.NewInputGetLogs() // InputGetLogs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetLogs(context.Background()).InputGetLogs(inputGetLogs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetLogs`: OutputGetLogs
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputGetLogs** | [**InputGetLogs**](InputGetLogs.md) |  | 

### Return type

[**OutputGetLogs**](OutputGetLogs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetProof

> OutputGetProof QanGetProof(ctx).InputGetProof(inputGetProof).Execute()

Returns the account and storage values of the specified account including the Merkle-proof.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	inputGetProof := *openapiclient.NewInputGetProof("Address_example", []string{"StorageKeys_example"}) // InputGetProof | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetProof(context.Background()).InputGetProof(inputGetProof).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetProof``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetProof`: OutputGetProof
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetProof`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanGetProofRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputGetProof** | [**InputGetProof**](InputGetProof.md) |  | 

### Return type

[**OutputGetProof**](OutputGetProof.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetStorageAt

> OutputGetStorageAt QanGetStorageAt(ctx).InputGetStorageAt(inputGetStorageAt).Execute()

Returns the value from a storage position at a given address.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	inputGetStorageAt := *openapiclient.NewInputGetStorageAt("Address_example", "BlockNumber_example", "Position_example") // InputGetStorageAt | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetStorageAt(context.Background()).InputGetStorageAt(inputGetStorageAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetStorageAt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetStorageAt`: OutputGetStorageAt
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetStorageAt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanGetStorageAtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputGetStorageAt** | [**InputGetStorageAt**](InputGetStorageAt.md) |  | 

### Return type

[**OutputGetStorageAt**](OutputGetStorageAt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetTransactionByBlockHashAndIndex

> OutputGetTransactionByBlockHashAndIndex QanGetTransactionByBlockHashAndIndex(ctx, blockHash, index).Execute()

Returns information about a transaction given a blockhash and transaction index position.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	blockHash := "0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd" // string | 
	index := "0" // string | An integer of the transaction index position

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetTransactionByBlockHashAndIndex(context.Background(), blockHash, index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetTransactionByBlockHashAndIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetTransactionByBlockHashAndIndex`: OutputGetTransactionByBlockHashAndIndex
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetTransactionByBlockHashAndIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHash** | **string** |  | 
**index** | **string** | An integer of the transaction index position | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetTransactionByBlockHashAndIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputGetTransactionByBlockHashAndIndex**](OutputGetTransactionByBlockHashAndIndex.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetTransactionByBlockNumberAndIndex

> OutputGetTransactionByBlockNumberAndIndex QanGetTransactionByBlockNumberAndIndex(ctx, blockNumber, index).Execute()

Returns information about a transaction given a block number and transaction index position.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	blockNumber := "latest" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending
	index := "0" // string | An integer of the transaction index position

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetTransactionByBlockNumberAndIndex(context.Background(), blockNumber, index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetTransactionByBlockNumberAndIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetTransactionByBlockNumberAndIndex`: OutputGetTransactionByBlockNumberAndIndex
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetTransactionByBlockNumberAndIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending | 
**index** | **string** | An integer of the transaction index position | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetTransactionByBlockNumberAndIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputGetTransactionByBlockNumberAndIndex**](OutputGetTransactionByBlockNumberAndIndex.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetTransactionByHash

> OutputGetTransactionByHash QanGetTransactionByHash(ctx, hash).Execute()

Returns the information about a transaction from a transaction hash.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	hash := "0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060" // string | The hash of a transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetTransactionByHash(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetTransactionByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetTransactionByHash`: OutputGetTransactionByHash
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetTransactionByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | The hash of a transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetTransactionByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetTransactionByHash**](OutputGetTransactionByHash.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetTransactionCount

> OutputGetTransactionCount QanGetTransactionCount(ctx, address, blockNumber).Execute()

Returns the number of transactions sent from an address.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	address := "0xa1e4380a3b1f749673e270229993ee55f35663b4" // string | The address from which the transaction count to be checked
	blockNumber := "latest" // string | The block number in hexadecimal or decimal format or the string latest, earliest, pending

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetTransactionCount(context.Background(), address, blockNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetTransactionCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetTransactionCount`: OutputGetTransactionCount
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetTransactionCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | The address from which the transaction count to be checked | 
**blockNumber** | **string** | The block number in hexadecimal or decimal format or the string latest, earliest, pending | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetTransactionCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputGetTransactionCount**](OutputGetTransactionCount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanGetTransactionReceipt

> OutputGetTransactionReceipt QanGetTransactionReceipt(ctx, hash).Execute()

Returns the receipt of a transaction by transaction hash.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	hash := "0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd" // string | The hash of a transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanGetTransactionReceipt(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanGetTransactionReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanGetTransactionReceipt`: OutputGetTransactionReceipt
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanGetTransactionReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | The hash of a transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanGetTransactionReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputGetTransactionReceipt**](OutputGetTransactionReceipt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanMaxPriorityFeePerGas

> OutputMaxPriorityFeePerGas QanMaxPriorityFeePerGas(ctx).Execute()

Get the priority fee needed to be included in a block.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanMaxPriorityFeePerGas(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanMaxPriorityFeePerGas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanMaxPriorityFeePerGas`: OutputMaxPriorityFeePerGas
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanMaxPriorityFeePerGas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanMaxPriorityFeePerGasRequest struct via the builder pattern


### Return type

[**OutputMaxPriorityFeePerGas**](OutputMaxPriorityFeePerGas.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanNewBlockFilter

> OutputNewBlockFilter QanNewBlockFilter(ctx).Execute()

Creates a filter in the node, to notify when a new block arrives.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanNewBlockFilter(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanNewBlockFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanNewBlockFilter`: OutputNewBlockFilter
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanNewBlockFilter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanNewBlockFilterRequest struct via the builder pattern


### Return type

[**OutputNewBlockFilter**](OutputNewBlockFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanNewFilter

> OutputNewFilter QanNewFilter(ctx).InputNewFilter(inputNewFilter).Execute()

Creates a filter object, based on filter options, to notify when the state changes (logs).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	inputNewFilter := *openapiclient.NewInputNewFilter(*openapiclient.NewFilterObject("Address_example", "FromBlock_example", "ToBlock_example", []string{"Topics_example"})) // InputNewFilter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanNewFilter(context.Background()).InputNewFilter(inputNewFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanNewFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanNewFilter`: OutputNewFilter
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanNewFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanNewFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputNewFilter** | [**InputNewFilter**](InputNewFilter.md) |  | 

### Return type

[**OutputNewFilter**](OutputNewFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanNewPendingTransactionFilter

> OutputNewPendingTransactionFilter QanNewPendingTransactionFilter(ctx).Execute()

Creates a filter in the node to notify when new pending transactions arrive.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanNewPendingTransactionFilter(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanNewPendingTransactionFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanNewPendingTransactionFilter`: OutputNewPendingTransactionFilter
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanNewPendingTransactionFilter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanNewPendingTransactionFilterRequest struct via the builder pattern


### Return type

[**OutputNewPendingTransactionFilter**](OutputNewPendingTransactionFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanSendRawTransaction

> OutputSendRawTransaction QanSendRawTransaction(ctx).InputSendRawTransaction(inputSendRawTransaction).Execute()

Creates new message call transaction or a contract creation for signed transactions.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	inputSendRawTransaction := *openapiclient.NewInputSendRawTransaction("Data_example") // InputSendRawTransaction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanSendRawTransaction(context.Background()).InputSendRawTransaction(inputSendRawTransaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanSendRawTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanSendRawTransaction`: OutputSendRawTransaction
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanSendRawTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQanSendRawTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputSendRawTransaction** | [**InputSendRawTransaction**](InputSendRawTransaction.md) |  | 

### Return type

[**OutputSendRawTransaction**](OutputSendRawTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanSyncing

> OutputSyncing QanSyncing(ctx).Execute()

Returns an object with the sync status of the node if the node is out-of-sync and is syncing. Returns null when the node is already in sync.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanSyncing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanSyncing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanSyncing`: OutputSyncing
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanSyncing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQanSyncingRequest struct via the builder pattern


### Return type

[**OutputSyncing**](OutputSyncing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanUninstallFilter

> OutputUninstallFilter QanUninstallFilter(ctx, filterId).Execute()

Uninstalls a filter with the given filter id.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	filterId := "filterId_example" // string | The filter ID that needs to be uninstalled. It should always be called when watch is no longer needed. Additionally, Filters timeout when they aren't requested with getFilterChanges for a period of time

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanUninstallFilter(context.Background(), filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanUninstallFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanUninstallFilter`: OutputUninstallFilter
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanUninstallFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **string** | The filter ID that needs to be uninstalled. It should always be called when watch is no longer needed. Additionally, Filters timeout when they aren&#39;t requested with getFilterChanges for a period of time | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanUninstallFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputUninstallFilter**](OutputUninstallFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QanXlinkValid

> OutputXlinkValid QanXlinkValid(ctx, address).Execute()

Returns the xlink validity time of the account of given address.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/QANplatform/sdk-go"
)

func main() {
	address := "address_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QANAPI.QanXlinkValid(context.Background(), address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QANAPI.QanXlinkValid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QanXlinkValid`: OutputXlinkValid
	fmt.Fprintf(os.Stdout, "Response from `QANAPI.QanXlinkValid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQanXlinkValidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OutputXlinkValid**](OutputXlinkValid.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

