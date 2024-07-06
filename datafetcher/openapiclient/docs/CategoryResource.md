# CategoryResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this resource: &#x60;categories&#x60; | 
**Id** | **string** | The unique identifier for this category. This is a human-readable but URL-safe value.  | 
**Attributes** | [**CategoryResourceAttributes**](CategoryResourceAttributes.md) |  | 
**Relationships** | [**CategoryResourceRelationships**](CategoryResourceRelationships.md) |  | 
**Links** | Pointer to [**AccountResourceLinks**](AccountResourceLinks.md) |  | [optional] 

## Methods

### NewCategoryResource

`func NewCategoryResource(type_ string, id string, attributes CategoryResourceAttributes, relationships CategoryResourceRelationships, ) *CategoryResource`

NewCategoryResource instantiates a new CategoryResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryResourceWithDefaults

`func NewCategoryResourceWithDefaults() *CategoryResource`

NewCategoryResourceWithDefaults instantiates a new CategoryResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CategoryResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CategoryResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CategoryResource) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CategoryResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CategoryResource) GetAttributes() CategoryResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CategoryResource) GetAttributesOk() (*CategoryResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CategoryResource) SetAttributes(v CategoryResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CategoryResource) GetRelationships() CategoryResourceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CategoryResource) GetRelationshipsOk() (*CategoryResourceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CategoryResource) SetRelationships(v CategoryResourceRelationships)`

SetRelationships sets Relationships field to given value.


### GetLinks

`func (o *CategoryResource) GetLinks() AccountResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CategoryResource) GetLinksOk() (*AccountResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CategoryResource) SetLinks(v AccountResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CategoryResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


