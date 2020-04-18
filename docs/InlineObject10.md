# InlineObject10

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorId** | Pointer to **int32** | The identifier of the factor to enroll the user with | 
**DisplayName** | Pointer to **string** | A name for the users device | 
**Number** | Pointer to **string** | The phone number of the user in E.164 format. | 

## Methods

### NewInlineObject10

`func NewInlineObject10(factorId int32, displayName string, number string, ) *InlineObject10`

NewInlineObject10 instantiates a new InlineObject10 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject10WithDefaults

`func NewInlineObject10WithDefaults() *InlineObject10`

NewInlineObject10WithDefaults instantiates a new InlineObject10 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorId

`func (o *InlineObject10) GetFactorId() int32`

GetFactorId returns the FactorId field if non-nil, zero value otherwise.

### GetFactorIdOk

`func (o *InlineObject10) GetFactorIdOk() (*int32, bool)`

GetFactorIdOk returns a tuple with the FactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorId

`func (o *InlineObject10) SetFactorId(v int32)`

SetFactorId sets FactorId field to given value.


### GetDisplayName

`func (o *InlineObject10) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InlineObject10) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InlineObject10) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetNumber

`func (o *InlineObject10) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineObject10) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineObject10) SetNumber(v string)`

SetNumber sets Number field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


