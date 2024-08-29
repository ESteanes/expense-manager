# AttachmentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this resource: &#x60;attachments&#x60; | 
**Id** | **string** | The unique identifier for this attachment.  | 
**Attributes** | [**AttachmentResourceAttributes**](AttachmentResourceAttributes.md) |  | 
**Relationships** | [**AttachmentResourceRelationships**](AttachmentResourceRelationships.md) |  | 
**Links** | Pointer to [**AccountResourceLinks**](AccountResourceLinks.md) |  | [optional] 

## Methods

### NewAttachmentResource

`func NewAttachmentResource(type_ string, id string, attributes AttachmentResourceAttributes, relationships AttachmentResourceRelationships, ) *AttachmentResource`

NewAttachmentResource instantiates a new AttachmentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentResourceWithDefaults

`func NewAttachmentResourceWithDefaults() *AttachmentResource`

NewAttachmentResourceWithDefaults instantiates a new AttachmentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttachmentResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentResource) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AttachmentResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AttachmentResource) GetAttributes() AttachmentResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AttachmentResource) GetAttributesOk() (*AttachmentResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AttachmentResource) SetAttributes(v AttachmentResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AttachmentResource) GetRelationships() AttachmentResourceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AttachmentResource) GetRelationshipsOk() (*AttachmentResourceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AttachmentResource) SetRelationships(v AttachmentResourceRelationships)`

SetRelationships sets Relationships field to given value.


### GetLinks

`func (o *AttachmentResource) GetLinks() AccountResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AttachmentResource) GetLinksOk() (*AccountResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AttachmentResource) SetLinks(v AccountResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AttachmentResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


