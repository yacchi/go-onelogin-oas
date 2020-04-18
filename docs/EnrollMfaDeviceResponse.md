# EnrollMfaDeviceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to [**[]MfaDevice**](MfaDevice.md) |  | [optional] 

## Methods

### NewEnrollMfaDeviceResponse

`func NewEnrollMfaDeviceResponse() *EnrollMfaDeviceResponse`

NewEnrollMfaDeviceResponse instantiates a new EnrollMfaDeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollMfaDeviceResponseWithDefaults

`func NewEnrollMfaDeviceResponseWithDefaults() *EnrollMfaDeviceResponse`

NewEnrollMfaDeviceResponseWithDefaults instantiates a new EnrollMfaDeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EnrollMfaDeviceResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnrollMfaDeviceResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnrollMfaDeviceResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnrollMfaDeviceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *EnrollMfaDeviceResponse) GetData() []MfaDevice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnrollMfaDeviceResponse) GetDataOk() (*[]MfaDevice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnrollMfaDeviceResponse) SetData(v []MfaDevice)`

SetData sets Data field to given value.

### HasData

`func (o *EnrollMfaDeviceResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


