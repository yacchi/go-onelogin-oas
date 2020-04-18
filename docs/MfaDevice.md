# MfaDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**StateToken** | Pointer to **string** |  | [optional] 
**AuthFactorName** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**TypeDisplayName** | Pointer to **string** |  | [optional] 
**NeedsTrigger** | Pointer to **bool** |  | [optional] 
**UserDisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 

## Methods

### NewMfaDevice

`func NewMfaDevice() *MfaDevice`

NewMfaDevice instantiates a new MfaDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaDeviceWithDefaults

`func NewMfaDeviceWithDefaults() *MfaDevice`

NewMfaDeviceWithDefaults instantiates a new MfaDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *MfaDevice) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MfaDevice) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MfaDevice) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MfaDevice) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefault

`func (o *MfaDevice) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *MfaDevice) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *MfaDevice) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *MfaDevice) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetStateToken

`func (o *MfaDevice) GetStateToken() string`

GetStateToken returns the StateToken field if non-nil, zero value otherwise.

### GetStateTokenOk

`func (o *MfaDevice) GetStateTokenOk() (*string, bool)`

GetStateTokenOk returns a tuple with the StateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateToken

`func (o *MfaDevice) SetStateToken(v string)`

SetStateToken sets StateToken field to given value.

### HasStateToken

`func (o *MfaDevice) HasStateToken() bool`

HasStateToken returns a boolean if a field has been set.

### GetAuthFactorName

`func (o *MfaDevice) GetAuthFactorName() string`

GetAuthFactorName returns the AuthFactorName field if non-nil, zero value otherwise.

### GetAuthFactorNameOk

`func (o *MfaDevice) GetAuthFactorNameOk() (*string, bool)`

GetAuthFactorNameOk returns a tuple with the AuthFactorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFactorName

`func (o *MfaDevice) SetAuthFactorName(v string)`

SetAuthFactorName sets AuthFactorName field to given value.

### HasAuthFactorName

`func (o *MfaDevice) HasAuthFactorName() bool`

HasAuthFactorName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *MfaDevice) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MfaDevice) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MfaDevice) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *MfaDevice) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetTypeDisplayName

`func (o *MfaDevice) GetTypeDisplayName() string`

GetTypeDisplayName returns the TypeDisplayName field if non-nil, zero value otherwise.

### GetTypeDisplayNameOk

`func (o *MfaDevice) GetTypeDisplayNameOk() (*string, bool)`

GetTypeDisplayNameOk returns a tuple with the TypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplayName

`func (o *MfaDevice) SetTypeDisplayName(v string)`

SetTypeDisplayName sets TypeDisplayName field to given value.

### HasTypeDisplayName

`func (o *MfaDevice) HasTypeDisplayName() bool`

HasTypeDisplayName returns a boolean if a field has been set.

### GetNeedsTrigger

`func (o *MfaDevice) GetNeedsTrigger() bool`

GetNeedsTrigger returns the NeedsTrigger field if non-nil, zero value otherwise.

### GetNeedsTriggerOk

`func (o *MfaDevice) GetNeedsTriggerOk() (*bool, bool)`

GetNeedsTriggerOk returns a tuple with the NeedsTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsTrigger

`func (o *MfaDevice) SetNeedsTrigger(v bool)`

SetNeedsTrigger sets NeedsTrigger field to given value.

### HasNeedsTrigger

`func (o *MfaDevice) HasNeedsTrigger() bool`

HasNeedsTrigger returns a boolean if a field has been set.

### GetUserDisplayName

`func (o *MfaDevice) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *MfaDevice) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *MfaDevice) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *MfaDevice) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### GetId

`func (o *MfaDevice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MfaDevice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MfaDevice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MfaDevice) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


