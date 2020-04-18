# InlineObject6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsernameOrEmail** | Pointer to **string** | Set to the username or email of the user that you want to log in. | 
**Password** | Pointer to **string** | Set to the password of the user that you want to log in. | 
**Subdomain** | Pointer to **string** | Set to the subdomain of the user that you want to log in. | 
**ReturnToUrl** | Pointer to **string** | Leave this value blank for now. Intended for future use with multi-factor authentication functionality. | [optional] 
**IpAddress** | Pointer to **string** | Leave this value blank for now. Intended for future use with multi-factor authentication functionality. It will be used to set to the IP address of the user accessing your login page. | [optional] 
**BrowserId** | Pointer to **string** | Leave this value blank for now. Intended for future use with multi-factor authentication functionality. It will be used to set to the ID of the browser being used by the user to access your login page. | [optional] 

## Methods

### NewInlineObject6

`func NewInlineObject6(usernameOrEmail string, password string, subdomain string, ) *InlineObject6`

NewInlineObject6 instantiates a new InlineObject6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject6WithDefaults

`func NewInlineObject6WithDefaults() *InlineObject6`

NewInlineObject6WithDefaults instantiates a new InlineObject6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsernameOrEmail

`func (o *InlineObject6) GetUsernameOrEmail() string`

GetUsernameOrEmail returns the UsernameOrEmail field if non-nil, zero value otherwise.

### GetUsernameOrEmailOk

`func (o *InlineObject6) GetUsernameOrEmailOk() (*string, bool)`

GetUsernameOrEmailOk returns a tuple with the UsernameOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameOrEmail

`func (o *InlineObject6) SetUsernameOrEmail(v string)`

SetUsernameOrEmail sets UsernameOrEmail field to given value.


### GetPassword

`func (o *InlineObject6) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject6) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject6) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSubdomain

`func (o *InlineObject6) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *InlineObject6) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *InlineObject6) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.


### GetReturnToUrl

`func (o *InlineObject6) GetReturnToUrl() string`

GetReturnToUrl returns the ReturnToUrl field if non-nil, zero value otherwise.

### GetReturnToUrlOk

`func (o *InlineObject6) GetReturnToUrlOk() (*string, bool)`

GetReturnToUrlOk returns a tuple with the ReturnToUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnToUrl

`func (o *InlineObject6) SetReturnToUrl(v string)`

SetReturnToUrl sets ReturnToUrl field to given value.

### HasReturnToUrl

`func (o *InlineObject6) HasReturnToUrl() bool`

HasReturnToUrl returns a boolean if a field has been set.

### GetIpAddress

`func (o *InlineObject6) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InlineObject6) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InlineObject6) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InlineObject6) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetBrowserId

`func (o *InlineObject6) GetBrowserId() string`

GetBrowserId returns the BrowserId field if non-nil, zero value otherwise.

### GetBrowserIdOk

`func (o *InlineObject6) GetBrowserIdOk() (*string, bool)`

GetBrowserIdOk returns a tuple with the BrowserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserId

`func (o *InlineObject6) SetBrowserId(v string)`

SetBrowserId sets BrowserId field to given value.

### HasBrowserId

`func (o *InlineObject6) HasBrowserId() bool`

HasBrowserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


