# ListAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AccountResource**](AccountResource.md) | The list of accounts returned in this response.  | 
**Links** | [**ListAccountsResponseLinks**](ListAccountsResponseLinks.md) |  | 

## Methods

### NewListAccountsResponse

`func NewListAccountsResponse(data []AccountResource, links ListAccountsResponseLinks, ) *ListAccountsResponse`

NewListAccountsResponse instantiates a new ListAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccountsResponseWithDefaults

`func NewListAccountsResponseWithDefaults() *ListAccountsResponse`

NewListAccountsResponseWithDefaults instantiates a new ListAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAccountsResponse) GetData() []AccountResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAccountsResponse) GetDataOk() (*[]AccountResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAccountsResponse) SetData(v []AccountResource)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListAccountsResponse) GetLinks() ListAccountsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListAccountsResponse) GetLinksOk() (*ListAccountsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListAccountsResponse) SetLinks(v ListAccountsResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


