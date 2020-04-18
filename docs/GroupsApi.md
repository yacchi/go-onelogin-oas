# \GroupsApi

All URIs are relative to *https://api.us.onelogin.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroup**](GroupsApi.md#GetGroup) | **Get** /1/groups/{id} | Get Group by ID
[**GetGroups**](GroupsApi.md#GetGroups) | **Get** /1/groups | Get Groups



## GetGroup

> models.GroupResponse GetGroup(ctx, id).Execute()

Get Group by ID



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.GroupResponse**](GroupResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> models.GroupsResponse GetGroups(ctx).Execute()

Get Groups



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


### Return type

[**models.GroupsResponse**](GroupsResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

