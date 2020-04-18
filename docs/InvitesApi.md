# \InvitesApi

All URIs are relative to *https://api.us.onelogin.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInviteLink**](InvitesApi.md#CreateInviteLink) | **Post** /1/invites/get_invite_link | Generate Invite Link
[**SendInviteLink**](InvitesApi.md#SendInviteLink) | **Post** /1/invites/send_invite_link | Send Invite Link



## CreateInviteLink

> models.CustomAttributesResponse CreateInviteLink(ctx).InlineObject14(inlineObject14).Execute()

Generate Invite Link



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInviteLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject14** | [**models.InlineObject14**](InlineObject14.md) |  | 

### Return type

[**models.CustomAttributesResponse**](CustomAttributesResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendInviteLink

> models.Status SendInviteLink(ctx).InlineObject15(inlineObject15).Execute()

Send Invite Link



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendInviteLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject15** | [**models.InlineObject15**](InlineObject15.md) |  | 

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

