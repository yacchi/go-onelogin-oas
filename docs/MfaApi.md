# \MfaApi

All URIs are relative to *https://api.us.onelogin.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUserMFADevice**](MfaApi.md#ActivateUserMFADevice) | **Post** /1/users/{id}/otp_devices/{device_id}/trigger | Activate an authentication factor
[**DeleteUserMFADevice**](MfaApi.md#DeleteUserMFADevice) | **Delete** /1/users/{id}/otp_devices/{device_id} | Remove an authentication device
[**EnrollUserMFADevice**](MfaApi.md#EnrollUserMFADevice) | **Post** /1/users/{id}/otp_devices | Enroll a user with a given authentication factor.
[**GetUserAvailableMFAFactors**](MfaApi.md#GetUserAvailableMFAFactors) | **Get** /1/users/{id}/auth_factors | Get a list of MFA factors available to user
[**GetUserEnrolledMFADevices**](MfaApi.md#GetUserEnrolledMFADevices) | **Get** /1/users/{id}/otp_devices | Get enrolled authentication devices
[**VerifyUserMFADevice**](MfaApi.md#VerifyUserMFADevice) | **Post** /1/users/{id}/otp_devices/{device_id}/verify | Verify an authentication device



## ActivateUserMFADevice

> models.EnrollMfaDeviceResponse ActivateUserMFADevice(ctx, id, deviceId).Execute()

Activate an authentication factor



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 
**deviceId** | **int32** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateUserMFADeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**models.EnrollMfaDeviceResponse**](EnrollMfaDeviceResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserMFADevice

> map[string]interface{} DeleteUserMFADevice(ctx, id, deviceId).Execute()

Remove an authentication device



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 
**deviceId** | **int32** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserMFADeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollUserMFADevice

> models.EnrollMfaDeviceResponse EnrollUserMFADevice(ctx, id).InlineObject10(inlineObject10).Execute()

Enroll a user with a given authentication factor.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollUserMFADeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject10** | [**models.InlineObject10**](InlineObject10.md) |  | 

### Return type

[**models.EnrollMfaDeviceResponse**](EnrollMfaDeviceResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAvailableMFAFactors

> models.UserMfaFactorsResponse GetUserAvailableMFAFactors(ctx, id).Execute()

Get a list of MFA factors available to user



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAvailableMFAFactorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.UserMfaFactorsResponse**](UserMfaFactorsResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserEnrolledMFADevices

> models.EnrolledMfaDevicesResponse GetUserEnrolledMFADevices(ctx, id).Execute()

Get enrolled authentication devices



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserEnrolledMFADevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.EnrolledMfaDevicesResponse**](EnrolledMfaDevicesResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyUserMFADevice

> models.Status VerifyUserMFADevice(ctx, id, deviceId).InlineObject11(inlineObject11).Execute()

Verify an authentication device



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 
**deviceId** | **int32** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyUserMFADeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject11** | [**models.InlineObject11**](InlineObject11.md) |  | 

### Return type

[**models.Status**](Status.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

