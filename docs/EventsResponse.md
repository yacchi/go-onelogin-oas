# EventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]Event**](Event.md) |  | [optional] 

## Methods

### NewEventsResponse

`func NewEventsResponse() *EventsResponse`

NewEventsResponse instantiates a new EventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsResponseWithDefaults

`func NewEventsResponseWithDefaults() *EventsResponse`

NewEventsResponseWithDefaults instantiates a new EventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EventsResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventsResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventsResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPagination

`func (o *EventsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *EventsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *EventsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *EventsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *EventsResponse) GetData() []Event`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventsResponse) GetDataOk() (*[]Event, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventsResponse) SetData(v []Event)`

SetData sets Data field to given value.

### HasData

`func (o *EventsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


