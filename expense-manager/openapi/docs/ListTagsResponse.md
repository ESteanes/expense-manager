# ListTagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TagResource**](TagResource.md) | The list of tags returned in this response.  | 
**Links** | [**ListAccountsResponseLinks**](ListAccountsResponseLinks.md) |  | 

## Methods

### NewListTagsResponse

`func NewListTagsResponse(data []TagResource, links ListAccountsResponseLinks, ) *ListTagsResponse`

NewListTagsResponse instantiates a new ListTagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTagsResponseWithDefaults

`func NewListTagsResponseWithDefaults() *ListTagsResponse`

NewListTagsResponseWithDefaults instantiates a new ListTagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListTagsResponse) GetData() []TagResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListTagsResponse) GetDataOk() (*[]TagResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListTagsResponse) SetData(v []TagResource)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListTagsResponse) GetLinks() ListAccountsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListTagsResponse) GetLinksOk() (*ListAccountsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListTagsResponse) SetLinks(v ListAccountsResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


