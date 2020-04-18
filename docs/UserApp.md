# UserApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the app that can be accessed by the user. | [optional] 
**Name** | Pointer to **string** | Constant name for the Event Type | [optional] 
**Icon** | Pointer to **string** | Template for the Event description  | [optional] 
**Provisioned** | Pointer to **int32** | Indicates whether a username and password has been stored on the login for the app and user. Valid values are: - 0 (no) - 1 (yes) | [optional] 
**Extension** | Pointer to **bool** | Indicates whether the app requires the OneLogin browser extension to login. Valid values are: - true: The app requires the OneLogin browser extension. - false: The app does not require the OneLogin browser extension. | [optional] 
**LoginId** | Pointer to **string** | Template for the Event description  | [optional] 
**Personal** | Pointer to **bool** | Indicates whether the app is a user’s personal app. Valid values are: - true: The app is a user’s personal app. - false: The app is not a user’s personal app.         | [optional] 

## Methods

### NewUserApp

`func NewUserApp() *UserApp`

NewUserApp instantiates a new UserApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAppWithDefaults

`func NewUserAppWithDefaults() *UserApp`

NewUserAppWithDefaults instantiates a new UserApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserApp) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserApp) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserApp) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIcon

`func (o *UserApp) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UserApp) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UserApp) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *UserApp) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetProvisioned

`func (o *UserApp) GetProvisioned() int32`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *UserApp) GetProvisionedOk() (*int32, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *UserApp) SetProvisioned(v int32)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *UserApp) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetExtension

`func (o *UserApp) GetExtension() bool`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *UserApp) GetExtensionOk() (*bool, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *UserApp) SetExtension(v bool)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *UserApp) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetLoginId

`func (o *UserApp) GetLoginId() string`

GetLoginId returns the LoginId field if non-nil, zero value otherwise.

### GetLoginIdOk

`func (o *UserApp) GetLoginIdOk() (*string, bool)`

GetLoginIdOk returns a tuple with the LoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginId

`func (o *UserApp) SetLoginId(v string)`

SetLoginId sets LoginId field to given value.

### HasLoginId

`func (o *UserApp) HasLoginId() bool`

HasLoginId returns a boolean if a field has been set.

### GetPersonal

`func (o *UserApp) GetPersonal() bool`

GetPersonal returns the Personal field if non-nil, zero value otherwise.

### GetPersonalOk

`func (o *UserApp) GetPersonalOk() (*bool, bool)`

GetPersonalOk returns a tuple with the Personal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonal

`func (o *UserApp) SetPersonal(v bool)`

SetPersonal sets Personal field to given value.

### HasPersonal

`func (o *UserApp) HasPersonal() bool`

HasPersonal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


