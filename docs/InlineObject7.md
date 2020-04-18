# InlineObject7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int32** | Provide the MFA device_id you are submitting for verification. | 
**StateToken** | Pointer to **string** | Provide the state_token associated with the MFA device_id you are submitting for verification. | 
**OtpToken** | Pointer to **string** | Provide the OTP value for the MFA factor you are submitting for verification. | [optional] 
**DoNotNotify** | Pointer to **bool** | When verifying MFA via Protect Push, set this to true to stop additional push notifications being sent to the OneLogin Protect device. | [optional] [default to false]

## Methods

### NewInlineObject7

`func NewInlineObject7(deviceId int32, stateToken string, ) *InlineObject7`

NewInlineObject7 instantiates a new InlineObject7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject7WithDefaults

`func NewInlineObject7WithDefaults() *InlineObject7`

NewInlineObject7WithDefaults instantiates a new InlineObject7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *InlineObject7) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *InlineObject7) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *InlineObject7) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.


### GetStateToken

`func (o *InlineObject7) GetStateToken() string`

GetStateToken returns the StateToken field if non-nil, zero value otherwise.

### GetStateTokenOk

`func (o *InlineObject7) GetStateTokenOk() (*string, bool)`

GetStateTokenOk returns a tuple with the StateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateToken

`func (o *InlineObject7) SetStateToken(v string)`

SetStateToken sets StateToken field to given value.


### GetOtpToken

`func (o *InlineObject7) GetOtpToken() string`

GetOtpToken returns the OtpToken field if non-nil, zero value otherwise.

### GetOtpTokenOk

`func (o *InlineObject7) GetOtpTokenOk() (*string, bool)`

GetOtpTokenOk returns a tuple with the OtpToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpToken

`func (o *InlineObject7) SetOtpToken(v string)`

SetOtpToken sets OtpToken field to given value.

### HasOtpToken

`func (o *InlineObject7) HasOtpToken() bool`

HasOtpToken returns a boolean if a field has been set.

### GetDoNotNotify

`func (o *InlineObject7) GetDoNotNotify() bool`

GetDoNotNotify returns the DoNotNotify field if non-nil, zero value otherwise.

### GetDoNotNotifyOk

`func (o *InlineObject7) GetDoNotNotifyOk() (*bool, bool)`

GetDoNotNotifyOk returns a tuple with the DoNotNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotNotify

`func (o *InlineObject7) SetDoNotNotify(v bool)`

SetDoNotNotify sets DoNotNotify field to given value.

### HasDoNotNotify

`func (o *InlineObject7) HasDoNotNotify() bool`

HasDoNotNotify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


