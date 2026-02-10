# WebhookTransactionCreatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The event type | [optional] 
**Data** | Pointer to [**WebhookTransactionCreatedEventData**](WebhookTransactionCreatedEventData.md) |  | [optional] 

## Methods

### NewWebhookTransactionCreatedEvent

`func NewWebhookTransactionCreatedEvent() *WebhookTransactionCreatedEvent`

NewWebhookTransactionCreatedEvent instantiates a new WebhookTransactionCreatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookTransactionCreatedEventWithDefaults

`func NewWebhookTransactionCreatedEventWithDefaults() *WebhookTransactionCreatedEvent`

NewWebhookTransactionCreatedEventWithDefaults instantiates a new WebhookTransactionCreatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookTransactionCreatedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookTransactionCreatedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookTransactionCreatedEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WebhookTransactionCreatedEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *WebhookTransactionCreatedEvent) GetData() WebhookTransactionCreatedEventData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookTransactionCreatedEvent) GetDataOk() (*WebhookTransactionCreatedEventData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookTransactionCreatedEvent) SetData(v WebhookTransactionCreatedEventData)`

SetData sets Data field to given value.

### HasData

`func (o *WebhookTransactionCreatedEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


