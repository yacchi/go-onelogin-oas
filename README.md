# Go API client for github.com/yacchi/go-onelogin-oas

This is an administrative API for OneLogin customers

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.
By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.0-oas3
- Package version: 0.2.1
- Build package: MyGoClientCodegen

## Installation

Install the following dependencies:

Put the package under your project folder and add the following in import:

```golang
import "github.com/yacchi/go-onelogin-oas"
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `onelogin.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), onelogin.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `onelogin.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), onelogin.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `onelogin.ContextOperationServerIndices` and `onelogin.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), onelogin.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), onelogin.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.us.onelogin.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApp**](docs/AppsApi.md#createapp) | **Post** /2/apps | Create an App
[**DeleteApp**](docs/AppsApi.md#deleteapp) | **Delete** /2/apps/{id} | Delete an app
[**DeleteAppParameter**](docs/AppsApi.md#deleteappparameter) | **Delete** /2/apps/{id}/parameters/{parameter_id} | Delete an app parameter
[**GetApp**](docs/AppsApi.md#getapp) | **Get** /2/apps/{id} | Get an App
[**GetApps**](docs/AppsApi.md#getapps) | **Get** /2/apps | Get Apps
[**UpdateApp**](docs/AppsApi.md#updateapp) | **Put** /2/apps/{id} | Update an App
[**GetConnectors**](docs/ConnectorsApi.md#getconnectors) | **Get** /2/connectors | Get Connectors
[**GetEvent**](docs/EventsApi.md#getevent) | **Get** /1/events/{id} | Get Event by ID
[**GetEventTypes**](docs/EventsApi.md#geteventtypes) | **Get** /1/events/types | Get Event Types
[**GetEvents**](docs/EventsApi.md#getevents) | **Get** /1/events | Get Events
[**GetGroup**](docs/GroupsApi.md#getgroup) | **Get** /1/groups/{id} | Get Group by ID
[**GetGroups**](docs/GroupsApi.md#getgroups) | **Get** /1/groups | Get Groups
[**CreateInviteLink**](docs/InvitesApi.md#createinvitelink) | **Post** /1/invites/get_invite_link | Generate Invite Link
[**SendInviteLink**](docs/InvitesApi.md#sendinvitelink) | **Post** /1/invites/send_invite_link | Send Invite Link
[**AuthenticateUser**](docs/LoginApi.md#authenticateuser) | **Post** /1/login/auth | Authenticate a user
[**VerifyLoginMFAToken**](docs/LoginApi.md#verifyloginmfatoken) | **Post** /1/login/verify_factor | Verify an MFA token
[**ActivateUserMFADevice**](docs/MfaApi.md#activateusermfadevice) | **Post** /1/users/{id}/otp_devices/{device_id}/trigger | Activate an authentication factor
[**DeleteUserMFADevice**](docs/MfaApi.md#deleteusermfadevice) | **Delete** /1/users/{id}/otp_devices/{device_id} | Remove an authentication device
[**EnrollUserMFADevice**](docs/MfaApi.md#enrollusermfadevice) | **Post** /1/users/{id}/otp_devices | Enroll a user with a given authentication factor.
[**GetUserAvailableMFAFactors**](docs/MfaApi.md#getuseravailablemfafactors) | **Get** /1/users/{id}/auth_factors | Get a list of MFA factors available to user
[**GetUserEnrolledMFADevices**](docs/MfaApi.md#getuserenrolledmfadevices) | **Get** /1/users/{id}/otp_devices | Get enrolled authentication devices
[**VerifyUserMFADevice**](docs/MfaApi.md#verifyusermfadevice) | **Post** /1/users/{id}/otp_devices/{device_id}/verify | Verify an authentication device
[**AddPrivilegeRoles**](docs/PrivilegesApi.md#addprivilegeroles) | **Post** /1/privileges/{id}/roles | Assign roles
[**AddPrivilegeUsers**](docs/PrivilegesApi.md#addprivilegeusers) | **Post** /1/privileges/{id}/users | Assign users
[**CreatePrivilege**](docs/PrivilegesApi.md#createprivilege) | **Post** /1/privileges | Creates privilege
[**DeletePrivilege**](docs/PrivilegesApi.md#deleteprivilege) | **Delete** /1/privileges/{id} | Delete privilege
[**GetPrivilege**](docs/PrivilegesApi.md#getprivilege) | **Get** /1/privileges/{id} | Get privilege
[**GetPrivilegeRoles**](docs/PrivilegesApi.md#getprivilegeroles) | **Get** /1/privileges/{id}/roles | Get roles
[**GetPrivilegeUsers**](docs/PrivilegesApi.md#getprivilegeusers) | **Get** /1/privileges/{id}/users | Get privilege users
[**GetPrivileges**](docs/PrivilegesApi.md#getprivileges) | **Get** /1/privileges | Get Privileges
[**RemovePrivilegeRole**](docs/PrivilegesApi.md#removeprivilegerole) | **Delete** /1/privileges/{id}/roles/{role_id} | Remove a role
[**RemovePrivlegeUser**](docs/PrivilegesApi.md#removeprivlegeuser) | **Delete** /1/privileges/{id}/users/{user_id} | Remove a user
[**UpdatePrivilege**](docs/PrivilegesApi.md#updateprivilege) | **Put** /1/privileges/{id} | Update privilege
[**GetRole**](docs/RolesApi.md#getrole) | **Get** /1/roles/{id} | Get Role by ID
[**GetRoles**](docs/RolesApi.md#getroles) | **Get** /1/roles | Get Roles
[**CreateSAMLAssertion**](docs/SamlApi.md#createsamlassertion) | **Post** /1/saml_assertion | Generate SAML assertion
[**VerifySAMLAssertionMFAToken**](docs/SamlApi.md#verifysamlassertionmfatoken) | **Post** /1/saml_assertion/verify_factor | Verify an MFA token
[**AddUserRoles**](docs/UsersApi.md#adduserroles) | **Put** /1/users/{id}/add_roles | Assign Role to User
[**Call1UsersPost**](docs/UsersApi.md#call1userspost) | **Post** /1/users | Create a User
[**CreateUserTempMFAToken**](docs/UsersApi.md#createusertempmfatoken) | **Post** /1/users/{id}/mfa_token | Generate Temp MFA Token
[**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /1/users/{id} | Delete a user account
[**GetCustomAttributes**](docs/UsersApi.md#getcustomattributes) | **Get** /1/users/custom_attributes | Get Custom Attributes
[**GetUser**](docs/UsersApi.md#getuser) | **Get** /1/users/{id} | Get a User
[**GetUserApps**](docs/UsersApi.md#getuserapps) | **Get** /1/users/{id}/apps | Get User Apps
[**GetUserRoles**](docs/UsersApi.md#getuserroles) | **Get** /1/users/{id}/roles | Get User Roles
[**GetUsers**](docs/UsersApi.md#getusers) | **Get** /1/users | Get Users
[**LockUser**](docs/UsersApi.md#lockuser) | **Put** /1/users/{id}/lock_user | Lock a user account
[**LogoutUser**](docs/UsersApi.md#logoutuser) | **Put** /1/users/{id}/logout | Log a user out of any and all sessions
[**RemoveUserRoles**](docs/UsersApi.md#removeuserroles) | **Put** /1/users/{id}/remove_roles | Remove Roles from User
[**UpdateUser**](docs/UsersApi.md#updateuser) | **Put** /1/users/{id} | Update a User
[**UpdateUserCustomAttributes**](docs/UsersApi.md#updateusercustomattributes) | **Put** /1/users/{id}/set_custom_attributes | Set a custom attribute
[**UpdateUserPassword**](docs/UsersApi.md#updateuserpassword) | **Put** /1/users/set_password_clear_text/{id} | Set a the password for a user
[**UpdateUserPasswordSalted**](docs/UsersApi.md#updateuserpasswordsalted) | **Put** /1/users/set_password_using_salt/{id} | Set a pre salted password for a user
[**UpdateUserState**](docs/UsersApi.md#updateuserstate) | **Put** /1/users/{id}/set_state | Set the State for a user.


## Documentation For Models

 - [Action](docs/Action.md)
 - [App](docs/App.md)
 - [AppConfiguration](docs/AppConfiguration.md)
 - [AppInfo](docs/AppInfo.md)
 - [AppParameters](docs/AppParameters.md)
 - [AppProvisioning](docs/AppProvisioning.md)
 - [AppSso](docs/AppSso.md)
 - [AppSsoCertificate](docs/AppSsoCertificate.md)
 - [AssignPrivilegeRolesResponse](docs/AssignPrivilegeRolesResponse.md)
 - [Connector](docs/Connector.md)
 - [CreatePrivilegeResponse](docs/CreatePrivilegeResponse.md)
 - [CreateUserResponse](docs/CreateUserResponse.md)
 - [CustomAttributesResponse](docs/CustomAttributesResponse.md)
 - [EnrollMfaDeviceResponse](docs/EnrollMfaDeviceResponse.md)
 - [EnrolledMfaDevicesResponse](docs/EnrolledMfaDevicesResponse.md)
 - [EnrolledMfaDevicesResponseData](docs/EnrolledMfaDevicesResponseData.md)
 - [Event](docs/Event.md)
 - [EventResponse](docs/EventResponse.md)
 - [EventType](docs/EventType.md)
 - [EventTypesResponse](docs/EventTypesResponse.md)
 - [EventsResponse](docs/EventsResponse.md)
 - [GenerateInviteLinkResponse](docs/GenerateInviteLinkResponse.md)
 - [GenerateMfaTokenResponse](docs/GenerateMfaTokenResponse.md)
 - [Group](docs/Group.md)
 - [GroupResponse](docs/GroupResponse.md)
 - [GroupsResponse](docs/GroupsResponse.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject10](docs/InlineObject10.md)
 - [InlineObject11](docs/InlineObject11.md)
 - [InlineObject12](docs/InlineObject12.md)
 - [InlineObject13](docs/InlineObject13.md)
 - [InlineObject14](docs/InlineObject14.md)
 - [InlineObject15](docs/InlineObject15.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineObject3](docs/InlineObject3.md)
 - [InlineObject4](docs/InlineObject4.md)
 - [InlineObject5](docs/InlineObject5.md)
 - [InlineObject6](docs/InlineObject6.md)
 - [InlineObject7](docs/InlineObject7.md)
 - [InlineObject8](docs/InlineObject8.md)
 - [InlineObject9](docs/InlineObject9.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [LoginVerifyMfaResponse](docs/LoginVerifyMfaResponse.md)
 - [LoginVerifyMfaResponseData](docs/LoginVerifyMfaResponseData.md)
 - [MfaDevice](docs/MfaDevice.md)
 - [Pagination](docs/Pagination.md)
 - [Privilege](docs/Privilege.md)
 - [PrivilegePrivilege](docs/PrivilegePrivilege.md)
 - [PrivilegeRolesResponse](docs/PrivilegeRolesResponse.md)
 - [PrivilegeUsersResponse](docs/PrivilegeUsersResponse.md)
 - [Response](docs/Response.md)
 - [ResponseDevices](docs/ResponseDevices.md)
 - [ResponseUser](docs/ResponseUser.md)
 - [Role](docs/Role.md)
 - [RoleReponse](docs/RoleReponse.md)
 - [RolesResponse](docs/RolesResponse.md)
 - [SamlAssertionResponse](docs/SamlAssertionResponse.md)
 - [Statement](docs/Statement.md)
 - [Status](docs/Status.md)
 - [User](docs/User.md)
 - [UserApp](docs/UserApp.md)
 - [UserAppsResponse](docs/UserAppsResponse.md)
 - [UserLoginResponse](docs/UserLoginResponse.md)
 - [UserMfaFactorsResponse](docs/UserMfaFactorsResponse.md)
 - [UserMfaFactorsResponseData](docs/UserMfaFactorsResponseData.md)
 - [UserMfaFactorsResponseDataAuthFactors](docs/UserMfaFactorsResponseDataAuthFactors.md)
 - [UserResponse](docs/UserResponse.md)
 - [UserRolesResponse](docs/UserRolesResponse.md)
 - [UsersResponse](docs/UsersResponse.md)


## Documentation For Authorization



### application


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), onelogin.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, onelogin.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Licence
[MIT](./LICENSE)


