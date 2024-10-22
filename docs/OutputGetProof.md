# OutputGetProof

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**AccountProof** | **string** | An array of rlp-serialized MerkleTree-Nodes which starts with the stateRoot-Node and follows the path of the SHA3 address as key | 
**Address** | **string** | The address associated with the account | 
**Balance** | **string** | The current balance of the account in wei | 
**CodeHash** | **string** | A 32 byte hash of the code of the account | 
**Nonce** | **NullableString** | The hash of the generated proof-of-work. Null if pending | 
**StorageHash** | **string** | A 32 byte SHA3 of the storageRoot. All storage will deliver a MerkleProof starting with this rootHash | 
**StorageProof** | [**[]ResponseStorageEntry**](ResponseStorageEntry.md) | An array of storage-entries as requested. Each entry is an object with the following fields: | 

## Methods

### NewOutputGetProof

`func NewOutputGetProof(accountProof string, address string, balance string, codeHash string, nonce NullableString, storageHash string, storageProof []ResponseStorageEntry, ) *OutputGetProof`

NewOutputGetProof instantiates a new OutputGetProof object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputGetProofWithDefaults

`func NewOutputGetProofWithDefaults() *OutputGetProof`

NewOutputGetProofWithDefaults instantiates a new OutputGetProof object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OutputGetProof) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputGetProof) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputGetProof) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OutputGetProof) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAccountProof

`func (o *OutputGetProof) GetAccountProof() string`

GetAccountProof returns the AccountProof field if non-nil, zero value otherwise.

### GetAccountProofOk

`func (o *OutputGetProof) GetAccountProofOk() (*string, bool)`

GetAccountProofOk returns a tuple with the AccountProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountProof

`func (o *OutputGetProof) SetAccountProof(v string)`

SetAccountProof sets AccountProof field to given value.


### GetAddress

`func (o *OutputGetProof) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OutputGetProof) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OutputGetProof) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalance

`func (o *OutputGetProof) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *OutputGetProof) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *OutputGetProof) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetCodeHash

`func (o *OutputGetProof) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *OutputGetProof) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *OutputGetProof) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.


### GetNonce

`func (o *OutputGetProof) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *OutputGetProof) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *OutputGetProof) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### SetNonceNil

`func (o *OutputGetProof) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *OutputGetProof) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
### GetStorageHash

`func (o *OutputGetProof) GetStorageHash() string`

GetStorageHash returns the StorageHash field if non-nil, zero value otherwise.

### GetStorageHashOk

`func (o *OutputGetProof) GetStorageHashOk() (*string, bool)`

GetStorageHashOk returns a tuple with the StorageHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageHash

`func (o *OutputGetProof) SetStorageHash(v string)`

SetStorageHash sets StorageHash field to given value.


### GetStorageProof

`func (o *OutputGetProof) GetStorageProof() []ResponseStorageEntry`

GetStorageProof returns the StorageProof field if non-nil, zero value otherwise.

### GetStorageProofOk

`func (o *OutputGetProof) GetStorageProofOk() (*[]ResponseStorageEntry, bool)`

GetStorageProofOk returns a tuple with the StorageProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProof

`func (o *OutputGetProof) SetStorageProof(v []ResponseStorageEntry)`

SetStorageProof sets StorageProof field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


