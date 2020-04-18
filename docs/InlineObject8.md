# InlineObject8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsernameOrEmail** | Pointer to **string** | Set to the username or email of the user that you want to log in. | 
**Password** | Pointer to **string** | Set to the password of the user that you want to log in. | 
**Subdomain** | Pointer to **string** | Set to the subdomain of the user that you want to log in. | 
**AppId** | Pointer to **int32** | App ID of the app for which you want to generate a SAML token. This is the app ID in OneLogin. | 
**IpAddress** | Pointer to **string** | If you are using this API in a scenario in which MFA is required and you’ll need to be able to honor IP address whitelisting defined in MFA policies, provide this parameter and set its value to the whitelisted IP address that needs to be bypassed. | [optional] 

## Methods

### NewInlineObject8

`func NewInlineObject8(usernameOrEmail string, password string, subdomain string, appId int32, ) *InlineObject8`

NewInlineObject8 instantiates a new InlineObject8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject8WithDefaults

`func NewInlineObject8WithDefaults() *InlineObject8`

NewInlineObject8WithDefaults instantiates a new InlineObject8 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsernameOrEmail

`func (o *InlineObject8) GetUsernameOrEmail() string`

GetUsernameOrEmail returns the UsernameOrEmail field if non-nil, zero value otherwise.

### GetUsernameOrEmailOk

`func (o *InlineObject8) GetUsernameOrEmailOk() (*string, bool)`

GetUsernameOrEmailOk returns a tuple with the UsernameOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameOrEmail

`func (o *InlineObject8) SetUsernameOrEmail(v string)`

SetUsernameOrEmail sets UsernameOrEmail field to given value.


### GetPassword

`func (o *InlineObject8) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject8) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject8) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSubdomain

`func (o *InlineObject8) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *InlineObject8) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *InlineObject8) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.


### GetAppId

`func (o *InlineObject8) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *InlineObject8) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *InlineObject8) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetIpAddress

`func (o *InlineObject8) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InlineObject8) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InlineObject8) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InlineObject8) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


