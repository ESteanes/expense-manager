# WebhookResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL that this webhook is configured to &#x60;POST&#x60; events to.  | 
**Description** | **NullableString** | An optional description that was provided at the time the webhook was created.  | 
**SecretKey** | Pointer to **string** | A shared secret key used to sign all webhook events sent to the configured webhook URL. This field is returned only once, upon the initial creation of the webhook. If lost, create a new webhook and delete this webhook.  The webhook URL receives a request with a &#x60;X-Up-Authenticity-Signature&#x60; header, which is the SHA-256 HMAC of the entire raw request body signed using this &#x60;secretKey&#x60;. It is advised to compute and check this signature to verify the authenticity of requests sent to the webhook URL. See [Handling webhook events](#callback_post_webhookURL) for full details.  | [optional] 
**CreatedAt** | **time.Time** | The date-time at which this webhook was created.  | 

## Methods

### NewWebhookResourceAttributes

`func NewWebhookResourceAttributes(url string, description NullableString, createdAt time.Time, ) *WebhookResourceAttributes`

NewWebhookResourceAttributes instantiates a new WebhookResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResourceAttributesWithDefaults

`func NewWebhookResourceAttributesWithDefaults() *WebhookResourceAttributes`

NewWebhookResourceAttributesWithDefaults instantiates a new WebhookResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookResourceAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookResourceAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookResourceAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *WebhookResourceAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookResourceAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookResourceAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WebhookResourceAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebhookResourceAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSecretKey

`func (o *WebhookResourceAttributes) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *WebhookResourceAttributes) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *WebhookResourceAttributes) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *WebhookResourceAttributes) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookResourceAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookResourceAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookResourceAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


