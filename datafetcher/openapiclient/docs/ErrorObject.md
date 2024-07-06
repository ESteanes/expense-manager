# ErrorObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The HTTP status code associated with this error. This can also be obtained from the response headers. The status indicates the broad type of error according to HTTP semantics.  | 
**Title** | **string** | A short description of this error. This should be stable across multiple occurrences of this type of error and typically expands on the reason for the status code.  | 
**Detail** | **string** | A detailed description of this error. This should be considered unique to individual occurrences of an error and subject to change. It is useful for debugging purposes.  | 
**Source** | Pointer to [**ErrorObjectSource**](ErrorObjectSource.md) |  | [optional] 

## Methods

### NewErrorObject

`func NewErrorObject(status string, title string, detail string, ) *ErrorObject`

NewErrorObject instantiates a new ErrorObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorObjectWithDefaults

`func NewErrorObjectWithDefaults() *ErrorObject`

NewErrorObjectWithDefaults instantiates a new ErrorObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ErrorObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorObject) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *ErrorObject) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorObject) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorObject) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ErrorObject) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorObject) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorObject) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetSource

`func (o *ErrorObject) GetSource() ErrorObjectSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ErrorObject) GetSourceOk() (*ErrorObjectSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ErrorObject) SetSource(v ErrorObjectSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ErrorObject) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


