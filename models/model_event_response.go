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

// EventResponse struct for EventResponse
type EventResponse struct {
	Status *Status `json:"status,omitempty"`
	Data *Event `json:"data,omitempty"`
}

// NewEventResponse instantiates a new EventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponse() *EventResponse {
	this := EventResponse{}
	return &this
}

// NewEventResponseWithDefaults instantiates a new EventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseWithDefaults() *EventResponse {
	this := EventResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EventResponse) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *EventResponse) SetStatus(v Status) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EventResponse) GetData() Event {
	if o == nil || o.Data == nil {
		var ret Event
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetDataOk() (*Event, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EventResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Event and assigns it to the Data field.
func (o *EventResponse) SetData(v Event) {
	o.Data = &v
}

func (o EventResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableEventResponse struct {
	value *EventResponse
	isSet bool
}

func (v NullableEventResponse) Get() *EventResponse {
	return v.value
}

func (v *NullableEventResponse) Set(val *EventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponse(val *EventResponse) *NullableEventResponse {
	return &NullableEventResponse{value: val, isSet: true}
}

func (v NullableEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
