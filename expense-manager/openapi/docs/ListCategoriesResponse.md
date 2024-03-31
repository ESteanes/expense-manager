# ListCategoriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CategoryResource**](CategoryResource.md) | The list of categories returned in this response.  | 

## Methods

### NewListCategoriesResponse

`func NewListCategoriesResponse(data []CategoryResource, ) *ListCategoriesResponse`

NewListCategoriesResponse instantiates a new ListCategoriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCategoriesResponseWithDefaults

`func NewListCategoriesResponseWithDefaults() *ListCategoriesResponse`

NewListCategoriesResponseWithDefaults instantiates a new ListCategoriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCategoriesResponse) GetData() []CategoryResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCategoriesResponse) GetDataOk() (*[]CategoryResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCategoriesResponse) SetData(v []CategoryResource)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


