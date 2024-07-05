# WebhookInputResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL that this webhook should post events to. This must be a valid HTTP or HTTPS URL that does not exceed 300 characters in length.  | 
**Description** | Pointer to **NullableString** | An optional description for this webhook, up to 64 characters in length.  | [optional] 

## Methods

### NewWebhookInputResourceAttributes

`func NewWebhookInputResourceAttributes(url string, ) *WebhookInputResourceAttributes`

NewWebhookInputResourceAttributes instantiates a new WebhookInputResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookInputResourceAttributesWithDefaults

`func NewWebhookInputResourceAttributesWithDefaults() *WebhookInputResourceAttributes`

NewWebhookInputResourceAttributesWithDefaults instantiates a new WebhookInputResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookInputResourceAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookInputResourceAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookInputResourceAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *WebhookInputResourceAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookInputResourceAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookInputResourceAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookInputResourceAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebhookInputResourceAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebhookInputResourceAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


