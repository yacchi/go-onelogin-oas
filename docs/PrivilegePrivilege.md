# PrivilegePrivilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version can be anything. Recommended to be Date/Time format | 
**Statement** | Pointer to [**[]Statement**](Statement.md) |  | 

## Methods

### NewPrivilegePrivilege

`func NewPrivilegePrivilege(version string, statement []Statement, ) *PrivilegePrivilege`

NewPrivilegePrivilege instantiates a new PrivilegePrivilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegePrivilegeWithDefaults

`func NewPrivilegePrivilegeWithDefaults() *PrivilegePrivilege`

NewPrivilegePrivilegeWithDefaults instantiates a new PrivilegePrivilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *PrivilegePrivilege) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PrivilegePrivilege) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PrivilegePrivilege) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetStatement

`func (o *PrivilegePrivilege) GetStatement() []Statement`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *PrivilegePrivilege) GetStatementOk() (*[]Statement, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *PrivilegePrivilege) SetStatement(v []Statement)`

SetStatement sets Statement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


