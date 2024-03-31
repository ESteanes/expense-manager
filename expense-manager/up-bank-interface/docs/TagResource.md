# TagResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this resource: &#x60;tags&#x60; | 
**Id** | **string** | The label of the tag, which also acts as the tag’s unique identifier.  | 
**Relationships** | [**AccountResourceRelationships**](AccountResourceRelationships.md) |  | 

## Methods

### NewTagResource

`func NewTagResource(type_ string, id string, relationships AccountResourceRelationships, ) *TagResource`

NewTagResource instantiates a new TagResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagResourceWithDefaults

`func NewTagResourceWithDefaults() *TagResource`

NewTagResourceWithDefaults instantiates a new TagResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TagResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagResource) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *TagResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagResource) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *TagResource) GetRelationships() AccountResourceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TagResource) GetRelationshipsOk() (*AccountResourceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TagResource) SetRelationships(v AccountResourceRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


