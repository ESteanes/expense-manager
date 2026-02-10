# Merchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Emoji** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMerchant

`func NewMerchant() *Merchant`

NewMerchant instantiates a new Merchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantWithDefaults

`func NewMerchantWithDefaults() *Merchant`

NewMerchantWithDefaults instantiates a new Merchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Merchant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Merchant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Merchant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Merchant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *Merchant) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Merchant) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Merchant) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Merchant) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *Merchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Merchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Merchant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Merchant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogo

`func (o *Merchant) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Merchant) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Merchant) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Merchant) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetEmoji

`func (o *Merchant) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *Merchant) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *Merchant) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *Merchant) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetCategory

`func (o *Merchant) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Merchant) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Merchant) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Merchant) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetAddress

`func (o *Merchant) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Merchant) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Merchant) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Merchant) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCreated

`func (o *Merchant) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Merchant) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Merchant) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Merchant) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


