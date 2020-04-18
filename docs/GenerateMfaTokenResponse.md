# GenerateMfaTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reusable** | Pointer to **bool** |  | [optional] 
**MfaToken** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewGenerateMfaTokenResponse

`func NewGenerateMfaTokenResponse() *GenerateMfaTokenResponse`

NewGenerateMfaTokenResponse instantiates a new GenerateMfaTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateMfaTokenResponseWithDefaults

`func NewGenerateMfaTokenResponseWithDefaults() *GenerateMfaTokenResponse`

NewGenerateMfaTokenResponseWithDefaults instantiates a new GenerateMfaTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReusable

`func (o *GenerateMfaTokenResponse) GetReusable() bool`

GetReusable returns the Reusable field if non-nil, zero value otherwise.

### GetReusableOk

`func (o *GenerateMfaTokenResponse) GetReusableOk() (*bool, bool)`

GetReusableOk returns a tuple with the Reusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusable

`func (o *GenerateMfaTokenResponse) SetReusable(v bool)`

SetReusable sets Reusable field to given value.

### HasReusable

`func (o *GenerateMfaTokenResponse) HasReusable() bool`

HasReusable returns a boolean if a field has been set.

### GetMfaToken

`func (o *GenerateMfaTokenResponse) GetMfaToken() string`

GetMfaToken returns the MfaToken field if non-nil, zero value otherwise.

### GetMfaTokenOk

`func (o *GenerateMfaTokenResponse) GetMfaTokenOk() (*string, bool)`

GetMfaTokenOk returns a tuple with the MfaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaToken

`func (o *GenerateMfaTokenResponse) SetMfaToken(v string)`

SetMfaToken sets MfaToken field to given value.

### HasMfaToken

`func (o *GenerateMfaTokenResponse) HasMfaToken() bool`

HasMfaToken returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GenerateMfaTokenResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GenerateMfaTokenResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GenerateMfaTokenResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GenerateMfaTokenResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


