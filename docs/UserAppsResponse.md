# UserAppsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]UserApp**](UserApp.md) |  | 

## Methods

### NewUserAppsResponse

`func NewUserAppsResponse(data []UserApp, ) *UserAppsResponse`

NewUserAppsResponse instantiates a new UserAppsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAppsResponseWithDefaults

`func NewUserAppsResponseWithDefaults() *UserAppsResponse`

NewUserAppsResponseWithDefaults instantiates a new UserAppsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserAppsResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserAppsResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserAppsResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserAppsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPagination

`func (o *UserAppsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *UserAppsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *UserAppsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *UserAppsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *UserAppsResponse) GetData() []UserApp`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserAppsResponse) GetDataOk() (*[]UserApp, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserAppsResponse) SetData(v []UserApp)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


