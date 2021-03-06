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

// UserMfaFactorsResponseDataAuthFactors struct for UserMfaFactorsResponseDataAuthFactors
type UserMfaFactorsResponseDataAuthFactors struct {
	FactorId *int32  `json:"factor_id,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// NewUserMfaFactorsResponseDataAuthFactors instantiates a new UserMfaFactorsResponseDataAuthFactors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMfaFactorsResponseDataAuthFactors() *UserMfaFactorsResponseDataAuthFactors {
	this := UserMfaFactorsResponseDataAuthFactors{}
	return &this
}

// NewUserMfaFactorsResponseDataAuthFactorsWithDefaults instantiates a new UserMfaFactorsResponseDataAuthFactors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMfaFactorsResponseDataAuthFactorsWithDefaults() *UserMfaFactorsResponseDataAuthFactors {
	this := UserMfaFactorsResponseDataAuthFactors{}
	return &this
}

// GetFactorId returns the FactorId field value if set, zero value otherwise.
func (o *UserMfaFactorsResponseDataAuthFactors) GetFactorId() int32 {
	if o == nil || o.FactorId == nil {
		var ret int32
		return ret
	}
	return *o.FactorId
}

// GetFactorIdOk returns a tuple with the FactorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMfaFactorsResponseDataAuthFactors) GetFactorIdOk() (*int32, bool) {
	if o == nil || o.FactorId == nil {
		return nil, false
	}
	return o.FactorId, true
}

// HasFactorId returns a boolean if a field has been set.
func (o *UserMfaFactorsResponseDataAuthFactors) HasFactorId() bool {
	if o != nil && o.FactorId != nil {
		return true
	}

	return false
}

// SetFactorId gets a reference to the given int32 and assigns it to the FactorId field.
func (o *UserMfaFactorsResponseDataAuthFactors) SetFactorId(v int32) {
	o.FactorId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserMfaFactorsResponseDataAuthFactors) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMfaFactorsResponseDataAuthFactors) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserMfaFactorsResponseDataAuthFactors) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserMfaFactorsResponseDataAuthFactors) SetName(v string) {
	o.Name = &v
}

func (o UserMfaFactorsResponseDataAuthFactors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FactorId != nil {
		toSerialize["factor_id"] = o.FactorId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUserMfaFactorsResponseDataAuthFactors struct {
	value *UserMfaFactorsResponseDataAuthFactors
	isSet bool
}

func (v NullableUserMfaFactorsResponseDataAuthFactors) Get() *UserMfaFactorsResponseDataAuthFactors {
	return v.value
}

func (v *NullableUserMfaFactorsResponseDataAuthFactors) Set(val *UserMfaFactorsResponseDataAuthFactors) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMfaFactorsResponseDataAuthFactors) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMfaFactorsResponseDataAuthFactors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMfaFactorsResponseDataAuthFactors(val *UserMfaFactorsResponseDataAuthFactors) *NullableUserMfaFactorsResponseDataAuthFactors {
	return &NullableUserMfaFactorsResponseDataAuthFactors{value: val, isSet: true}
}

func (v NullableUserMfaFactorsResponseDataAuthFactors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMfaFactorsResponseDataAuthFactors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
