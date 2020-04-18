# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**AuthMethod** | Pointer to **int32** |  | [optional] 
**PolicyId** | Pointer to **int32** |  | [optional] 
**AllowAssumedSignin** | Pointer to **bool** |  | [optional] 
**TabId** | Pointer to **int32** |  | [optional] 
**ConnectorId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Provisioning** | Pointer to [**AppProvisioning**](App_provisioning.md) |  | [optional] 
**Sso** | Pointer to [**AppSso**](App_sso.md) |  | [optional] 
**Configuration** | Pointer to [**AppConfiguration**](App_configuration.md) |  | [optional] 
**Parameters** | Pointer to [**map[string]AppParameters**](App_parameters.md) |  | [optional] 

## Methods

### NewApp

`func NewApp() *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *App) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *App) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *App) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *App) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *App) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *App) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisible

`func (o *App) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *App) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *App) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *App) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetDescription

`func (o *App) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *App) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *App) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *App) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotes

`func (o *App) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *App) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *App) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *App) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIconUrl

`func (o *App) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *App) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *App) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *App) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetAuthMethod

`func (o *App) GetAuthMethod() int32`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *App) GetAuthMethodOk() (*int32, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *App) SetAuthMethod(v int32)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *App) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetPolicyId

`func (o *App) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *App) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *App) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *App) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetAllowAssumedSignin

`func (o *App) GetAllowAssumedSignin() bool`

GetAllowAssumedSignin returns the AllowAssumedSignin field if non-nil, zero value otherwise.

### GetAllowAssumedSigninOk

`func (o *App) GetAllowAssumedSigninOk() (*bool, bool)`

GetAllowAssumedSigninOk returns a tuple with the AllowAssumedSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAssumedSignin

`func (o *App) SetAllowAssumedSignin(v bool)`

SetAllowAssumedSignin sets AllowAssumedSignin field to given value.

### HasAllowAssumedSignin

`func (o *App) HasAllowAssumedSignin() bool`

HasAllowAssumedSignin returns a boolean if a field has been set.

### GetTabId

`func (o *App) GetTabId() int32`

GetTabId returns the TabId field if non-nil, zero value otherwise.

### GetTabIdOk

`func (o *App) GetTabIdOk() (*int32, bool)`

GetTabIdOk returns a tuple with the TabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabId

`func (o *App) SetTabId(v int32)`

SetTabId sets TabId field to given value.

### HasTabId

`func (o *App) HasTabId() bool`

HasTabId returns a boolean if a field has been set.

### GetConnectorId

`func (o *App) GetConnectorId() int32`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *App) GetConnectorIdOk() (*int32, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *App) SetConnectorId(v int32)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *App) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *App) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *App) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *App) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *App) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *App) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *App) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *App) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *App) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProvisioning

`func (o *App) GetProvisioning() AppProvisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *App) GetProvisioningOk() (*AppProvisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *App) SetProvisioning(v AppProvisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *App) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetSso

`func (o *App) GetSso() AppSso`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *App) GetSsoOk() (*AppSso, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *App) SetSso(v AppSso)`

SetSso sets Sso field to given value.

### HasSso

`func (o *App) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetConfiguration

`func (o *App) GetConfiguration() AppConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *App) GetConfigurationOk() (*AppConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *App) SetConfiguration(v AppConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *App) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetParameters

`func (o *App) GetParameters() map[string]AppParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *App) GetParametersOk() (*map[string]AppParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *App) SetParameters(v map[string]AppParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *App) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


