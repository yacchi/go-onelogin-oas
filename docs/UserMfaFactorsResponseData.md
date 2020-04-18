# UserMfaFactorsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthFactors** | Pointer to [**[]UserMfaFactorsResponseDataAuthFactors**](UserMfaFactorsResponse_data_auth_factors.md) |  | [optional] 

## Methods

### NewUserMfaFactorsResponseData

`func NewUserMfaFactorsResponseData() *UserMfaFactorsResponseData`

NewUserMfaFactorsResponseData instantiates a new UserMfaFactorsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMfaFactorsResponseDataWithDefaults

`func NewUserMfaFactorsResponseDataWithDefaults() *UserMfaFactorsResponseData`

NewUserMfaFactorsResponseDataWithDefaults instantiates a new UserMfaFactorsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthFactors

`func (o *UserMfaFactorsResponseData) GetAuthFactors() []UserMfaFactorsResponseDataAuthFactors`

GetAuthFactors returns the AuthFactors field if non-nil, zero value otherwise.

### GetAuthFactorsOk

`func (o *UserMfaFactorsResponseData) GetAuthFactorsOk() (*[]UserMfaFactorsResponseDataAuthFactors, bool)`

GetAuthFactorsOk returns a tuple with the AuthFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFactors

`func (o *UserMfaFactorsResponseData) SetAuthFactors(v []UserMfaFactorsResponseDataAuthFactors)`

SetAuthFactors sets AuthFactors field to given value.

### HasAuthFactors

`func (o *UserMfaFactorsResponseData) HasAuthFactors() bool`

HasAuthFactors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


