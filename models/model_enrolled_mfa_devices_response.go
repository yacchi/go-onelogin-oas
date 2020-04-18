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

// EnrolledMfaDevicesResponse struct for EnrolledMfaDevicesResponse
type EnrolledMfaDevicesResponse struct {
	Status *Status                         `json:"status,omitempty"`
	Data   *EnrolledMfaDevicesResponseData `json:"data,omitempty"`
}

// NewEnrolledMfaDevicesResponse instantiates a new EnrolledMfaDevicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrolledMfaDevicesResponse() *EnrolledMfaDevicesResponse {
	this := EnrolledMfaDevicesResponse{}
	return &this
}

// NewEnrolledMfaDevicesResponseWithDefaults instantiates a new EnrolledMfaDevicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrolledMfaDevicesResponseWithDefaults() *EnrolledMfaDevicesResponse {
	this := EnrolledMfaDevicesResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnrolledMfaDevicesResponse) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolledMfaDevicesResponse) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnrolledMfaDevicesResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *EnrolledMfaDevicesResponse) SetStatus(v Status) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EnrolledMfaDevicesResponse) GetData() EnrolledMfaDevicesResponseData {
	if o == nil || o.Data == nil {
		var ret EnrolledMfaDevicesResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrolledMfaDevicesResponse) GetDataOk() (*EnrolledMfaDevicesResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EnrolledMfaDevicesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given EnrolledMfaDevicesResponseData and assigns it to the Data field.
func (o *EnrolledMfaDevicesResponse) SetData(v EnrolledMfaDevicesResponseData) {
	o.Data = &v
}

func (o EnrolledMfaDevicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableEnrolledMfaDevicesResponse struct {
	value *EnrolledMfaDevicesResponse
	isSet bool
}

func (v NullableEnrolledMfaDevicesResponse) Get() *EnrolledMfaDevicesResponse {
	return v.value
}

func (v *NullableEnrolledMfaDevicesResponse) Set(val *EnrolledMfaDevicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrolledMfaDevicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrolledMfaDevicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrolledMfaDevicesResponse(val *EnrolledMfaDevicesResponse) *NullableEnrolledMfaDevicesResponse {
	return &NullableEnrolledMfaDevicesResponse{value: val, isSet: true}
}

func (v NullableEnrolledMfaDevicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrolledMfaDevicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
