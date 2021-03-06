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

// EventTypesResponse struct for EventTypesResponse
type EventTypesResponse struct {
	Status *Status      `json:"status,omitempty"`
	Data   *[]EventType `json:"data,omitempty"`
}

// NewEventTypesResponse instantiates a new EventTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypesResponse() *EventTypesResponse {
	this := EventTypesResponse{}
	return &this
}

// NewEventTypesResponseWithDefaults instantiates a new EventTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypesResponseWithDefaults() *EventTypesResponse {
	this := EventTypesResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EventTypesResponse) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypesResponse) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventTypesResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *EventTypesResponse) SetStatus(v Status) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EventTypesResponse) GetData() []EventType {
	if o == nil || o.Data == nil {
		var ret []EventType
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypesResponse) GetDataOk() (*[]EventType, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EventTypesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []EventType and assigns it to the Data field.
func (o *EventTypesResponse) SetData(v []EventType) {
	o.Data = &v
}

func (o EventTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableEventTypesResponse struct {
	value *EventTypesResponse
	isSet bool
}

func (v NullableEventTypesResponse) Get() *EventTypesResponse {
	return v.value
}

func (v *NullableEventTypesResponse) Set(val *EventTypesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypesResponse(val *EventTypesResponse) *NullableEventTypesResponse {
	return &NullableEventTypesResponse{value: val, isSet: true}
}

func (v NullableEventTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
