# InlineObject11

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpToken** | Pointer to **int32** | OTP code provided by the device or SMS message sent to user. | [optional] 
**StateToken** | Pointer to **string** | The state_token is returned after a successful request to Enroll a Factor or Activate a Factor. The state_token MUST be provided if the needs_trigger attribute from the proceeding calls is set to true. | [optional] 

## Methods

### NewInlineObject11

`func NewInlineObject11() *InlineObject11`

NewInlineObject11 instantiates a new InlineObject11 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject11WithDefaults

`func NewInlineObject11WithDefaults() *InlineObject11`

NewInlineObject11WithDefaults instantiates a new InlineObject11 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpToken

`func (o *InlineObject11) GetOtpToken() int32`

GetOtpToken returns the OtpToken field if non-nil, zero value otherwise.

### GetOtpTokenOk

`func (o *InlineObject11) GetOtpTokenOk() (*int32, bool)`

GetOtpTokenOk returns a tuple with the OtpToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpToken

`func (o *InlineObject11) SetOtpToken(v int32)`

SetOtpToken sets OtpToken field to given value.

### HasOtpToken

`func (o *InlineObject11) HasOtpToken() bool`

HasOtpToken returns a boolean if a field has been set.

### GetStateToken

`func (o *InlineObject11) GetStateToken() string`

GetStateToken returns the StateToken field if non-nil, zero value otherwise.

### GetStateTokenOk

`func (o *InlineObject11) GetStateTokenOk() (*string, bool)`

GetStateTokenOk returns a tuple with the StateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateToken

`func (o *InlineObject11) SetStateToken(v string)`

SetStateToken sets StateToken field to given value.

### HasStateToken

`func (o *InlineObject11) HasStateToken() bool`

HasStateToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


