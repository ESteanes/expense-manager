# ListAccountsResponseLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prev** | **NullableString** | The link to the previous page in the results. If this value is &#x60;null&#x60; there is no previous page.  | 
**Next** | **NullableString** | The link to the next page in the results. If this value is &#x60;null&#x60; there is no next page.  | 

## Methods

### NewListAccountsResponseLinks

`func NewListAccountsResponseLinks(prev NullableString, next NullableString, ) *ListAccountsResponseLinks`

NewListAccountsResponseLinks instantiates a new ListAccountsResponseLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccountsResponseLinksWithDefaults

`func NewListAccountsResponseLinksWithDefaults() *ListAccountsResponseLinks`

NewListAccountsResponseLinksWithDefaults instantiates a new ListAccountsResponseLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrev

`func (o *ListAccountsResponseLinks) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *ListAccountsResponseLinks) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *ListAccountsResponseLinks) SetPrev(v string)`

SetPrev sets Prev field to given value.


### SetPrevNil

`func (o *ListAccountsResponseLinks) SetPrevNil(b bool)`

 SetPrevNil sets the value for Prev to be an explicit nil

### UnsetPrev
`func (o *ListAccountsResponseLinks) UnsetPrev()`

UnsetPrev ensures that no value is present for Prev, not even an explicit nil
### GetNext

`func (o *ListAccountsResponseLinks) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListAccountsResponseLinks) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListAccountsResponseLinks) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *ListAccountsResponseLinks) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *ListAccountsResponseLinks) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


