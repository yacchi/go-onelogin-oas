# \LoginApi

All URIs are relative to *https://api.us.onelogin.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateUser**](LoginApi.md#AuthenticateUser) | **Post** /1/login/auth | Authenticate a user
[**VerifyLoginMFAToken**](LoginApi.md#VerifyLoginMFAToken) | **Post** /1/login/verify_factor | Verify an MFA token



## AuthenticateUser

> models.UserLoginResponse AuthenticateUser(ctx).CustomAllowedOriginHeader1(customAllowedOriginHeader1).InlineObject6(inlineObject6).Execute()

Authenticate a user



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customAllowedOriginHeader1** | **string** | Required for CORS requests only. Set to the Origin URI from which you are allowed to send a request using CORS. | 
 **inlineObject6** | [**models.InlineObject6**](InlineObject6.md) |  | 

### Return type

[**models.UserLoginResponse**](UserLoginResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyLoginMFAToken

> models.LoginVerifyMfaResponse VerifyLoginMFAToken(ctx).CustomAllowedOriginHeader1(customAllowedOriginHeader1).InlineObject7(inlineObject7).Execute()

Verify an MFA token



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyLoginMFATokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customAllowedOriginHeader1** | **string** | Required for CORS requests only. Set to the Origin URI from which you are allowed to send a request using CORS. | 
 **inlineObject7** | [**models.InlineObject7**](InlineObject7.md) |  | 

### Return type

[**models.LoginVerifyMfaResponse**](LoginVerifyMfaResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

