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

// ResponseDevices struct for ResponseDevices
type ResponseDevices struct {
	DeviceId   *int32  `json:"device_id,omitempty"`
	DeviceType *string `json:"device_type,omitempty"`
}

// NewResponseDevices instantiates a new ResponseDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDevices() *ResponseDevices {
	this := ResponseDevices{}
	return &this
}

// NewResponseDevicesWithDefaults instantiates a new ResponseDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDevicesWithDefaults() *ResponseDevices {
	this := ResponseDevices{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ResponseDevices) GetDeviceId() int32 {
	if o == nil || o.DeviceId == nil {
		var ret int32
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDevices) GetDeviceIdOk() (*int32, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ResponseDevices) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int32 and assigns it to the DeviceId field.
func (o *ResponseDevices) SetDeviceId(v int32) {
	o.DeviceId = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *ResponseDevices) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDevices) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *ResponseDevices) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *ResponseDevices) SetDeviceType(v string) {
	o.DeviceType = &v
}

func (o ResponseDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	if o.DeviceType != nil {
		toSerialize["device_type"] = o.DeviceType
	}
	return json.Marshal(toSerialize)
}

type NullableResponseDevices struct {
	value *ResponseDevices
	isSet bool
}

func (v NullableResponseDevices) Get() *ResponseDevices {
	return v.value
}

func (v *NullableResponseDevices) Set(val *ResponseDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDevices(val *ResponseDevices) *NullableResponseDevices {
	return &NullableResponseDevices{value: val, isSet: true}
}

func (v NullableResponseDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
