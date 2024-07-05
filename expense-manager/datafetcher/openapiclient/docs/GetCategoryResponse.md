# GetCategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CategoryResource**](CategoryResource.md) | The category returned in this response.  | 

## Methods

### NewGetCategoryResponse

`func NewGetCategoryResponse(data CategoryResource, ) *GetCategoryResponse`

NewGetCategoryResponse instantiates a new GetCategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCategoryResponseWithDefaults

`func NewGetCategoryResponseWithDefaults() *GetCategoryResponse`

NewGetCategoryResponseWithDefaults instantiates a new GetCategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCategoryResponse) GetData() CategoryResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCategoryResponse) GetDataOk() (*CategoryResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCategoryResponse) SetData(v CategoryResource)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


