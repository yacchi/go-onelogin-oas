# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** |  | 
**EventTypeId** | Pointer to **int32** |  | 
**Id** | Pointer to **int32** | The Event ID | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date/time the Event was created | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Ipaddr** | Pointer to **string** |  | [optional] 
**ActorUserId** | Pointer to **int32** |  | [optional] 
**AssumingActingUserId** | Pointer to **int32** |  | [optional] 
**RoleId** | Pointer to **int32** |  | [optional] 
**AppId** | Pointer to **int32** |  | [optional] 
**GroupId** | Pointer to **int32** |  | [optional] 
**OtpDeviceId** | Pointer to **int32** |  | [optional] 
**PolicyId** | Pointer to **int32** |  | [optional] 
**ActorSystem** | Pointer to **string** |  | [optional] 
**CustomMessage** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**ActorUserName** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**OtpDeviceName** | Pointer to **string** |  | [optional] 
**OperationName** | Pointer to **string** |  | [optional] 
**DirectorySyncRunId** | Pointer to **int32** |  | [optional] 
**DirectoryId** | Pointer to **int32** |  | [optional] 
**Resolution** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ResourcseTypeId** | Pointer to **int32** |  | [optional] 
**ErrorDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewEvent

`func NewEvent(accountId int32, eventTypeId int32, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Event) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Event) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Event) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetEventTypeId

`func (o *Event) GetEventTypeId() int32`

GetEventTypeId returns the EventTypeId field if non-nil, zero value otherwise.

### GetEventTypeIdOk

`func (o *Event) GetEventTypeIdOk() (*int32, bool)`

GetEventTypeIdOk returns a tuple with the EventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeId

`func (o *Event) SetEventTypeId(v int32)`

SetEventTypeId sets EventTypeId field to given value.


### GetId

`func (o *Event) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Event) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Event) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Event) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Event) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Event) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *Event) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Event) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Event) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Event) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetNotes

`func (o *Event) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Event) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Event) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Event) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIpaddr

`func (o *Event) GetIpaddr() string`

GetIpaddr returns the Ipaddr field if non-nil, zero value otherwise.

### GetIpaddrOk

`func (o *Event) GetIpaddrOk() (*string, bool)`

GetIpaddrOk returns a tuple with the Ipaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddr

`func (o *Event) SetIpaddr(v string)`

SetIpaddr sets Ipaddr field to given value.

### HasIpaddr

`func (o *Event) HasIpaddr() bool`

HasIpaddr returns a boolean if a field has been set.

### GetActorUserId

`func (o *Event) GetActorUserId() int32`

GetActorUserId returns the ActorUserId field if non-nil, zero value otherwise.

### GetActorUserIdOk

`func (o *Event) GetActorUserIdOk() (*int32, bool)`

GetActorUserIdOk returns a tuple with the ActorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUserId

`func (o *Event) SetActorUserId(v int32)`

SetActorUserId sets ActorUserId field to given value.

### HasActorUserId

`func (o *Event) HasActorUserId() bool`

HasActorUserId returns a boolean if a field has been set.

### GetAssumingActingUserId

`func (o *Event) GetAssumingActingUserId() int32`

GetAssumingActingUserId returns the AssumingActingUserId field if non-nil, zero value otherwise.

### GetAssumingActingUserIdOk

`func (o *Event) GetAssumingActingUserIdOk() (*int32, bool)`

GetAssumingActingUserIdOk returns a tuple with the AssumingActingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumingActingUserId

`func (o *Event) SetAssumingActingUserId(v int32)`

SetAssumingActingUserId sets AssumingActingUserId field to given value.

### HasAssumingActingUserId

`func (o *Event) HasAssumingActingUserId() bool`

HasAssumingActingUserId returns a boolean if a field has been set.

### GetRoleId

`func (o *Event) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *Event) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *Event) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *Event) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetAppId

`func (o *Event) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Event) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Event) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Event) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetGroupId

`func (o *Event) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Event) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Event) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Event) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetOtpDeviceId

`func (o *Event) GetOtpDeviceId() int32`

GetOtpDeviceId returns the OtpDeviceId field if non-nil, zero value otherwise.

### GetOtpDeviceIdOk

`func (o *Event) GetOtpDeviceIdOk() (*int32, bool)`

GetOtpDeviceIdOk returns a tuple with the OtpDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpDeviceId

`func (o *Event) SetOtpDeviceId(v int32)`

SetOtpDeviceId sets OtpDeviceId field to given value.

