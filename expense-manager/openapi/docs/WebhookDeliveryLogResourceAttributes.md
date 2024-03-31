# WebhookDeliveryLogResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | [**WebhookDeliveryLogResourceAttributesRequest**](WebhookDeliveryLogResourceAttributesRequest.md) |  | 
**Response** | [**NullableWebhookDeliveryLogResourceAttributesResponse**](WebhookDeliveryLogResourceAttributesResponse.md) |  | 
**DeliveryStatus** | [**WebhookDeliveryStatusEnum**](WebhookDeliveryStatusEnum.md) | The success or failure status of this delivery attempt.  | 
**CreatedAt** | **time.Time** | The date-time at which this log entry was created.  | 

## Methods

### NewWebhookDeliveryLogResourceAttributes

`func NewWebhookDeliveryLogResourceAttributes(request WebhookDeliveryLogResourceAttributesRequest, response NullableWebhookDeliveryLogResourceAttributesResponse, deliveryStatus WebhookDeliveryStatusEnum, createdAt time.Time, ) *WebhookDeliveryLogResourceAttributes`

NewWebhookDeliveryLogResourceAttributes instantiates a new WebhookDeliveryLogResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryLogResourceAttributesWithDefaults

`func NewWebhookDeliveryLogResourceAttributesWithDefaults() *WebhookDeliveryLogResourceAttributes`

NewWebhookDeliveryLogResourceAttributesWithDefaults instantiates a new WebhookDeliveryLogResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *WebhookDeliveryLogResourceAttributes) GetRequest() WebhookDeliveryLogResourceAttributesRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *WebhookDeliveryLogResourceAttributes) GetRequestOk() (*WebhookDeliveryLogResourceAttributesRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *WebhookDeliveryLogResourceAttributes) SetRequest(v WebhookDeliveryLogResourceAttributesRequest)`

SetRequest sets Request field to given value.


### GetResponse

`func (o *WebhookDeliveryLogResourceAttributes) GetResponse() WebhookDeliveryLogResourceAttributesResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebhookDeliveryLogResourceAttributes) GetResponseOk() (*WebhookDeliveryLogResourceAttributesResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WebhookDeliveryLogResourceAttributes) SetResponse(v WebhookDeliveryLogResourceAttributesResponse)`

SetResponse sets Response field to given value.


### SetResponseNil

`func (o *WebhookDeliveryLogResourceAttributes) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *WebhookDeliveryLogResourceAttributes) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetDeliveryStatus

`func (o *WebhookDeliveryLogResourceAttributes) GetDeliveryStatus() WebhookDeliveryStatusEnum`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *WebhookDeliveryLogResourceAttributes) GetDeliveryStatusOk() (*WebhookDeliveryStatusEnum, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *WebhookDeliveryLogResourceAttributes) SetDeliveryStatus(v WebhookDeliveryStatusEnum)`

SetDeliveryStatus sets DeliveryStatus field to given value.


### GetCreatedAt

`func (o *WebhookDeliveryLogResourceAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookDeliveryLogResourceAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookDeliveryLogResourceAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


