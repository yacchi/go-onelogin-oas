# \SamlApi

All URIs are relative to *https://virtserver.swaggerhub.com/OneLogin-Auth/onelogin-api/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSAMLAssertion**](SamlApi.md#CreateSAMLAssertion) | **Post** /1/saml_assertion | Generate SAML assertion
[**VerifySAMLAssertionMFAToken**](SamlApi.md#VerifySAMLAssertionMFAToken) | **Post** /1/saml_assertion/verify_factor | Verify an MFA token



## CreateSAMLAssertion

> models.SamlAssertionResponse CreateSAMLAssertion(ctx).InlineObject8(inlineObject8).Execute()

Generate SAML assertion



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSAMLAssertionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject8** | [**models.InlineObject8**](InlineObject8.md) |  | 

### Return type

[**models.SamlAssertionResponse**](SamlAssertionResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifySAMLAssertionMFAToken

> models.SamlAssertionResponse VerifySAMLAssertionMFAToken(ctx).InlineObject9(inlineObject9).Execute()

Verify an MFA token



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifySAMLAssertionMFATokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject9** | [**models.InlineObject9**](InlineObject9.md) |  | 

### Return type

[**models.SamlAssertionResponse**](SamlAssertionResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

