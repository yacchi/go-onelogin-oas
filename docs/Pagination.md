# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeforeCursor** | Pointer to **string** |  | [optional] 
**AfterCursor** | Pointer to **string** |  | [optional] 
**PreviousLink** | Pointer to **string** |  | [optional] 
**NextLink** | Pointer to **string** |  | [optional] 

## Methods

### NewPagination

`func NewPagination() *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeforeCursor

`func (o *Pagination) GetBeforeCursor() string`

GetBeforeCursor returns the BeforeCursor field if non-nil, zero value otherwise.

### GetBeforeCursorOk

`func (o *Pagination) GetBeforeCursorOk() (*string, bool)`

GetBeforeCursorOk returns a tuple with the BeforeCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeCursor

`func (o *Pagination) SetBeforeCursor(v string)`

SetBeforeCursor sets BeforeCursor field to given value.

### HasBeforeCursor

`func (o *Pagination) HasBeforeCursor() bool`

HasBeforeCursor returns a boolean if a field has been set.

### GetAfterCursor

`func (o *Pagination) GetAfterCursor() string`

GetAfterCursor returns the AfterCursor field if non-nil, zero value otherwise.

### GetAfterCursorOk

`func (o *Pagination) GetAfterCursorOk() (*string, bool)`

GetAfterCursorOk returns a tuple with the AfterCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterCursor

`func (o *Pagination) SetAfterCursor(v string)`

SetAfterCursor sets AfterCursor field to given value.

### HasAfterCursor

`func (o *Pagination) HasAfterCursor() bool`

HasAfterCursor returns a boolean if a field has been set.

### GetPreviousLink

`func (o *Pagination) GetPreviousLink() string`

GetPreviousLink returns the PreviousLink field if non-nil, zero value otherwise.

### GetPreviousLinkOk

`func (o *Pagination) GetPreviousLinkOk() (*string, bool)`

GetPreviousLinkOk returns a tuple with the PreviousLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousLink

`func (o *Pagination) SetPreviousLink(v string)`

SetPreviousLink sets PreviousLink field to given value.

### HasPreviousLink

`func (o *Pagination) HasPreviousLink() bool`

HasPreviousLink returns a boolean if a field has been set.

### GetNextLink

`func (o *Pagination) GetNextLink() string`

GetNextLink returns the NextLink field if non-nil, zero value otherwise.

### GetNextLinkOk

`func (o *Pagination) GetNextLinkOk() (*string, bool)`

GetNextLinkOk returns a tuple with the NextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLink

`func (o *Pagination) SetNextLink(v string)`

SetNextLink sets NextLink field to given value.

### HasNextLink

`func (o *Pagination) HasNextLink() bool`

HasNextLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


