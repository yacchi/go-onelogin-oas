# AppParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**UserAttributeMappings** | Pointer to **string** |  | [optional] 
**UserAttributeMacros** | Pointer to **string** |  | [optional] 
**AttributesTransformations** | Pointer to **string** |  | [optional] 
**SkipIfBlank** | Pointer to **bool** |  | [optional] 
**Values** | Pointer to **string** |  | [optional] 
**DefaultValues** | Pointer to **string** |  | [optional] 
**ProvisionedEntitlements** | Pointer to **bool** |  | [optional] 
**SafeEntitlementsEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewAppParameters

`func NewAppParameters() *AppParameters`

NewAppParameters instantiates a new AppParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppParametersWithDefaults

`func NewAppParametersWithDefaults() *AppParameters`

NewAppParametersWithDefaults instantiates a new AppParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppParameters) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppParameters) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppParameters) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AppParameters) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *AppParameters) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AppParameters) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AppParameters) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AppParameters) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUserAttributeMappings

`func (o *AppParameters) GetUserAttributeMappings() string`

GetUserAttributeMappings returns the UserAttributeMappings field if non-nil, zero value otherwise.

### GetUserAttributeMappingsOk

`func (o *AppParameters) GetUserAttributeMappingsOk() (*string, bool)`

GetUserAttributeMappingsOk returns a tuple with the UserAttributeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeMappings

`func (o *AppParameters) SetUserAttributeMappings(v string)`

SetUserAttributeMappings sets UserAttributeMappings field to given value.

### HasUserAttributeMappings

`func (o *AppParameters) HasUserAttributeMappings() bool`

HasUserAttributeMappings returns a boolean if a field has been set.

### GetUserAttributeMacros

`func (o *AppParameters) GetUserAttributeMacros() string`

GetUserAttributeMacros returns the UserAttributeMacros field if non-nil, zero value otherwise.

### GetUserAttributeMacrosOk

`func (o *AppParameters) GetUserAttributeMacrosOk() (*string, bool)`

GetUserAttributeMacrosOk returns a tuple with the UserAttributeMacros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeMacros

`func (o *AppParameters) SetUserAttributeMacros(v string)`

SetUserAttributeMacros sets UserAttributeMacros field to given value.

### HasUserAttributeMacros

`func (o *AppParameters) HasUserAttributeMacros() bool`

HasUserAttributeMacros returns a boolean if a field has been set.

### GetAttributesTransformations

`func (o *AppParameters) GetAttributesTransformations() string`

GetAttributesTransformations returns the AttributesTransformations field if non-nil, zero value otherwise.

### GetAttributesTransformationsOk

`func (o *AppParameters) GetAttributesTransformationsOk() (*string, bool)`

GetAttributesTransformationsOk returns a tuple with the AttributesTransformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesTransformations

`func (o *AppParameters) SetAttributesTransformations(v string)`

SetAttributesTransformations sets AttributesTransformations field to given value.

### HasAttributesTransformations

`func (o *AppParameters) HasAttributesTransformations() bool`

HasAttributesTransformations returns a boolean if a field has been set.

### GetSkipIfBlank

`func (o *AppParameters) GetSkipIfBlank() bool`

GetSkipIfBlank returns the SkipIfBlank field if non-nil, zero value otherwise.

### GetSkipIfBlankOk

`func (o *AppParameters) GetSkipIfBlankOk() (*bool, bool)`

GetSkipIfBlankOk returns a tuple with the SkipIfBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIfBlank

`func (o *AppParameters) SetSkipIfBlank(v bool)`

SetSkipIfBlank sets SkipIfBlank field to given value.

### HasSkipIfBlank

`func (o *AppParameters) HasSkipIfBlank() bool`

HasSkipIfBlank returns a boolean if a field has been set.

### GetValues

`func (o *AppParameters) GetValues() string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AppParameters) GetValuesOk() (*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AppParameters) SetValues(v string)`

SetValues sets Values field to given value.

### HasValues

`func (o *AppParameters) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetDefaultValues

`func (o *AppParameters) GetDefaultValues() string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *AppParameters) GetDefaultValuesOk() (*string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *AppParameters) SetDefaultValues(v string)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *AppParameters) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetProvisionedEntitlements

`func (o *AppParameters) GetProvisionedEntitlements() bool`

GetProvisionedEntitlements returns the ProvisionedEntitlements field if non-nil, zero value otherwise.

### GetProvisionedEntitlementsOk

`func (o *AppParameters) GetProvisionedEntitlementsOk() (*bool, bool)`

GetProvisionedEntitlementsOk returns a tuple with the ProvisionedEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedEntitlements

`func (o *AppParameters) SetProvisionedEntitlements(v bool)`

SetProvisionedEntitlements sets ProvisionedEntitlements field to given value.

### HasProvisionedEntitlements

`func (o *AppParameters) HasProvisionedEntitlements() bool`

HasProvisionedEntitlements returns a boolean if a field has been set.

### GetSafeEntitlementsEnabled

`func (o *AppParameters) GetSafeEntitlementsEnabled() bool`

GetSafeEntitlementsEnabled returns the SafeEntitlementsEnabled field if non-nil, zero value otherwise.

### GetSafeEntitlementsEnabledOk

`func (o *AppParameters) GetSafeEntitlementsEnabledOk() (*bool, bool)`

GetSafeEntitlementsEnabledOk returns a tuple with the SafeEntitlementsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeEntitlementsEnabled

`func (o *AppParameters) SetSafeEntitlementsEnabled(v bool)`

SetSafeEntitlementsEnabled sets SafeEntitlementsEnabled field to given value.

### HasSafeEntitlementsEnabled

`func (o *AppParameters) HasSafeEntitlementsEnabled() bool`

HasSafeEntitlementsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


