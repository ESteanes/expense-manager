# AttachmentUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | Pointer to **string** | The URL of the file once it has been uploaded | [optional] 
**UploadUrl** | Pointer to **string** | The URL to POST the file to when uploading | [optional] 

## Methods

### NewAttachmentUploadResponse

`func NewAttachmentUploadResponse() *AttachmentUploadResponse`

NewAttachmentUploadResponse instantiates a new AttachmentUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentUploadResponseWithDefaults

`func NewAttachmentUploadResponseWithDefaults() *AttachmentUploadResponse`

NewAttachmentUploadResponseWithDefaults instantiates a new AttachmentUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *AttachmentUploadResponse) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *AttachmentUploadResponse) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *AttachmentUploadResponse) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *AttachmentUploadResponse) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetUploadUrl

`func (o *AttachmentUploadResponse) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *AttachmentUploadResponse) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *AttachmentUploadResponse) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *AttachmentUploadResponse) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


