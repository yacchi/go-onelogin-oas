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
)

// Action the model 'Action'
type Action string

// List of Action
const (
	STAR                         Action = "*"
	PRIVILEGESLIST               Action = "privileges:List"
	PRIVILEGESGET                Action = "privileges:Get"
	PRIVILEGESCREATE             Action = "privileges:Create"
	PRIVILEGESUPDATE             Action = "privileges:Update"
	PRIVILEGESDELETE             Action = "privileges:Delete"
	PRIVILEGESLIST_USERS         Action = "privileges:ListUsers"
	PRIVILEGESLIST_ROLES         Action = "privileges:ListRoles"
	PRIVILEGESADD_USER           Action = "privileges:AddUser"
	PRIVILEGESREMOVE_USER        Action = "privileges:RemoveUser"
	PRIVILEGESADD_ROLE           Action = "privileges:AddRole"
	PRIVILEGESREMOVE_ROLE        Action = "privileges:RemoveRole"
	ROLESLIST                    Action = "roles:List"
	ROLESGET                     Action = "roles:Get"
	ROLESCREATE                  Action = "roles:Create"
	ROLESUPDATE                  Action = "roles:Update"
	ROLESDELETE                  Action = "roles:Delete"
	ROLESADD_USER                Action = "roles:AddUser"
	ROLESADD_APP                 Action = "roles:AddApp"
	ROLESREMOVE_APP              Action = "roles:RemoveApp"
	DIRECTORIESLIST              Action = "directories:List"
	DIRECTORIESGET               Action = "directories:Get"
	DIRECTORIESCREATE            Action = "directories:Create"
	DIRECTORIESUPDATE            Action = "directories:Update"
	DIRECTORIESDELETE            Action = "directories:Delete"
	DIRECTORIESSYNC_USERS        Action = "directories:SyncUsers"
	DIRECTORIESREFRESH_SCHEMA    Action = "directories:RefreshSchema"
	TRUSTEDIDPLIST               Action = "trustedidp:List"
	TRUSTEDIDPGET                Action = "trustedidp:Get"
	TRUSTEDIDPCREATE             Action = "trustedidp:Create"
	TRUSTEDIDPUPDATE             Action = "trustedidp:Update"
	TRUSTEDIDPDELETE             Action = "trustedidp:Delete"
	POLICIESLIST                 Action = "policies:List"
	POLICIESUSERGET              Action = "policies:user:Get"
	POLICIESUSERCREATE           Action = "policies:user:Create"
	POLICIESUSERUPDATE           Action = "policies:user:Update"
	POLICIESUSERDELETE           Action = "policies:user:Delete"
	POLICIESAPPGET               Action = "policies:app:Get"
	POLICIESAPPCREATE            Action = "policies:app:Create"
	POLICIESAPPUPDATE            Action = "policies:app:Update"
	POLICIESAPPDELETE            Action = "policies:app:Delete"
	USERSLIST                    Action = "users:List"
	USERSGET                     Action = "users:Get"
	USERSCREATE                  Action = "users:Create"
	USERSUPDATE                  Action = "users:Update"
	USERSDELETE                  Action = "users:Delete"
	USERSUNLOCK                  Action = "users:Unlock"
	USERSRESET_PASSWORD          Action = "users:ResetPassword"
	USERSFORCE_LOGOUT            Action = "users:ForceLogout"
	USERSINVITE                  Action = "users:Invite"
	USERSREAPPLY_MAPPINGS        Action = "users:ReapplyMappings"
	USERSADD_ROLE                Action = "users:AddRole"
	USERSREMOVE_ROLE             Action = "users:RemoveRole"
	USERSGENERATE_TEMP_MFA_TOKEN Action = "users:GenerateTempMfaToken"
	USERSADD_APP                 Action = "users:AddApp"
	USERSREMOVE_APP              Action = "users:RemoveApp"
	APPSLIST                     Action = "apps:List"
	APPSGET                      Action = "apps:Get"
	APPSCREATE                   Action = "apps:Create"
	APPSUPDATE_CONFIGURATION     Action = "apps:UpdateConfiguration"
	APPSUPDATE_SSO               Action = "apps:UpdateSSO"
	APPSUPDATE_PARAMETERS        Action = "apps:UpdateParameters"
	APPSDELETE                   Action = "apps:Delete"
	APPSADD_USER                 Action = "apps:AddUser"
	APPSREMOVE_USER              Action = "apps:RemoveUser"
	REPORTSLIST                  Action = "reports:List"
	REPORTSGET                   Action = "reports:Get"
	REPORTSCREATE                Action = "reports:Create"
	REPORTSRUN                   Action = "reports:Run"
	MAPPINGSREAPPLY_ALL          Action = "mappings:reapplyAll"
	EVENTSLIST                   Action = "events:List"
	EVENTSGET                    Action = "events:Get"
)

// Ptr returns reference to Action value
func (v Action) Ptr() *Action {
	return &v
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
