# PrivilegeUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 
**BeforeCursor** | Pointer to **string** |  | [optional] 
**PreviousLink** | Pointer to **string** |  | [optional] 
**AfterCursor** | Pointer to **string** |  | [optional] 
**NextLink** | Pointer to **string** |  | [optional] 

## Methods

### NewPrivilegeUsersResponse

`func NewPrivilegeUsersResponse() *PrivilegeUsersResponse`

NewPrivilegeUsersResponse instantiates a new PrivilegeUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeUsersResponseWithDefaults

`func NewPrivilegeUsersResponseWithDefaults() *PrivilegeUsersResponse`

NewPrivilegeUsersResponseWithDefaults instantiates a new PrivilegeUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PrivilegeUsersResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PrivilegeUsersResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PrivilegeUsersResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PrivilegeUsersResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsers

`func (o *PrivilegeUsersResponse) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PrivilegeUsersResponse) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PrivilegeUsersResponse) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PrivilegeUsersResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetBeforeCursor

`func (o *PrivilegeUsersResponse) GetBeforeCursor() string`

GetBeforeCursor returns the BeforeCursor field if non-nil, zero value otherwise.

### GetBeforeCursorOk

`func (o *PrivilegeUsersResponse) GetBeforeCursorOk() (*string, bool)`

GetBeforeCursorOk returns a tuple with the BeforeCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeCursor

`func (o *PrivilegeUsersResponse) SetBeforeCursor(v string)`

SetBeforeCursor sets BeforeCursor field to given value.

### HasBeforeCursor

`func (o *PrivilegeUsersResponse) HasBeforeCursor() bool`

HasBeforeCursor returns a boolean if a field has been set.

### GetPreviousLink

`func (o *PrivilegeUsersResponse) GetPreviousLink() string`

GetPreviousLink returns the PreviousLink field if non-nil, zero value otherwise.

### GetPreviousLinkOk

`func (o *PrivilegeUsersResponse) GetPreviousLinkOk() (*string, bool)`

GetPreviousLinkOk returns a tuple with the PreviousLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousLink

`func (o *PrivilegeUsersResponse) SetPreviousLink(v string)`

SetPreviousLink sets PreviousLink field to given value.

### HasPreviousLink

`func (o *PrivilegeUsersResponse) HasPreviousLink() bool`

HasPreviousLink returns a boolean if a field has been set.

### GetAfterCursor

`func (o *PrivilegeUsersResponse) GetAfterCursor() string`

GetAfterCursor returns the AfterCursor field if non-nil, zero value otherwise.

### GetAfterCursorOk

`func (o *PrivilegeUsersResponse) GetAfterCursorOk() (*string, bool)`

GetAfterCursorOk returns a tuple with the AfterCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterCursor

`func (o *PrivilegeUsersResponse) SetAfterCursor(v string)`

SetAfterCursor sets AfterCursor field to given value.

### HasAfterCursor

`func (o *PrivilegeUsersResponse) HasAfterCursor() bool`

HasAfterCursor returns a boolean if a field has been set.

### GetNextLink

`func (o *PrivilegeUsersResponse) GetNextLink() string`

GetNextLink returns the NextLink field if non-nil, zero value otherwise.

### GetNextLinkOk

`func (o *PrivilegeUsersResponse) GetNextLinkOk() (*string, bool)`

GetNextLinkOk returns a tuple with the NextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLink

`func (o *PrivilegeUsersResponse) SetNextLink(v string)`

SetNextLink sets NextLink field to given value.

### HasNextLink

`func (o *PrivilegeUsersResponse) HasNextLink() bool`

HasNextLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


