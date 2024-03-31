# PingResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the authenticated customer.  | 
**StatusEmoji** | **string** | A cute emoji that represents the response status.  | 

## Methods

### NewPingResponseMeta

`func NewPingResponseMeta(id string, statusEmoji string, ) *PingResponseMeta`

NewPingResponseMeta instantiates a new PingResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingResponseMetaWithDefaults

`func NewPingResponseMetaWithDefaults() *PingResponseMeta`

NewPingResponseMetaWithDefaults instantiates a new PingResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PingResponseMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingResponseMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingResponseMeta) SetId(v string)`

SetId sets Id field to given value.


### GetStatusEmoji

`func (o *PingResponseMeta) GetStatusEmoji() string`

GetStatusEmoji returns the StatusEmoji field if non-nil, zero value otherwise.

### GetStatusEmojiOk

`func (o *PingResponseMeta) GetStatusEmojiOk() (*string, bool)`

GetStatusEmojiOk returns a tuple with the StatusEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEmoji

`func (o *PingResponseMeta) SetStatusEmoji(v string)`

SetStatusEmoji sets StatusEmoji field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


