# WebhookEventCallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**WebhookEventResource**](WebhookEventResource.md) | The webhook event data sent to the subscribed webhook.  | 

## Methods

### NewWebhookEventCallback

`func NewWebhookEventCallback(data WebhookEventResource, ) *WebhookEventCallback`

NewWebhookEventCallback instantiates a new WebhookEventCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventCallbackWithDefaults

`func NewWebhookEventCallbackWithDefaults() *WebhookEventCallback`

NewWebhookEventCallbackWithDefaults instantiates a new WebhookEventCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WebhookEventCallback) GetData() WebhookEventResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookEventCallback) GetDataOk() (*WebhookEventResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookEventCallback) SetData(v WebhookEventResource)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


