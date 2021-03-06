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

// InlineObject struct for InlineObject
type InlineObject struct {
	// Unapproved: 0 Approved (licensed): 1 Rejected: 2 Unlicensed: 3
	State int32 `json:"state"`
}

// NewInlineObject instantiates a new InlineObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject(state int32) *InlineObject {
	this := InlineObject{}
	this.State = state
	return &this
}

// NewInlineObjectWithDefaults instantiates a new InlineObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectWithDefaults() *InlineObject {
	this := InlineObject{}
	return &this
}

// GetState returns the State field value
func (o *InlineObject) GetState() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *InlineObject) GetStateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *InlineObject) SetState(v int32) {
	o.State = v
}

func (o InlineObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject struct {
	value *InlineObject
	isSet bool
}

func (v NullableInlineObject) Get() *InlineObject {
	return v.value
}

func (v *NullableInlineObject) Set(val *InlineObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject(val *InlineObject) *NullableInlineObject {
	return &NullableInlineObject{value: val, isSet: true}
}

func (v NullableInlineObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
