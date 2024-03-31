# ListTransactionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TransactionResource**](TransactionResource.md) | The list of transactions returned in this response.  | 
**Links** | [**ListAccountsResponseLinks**](ListAccountsResponseLinks.md) |  | 

## Methods

### NewListTransactionsResponse

`func NewListTransactionsResponse(data []TransactionResource, links ListAccountsResponseLinks, ) *ListTransactionsResponse`

NewListTransactionsResponse instantiates a new ListTransactionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsResponseWithDefaults

`func NewListTransactionsResponseWithDefaults() *ListTransactionsResponse`

NewListTransactionsResponseWithDefaults instantiates a new ListTransactionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListTransactionsResponse) GetData() []TransactionResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListTransactionsResponse) GetDataOk() (*[]TransactionResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListTransactionsResponse) SetData(v []TransactionResource)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListTransactionsResponse) GetLinks() ListAccountsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListTransactionsResponse) GetLinksOk() (*ListAccountsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListTransactionsResponse) SetLinks(v ListAccountsResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


