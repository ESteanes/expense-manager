# WebhookEventResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**WebhookEventTypeEnum**](WebhookEventTypeEnum.md) | The type of this event. This can be used to determine what action to take in response to the event.  | 
**CreatedAt** | **time.Time** | The date-time at which this event was generated.  | 

## Methods

### NewWebhookEventResourceAttributes

`func NewWebhookEventResourceAttributes(eventType WebhookEventTypeEnum, createdAt time.Time, ) *WebhookEventResourceAttributes`

NewWebhookEventResourceAttributes instantiates a new WebhookEventResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventResourceAttributesWithDefaults

`func NewWebhookEventResourceAttributesWithDefaults() *WebhookEventResourceAttributes`

NewWebhookEventResourceAttributesWithDefaults instantiates a new WebhookEventResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *WebhookEventResourceAttributes) GetEventType() WebhookEventTypeEnum`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookEventResourceAttributes) GetEventTypeOk() (*WebhookEventTypeEnum, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookEventResourceAttributes) SetEventType(v WebhookEventTypeEnum)`

SetEventType sets EventType field to given value.


### GetCreatedAt

`func (o *WebhookEventResourceAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookEventResourceAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookEventResourceAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


