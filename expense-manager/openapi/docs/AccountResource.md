# AccountResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this resource: &#x60;accounts&#x60; | 
**Id** | **string** | The unique identifier for this account.  | 
**Attributes** | [**AccountResourceAttributes**](AccountResourceAttributes.md) |  | 
**Relationships** | [**AccountResourceRelationships**](AccountResourceRelationships.md) |  | 
**Links** | Pointer to [**AccountResourceLinks**](AccountResourceLinks.md) |  | [optional] 

## Methods

### NewAccountResource

`func NewAccountResource(type_ string, id string, attributes AccountResourceAttributes, relationships AccountResourceRelationships, ) *AccountResource`

NewAccountResource instantiates a new AccountResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResourceWithDefaults

`func NewAccountResourceWithDefaults() *AccountResource`

NewAccountResourceWithDefaults instantiates a new AccountResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccountResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountResource) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AccountResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AccountResource) GetAttributes() AccountResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AccountResource) GetAttributesOk() (*AccountResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AccountResource) SetAttributes(v AccountResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AccountResource) GetRelationships() AccountResourceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AccountResource) GetRelationshipsOk() (*AccountResourceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AccountResource) SetRelationships(v AccountResourceRelationships)`

SetRelationships sets Relationships field to given value.


### GetLinks

`func (o *AccountResource) GetLinks() AccountResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccountResource) GetLinksOk() (*AccountResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccountResource) SetLinks(v AccountResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AccountResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


