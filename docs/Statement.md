# Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | Pointer to **string** |  | 
**Action** | Pointer to [**[]Action**](Action.md) |  | 
**Scope** | Pointer to **[]string** |  | 

## Methods

### NewStatement

`func NewStatement(effect string, action []Action, scope []string, ) *Statement`

NewStatement instantiates a new Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementWithDefaults

`func NewStatementWithDefaults() *Statement`

NewStatementWithDefaults instantiates a new Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *Statement) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *Statement) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *Statement) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetAction

`func (o *Statement) GetAction() []Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Statement) GetActionOk() (*[]Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Statement) SetAction(v []Action)`

SetAction sets Action field to given value.


### GetScope

`func (o *Statement) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Statement) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Statement) SetScope(v []string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


