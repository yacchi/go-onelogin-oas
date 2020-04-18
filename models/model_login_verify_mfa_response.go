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
	"github.com/yacchi/go-onelogin-oas/models"
)

// LoginVerifyMfaResponse struct for LoginVerifyMfaResponse
type LoginVerifyMfaResponse struct {
	Status *Status `json:"status,omitempty"`
	Data *LoginVerifyMfaResponseData `json:"data,omitempty"`
}

// NewLoginVerifyMfaResponse instantiates a new LoginVerifyMfaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginVerifyMfaResponse() *LoginVerifyMfaResponse {
	this := LoginVerifyMfaResponse{}
	return &this
}

// NewLoginVerifyMfaResponseWithDefaults instantiates a new LoginVerifyMfaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginVerifyMfaResponseWithDefaults() *LoginVerifyMfaResponse {
	this := LoginVerifyMfaResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LoginVerifyMfaResponse) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginVerifyMfaResponse) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LoginVerifyMfaResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *LoginVerifyMfaResponse) SetStatus(v Status) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LoginVerifyMfaResponse) GetData() LoginVerifyMfaResponseData {
	if o == nil || o.Data == nil {
		var ret LoginVerifyMfaResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginVerifyMfaResponse) GetDataOk() (*LoginVerifyMfaResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LoginVerifyMfaResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given LoginVerifyMfaResponseData and assigns it to the Data field.
func (o *LoginVerifyMfaResponse) SetData(v LoginVerifyMfaResponseData) {
	o.Data = &v
}

func (o LoginVerifyMfaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLoginVerifyMfaResponse struct {
	value *LoginVerifyMfaResponse
	isSet bool
}

func (v NullableLoginVerifyMfaResponse) Get() *LoginVerifyMfaResponse {
	return v.value
}

func (v *NullableLoginVerifyMfaResponse) Set(val *LoginVerifyMfaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginVerifyMfaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginVerifyMfaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginVerifyMfaResponse(val *LoginVerifyMfaResponse) *NullableLoginVerifyMfaResponse {
	return &NullableLoginVerifyMfaResponse{value: val, isSet: true}
}

func (v NullableLoginVerifyMfaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginVerifyMfaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
