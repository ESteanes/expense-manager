# ErrorObjectSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameter** | Pointer to **string** | If this error relates to a query parameter, the name of the parameter.  | [optional] 
**Pointer** | Pointer to **string** | If this error relates to an attribute in the request body, a rfc-6901 JSON pointer to the attribute.  | [optional] 

## Methods

### NewErrorObjectSource

`func NewErrorObjectSource() *ErrorObjectSource`

NewErrorObjectSource instantiates a new ErrorObjectSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorObjectSourceWithDefaults

`func NewErrorObjectSourceWithDefaults() *ErrorObjectSource`

NewErrorObjectSourceWithDefaults instantiates a new ErrorObjectSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameter

`func (o *ErrorObjectSource) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ErrorObjectSource) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ErrorObjectSource) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *ErrorObjectSource) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetPointer

`func (o *ErrorObjectSource) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *ErrorObjectSource) GetPointerOk() (*string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *ErrorObjectSource) SetPointer(v string)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *ErrorObjectSource) HasPointer() bool`

HasPointer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


