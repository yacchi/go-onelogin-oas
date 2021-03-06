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

// CustomAttributesResponse struct for CustomAttributesResponse
type CustomAttributesResponse struct {
	Status *Status    `json:"status,omitempty"`
	Data   [][]string `json:"data"`
}

// NewCustomAttributesResponse instantiates a new CustomAttributesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAttributesResponse(data [][]string) *CustomAttributesResponse {
	this := CustomAttributesResponse{}
	this.Data = data
	return &this
}

// NewCustomAttributesResponseWithDefaults instantiates a new CustomAttributesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAttributesResponseWithDefaults() *CustomAttributesResponse {
	this := CustomAttributesResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomAttributesResponse) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributesResponse) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomAttributesResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *CustomAttributesResponse) SetStatus(v Status) {
	o.Status = &v
}

// GetData returns the Data field value
func (o *CustomAttributesResponse) GetData() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomAttributesResponse) GetDataOk() (*[][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomAttributesResponse) SetData(v [][]string) {
	o.Data = v
}

func (o CustomAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomAttributesResponse struct {
	value *CustomAttributesResponse
	isSet bool
}

func (v NullableCustomAttributesResponse) Get() *CustomAttributesResponse {
	return v.value
}

func (v *NullableCustomAttributesResponse) Set(val *CustomAttributesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAttributesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAttributesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAttributesResponse(val *CustomAttributesResponse) *NullableCustomAttributesResponse {
	return &NullableCustomAttributesResponse{value: val, isSet: true}
}

func (v NullableCustomAttributesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAttributesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
