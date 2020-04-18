# \RolesApi

All URIs are relative to *https://api.us.onelogin.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRole**](RolesApi.md#GetRole) | **Get** /1/roles/{id} | Get Role by ID
[**GetRoles**](RolesApi.md#GetRoles) | **Get** /1/roles | Get Roles



## GetRole

> models.RoleReponse GetRole(ctx, id).Execute()

Get Role by ID



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.RoleReponse**](RoleReponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoles

> models.RolesResponse GetRoles(ctx).Id(id).Name(name).Limit(limit).Sort(sort).Fields(fields).Execute()

Get Roles



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **name** | **string** |  | 
 **limit** | **int32** |  | 
 **sort** | **string** |  | 
 **fields** | **string** |  | 

### Return type

[**models.RolesResponse**](RolesResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

