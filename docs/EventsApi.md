# \EventsApi

All URIs are relative to *https://api.us.onelogin.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /1/events/{id} | Get Event by ID
[**GetEventTypes**](EventsApi.md#GetEventTypes) | **Get** /1/events/types | Get Event Types
[**GetEvents**](EventsApi.md#GetEvents) | **Get** /1/events | Get Events



## GetEvent

> models.EventResponse GetEvent(ctx, id).Execute()

Get Event by ID



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Event ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**models.EventResponse**](EventResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTypes

> models.EventTypesResponse GetEventTypes(ctx).Execute()

Get Event Types



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTypesRequest struct via the builder pattern


### Return type

[**models.EventTypesResponse**](EventTypesResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> models.EventsResponse GetEvents(ctx).ClientId(clientId).CreatedAt(createdAt).DirectoryId(directoryId).EventTypeId(eventTypeId).Id(id).Resolution(resolution).Since(since).Until(until).UserId(userId).Execute()

Get Events



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **createdAt** | **time.Time** |  | 
 **directoryId** | **int32** |  | 
 **eventTypeId** | **int32** |  | 
 **id** | **int32** |  | 
 **resolution** | **string** |  | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **userId** | **int32** |  | 

### Return type

[**models.EventsResponse**](EventsResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

