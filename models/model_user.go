// Code generated by "openapi-codegen"; DO NOT EDIT.
/*
 * OneLogin API
 *
 * This is an administrative API for OneLogin customers
 *
 * API version: 1.1.0-oas3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"encoding/json"
	"time"
)

// User struct for User
type User struct {
	// User’s unique ID in OneLogin
	Id int32 `json:"id"`
	// User’s email address, which he also uses to log in to OneLogin
	Email string `json:"email"`
	// If the user’s directory is set to authenticate using a user name value, this is the value used to sign in
	Username string `json:"username"`
	// User’s first name
	Firstname string `json:"firstname"`
	// User’s last name
	Lastname string `json:"lastname"`
	// Group to which the user belongs
	GroupId *int32 `json:"group_id,omitempty"`
	// Number of sequential invalid login attempts the user has made that is less than or equal to the Maximum invalid login attempts value defined on the Session page in OneLogin. When this number reaches this value, the user account will be locked for the amount of time defined by the Lock effective period field on the Session page and this value will be reset to 0.
	InvalidLoginAttempts *int32 `json:"invalid_login_attempts,omitempty"`
	// Date and time at which the user’s status was set to 1 (active)
	ActivatedAt *time.Time `json:"activated_at,omitempty"`
	// Date and time at which the user was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time at which the user’s information was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Date and time at which an invitation to OneLogin was sent to the user
	InvitationSentAt *time.Time `json:"invitation_sent_at,omitempty"`
	// Date and time at which the user’s password was last changed
	PasswordChangedAt *time.Time `json:"password_changed_at,omitempty"`
	// Date and time of the user’s last login
	LastLogin *time.Time `json:"last_login,omitempty"`
	// Date and time at which the user’s account will be unlocked
	LockedUntil *time.Time `json:"locked_until,omitempty"`
	Notes       *string    `json:"notes,omitempty"`
	// OpenID URL that can be configured in other applications that accept OpenID for sign-in
	OpenidName *string `json:"openid_name,omitempty"`
	// Represents a geographical, political, or cultural region. Some features may use the locale value to tailor the display of information, such as numbers, for the user based on locale-specific customs and conventions
	LocaleCode *string `json:"locale_code,omitempty"`
	// User’s phone number
	Phone *string `json:"phone,omitempty"`
	// Determines the user’s ability to log in to OneLogin    Possible values      0 - Unactivated   1 - Active Only users assigned this status can log in to OneLogin.   2 - Suspended   3 - Locked   4 - Password expired   5 - Awaiting password reset
	Status *int32 `json:"status,omitempty"`
	// Represents the user’s stage in a process (such as user account approval). User state determines the possible statuses a user account can be in. States include 0 - Unapproved 1 - Approved 2 - Rejected 3 - Unlicensed
	State *int32 `json:"state,omitempty"`
	// Synchronized from Active Directory
	DistinguishedName *string `json:"distinguished_name,omitempty"`
	// External ID that can be used to uniquely identify the user in another system
	ExternalId *string `json:"external_id,omitempty"`
	// ID of the directory (Active Directory, LDAP, for example) from which the user was created
	DirectoryId *int32 `json:"directory_id,omitempty"`
	// Synchronized from Active Directory
	MemberOf *string `json:"member_of,omitempty"`
	// Synchronized from Active Directory
	Samaccountname *string `json:"samaccountname,omitempty"`
	// Synchronized from Active Directory
	Userprincipalname *string `json:"userprincipalname,omitempty"`
	// ID of the user’s manager in Active Directory
	ManagerAdId *string `json:"manager_ad_id,omitempty"`
	// Role IDs to which the user is assigned
	RoleId []int32 `json:"role_id"`
	// Provides a list of custom attribute fields (also known as custom user fields) that are available for your account. The values returned correspond to the values you provided in the Shortname field when you defined the custom user field
	CustomAttributes map[string]string `json:"custom_attributes"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(id int32, email string, username string, firstname string, lastname string, roleId []int32, customAttributes map[string]string) *User {
	this := User{}
	this.Id = id
	this.Email = email
	this.Username = username
	this.Firstname = firstname
	this.Lastname = lastname
	this.RoleId = roleId
	this.CustomAttributes = customAttributes
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value
func (o *User) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *User) SetId(v int32) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

// GetUsername returns the Username field value
func (o *User) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *User) SetUsername(v string) {
	o.Username = v
}

// GetFirstname returns the Firstname field value
func (o *User) GetFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value
// and a boolean to check if the value has been set.
func (o *User) GetFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Firstname, true
}

// SetFirstname sets field value
func (o *User) SetFirstname(v string) {
	o.Firstname = v
}

// GetLastname returns the Lastname field value
func (o *User) GetLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value
// and a boolean to check if the value has been set.
func (o *User) GetLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lastname, true
}

// SetLastname sets field value
func (o *User) SetLastname(v string) {
	o.Lastname = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *User) GetGroupId() int32 {
	if o == nil || o.GroupId == nil {
		var ret int32
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGroupIdOk() (*int32, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *User) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.
func (o *User) SetGroupId(v int32) {
	o.GroupId = &v
}

// GetInvalidLoginAttempts returns the InvalidLoginAttempts field value if set, zero value otherwise.
func (o *User) GetInvalidLoginAttempts() int32 {
	if o == nil || o.InvalidLoginAttempts == nil {
		var ret int32
		return ret
	}
	return *o.InvalidLoginAttempts
}

// GetInvalidLoginAttemptsOk returns a tuple with the InvalidLoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetInvalidLoginAttemptsOk() (*int32, bool) {
	if o == nil || o.InvalidLoginAttempts == nil {
		return nil, false
	}
	return o.InvalidLoginAttempts, true
}

// HasInvalidLoginAttempts returns a boolean if a field has been set.
func (o *User) HasInvalidLoginAttempts() bool {
	if o != nil && o.InvalidLoginAttempts != nil {
		return true
	}

	return false
}

// SetInvalidLoginAttempts gets a reference to the given int32 and assigns it to the InvalidLoginAttempts field.
func (o *User) SetInvalidLoginAttempts(v int32) {
	o.InvalidLoginAttempts = &v
}

// GetActivatedAt returns the ActivatedAt field value if set, zero value otherwise.
func (o *User) GetActivatedAt() time.Time {
	if o == nil || o.ActivatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ActivatedAt
}

// GetActivatedAtOk returns a tuple with the ActivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetActivatedAtOk() (*time.Time, bool) {
	if o == nil || o.ActivatedAt == nil {
		return nil, false
	}
	return o.ActivatedAt, true
}

// HasActivatedAt returns a boolean if a field has been set.
func (o *User) HasActivatedAt() bool {
	if o != nil && o.ActivatedAt != nil {
		return true
	}

	return false
}

// SetActivatedAt gets a reference to the given time.Time and assigns it to the ActivatedAt field.
func (o *User) SetActivatedAt(v time.Time) {
	o.ActivatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *User) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *User) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *User) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetInvitationSentAt returns the InvitationSentAt field value if set, zero value otherwise.
func (o *User) GetInvitationSentAt() time.Time {
	if o == nil || o.InvitationSentAt == nil {
		var ret time.Time
		return ret
	}
	return *o.InvitationSentAt
}

// GetInvitationSentAtOk returns a tuple with the InvitationSentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetInvitationSentAtOk() (*time.Time, bool) {
	if o == nil || o.InvitationSentAt == nil {
		return nil, false
	}
	return o.InvitationSentAt, true
}

// HasInvitationSentAt returns a boolean if a field has been set.
func (o *User) HasInvitationSentAt() bool {
	if o != nil && o.InvitationSentAt != nil {
		return true
	}

	return false
}

// SetInvitationSentAt gets a reference to the given time.Time and assigns it to the InvitationSentAt field.
func (o *User) SetInvitationSentAt(v time.Time) {
	o.InvitationSentAt = &v
}

// GetPasswordChangedAt returns the PasswordChangedAt field value if set, zero value otherwise.
func (o *User) GetPasswordChangedAt() time.Time {
	if o == nil || o.PasswordChangedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PasswordChangedAt
}

// GetPasswordChangedAtOk returns a tuple with the PasswordChangedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordChangedAtOk() (*time.Time, bool) {
	if o == nil || o.PasswordChangedAt == nil {
		return nil, false
	}
	return o.PasswordChangedAt, true
}

// HasPasswordChangedAt returns a boolean if a field has been set.
func (o *User) HasPasswordChangedAt() bool {
	if o != nil && o.PasswordChangedAt != nil {
		return true
	}

	return false
}

// SetPasswordChangedAt gets a reference to the given time.Time and assigns it to the PasswordChangedAt field.
func (o *User) SetPasswordChangedAt(v time.Time) {
	o.PasswordChangedAt = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *User) GetLastLogin() time.Time {
	if o == nil || o.LastLogin == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginOk() (*time.Time, bool) {
	if o == nil || o.LastLogin == nil {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *User) HasLastLogin() bool {
	if o != nil && o.LastLogin != nil {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given time.Time and assigns it to the LastLogin field.
func (o *User) SetLastLogin(v time.Time) {
	o.LastLogin = &v
}

// GetLockedUntil returns the LockedUntil field value if set, zero value otherwise.
func (o *User) GetLockedUntil() time.Time {
	if o == nil || o.LockedUntil == nil {
		var ret time.Time
		return ret
	}
	return *o.LockedUntil
}

// GetLockedUntilOk returns a tuple with the LockedUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLockedUntilOk() (*time.Time, bool) {
	if o == nil || o.LockedUntil == nil {
		return nil, false
	}
	return o.LockedUntil, true
}

// HasLockedUntil returns a boolean if a field has been set.
func (o *User) HasLockedUntil() bool {
	if o != nil && o.LockedUntil != nil {
		return true
	}

	return false
}

// SetLockedUntil gets a reference to the given time.Time and assigns it to the LockedUntil field.
func (o *User) SetLockedUntil(v time.Time) {
	o.LockedUntil = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *User) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *User) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *User) SetNotes(v string) {
	o.Notes = &v
}

// GetOpenidName returns the OpenidName field value if set, zero value otherwise.
func (o *User) GetOpenidName() string {
	if o == nil || o.OpenidName == nil {
		var ret string
		return ret
	}
	return *o.OpenidName
}

// GetOpenidNameOk returns a tuple with the OpenidName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetOpenidNameOk() (*string, bool) {
	if o == nil || o.OpenidName == nil {
		return nil, false
	}
	return o.OpenidName, true
}

// HasOpenidName returns a boolean if a field has been set.
func (o *User) HasOpenidName() bool {
	if o != nil && o.OpenidName != nil {
		return true
	}

	return false
}

// SetOpenidName gets a reference to the given string and assigns it to the OpenidName field.
func (o *User) SetOpenidName(v string) {
	o.OpenidName = &v
}

// GetLocaleCode returns the LocaleCode field value if set, zero value otherwise.
func (o *User) GetLocaleCode() string {
	if o == nil || o.LocaleCode == nil {
		var ret string
		return ret
	}
	return *o.LocaleCode
}

// GetLocaleCodeOk returns a tuple with the LocaleCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLocaleCodeOk() (*string, bool) {
	if o == nil || o.LocaleCode == nil {
		return nil, false
	}
	return o.LocaleCode, true
}

// HasLocaleCode returns a boolean if a field has been set.
func (o *User) HasLocaleCode() bool {
	if o != nil && o.LocaleCode != nil {
		return true
	}

	return false
}

// SetLocaleCode gets a reference to the given string and assigns it to the LocaleCode field.
func (o *User) SetLocaleCode(v string) {
	o.LocaleCode = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *User) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *User) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *User) SetPhone(v string) {
	o.Phone = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *User) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *User) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *User) SetStatus(v int32) {
	o.Status = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *User) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *User) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *User) SetState(v int32) {
	o.State = &v
}

// GetDistinguishedName returns the DistinguishedName field value if set, zero value otherwise.
func (o *User) GetDistinguishedName() string {
	if o == nil || o.DistinguishedName == nil {
		var ret string
		return ret
	}
	return *o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDistinguishedNameOk() (*string, bool) {
	if o == nil || o.DistinguishedName == nil {
		return nil, false
	}
	return o.DistinguishedName, true
}

// HasDistinguishedName returns a boolean if a field has been set.
func (o *User) HasDistinguishedName() bool {
	if o != nil && o.DistinguishedName != nil {
		return true
	}

	return false
}

// SetDistinguishedName gets a reference to the given string and assigns it to the DistinguishedName field.
func (o *User) SetDistinguishedName(v string) {
	o.DistinguishedName = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *User) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *User) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *User) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetDirectoryId returns the DirectoryId field value if set, zero value otherwise.
func (o *User) GetDirectoryId() int32 {
	if o == nil || o.DirectoryId == nil {
		var ret int32
		return ret
	}
	return *o.DirectoryId
}

// GetDirectoryIdOk returns a tuple with the DirectoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDirectoryIdOk() (*int32, bool) {
	if o == nil || o.DirectoryId == nil {
		return nil, false
	}
	return o.DirectoryId, true
}

// HasDirectoryId returns a boolean if a field has been set.
func (o *User) HasDirectoryId() bool {
	if o != nil && o.DirectoryId != nil {
		return true
	}

	return false
}

// SetDirectoryId gets a reference to the given int32 and assigns it to the DirectoryId field.
func (o *User) SetDirectoryId(v int32) {
	o.DirectoryId = &v
}

// GetMemberOf returns the MemberOf field value if set, zero value otherwise.
func (o *User) GetMemberOf() string {
	if o == nil || o.MemberOf == nil {
		var ret string
		return ret
	}
	return *o.MemberOf
}

// GetMemberOfOk returns a tuple with the MemberOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetMemberOfOk() (*string, bool) {
	if o == nil || o.MemberOf == nil {
		return nil, false
	}
	return o.MemberOf, true
}

// HasMemberOf returns a boolean if a field has been set.
func (o *User) HasMemberOf() bool {
	if o != nil && o.MemberOf != nil {
		return true
	}

	return false
}

// SetMemberOf gets a reference to the given string and assigns it to the MemberOf field.
func (o *User) SetMemberOf(v string) {
	o.MemberOf = &v
}

// GetSamaccountname returns the Samaccountname field value if set, zero value otherwise.
func (o *User) GetSamaccountname() string {
	if o == nil || o.Samaccountname == nil {
		var ret string
		return ret
	}
	return *o.Samaccountname
}

// GetSamaccountnameOk returns a tuple with the Samaccountname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSamaccountnameOk() (*string, bool) {
	if o == nil || o.Samaccountname == nil {
		return nil, false
	}
	return o.Samaccountname, true
}

// HasSamaccountname returns a boolean if a field has been set.
func (o *User) HasSamaccountname() bool {
	if o != nil && o.Samaccountname != nil {
		return true
	}

	return false
}

// SetSamaccountname gets a reference to the given string and assigns it to the Samaccountname field.
func (o *User) SetSamaccountname(v string) {
	o.Samaccountname = &v
}

// GetUserprincipalname returns the Userprincipalname field value if set, zero value otherwise.
func (o *User) GetUserprincipalname() string {
	if o == nil || o.Userprincipalname == nil {
		var ret string
		return ret
	}
	return *o.Userprincipalname
}

// GetUserprincipalnameOk returns a tuple with the Userprincipalname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUserprincipalnameOk() (*string, bool) {
	if o == nil || o.Userprincipalname == nil {
		return nil, false
	}
	return o.Userprincipalname, true
}

// HasUserprincipalname returns a boolean if a field has been set.
func (o *User) HasUserprincipalname() bool {
	if o != nil && o.Userprincipalname != nil {
		return true
	}

	return false
}

// SetUserprincipalname gets a reference to the given string and assigns it to the Userprincipalname field.
func (o *User) SetUserprincipalname(v string) {
	o.Userprincipalname = &v
}

// GetManagerAdId returns the ManagerAdId field value if set, zero value otherwise.
func (o *User) GetManagerAdId() string {
	if o == nil || o.ManagerAdId == nil {
		var ret string
		return ret
	}
	return *o.ManagerAdId
}

// GetManagerAdIdOk returns a tuple with the ManagerAdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetManagerAdIdOk() (*string, bool) {
	if o == nil || o.ManagerAdId == nil {
		return nil, false
	}
	return o.ManagerAdId, true
}

// HasManagerAdId returns a boolean if a field has been set.
func (o *User) HasManagerAdId() bool {
	if o != nil && o.ManagerAdId != nil {
		return true
	}

	return false
}

// SetManagerAdId gets a reference to the given string and assigns it to the ManagerAdId field.
func (o *User) SetManagerAdId(v string) {
	o.ManagerAdId = &v
}

// GetRoleId returns the RoleId field value
func (o *User) GetRoleId() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *User) GetRoleIdOk() (*[]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *User) SetRoleId(v []int32) {
	o.RoleId = v
}

// GetCustomAttributes returns the CustomAttributes field value
func (o *User) GetCustomAttributes() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value
// and a boolean to check if the value has been set.
func (o *User) GetCustomAttributesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomAttributes, true
}

// SetCustomAttributes sets field value
func (o *User) SetCustomAttributes(v map[string]string) {
	o.CustomAttributes = v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["firstname"] = o.Firstname
	}
	if true {
		toSerialize["lastname"] = o.Lastname
	}
	if o.GroupId != nil {
		toSerialize["group_id"] = o.GroupId
	}
	if o.InvalidLoginAttempts != nil {
		toSerialize["invalid_login_attempts"] = o.InvalidLoginAttempts
	}
	if o.ActivatedAt != nil {
		toSerialize["activated_at"] = o.ActivatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.InvitationSentAt != nil {
		toSerialize["invitation_sent_at"] = o.InvitationSentAt
	}
	if o.PasswordChangedAt != nil {
		toSerialize["password_changed_at"] = o.PasswordChangedAt
	}
	if o.LastLogin != nil {
		toSerialize["last_login"] = o.LastLogin
	}
	if o.LockedUntil != nil {
		toSerialize["locked_until"] = o.LockedUntil
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.OpenidName != nil {
		toSerialize["openid_name"] = o.OpenidName
	}
	if o.LocaleCode != nil {
		toSerialize["locale_code"] = o.LocaleCode
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.DistinguishedName != nil {
		toSerialize["distinguished_name"] = o.DistinguishedName
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.DirectoryId != nil {
		toSerialize["directory_id"] = o.DirectoryId
	}
	if o.MemberOf != nil {
		toSerialize["member_of"] = o.MemberOf
	}
	if o.Samaccountname != nil {
		toSerialize["samaccountname"] = o.Samaccountname
	}
	if o.Userprincipalname != nil {
		toSerialize["userprincipalname"] = o.Userprincipalname
	}
	if o.ManagerAdId != nil {
		toSerialize["manager_ad_id"] = o.ManagerAdId
	}
	if true {
		toSerialize["role_id"] = o.RoleId
	}
	if true {
		toSerialize["custom_attributes"] = o.CustomAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
