# WebhookDeliveryLogResourceAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** | The HTTP status code received in the response.  | 
**Body** | **string** | The payload that was received in the response body.  | 

## Methods

### NewWebhookDeliveryLogResourceAttributesResponse

`func NewWebhookDeliveryLogResourceAttributesResponse(statusCode int32, body string, ) *WebhookDeliveryLogResourceAttributesResponse`

NewWebhookDeliveryLogResourceAttributesResponse instantiates a new WebhookDeliveryLogResourceAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryLogResourceAttributesResponseWithDefaults

`func NewWebhookDeliveryLogResourceAttributesResponseWithDefaults() *WebhookDeliveryLogResourceAttributesResponse`

NewWebhookDeliveryLogResourceAttributesResponseWithDefaults instantiates a new WebhookDeliveryLogResourceAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *WebhookDeliveryLogResourceAttributesResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *WebhookDeliveryLogResourceAttributesResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *WebhookDeliveryLogResourceAttributesResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetBody

`func (o *WebhookDeliveryLogResourceAttributesResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WebhookDeliveryLogResourceAttributesResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WebhookDeliveryLogResourceAttributesResponse) SetBody(v string)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


