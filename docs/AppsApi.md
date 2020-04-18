# \AppsApi

All URIs are relative to *https://virtserver.swaggerhub.com/OneLogin-Auth/onelogin-api/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApp**](AppsApi.md#CreateApp) | **Post** /2/apps | Create an App
[**DeleteApp**](AppsApi.md#DeleteApp) | **Delete** /2/apps/{id} | Delete an app
[**DeleteAppParameter**](AppsApi.md#DeleteAppParameter) | **Delete** /2/apps/{id}/parameters/{parameter_id} | Delete an app parameter
[**GetApp**](AppsApi.md#GetApp) | **Get** /2/apps/{id} | Get an App
[**GetApps**](AppsApi.md#GetApps) | **Get** /2/apps | Get Apps
[**UpdateApp**](AppsApi.md#UpdateApp) | **Put** /2/apps/{id} | Update an App



## CreateApp

> models.App CreateApp(ctx).App(app).Execute()

Create an App



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | [**models.App**](App.md) | The app to create | 

### Return type

[**models.App**](App.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApp

> DeleteApp(ctx, id).Execute()

Delete an app



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppRequest struct via the builder pattern


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


## DeleteAppParameter

> DeleteAppParameter(ctx, id, parameterId).Execute()

Delete an app parameter



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | App ID | 
**parameterId** | **int32** | Parameter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppParameterRequest struct via the builder pattern


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


## GetApp

> models.App GetApp(ctx, id).Execute()

Get an App



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.App**](App.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApps

> []models.AppInfo GetApps(ctx).Name(name).AuthMethod(authMethod).ConnectorId(connectorId).Execute()

Get Apps



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **authMethod** | **int32** |  | 
 **connectorId** | **int32** |  | 

### Return type

[**[]models.AppInfo**](AppInfo.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApp

> models.App UpdateApp(ctx, id).App(app).Execute()

Update an App



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **app** | [**models.App**](App.md) | The app to update | 

### Return type

[**models.App**](App.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

