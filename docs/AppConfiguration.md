# AppConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUri** | Pointer to **string** |  | [optional] 
**RefreshTokenExpirationMinutes** | Pointer to **int32** |  | [optional] 
**LoginUrl** | Pointer to **string** |  | [optional] 
**OidcApplicationType** | Pointer to **int32** |  | [optional] 
**TokenEndpointAuthMethod** | Pointer to **int32** |  | [optional] 
**AccessTokenExpirationMinutes** | Pointer to **int32** |  | [optional] 
**ProviderArn** | Pointer to **string** |  | [optional] 
**SignatureAlgorithm** | Pointer to **string** |  | [optional] 

## Methods

### NewAppConfiguration

`func NewAppConfiguration() *AppConfiguration`

NewAppConfiguration instantiates a new AppConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConfigurationWithDefaults

`func NewAppConfigurationWithDefaults() *AppConfiguration`

NewAppConfigurationWithDefaults instantiates a new AppConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUri

`func (o *AppConfiguration) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *AppConfiguration) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *AppConfiguration) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *AppConfiguration) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetRefreshTokenExpirationMinutes

`func (o *AppConfiguration) GetRefreshTokenExpirationMinutes() int32`

GetRefreshTokenExpirationMinutes returns the RefreshTokenExpirationMinutes field if non-nil, zero value otherwise.

### GetRefreshTokenExpirationMinutesOk

`func (o *AppConfiguration) GetRefreshTokenExpirationMinutesOk() (*int32, bool)`

GetRefreshTokenExpirationMinutesOk returns a tuple with the RefreshTokenExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpirationMinutes

`func (o *AppConfiguration) SetRefreshTokenExpirationMinutes(v int32)`

SetRefreshTokenExpirationMinutes sets RefreshTokenExpirationMinutes field to given value.

### HasRefreshTokenExpirationMinutes

`func (o *AppConfiguration) HasRefreshTokenExpirationMinutes() bool`

HasRefreshTokenExpirationMinutes returns a boolean if a field has been set.

### GetLoginUrl

`func (o *AppConfiguration) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *AppConfiguration) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *AppConfiguration) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *AppConfiguration) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetOidcApplicationType

`func (o *AppConfiguration) GetOidcApplicationType() int32`

GetOidcApplicationType returns the OidcApplicationType field if non-nil, zero value otherwise.

### GetOidcApplicationTypeOk

`func (o *AppConfiguration) GetOidcApplicationTypeOk() (*int32, bool)`

GetOidcApplicationTypeOk returns a tuple with the OidcApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcApplicationType

`func (o *AppConfiguration) SetOidcApplicationType(v int32)`

SetOidcApplicationType sets OidcApplicationType field to given value.

### HasOidcApplicationType

`func (o *AppConfiguration) HasOidcApplicationType() bool`

HasOidcApplicationType returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *AppConfiguration) GetTokenEndpointAuthMethod() int32`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *AppConfiguration) GetTokenEndpointAuthMethodOk() (*int32, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *AppConfiguration) SetTokenEndpointAuthMethod(v int32)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.

### HasTokenEndpointAuthMethod

`func (o *AppConfiguration) HasTokenEndpointAuthMethod() bool`

HasTokenEndpointAuthMethod returns a boolean if a field has been set.

### GetAccessTokenExpirationMinutes

`func (o *AppConfiguration) GetAccessTokenExpirationMinutes() int32`

GetAccessTokenExpirationMinutes returns the AccessTokenExpirationMinutes field if non-nil, zero value otherwise.

### GetAccessTokenExpirationMinutesOk

`func (o *AppConfiguration) GetAccessTokenExpirationMinutesOk() (*int32, bool)`

GetAccessTokenExpirationMinutesOk returns a tuple with the AccessTokenExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpirationMinutes

`func (o *AppConfiguration) SetAccessTokenExpirationMinutes(v int32)`

SetAccessTokenExpirationMinutes sets AccessTokenExpirationMinutes field to given value.

### HasAccessTokenExpirationMinutes

`func (o *AppConfiguration) HasAccessTokenExpirationMinutes() bool`

HasAccessTokenExpirationMinutes returns a boolean if a field has been set.

### GetProviderArn

`func (o *AppConfiguration) GetProviderArn() string`

GetProviderArn returns the ProviderArn field if non-nil, zero value otherwise.

### GetProviderArnOk

`func (o *AppConfiguration) GetProviderArnOk() (*string, bool)`

GetProviderArnOk returns a tuple with the ProviderArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderArn

`func (o *AppConfiguration) SetProviderArn(v string)`

SetProviderArn sets ProviderArn field to given value.

### HasProviderArn

`func (o *AppConfiguration) HasProviderArn() bool`

HasProviderArn returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *AppConfiguration) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *AppConfiguration) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *AppConfiguration) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *AppConfiguration) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


