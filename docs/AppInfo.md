# AppInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**AuthMethod** | Pointer to **int32** |  | [optional] 
**ConnectorId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewAppInfo

`func NewAppInfo() *AppInfo`

NewAppInfo instantiates a new AppInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoWithDefaults

`func NewAppInfoWithDefaults() *AppInfo`

NewAppInfoWithDefaults instantiates a new AppInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AppInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AppInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisible

`func (o *AppInfo) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *AppInfo) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *AppInfo) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *AppInfo) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetDescription

`func (o *AppInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *AppInfo) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *AppInfo) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *AppInfo) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *AppInfo) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetAuthMethod

`func (o *AppInfo) GetAuthMethod() int32`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *AppInfo) GetAuthMethodOk() (*int32, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *AppInfo) SetAuthMethod(v int32)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *AppInfo) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetConnectorId

`func (o *AppInfo) GetConnectorId() int32`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *AppInfo) GetConnectorIdOk() (*int32, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *AppInfo) SetConnectorId(v int32)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *AppInfo) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AppInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AppInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AppInfo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppInfo) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppInfo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AppInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


