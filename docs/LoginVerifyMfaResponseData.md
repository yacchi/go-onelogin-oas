# LoginVerifyMfaResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ResponseUser**](Response_user.md) |  | [optional] 
**ReturnToUrl** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**SessionToken** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginVerifyMfaResponseData

`func NewLoginVerifyMfaResponseData() *LoginVerifyMfaResponseData`

NewLoginVerifyMfaResponseData instantiates a new LoginVerifyMfaResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginVerifyMfaResponseDataWithDefaults

`func NewLoginVerifyMfaResponseDataWithDefaults() *LoginVerifyMfaResponseData`

NewLoginVerifyMfaResponseDataWithDefaults instantiates a new LoginVerifyMfaResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LoginVerifyMfaResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoginVerifyMfaResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoginVerifyMfaResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LoginVerifyMfaResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *LoginVerifyMfaResponseData) GetUser() ResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoginVerifyMfaResponseData) GetUserOk() (*ResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoginVerifyMfaResponseData) SetUser(v ResponseUser)`

SetUser sets User field to given value.

### HasUser

`func (o *LoginVerifyMfaResponseData) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetReturnToUrl

`func (o *LoginVerifyMfaResponseData) GetReturnToUrl() string`

GetReturnToUrl returns the ReturnToUrl field if non-nil, zero value otherwise.

### GetReturnToUrlOk

`func (o *LoginVerifyMfaResponseData) GetReturnToUrlOk() (*string, bool)`

GetReturnToUrlOk returns a tuple with the ReturnToUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnToUrl

`func (o *LoginVerifyMfaResponseData) SetReturnToUrl(v string)`

SetReturnToUrl sets ReturnToUrl field to given value.

### HasReturnToUrl

`func (o *LoginVerifyMfaResponseData) HasReturnToUrl() bool`

HasReturnToUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *LoginVerifyMfaResponseData) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *LoginVerifyMfaResponseData) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *LoginVerifyMfaResponseData) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *LoginVerifyMfaResponseData) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSessionToken

`func (o *LoginVerifyMfaResponseData) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *LoginVerifyMfaResponseData) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *LoginVerifyMfaResponseData) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *LoginVerifyMfaResponseData) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