### HasOtpDeviceId

`func (o *Event) HasOtpDeviceId() bool`

HasOtpDeviceId returns a boolean if a field has been set.

### GetPolicyId

`func (o *Event) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *Event) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *Event) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *Event) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetActorSystem

`func (o *Event) GetActorSystem() string`

GetActorSystem returns the ActorSystem field if non-nil, zero value otherwise.

### GetActorSystemOk

`func (o *Event) GetActorSystemOk() (*string, bool)`

GetActorSystemOk returns a tuple with the ActorSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorSystem

`func (o *Event) SetActorSystem(v string)`

SetActorSystem sets ActorSystem field to given value.

### HasActorSystem

`func (o *Event) HasActorSystem() bool`

HasActorSystem returns a boolean if a field has been set.

### GetCustomMessage

`func (o *Event) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *Event) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *Event) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *Event) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetRoleName

`func (o *Event) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *Event) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *Event) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *Event) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetAppName

`func (o *Event) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *Event) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *Event) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *Event) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetGroupName

`func (o *Event) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *Event) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *Event) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *Event) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetActorUserName

`func (o *Event) GetActorUserName() string`

GetActorUserName returns the ActorUserName field if non-nil, zero value otherwise.

### GetActorUserNameOk

`func (o *Event) GetActorUserNameOk() (*string, bool)`

GetActorUserNameOk returns a tuple with the ActorUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUserName

`func (o *Event) SetActorUserName(v string)`

SetActorUserName sets ActorUserName field to given value.

### HasActorUserName

`func (o *Event) HasActorUserName() bool`

HasActorUserName returns a boolean if a field has been set.

### GetUserName

`func (o *Event) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Event) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Event) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *Event) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPolicyName

`func (o *Event) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *Event) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *Event) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *Event) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetOtpDeviceName

`func (o *Event) GetOtpDeviceName() string`

GetOtpDeviceName returns the OtpDeviceName field if non-nil, zero value otherwise.

### GetOtpDeviceNameOk

`func (o *Event) GetOtpDeviceNameOk() (*string, bool)`

GetOtpDeviceNameOk returns a tuple with the OtpDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpDeviceName

`func (o *Event) SetOtpDeviceName(v string)`

SetOtpDeviceName sets OtpDeviceName field to given value.

### HasOtpDeviceName

`func (o *Event) HasOtpDeviceName() bool`

HasOtpDeviceName returns a boolean if a field has been set.

### GetOperationName

`func (o *Event) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *Event) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *Event) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *Event) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.

### GetDirectorySyncRunId

`func (o *Event) GetDirectorySyncRunId() int32`

GetDirectorySyncRunId returns the DirectorySyncRunId field if non-nil, zero value otherwise.

### GetDirectorySyncRunIdOk

`func (o *Event) GetDirectorySyncRunIdOk() (*int32, bool)`

GetDirectorySyncRunIdOk returns a tuple with the DirectorySyncRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorySyncRunId

`func (o *Event) SetDirectorySyncRunId(v int32)`

SetDirectorySyncRunId sets DirectorySyncRunId field to given value.

### HasDirectorySyncRunId

`func (o *Event) HasDirectorySyncRunId() bool`

HasDirectorySyncRunId returns a boolean if a field has been set.

### GetDirectoryId

`func (o *Event) GetDirectoryId() int32`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *Event) GetDirectoryIdOk() (*int32, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *Event) SetDirectoryId(v int32)`

SetDirectoryId sets DirectoryId field to given value.

### HasDirectoryId

`func (o *Event) HasDirectoryId() bool`

HasDirectoryId returns a boolean if a field has been set.

### GetResolution

`func (o *Event) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *Event) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *Event) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *Event) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetClientId

`func (o *Event) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Event) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Event) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Event) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetResourcseTypeId

`func (o *Event) GetResourcseTypeId() int32`

GetResourcseTypeId returns the ResourcseTypeId field if non-nil, zero value otherwise.

### GetResourcseTypeIdOk

`func (o *Event) GetResourcseTypeIdOk() (*int32, bool)`

GetResourcseTypeIdOk returns a tuple with the ResourcseTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcseTypeId

`func (o *Event) SetResourcseTypeId(v int32)`

SetResourcseTypeId sets ResourcseTypeId field to given value.

### HasResourcseTypeId

`func (o *Event) HasResourcseTypeId() bool`

HasResourcseTypeId returns a boolean if a field has been set.

### GetErrorDescription

`func (o *Event) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *Event) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *Event) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *Event) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


