# AttachmentResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **NullableTime** | The date-time when the file was created.  | 
**FileURL** | **NullableString** | A temporary link to download the file.  | 
**FileURLExpiresAt** | **time.Time** | The date-time at which the &#x60;fileURL&#x60; link expires.  | 
**FileExtension** | **NullableString** | File extension for the uploaded attachment.  | 
**FileContentType** | **NullableString** | Content type for the uploaded attachment.  | 

## Methods

### NewAttachmentResourceAttributes

`func NewAttachmentResourceAttributes(createdAt NullableTime, fileURL NullableString, fileURLExpiresAt time.Time, fileExtension NullableString, fileContentType NullableString, ) *AttachmentResourceAttributes`

NewAttachmentResourceAttributes instantiates a new AttachmentResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentResourceAttributesWithDefaults

`func NewAttachmentResourceAttributesWithDefaults() *AttachmentResourceAttributes`

NewAttachmentResourceAttributesWithDefaults instantiates a new AttachmentResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AttachmentResourceAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AttachmentResourceAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AttachmentResourceAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *AttachmentResourceAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AttachmentResourceAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetFileURL

`func (o *AttachmentResourceAttributes) GetFileURL() string`

GetFileURL returns the FileURL field if non-nil, zero value otherwise.

### GetFileURLOk

`func (o *AttachmentResourceAttributes) GetFileURLOk() (*string, bool)`

GetFileURLOk returns a tuple with the FileURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileURL

`func (o *AttachmentResourceAttributes) SetFileURL(v string)`

SetFileURL sets FileURL field to given value.


### SetFileURLNil

`func (o *AttachmentResourceAttributes) SetFileURLNil(b bool)`

 SetFileURLNil sets the value for FileURL to be an explicit nil

### UnsetFileURL
`func (o *AttachmentResourceAttributes) UnsetFileURL()`

UnsetFileURL ensures that no value is present for FileURL, not even an explicit nil
### GetFileURLExpiresAt

`func (o *AttachmentResourceAttributes) GetFileURLExpiresAt() time.Time`

GetFileURLExpiresAt returns the FileURLExpiresAt field if non-nil, zero value otherwise.

### GetFileURLExpiresAtOk

`func (o *AttachmentResourceAttributes) GetFileURLExpiresAtOk() (*time.Time, bool)`

GetFileURLExpiresAtOk returns a tuple with the FileURLExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileURLExpiresAt

`func (o *AttachmentResourceAttributes) SetFileURLExpiresAt(v time.Time)`

SetFileURLExpiresAt sets FileURLExpiresAt field to given value.


### GetFileExtension

`func (o *AttachmentResourceAttributes) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *AttachmentResourceAttributes) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *AttachmentResourceAttributes) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.


### SetFileExtensionNil

`func (o *AttachmentResourceAttributes) SetFileExtensionNil(b bool)`

 SetFileExtensionNil sets the value for FileExtension to be an explicit nil

### UnsetFileExtension
`func (o *AttachmentResourceAttributes) UnsetFileExtension()`

UnsetFileExtension ensures that no value is present for FileExtension, not even an explicit nil
### GetFileContentType

`func (o *AttachmentResourceAttributes) GetFileContentType() string`

GetFileContentType returns the FileContentType field if non-nil, zero value otherwise.

### GetFileContentTypeOk

`func (o *AttachmentResourceAttributes) GetFileContentTypeOk() (*string, bool)`

GetFileContentTypeOk returns a tuple with the FileContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContentType

`func (o *AttachmentResourceAttributes) SetFileContentType(v string)`

SetFileContentType sets FileContentType field to given value.


### SetFileContentTypeNil

`func (o *AttachmentResourceAttributes) SetFileContentTypeNil(b bool)`

 SetFileContentTypeNil sets the value for FileContentType to be an explicit nil

### UnsetFileContentType
`func (o *AttachmentResourceAttributes) UnsetFileContentType()`

UnsetFileContentType ensures that no value is present for FileContentType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


