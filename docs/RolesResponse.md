# RolesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]Role**](Role.md) |  | [optional] 

## Methods

### NewRolesResponse

`func NewRolesResponse() *RolesResponse`

NewRolesResponse instantiates a new RolesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesResponseWithDefaults

`func NewRolesResponseWithDefaults() *RolesResponse`

NewRolesResponseWithDefaults instantiates a new RolesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RolesResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RolesResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RolesResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RolesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPagination

`func (o *RolesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RolesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RolesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RolesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *RolesResponse) GetData() []Role`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RolesResponse) GetDataOk() (*[]Role, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RolesResponse) SetData(v []Role)`

SetData sets Data field to given value.

### HasData

`func (o *RolesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


