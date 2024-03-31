# GetTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TransactionResource**](TransactionResource.md) | The transaction returned in this response.  | 

## Methods

### NewGetTransactionResponse

`func NewGetTransactionResponse(data TransactionResource, ) *GetTransactionResponse`

NewGetTransactionResponse instantiates a new GetTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionResponseWithDefaults

`func NewGetTransactionResponseWithDefaults() *GetTransactionResponse`

NewGetTransactionResponseWithDefaults instantiates a new GetTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTransactionResponse) GetData() TransactionResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTransactionResponse) GetDataOk() (*TransactionResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTransactionResponse) SetData(v TransactionResource)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


