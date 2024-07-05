# WebhookEventResourceRelationshipsTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**WebhookEventResourceRelationshipsTransactionData**](WebhookEventResourceRelationshipsTransactionData.md) |  | 
**Links** | Pointer to [**AccountResourceRelationshipsTransactionsLinks**](AccountResourceRelationshipsTransactionsLinks.md) |  | [optional] 

## Methods

### NewWebhookEventResourceRelationshipsTransaction

`func NewWebhookEventResourceRelationshipsTransaction(data WebhookEventResourceRelationshipsTransactionData, ) *WebhookEventResourceRelationshipsTransaction`

NewWebhookEventResourceRelationshipsTransaction instantiates a new WebhookEventResourceRelationshipsTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventResourceRelationshipsTransactionWithDefaults

`func NewWebhookEventResourceRelationshipsTransactionWithDefaults() *WebhookEventResourceRelationshipsTransaction`

NewWebhookEventResourceRelationshipsTransactionWithDefaults instantiates a new WebhookEventResourceRelationshipsTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WebhookEventResourceRelationshipsTransaction) GetData() WebhookEventResourceRelationshipsTransactionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookEventResourceRelationshipsTransaction) GetDataOk() (*WebhookEventResourceRelationshipsTransactionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookEventResourceRelationshipsTransaction) SetData(v WebhookEventResourceRelationshipsTransactionData)`

SetData sets Data field to given value.


### GetLinks

`func (o *WebhookEventResourceRelationshipsTransaction) GetLinks() AccountResourceRelationshipsTransactionsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WebhookEventResourceRelationshipsTransaction) GetLinksOk() (*AccountResourceRelationshipsTransactionsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WebhookEventResourceRelationshipsTransaction) SetLinks(v AccountResourceRelationshipsTransactionsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WebhookEventResourceRelationshipsTransaction) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


