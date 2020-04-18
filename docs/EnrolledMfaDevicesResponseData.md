# EnrolledMfaDevicesResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpDevices** | Pointer to [**[]MfaDevice**](MfaDevice.md) |  | [optional] 

## Methods

### NewEnrolledMfaDevicesResponseData

`func NewEnrolledMfaDevicesResponseData() *EnrolledMfaDevicesResponseData`

NewEnrolledMfaDevicesResponseData instantiates a new EnrolledMfaDevicesResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrolledMfaDevicesResponseDataWithDefaults

`func NewEnrolledMfaDevicesResponseDataWithDefaults() *EnrolledMfaDevicesResponseData`

NewEnrolledMfaDevicesResponseDataWithDefaults instantiates a new EnrolledMfaDevicesResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpDevices

`func (o *EnrolledMfaDevicesResponseData) GetOtpDevices() []MfaDevice`

GetOtpDevices returns the OtpDevices field if non-nil, zero value otherwise.

### GetOtpDevicesOk

`func (o *EnrolledMfaDevicesResponseData) GetOtpDevicesOk() (*[]MfaDevice, bool)`

GetOtpDevicesOk returns a tuple with the OtpDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpDevices

`func (o *EnrolledMfaDevicesResponseData) SetOtpDevices(v []MfaDevice)`

SetOtpDevices sets OtpDevices field to given value.

### HasOtpDevices

`func (o *EnrolledMfaDevicesResponseData) HasOtpDevices() bool`

HasOtpDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


