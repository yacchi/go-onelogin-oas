# InlineObject5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The new password  | 
**PasswordConfirmation** | Pointer to **string** | The new password repeated | 
**PasswordAlgorithm** | Pointer to **string** | Set to salt+sha256. | 
**PasswordSalt** | Pointer to **string** | If your password hash has been salted then you can provide the salt used in this param. | [optional] 

## Methods

### NewInlineObject5

`func NewInlineObject5(password string, passwordConfirmation string, passwordAlgorithm string, ) *InlineObject5`

NewInlineObject5 instantiates a new InlineObject5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject5WithDefaults

`func NewInlineObject5WithDefaults() *InlineObject5`

NewInlineObject5WithDefaults instantiates a new InlineObject5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *InlineObject5) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject5) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject5) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPasswordConfirmation

`func (o *InlineObject5) GetPasswordConfirmation() string`

GetPasswordConfirmation returns the PasswordConfirmation field if non-nil, zero value otherwise.

### GetPasswordConfirmationOk

`func (o *InlineObject5) GetPasswordConfirmationOk() (*string, bool)`

GetPasswordConfirmationOk returns a tuple with the PasswordConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordConfirmation

`func (o *InlineObject5) SetPasswordConfirmation(v string)`

SetPasswordConfirmation sets PasswordConfirmation field to given value.


### GetPasswordAlgorithm

`func (o *InlineObject5) GetPasswordAlgorithm() string`

GetPasswordAlgorithm returns the PasswordAlgorithm field if non-nil, zero value otherwise.

### GetPasswordAlgorithmOk

`func (o *InlineObject5) GetPasswordAlgorithmOk() (*string, bool)`

GetPasswordAlgorithmOk returns a tuple with the PasswordAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordAlgorithm

`func (o *InlineObject5) SetPasswordAlgorithm(v string)`

SetPasswordAlgorithm sets PasswordAlgorithm field to given value.


### GetPasswordSalt

`func (o *InlineObject5) GetPasswordSalt() string`

GetPasswordSalt returns the PasswordSalt field if non-nil, zero value otherwise.

### GetPasswordSaltOk

`func (o *InlineObject5) GetPasswordSaltOk() (*string, bool)`

GetPasswordSaltOk returns a tuple with the PasswordSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSalt

`func (o *InlineObject5) SetPasswordSalt(v string)`

SetPasswordSalt sets PasswordSalt field to given value.

### HasPasswordSalt

`func (o *InlineObject5) HasPasswordSalt() bool`

HasPasswordSalt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


