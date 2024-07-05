# WebhookDeliveryLogResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this resource: &#x60;webhook-delivery-logs&#x60; | 
**Id** | **string** | The unique identifier for this log entry.  | 
**Attributes** | [**WebhookDeliveryLogResourceAttributes**](WebhookDeliveryLogResourceAttributes.md) |  | 
**Relationships** | [**WebhookDeliveryLogResourceRelationships**](WebhookDeliveryLogResourceRelationships.md) |  | 

## Methods

### NewWebhookDeliveryLogResource

`func NewWebhookDeliveryLogResource(type_ string, id string, attributes WebhookDeliveryLogResourceAttributes, relationships WebhookDeliveryLogResourceRelationships, ) *WebhookDeliveryLogResource`

NewWebhookDeliveryLogResource instantiates a new WebhookDeliveryLogResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryLogResourceWithDefaults

`func NewWebhookDeliveryLogResourceWithDefaults() *WebhookDeliveryLogResource`

NewWebhookDeliveryLogResourceWithDefaults instantiates a new WebhookDeliveryLogResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookDeliveryLogResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookDeliveryLogResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookDeliveryLogResource) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *WebhookDeliveryLogResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookDeliveryLogResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookDeliveryLogResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *WebhookDeliveryLogResource) GetAttributes() WebhookDeliveryLogResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WebhookDeliveryLogResource) GetAttributesOk() (*WebhookDeliveryLogResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WebhookDeliveryLogResource) SetAttributes(v WebhookDeliveryLogResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *WebhookDeliveryLogResource) GetRelationships() WebhookDeliveryLogResourceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WebhookDeliveryLogResource) GetRelationshipsOk() (*WebhookDeliveryLogResourceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WebhookDeliveryLogResource) SetRelationships(v WebhookDeliveryLogResourceRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


