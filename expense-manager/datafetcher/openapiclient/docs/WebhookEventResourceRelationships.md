# WebhookEventResourceRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Webhook** | [**WebhookEventResourceRelationshipsWebhook**](WebhookEventResourceRelationshipsWebhook.md) |  | 
**Transaction** | Pointer to [**WebhookEventResourceRelationshipsTransaction**](WebhookEventResourceRelationshipsTransaction.md) |  | [optional] 

## Methods

### NewWebhookEventResourceRelationships

`func NewWebhookEventResourceRelationships(webhook WebhookEventResourceRelationshipsWebhook, ) *WebhookEventResourceRelationships`

NewWebhookEventResourceRelationships instantiates a new WebhookEventResourceRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventResourceRelationshipsWithDefaults

`func NewWebhookEventResourceRelationshipsWithDefaults() *WebhookEventResourceRelationships`

NewWebhookEventResourceRelationshipsWithDefaults instantiates a new WebhookEventResourceRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhook

`func (o *WebhookEventResourceRelationships) GetWebhook() WebhookEventResourceRelationshipsWebhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *WebhookEventResourceRelationships) GetWebhookOk() (*WebhookEventResourceRelationshipsWebhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *WebhookEventResourceRelationships) SetWebhook(v WebhookEventResourceRelationshipsWebhook)`

SetWebhook sets Webhook field to given value.


### GetTransaction

`func (o *WebhookEventResourceRelationships) GetTransaction() WebhookEventResourceRelationshipsTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *WebhookEventResourceRelationships) GetTransactionOk() (*WebhookEventResourceRelationshipsTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *WebhookEventResourceRelationships) SetTransaction(v WebhookEventResourceRelationshipsTransaction)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *WebhookEventResourceRelationships) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


