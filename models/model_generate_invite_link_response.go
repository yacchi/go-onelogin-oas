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

// GenerateInviteLinkResponse struct for GenerateInviteLinkResponse
type GenerateInviteLinkResponse struct {
	Status *Status   `json:"status,omitempty"`
	Data   *[]string `json:"data,omitempty"`
}

// NewGenerateInviteLinkResponse instantiates a new GenerateInviteLinkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateInviteLinkResponse() *GenerateInviteLinkResponse {
	this := GenerateInviteLinkResponse{}
	return &this
}

// NewGenerateInviteLinkResponseWithDefaults instantiates a new GenerateInviteLinkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateInviteLinkResponseWithDefaults() *GenerateInviteLinkResponse {
	this := GenerateInviteLinkResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GenerateInviteLinkResponse) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateInviteLinkResponse) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GenerateInviteLinkResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *GenerateInviteLinkResponse) SetStatus(v Status) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GenerateInviteLinkResponse) GetData() []string {
	if o == nil || o.Data == nil {
		var ret []string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateInviteLinkResponse) GetDataOk() (*[]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GenerateInviteLinkResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *GenerateInviteLinkResponse) SetData(v []string) {
	o.Data = &v
}

func (o GenerateInviteLinkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateInviteLinkResponse struct {
	value *GenerateInviteLinkResponse
	isSet bool
}

func (v NullableGenerateInviteLinkResponse) Get() *GenerateInviteLinkResponse {
	return v.value
}

func (v *NullableGenerateInviteLinkResponse) Set(val *GenerateInviteLinkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateInviteLinkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateInviteLinkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateInviteLinkResponse(val *GenerateInviteLinkResponse) *NullableGenerateInviteLinkResponse {
	return &NullableGenerateInviteLinkResponse{value: val, isSet: true}
}

func (v NullableGenerateInviteLinkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateInviteLinkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
