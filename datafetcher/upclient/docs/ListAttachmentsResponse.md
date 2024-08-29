# ListAttachmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AttachmentResource**](AttachmentResource.md) | The list of attachments returned in this response.  | 
**Links** | [**ListAccountsResponseLinks**](ListAccountsResponseLinks.md) |  | 

## Methods

### NewListAttachmentsResponse

`func NewListAttachmentsResponse(data []AttachmentResource, links ListAccountsResponseLinks, ) *ListAttachmentsResponse`

NewListAttachmentsResponse instantiates a new ListAttachmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAttachmentsResponseWithDefaults

`func NewListAttachmentsResponseWithDefaults() *ListAttachmentsResponse`

NewListAttachmentsResponseWithDefaults instantiates a new ListAttachmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAttachmentsResponse) GetData() []AttachmentResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAttachmentsResponse) GetDataOk() (*[]AttachmentResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAttachmentsResponse) SetData(v []AttachmentResource)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListAttachmentsResponse) GetLinks() ListAccountsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListAttachmentsResponse) GetLinksOk() (*ListAccountsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListAttachmentsResponse) SetLinks(v ListAccountsResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


