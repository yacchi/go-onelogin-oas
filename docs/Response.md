# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ResponseUser**](Response_user.md) |  | [optional] 
**ReturnToUrl** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**SessionToken** | Pointer to **string** |  | [optional] 
**StateToken** | Pointer to **string** |  | [optional] 
**CallbackUrl** | Pointer to **string** |  | [optional] 
**Devices** | Pointer to [**[]ResponseDevices**](Response_devices.md) |  | [optional] 

## Methods

### NewResponse

`func NewResponse() *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *Response) GetUser() ResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Response) GetUserOk() (*ResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Response) SetUser(v ResponseUser)`

SetUser sets User field to given value.

### HasUser

`func (o *Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetReturnToUrl

`func (o *Response) GetReturnToUrl() string`

GetReturnToUrl returns the ReturnToUrl field if non-nil, zero value otherwise.

### GetReturnToUrlOk

`func (o *Response) GetReturnToUrlOk() (*string, bool)`

GetReturnToUrlOk returns a tuple with the ReturnToUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnToUrl

`func (o *Response) SetReturnToUrl(v string)`

SetReturnToUrl sets ReturnToUrl field to given value.

### HasReturnToUrl

`func (o *Response) HasReturnToUrl() bool`

HasReturnToUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Response) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Response) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Response) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSessionToken

`func (o *Response) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *Response) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *Response) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *Response) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetStateToken

`func (o *Response) GetStateToken() string`

GetStateToken returns the StateToken field if non-nil, zero value otherwise.

### GetStateTokenOk

`func (o *Response) GetStateTokenOk() (*string, bool)`

GetStateTokenOk returns a tuple with the StateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateToken

`func (o *Response) SetStateToken(v string)`

SetStateToken sets StateToken field to given value.

### HasStateToken

`func (o *Response) HasStateToken() bool`

HasStateToken returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *Response) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *Response) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *Response) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *Response) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetDevices

`func (o *Response) GetDevices() []ResponseDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *Response) GetDevicesOk() (*[]ResponseDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *Response) SetDevices(v []ResponseDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *Response) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


