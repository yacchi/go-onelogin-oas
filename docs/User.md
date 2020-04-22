# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | User’s unique ID in OneLogin | 
**Email** | Pointer to **string** | User’s email address, which he also uses to log in to OneLogin | 
**Username** | Pointer to **string** | If the user’s directory is set to authenticate using a user name value, this is the value used to sign in | 
**Firstname** | Pointer to **string** | User’s first name | 
**Lastname** | Pointer to **string** | User’s last name | 
**GroupId** | Pointer to **int32** | Group to which the user belongs | [optional] 
**InvalidLoginAttempts** | Pointer to **int32** | Number of sequential invalid login attempts the user has made that is less than or equal to the Maximum invalid login attempts value defined on the Session page in OneLogin. When this number reaches this value, the user account will be locked for the amount of time defined by the Lock effective period field on the Session page and this value will be reset to 0. | [optional] 
**ActivatedAt** | Pointer to [**time.Time**](time.Time.md) | Date and time at which the user’s status was set to 1 (active)       | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Date and time at which the user was created  | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | Date and time at which the user’s information was last updated | [optional] 
**InvitationSentAt** | Pointer to [**time.Time**](time.Time.md) | Date and time at which an invitation to OneLogin was sent to the user  | [optional] 
**PasswordChangedAt** | Pointer to [**time.Time**](time.Time.md) | Date and time at which the user’s password was last changed | [optional] 
**LastLogin** | Pointer to [**time.Time**](time.Time.md) | Date and time of the user’s last login  | [optional] 
**LockedUntil** | Pointer to [**time.Time**](time.Time.md) | Date and time at which the user’s account will be unlocked  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**OpenidName** | Pointer to **string** | OpenID URL that can be configured in other applications that accept OpenID for sign-in | [optional] 
**LocaleCode** | Pointer to **string** | Represents a geographical, political, or cultural region. Some features may use the locale value to tailor the display of information, such as numbers, for the user based on locale-specific customs and conventions | [optional] 
**Phone** | Pointer to **string** | User’s phone number | [optional] 
**Status** | Pointer to **int32** | Determines the user’s ability to log in to OneLogin    Possible values      0 - Unactivated   1 - Active Only users assigned this status can log in to OneLogin.   2 - Suspended   3 - Locked   4 - Password expired   5 - Awaiting password reset | [optional] 
**State** | Pointer to **int32** | Represents the user’s stage in a process (such as user account approval). User state determines the possible statuses a user account can be in. States include 0 - Unapproved 1 - Approved 2 - Rejected 3 - Unlicensed | [optional] 
**DistinguishedName** | Pointer to **string** | Synchronized from Active Directory | [optional] 
**ExternalId** | Pointer to **string** | External ID that can be used to uniquely identify the user in another system | [optional] 
**DirectoryId** | Pointer to **int32** | ID of the directory (Active Directory, LDAP, for example) from which the user was created | [optional] 
**MemberOf** | Pointer to **string** | Synchronized from Active Directory | [optional] 
**Samaccountname** | Pointer to **string** | Synchronized from Active Directory | [optional] 
**Userprincipalname** | Pointer to **string** | Synchronized from Active Directory | [optional] 
**ManagerAdId** | Pointer to **string** | ID of the user’s manager in Active Directory | [optional] 
**RoleId** | Pointer to **[]int32** | Role IDs to which the user is assigned | 
**CustomAttributes** | Pointer to **map[string]string** | Provides a list of custom attribute fields (also known as custom user fields) that are available for your account. The values returned correspond to the values you provided in the Shortname field when you defined the custom user field | 

## Methods

### NewUser

`func NewUser(id int32, email string, username string, firstname string, lastname string, roleId []int32, customAttributes map[string]string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetFirstname

`func (o *User) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *User) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *User) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.


### GetLastname

`func (o *User) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *User) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *User) SetLastname(v string)`

SetLastname sets Lastname field to given value.


### GetGroupId

`func (o *User) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *User) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *User) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *User) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetInvalidLoginAttempts

`func (o *User) GetInvalidLoginAttempts() int32`

GetInvalidLoginAttempts returns the InvalidLoginAttempts field if non-nil, zero value otherwise.

### GetInvalidLoginAttemptsOk

`func (o *User) GetInvalidLoginAttemptsOk() (*int32, bool)`

GetInvalidLoginAttemptsOk returns a tuple with the InvalidLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidLoginAttempts

`func (o *User) SetInvalidLoginAttempts(v int32)`

SetInvalidLoginAttempts sets InvalidLoginAttempts field to given value.

### HasInvalidLoginAttempts

`func (o *User) HasInvalidLoginAttempts() bool`

HasInvalidLoginAttempts returns a boolean if a field has been set.

### GetActivatedAt

`func (o *User) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *User) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *User) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *User) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInvitationSentAt

`func (o *User) GetInvitationSentAt() time.Time`

GetInvitationSentAt returns the InvitationSentAt field if non-nil, zero value otherwise.

### GetInvitationSentAtOk

`func (o *User) GetInvitationSentAtOk() (*time.Time, bool)`

GetInvitationSentAtOk returns a tuple with the InvitationSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationSentAt

`func (o *User) SetInvitationSentAt(v time.Time)`

SetInvitationSentAt sets InvitationSentAt field to given value.

### HasInvitationSentAt

`func (o *User) HasInvitationSentAt() bool`

HasInvitationSentAt returns a boolean if a field has been set.

