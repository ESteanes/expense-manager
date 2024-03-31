# WebhookResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this resource: &#x60;webhooks&#x60; | 
**Id** | **string** | The unique identifier for this webhook.  | 
**Attributes** | [**WebhookResourceAttributes**](WebhookResourceAttributes.md) |  | 
**Relationships** | [**WebhookResourceRelationships**](WebhookResourceRelationships.md) |  | 
**Links** | Pointer to [**AccountResourceLinks**](AccountResourceLinks.md) |  | [optional] 

## Methods

### NewWebhookResource

`func NewWebhookResource(type_ string, id string, attributes WebhookResourceAttributes, relationships WebhookResourceRelationships, ) *WebhookResource`

NewWebhookResource instantiates a new WebhookResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResourceWithDefaults

`func NewWebhookResourceWithDefaults() *WebhookResource`

NewWebhookResourceWithDefaults instantiates a new WebhookResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookResource) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *WebhookResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *WebhookResource) GetAttributes() WebhookResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WebhookResource) GetAttributesOk() (*WebhookResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WebhookResource) SetAttributes(v WebhookResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *WebhookResource) GetRelationships() WebhookResourceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WebhookResource) GetRelationshipsOk() (*WebhookResourceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WebhookResource) SetRelationships(v WebhookResourceRelationships)`

SetRelationships sets Relationships field to given value.


### GetLinks

`func (o *WebhookResource) GetLinks() AccountResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WebhookResource) GetLinksOk() (*AccountResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WebhookResource) SetLinks(v AccountResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WebhookResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


