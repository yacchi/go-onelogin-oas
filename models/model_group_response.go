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

// GroupResponse struct for GroupResponse
type GroupResponse struct {
	Status *Status `json:"status,omitempty"`
	Data   []Group `json:"data"`
}

// NewGroupResponse instantiates a new GroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupResponse(data []Group) *GroupResponse {
	this := GroupResponse{}
	this.Data = data
	return &this
}

// NewGroupResponseWithDefaults instantiates a new GroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupResponseWithDefaults() *GroupResponse {
	this := GroupResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GroupResponse) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupResponse) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GroupResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *GroupResponse) SetStatus(v Status) {
	o.Status = &v
}

// GetData returns the Data field value
func (o *GroupResponse) GetData() []Group {
	if o == nil {
		var ret []Group
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GroupResponse) GetDataOk() (*[]Group, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GroupResponse) SetData(v []Group) {
	o.Data = v
}

func (o GroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGroupResponse struct {
	value *GroupResponse
	isSet bool
}

func (v NullableGroupResponse) Get() *GroupResponse {
	return v.value
}

func (v *NullableGroupResponse) Set(val *GroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupResponse(val *GroupResponse) *NullableGroupResponse {
	return &NullableGroupResponse{value: val, isSet: true}
}

func (v NullableGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
