# AppSso

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**MetadataUrl** | Pointer to **string** |  | [optional] 
**AcsUrl** | Pointer to **string** |  | [optional] 
**SlsUrl** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to [**AppSsoCertificate**](App_sso_certificate.md) |  | [optional] 

## Methods

### NewAppSso

`func NewAppSso() *AppSso`

NewAppSso instantiates a new AppSso object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSsoWithDefaults

`func NewAppSsoWithDefaults() *AppSso`

NewAppSsoWithDefaults instantiates a new AppSso object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AppSso) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AppSso) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AppSso) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AppSso) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AppSso) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AppSso) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AppSso) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AppSso) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetMetadataUrl

`func (o *AppSso) GetMetadataUrl() string`

GetMetadataUrl returns the MetadataUrl field if non-nil, zero value otherwise.

### GetMetadataUrlOk

`func (o *AppSso) GetMetadataUrlOk() (*string, bool)`

GetMetadataUrlOk returns a tuple with the MetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUrl

`func (o *AppSso) SetMetadataUrl(v string)`

SetMetadataUrl sets MetadataUrl field to given value.

### HasMetadataUrl

`func (o *AppSso) HasMetadataUrl() bool`

HasMetadataUrl returns a boolean if a field has been set.

### GetAcsUrl

`func (o *AppSso) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *AppSso) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *AppSso) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.

### HasAcsUrl

`func (o *AppSso) HasAcsUrl() bool`

HasAcsUrl returns a boolean if a field has been set.

### GetSlsUrl

`func (o *AppSso) GetSlsUrl() string`

GetSlsUrl returns the SlsUrl field if non-nil, zero value otherwise.

### GetSlsUrlOk

`func (o *AppSso) GetSlsUrlOk() (*string, bool)`

GetSlsUrlOk returns a tuple with the SlsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlsUrl

`func (o *AppSso) SetSlsUrl(v string)`

SetSlsUrl sets SlsUrl field to given value.

### HasSlsUrl

`func (o *AppSso) HasSlsUrl() bool`

HasSlsUrl returns a boolean if a field has been set.

### GetIssuer

`func (o *AppSso) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AppSso) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AppSso) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *AppSso) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetCertificate

`func (o *AppSso) GetCertificate() AppSsoCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *AppSso) GetCertificateOk() (*AppSsoCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *AppSso) SetCertificate(v AppSsoCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *AppSso) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