### GetPasswordChangedAt

`func (o *User) GetPasswordChangedAt() time.Time`

GetPasswordChangedAt returns the PasswordChangedAt field if non-nil, zero value otherwise.

### GetPasswordChangedAtOk

`func (o *User) GetPasswordChangedAtOk() (*time.Time, bool)`

GetPasswordChangedAtOk returns a tuple with the PasswordChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangedAt

`func (o *User) SetPasswordChangedAt(v time.Time)`

SetPasswordChangedAt sets PasswordChangedAt field to given value.

### HasPasswordChangedAt

`func (o *User) HasPasswordChangedAt() bool`

HasPasswordChangedAt returns a boolean if a field has been set.

### GetLastLogin

`func (o *User) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLockedUntil

`func (o *User) GetLockedUntil() time.Time`

GetLockedUntil returns the LockedUntil field if non-nil, zero value otherwise.

### GetLockedUntilOk

`func (o *User) GetLockedUntilOk() (*time.Time, bool)`

GetLockedUntilOk returns a tuple with the LockedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedUntil

`func (o *User) SetLockedUntil(v time.Time)`

SetLockedUntil sets LockedUntil field to given value.

### HasLockedUntil

`func (o *User) HasLockedUntil() bool`

HasLockedUntil returns a boolean if a field has been set.

### GetNotes

`func (o *User) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *User) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *User) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *User) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOpenidName

`func (o *User) GetOpenidName() string`

GetOpenidName returns the OpenidName field if non-nil, zero value otherwise.

### GetOpenidNameOk

`func (o *User) GetOpenidNameOk() (*string, bool)`

GetOpenidNameOk returns a tuple with the OpenidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenidName

`func (o *User) SetOpenidName(v string)`

SetOpenidName sets OpenidName field to given value.

### HasOpenidName

`func (o *User) HasOpenidName() bool`

HasOpenidName returns a boolean if a field has been set.

### GetLocaleCode

`func (o *User) GetLocaleCode() string`

GetLocaleCode returns the LocaleCode field if non-nil, zero value otherwise.

### GetLocaleCodeOk

`func (o *User) GetLocaleCodeOk() (*string, bool)`

GetLocaleCodeOk returns a tuple with the LocaleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaleCode

`func (o *User) SetLocaleCode(v string)`

SetLocaleCode sets LocaleCode field to given value.

### HasLocaleCode

`func (o *User) HasLocaleCode() bool`

HasLocaleCode returns a boolean if a field has been set.

### GetPhone

`func (o *User) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *User) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *User) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *User) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStatus

`func (o *User) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *User) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetState

`func (o *User) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *User) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *User) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *User) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *User) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *User) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *User) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetExternalId

`func (o *User) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *User) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *User) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *User) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDirectoryId

`func (o *User) GetDirectoryId() int32`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *User) GetDirectoryIdOk() (*int32, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *User) SetDirectoryId(v int32)`

SetDirectoryId sets DirectoryId field to given value.

### HasDirectoryId

`func (o *User) HasDirectoryId() bool`

HasDirectoryId returns a boolean if a field has been set.

### GetMemberOf

`func (o *User) GetMemberOf() string`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *User) GetMemberOfOk() (*string, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *User) SetMemberOf(v string)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *User) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetSamaccountname

`func (o *User) GetSamaccountname() string`

GetSamaccountname returns the Samaccountname field if non-nil, zero value otherwise.

### GetSamaccountnameOk

`func (o *User) GetSamaccountnameOk() (*string, bool)`

GetSamaccountnameOk returns a tuple with the Samaccountname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamaccountname

`func (o *User) SetSamaccountname(v string)`

SetSamaccountname sets Samaccountname field to given value.

### HasSamaccountname

`func (o *User) HasSamaccountname() bool`

HasSamaccountname returns a boolean if a field has been set.

### GetUserprincipalname

`func (o *User) GetUserprincipalname() string`

GetUserprincipalname returns the Userprincipalname field if non-nil, zero value otherwise.

### GetUserprincipalnameOk

`func (o *User) GetUserprincipalnameOk() (*string, bool)`

GetUserprincipalnameOk returns a tuple with the Userprincipalname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserprincipalname

`func (o *User) SetUserprincipalname(v string)`

SetUserprincipalname sets Userprincipalname field to given value.

### HasUserprincipalname

`func (o *User) HasUserprincipalname() bool`

HasUserprincipalname returns a boolean if a field has been set.

### GetManagerAdId

`func (o *User) GetManagerAdId() string`

GetManagerAdId returns the ManagerAdId field if non-nil, zero value otherwise.

### GetManagerAdIdOk

`func (o *User) GetManagerAdIdOk() (*string, bool)`

GetManagerAdIdOk returns a tuple with the ManagerAdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerAdId

`func (o *User) SetManagerAdId(v string)`

SetManagerAdId sets ManagerAdId field to given value.

### HasManagerAdId

`func (o *User) HasManagerAdId() bool`

HasManagerAdId returns a boolean if a field has been set.

### GetRoleId

`func (o *User) GetRoleId() []int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *User) GetRoleIdOk() (*[]int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *User) SetRoleId(v []int32)`

SetRoleId sets RoleId field to given value.


### GetCustomAttributes

`func (o *User) GetCustomAttributes() map[string]string`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *User) GetCustomAttributesOk() (*map[string]string, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *User) SetCustomAttributes(v map[string]string)`

SetCustomAttributes sets CustomAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


