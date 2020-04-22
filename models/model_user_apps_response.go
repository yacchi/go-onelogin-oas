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

// UserAppsResponse struct for UserAppsResponse
type UserAppsResponse struct {
	Status     *Status     `json:"status,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       []UserApp   `json:"data"`
}

// NewUserAppsResponse instantiates a new UserAppsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAppsResponse(data []UserApp) *UserAppsResponse {
	this := UserAppsResponse{}
	this.Data = data
	return &this
}

// NewUserAppsResponseWithDefaults instantiates a new UserAppsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAppsResponseWithDefaults() *UserAppsResponse {
	this := UserAppsResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserAppsResponse) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAppsResponse) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserAppsResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *UserAppsResponse) SetStatus(v Status) {
	o.Status = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *UserAppsResponse) GetPagination() Pagination {
	if o == nil || o.Pagination == nil {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAppsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *UserAppsResponse) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *UserAppsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value
func (o *UserAppsResponse) GetData() []UserApp {
	if o == nil {
		var ret []UserApp
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserAppsResponse) GetDataOk() (*[]UserApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UserAppsResponse) SetData(v []UserApp) {
	o.Data = v
}

func (o UserAppsResponse) MarshalJSON() ([]byte, error) {
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

type NullableUserAppsResponse struct {
	value *UserAppsResponse
	isSet bool
}

func (v NullableUserAppsResponse) Get() *UserAppsResponse {
	return v.value
}

func (v *NullableUserAppsResponse) Set(val *UserAppsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAppsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAppsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAppsResponse(val *UserAppsResponse) *NullableUserAppsResponse {
	return &NullableUserAppsResponse{value: val, isSet: true}
}

func (v NullableUserAppsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAppsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
