# \UsersApi

All URIs are relative to *https://api.us.onelogin.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserRoles**](UsersApi.md#AddUserRoles) | **Put** /1/users/{id}/add_roles | Assign Role to User
[**Call1UsersPost**](UsersApi.md#Call1UsersPost) | **Post** /1/users | Create a User
[**CreateUserTempMFAToken**](UsersApi.md#CreateUserTempMFAToken) | **Post** /1/users/{id}/mfa_token | Generate Temp MFA Token
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /1/users/{id} | Delete a user account
[**GetCustomAttributes**](UsersApi.md#GetCustomAttributes) | **Get** /1/users/custom_attributes | Get Custom Attributes
[**GetUser**](UsersApi.md#GetUser) | **Get** /1/users/{id} | Get a User
[**GetUserApps**](UsersApi.md#GetUserApps) | **Get** /1/users/{id}/apps | Get User Apps
[**GetUserRoles**](UsersApi.md#GetUserRoles) | **Get** /1/users/{id}/roles | Get User Roles
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /1/users | Get Users
[**LockUser**](UsersApi.md#LockUser) | **Put** /1/users/{id}/lock_user | Lock a user account
[**LogoutUser**](UsersApi.md#LogoutUser) | **Put** /1/users/{id}/logout | Log a user out of any and all sessions
[**RemoveUserRoles**](UsersApi.md#RemoveUserRoles) | **Put** /1/users/{id}/remove_roles | Remove Roles from User
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /1/users/{id} | Update a User
[**UpdateUserCustomAttributes**](UsersApi.md#UpdateUserCustomAttributes) | **Put** /1/users/{id}/set_custom_attributes | Set a custom attribute
[**UpdateUserPassword**](UsersApi.md#UpdateUserPassword) | **Put** /1/users/set_password_clear_text/{id} | Set a the password for a user
[**UpdateUserPasswordSalted**](UsersApi.md#UpdateUserPasswordSalted) | **Put** /1/users/set_password_using_salt/{id} | Set a pre salted password for a user
[**UpdateUserState**](UsersApi.md#UpdateUserState) | **Put** /1/users/{id}/set_state | Set the State for a user.



## AddUserRoles

> models.Status AddUserRoles(ctx, id).RequestBody(requestBody).Execute()

Assign Role to User



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**[]int32**](int32.md) | The roles to assign | 

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


## Call1UsersPost

> models.CreateUserResponse Call1UsersPost(ctx).User(user).Execute()

Create a User



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCall1UsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**models.User**](User.md) | The user to create | 

### Return type

[**models.CreateUserResponse**](CreateUserResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserTempMFAToken

> models.GenerateMfaTokenResponse CreateUserTempMFAToken(ctx, id).InlineObject3(inlineObject3).Execute()

Generate Temp MFA Token



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserTempMFATokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject3** | [**models.InlineObject3**](InlineObject3.md) |  | 

### Return type

[**models.GenerateMfaTokenResponse**](GenerateMfaTokenResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> models.Status DeleteUser(ctx, id).Execute()

Delete a user account



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.Status**](Status.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomAttributes

> models.CustomAttributesResponse GetCustomAttributes(ctx).Execute()

Get Custom Attributes



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomAttributesRequest struct via the builder pattern


### Return type

[**models.CustomAttributesResponse**](CustomAttributesResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> models.UserResponse GetUser(ctx, id).Execute()

Get a User



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.UserResponse**](UserResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApps

> models.UserAppsResponse GetUserApps(ctx, id).Execute()

Get User Apps



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.UserAppsResponse**](UserAppsResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRoles

> models.UserRolesResponse GetUserRoles(ctx, id).Execute()

Get User Roles



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.UserRolesResponse**](UserRolesResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> models.UsersResponse GetUsers(ctx).DirectoryId(directoryId).Email(email).ExternalId(externalId).Firstname(firstname).Id(id).ManagerAdId(managerAdId).RoleId(roleId).Samaccountname(samaccountname).Since(since).Until(until).Username(username).Userprincipalname(userprincipalname).Execute()

Get Users



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directoryId** | **int32** |  | 
 **email** | **string** |  | 
 **externalId** | **string** |  | 
 **firstname** | **string** |  | 
 **id** | **int32** |  | 
 **managerAdId** | **int32** |  | 
 **roleId** | **int32** |  | 
 **samaccountname** | **string** |  | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **username** | **string** |  | 
 **userprincipalname** | **string** |  | 

### Return type

[**models.UsersResponse**](UsersResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockUser

> models.Status LockUser(ctx, id).InlineObject1(inlineObject1).Execute()

Lock a user account



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject1** | [**models.InlineObject1**](InlineObject1.md) |  | 

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


## LogoutUser

> models.Status LogoutUser(ctx, id).Execute()

Log a user out of any and all sessions



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.Status**](Status.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserRoles

> models.Status RemoveUserRoles(ctx, id).RequestBody(requestBody).Execute()

Remove Roles from User



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**[]int32**](int32.md) | The roles to remove | 

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


## UpdateUser

> models.CreateUserResponse UpdateUser(ctx, id).User(user).Execute()

Update a User



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**models.User**](User.md) | The user to update | 

### Return type

[**models.CreateUserResponse**](CreateUserResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserCustomAttributes

> models.Status UpdateUserCustomAttributes(ctx, id).InlineObject2(inlineObject2).Execute()

Set a custom attribute



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserCustomAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject2** | [**models.InlineObject2**](InlineObject2.md) |  | 

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


## UpdateUserPassword

> models.Status UpdateUserPassword(ctx, id).InlineObject4(inlineObject4).Execute()

Set a the password for a user



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject4** | [**models.InlineObject4**](InlineObject4.md) |  | 

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


## UpdateUserPasswordSalted

> models.Status UpdateUserPasswordSalted(ctx, id).InlineObject5(inlineObject5).Execute()

Set a pre salted password for a user



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordSaltedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject5** | [**models.InlineObject5**](InlineObject5.md) |  | 

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


## UpdateUserState

> models.Status UpdateUserState(ctx, id).InlineObject(inlineObject).Execute()

Set the State for a user.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject** | [**models.InlineObject**](InlineObject.md) |  | 

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

