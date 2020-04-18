# \PrivilegesApi

All URIs are relative to *https://api.us.onelogin.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrivilegeRoles**](PrivilegesApi.md#AddPrivilegeRoles) | **Post** /1/privileges/{id}/roles | Assign roles
[**AddPrivilegeUsers**](PrivilegesApi.md#AddPrivilegeUsers) | **Post** /1/privileges/{id}/users | Assign users
[**CreatePrivilege**](PrivilegesApi.md#CreatePrivilege) | **Post** /1/privileges | Creates privilege
[**DeletePrivilege**](PrivilegesApi.md#DeletePrivilege) | **Delete** /1/privileges/{id} | Delete privilege
[**GetPrivilege**](PrivilegesApi.md#GetPrivilege) | **Get** /1/privileges/{id} | Get privilege
[**GetPrivilegeRoles**](PrivilegesApi.md#GetPrivilegeRoles) | **Get** /1/privileges/{id}/roles | Get roles
[**GetPrivilegeUsers**](PrivilegesApi.md#GetPrivilegeUsers) | **Get** /1/privileges/{id}/users | Get privilege users
[**GetPrivileges**](PrivilegesApi.md#GetPrivileges) | **Get** /1/privileges | Get Privileges
[**RemovePrivilegeRole**](PrivilegesApi.md#RemovePrivilegeRole) | **Delete** /1/privileges/{id}/roles/{role_id} | Remove a role
[**RemovePrivlegeUser**](PrivilegesApi.md#RemovePrivlegeUser) | **Delete** /1/privileges/{id}/users/{user_id} | Remove a user
[**UpdatePrivilege**](PrivilegesApi.md#UpdatePrivilege) | **Put** /1/privileges/{id} | Update privilege



## AddPrivilegeRoles

> models.AssignPrivilegeRolesResponse AddPrivilegeRoles(ctx, id).InlineObject12(inlineObject12).Execute()

Assign roles



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Privilege ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPrivilegeRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject12** | [**models.InlineObject12**](InlineObject12.md) |  | 

### Return type

[**models.AssignPrivilegeRolesResponse**](AssignPrivilegeRolesResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPrivilegeUsers

> models.AssignPrivilegeRolesResponse AddPrivilegeUsers(ctx, id).InlineObject13(inlineObject13).Execute()

Assign users



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Privilege ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPrivilegeUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject13** | [**models.InlineObject13**](InlineObject13.md) |  | 

### Return type

[**models.AssignPrivilegeRolesResponse**](AssignPrivilegeRolesResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivilege

> models.CreatePrivilegeResponse CreatePrivilege(ctx).Privilege(privilege).Execute()

Creates privilege



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privilege** | [**models.Privilege**](Privilege.md) |  | 

### Return type

[**models.CreatePrivilegeResponse**](CreatePrivilegeResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivilege

> DeletePrivilege(ctx, id).Execute()

Delete privilege



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Privilege ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivilege

> GetPrivilege(ctx, id).Execute()

Get privilege



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Privilege ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivilegeRoles

> models.PrivilegeRolesResponse GetPrivilegeRoles(ctx, id).Execute()

Get roles



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Privilege ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegeRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.PrivilegeRolesResponse**](PrivilegeRolesResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivilegeUsers

> models.PrivilegeUsersResponse GetPrivilegeUsers(ctx, id).Execute()

Get privilege users



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Privilege ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegeUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.PrivilegeUsersResponse**](PrivilegeUsersResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivileges

> []models.Privilege GetPrivileges(ctx).Execute()

Get Privileges



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegesRequest struct via the builder pattern


### Return type

[**[]models.Privilege**](Privilege.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePrivilegeRole

> RemovePrivilegeRole(ctx, id, roleId).Execute()

Remove a role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Privilege ID | 
**roleId** | **int32** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePrivilegeRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePrivlegeUser

> RemovePrivlegeUser(ctx, id, userId).Execute()

Remove a user



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Privilege ID | 
**userId** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePrivlegeUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivilege

> models.CreatePrivilegeResponse UpdatePrivilege(ctx, id).Privilege(privilege).Execute()

Update privilege



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Privilege ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privilege** | [**models.Privilege**](Privilege.md) |  | 

### Return type

[**models.CreatePrivilegeResponse**](CreatePrivilegeResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

