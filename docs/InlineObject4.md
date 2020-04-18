# InlineObject4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The new password  | 
**PasswordConfirmation** | Pointer to **string** | The new password repeated | 
**ValidatePolicy** | Pointer to **string** | Defaults to false. This will validate the password against the users OneLogin password policy. | [optional] 

## Methods

### NewInlineObject4

`func NewInlineObject4(password string, passwordConfirmation string, ) *InlineObject4`

NewInlineObject4 instantiates a new InlineObject4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject4WithDefaults

`func NewInlineObject4WithDefaults() *InlineObject4`

NewInlineObject4WithDefaults instantiates a new InlineObject4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *InlineObject4) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject4) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject4) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPasswordConfirmation

`func (o *InlineObject4) GetPasswordConfirmation() string`

GetPasswordConfirmation returns the PasswordConfirmation field if non-nil, zero value otherwise.

### GetPasswordConfirmationOk

`func (o *InlineObject4) GetPasswordConfirmationOk() (*string, bool)`

GetPasswordConfirmationOk returns a tuple with the PasswordConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordConfirmation

`func (o *InlineObject4) SetPasswordConfirmation(v string)`

SetPasswordConfirmation sets PasswordConfirmation field to given value.


### GetValidatePolicy

`func (o *InlineObject4) GetValidatePolicy() string`

GetValidatePolicy returns the ValidatePolicy field if non-nil, zero value otherwise.

### GetValidatePolicyOk

`func (o *InlineObject4) GetValidatePolicyOk() (*string, bool)`

GetValidatePolicyOk returns a tuple with the ValidatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatePolicy

`func (o *InlineObject4) SetValidatePolicy(v string)`

SetValidatePolicy sets ValidatePolicy field to given value.

### HasValidatePolicy

`func (o *InlineObject4) HasValidatePolicy() bool`

HasValidatePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


