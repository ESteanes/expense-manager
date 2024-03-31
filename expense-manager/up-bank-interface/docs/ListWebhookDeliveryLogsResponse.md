# ListWebhookDeliveryLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]WebhookDeliveryLogResource**](WebhookDeliveryLogResource.md) | The list of delivery logs returned in this response.  | 
**Links** | [**ListAccountsResponseLinks**](ListAccountsResponseLinks.md) |  | 

## Methods

### NewListWebhookDeliveryLogsResponse

`func NewListWebhookDeliveryLogsResponse(data []WebhookDeliveryLogResource, links ListAccountsResponseLinks, ) *ListWebhookDeliveryLogsResponse`

NewListWebhookDeliveryLogsResponse instantiates a new ListWebhookDeliveryLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhookDeliveryLogsResponseWithDefaults

`func NewListWebhookDeliveryLogsResponseWithDefaults() *ListWebhookDeliveryLogsResponse`

NewListWebhookDeliveryLogsResponseWithDefaults instantiates a new ListWebhookDeliveryLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListWebhookDeliveryLogsResponse) GetData() []WebhookDeliveryLogResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListWebhookDeliveryLogsResponse) GetDataOk() (*[]WebhookDeliveryLogResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListWebhookDeliveryLogsResponse) SetData(v []WebhookDeliveryLogResource)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListWebhookDeliveryLogsResponse) GetLinks() ListAccountsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListWebhookDeliveryLogsResponse) GetLinksOk() (*ListAccountsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListWebhookDeliveryLogsResponse) SetLinks(v ListAccountsResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


