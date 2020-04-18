# SamlAssertionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 

## Methods

### NewSamlAssertionResponse

`func NewSamlAssertionResponse() *SamlAssertionResponse`

NewSamlAssertionResponse instantiates a new SamlAssertionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlAssertionResponseWithDefaults

`func NewSamlAssertionResponseWithDefaults() *SamlAssertionResponse`

NewSamlAssertionResponseWithDefaults instantiates a new SamlAssertionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SamlAssertionResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SamlAssertionResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SamlAssertionResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SamlAssertionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *SamlAssertionResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SamlAssertionResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SamlAssertionResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SamlAssertionResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


