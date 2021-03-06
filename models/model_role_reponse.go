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

// RoleReponse struct for RoleReponse
type RoleReponse struct {
	Status *Status `json:"status,omitempty"`
	Data   []Role  `json:"data"`
}

// NewRoleReponse instantiates a new RoleReponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleReponse(data []Role) *RoleReponse {
	this := RoleReponse{}
	this.Data = data
	return &this
}

// NewRoleReponseWithDefaults instantiates a new RoleReponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleReponseWithDefaults() *RoleReponse {
	this := RoleReponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RoleReponse) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleReponse) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RoleReponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *RoleReponse) SetStatus(v Status) {
	o.Status = &v
}

// GetData returns the Data field value
func (o *RoleReponse) GetData() []Role {
	if o == nil {
		var ret []Role
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RoleReponse) GetDataOk() (*[]Role, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RoleReponse) SetData(v []Role) {
	o.Data = v
}

func (o RoleReponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableRoleReponse struct {
	value *RoleReponse
	isSet bool
}

func (v NullableRoleReponse) Get() *RoleReponse {
	return v.value
}

func (v *NullableRoleReponse) Set(val *RoleReponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleReponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleReponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleReponse(val *RoleReponse) *NullableRoleReponse {
	return &NullableRoleReponse{value: val, isSet: true}
}

func (v NullableRoleReponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleReponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
