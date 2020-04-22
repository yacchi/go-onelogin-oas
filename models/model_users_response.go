// Code generated by "openapi-codegen"; DO NOT EDIT.
/*
 * OneLogin API
 *
 * This is an administrative API for OneLogin customers
 *
 * API version: 1.1.0-oas3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"encoding/json"
)

// UsersResponse struct for UsersResponse
type UsersResponse struct {
	Status     *Status     `json:"status,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       []User      `json:"data"`
}

// NewUsersResponse instantiates a new UsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersResponse(data []User) *UsersResponse {
	this := UsersResponse{}
	this.Data = data
	return &this
}

// NewUsersResponseWithDefaults instantiates a new UsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersResponseWithDefaults() *UsersResponse {
	this := UsersResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UsersResponse) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersResponse) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UsersResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *UsersResponse) SetStatus(v Status) {
	o.Status = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *UsersResponse) GetPagination() Pagination {
	if o == nil || o.Pagination == nil {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *UsersResponse) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *UsersResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value
func (o *UsersResponse) GetData() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UsersResponse) GetDataOk() (*[]User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UsersResponse) SetData(v []User) {
	o.Data = v
}

func (o UsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUsersResponse struct {
	value *UsersResponse
	isSet bool
}

func (v NullableUsersResponse) Get() *UsersResponse {
	return v.value
}

func (v *NullableUsersResponse) Set(val *UsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersResponse(val *UsersResponse) *NullableUsersResponse {
	return &NullableUsersResponse{value: val, isSet: true}
}

func (v NullableUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
