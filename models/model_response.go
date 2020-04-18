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

// Response struct for Response
type Response struct {
	Status       *string            `json:"status,omitempty"`
	User         *ResponseUser      `json:"user,omitempty"`
	ReturnToUrl  *string            `json:"return_to_url,omitempty"`
	ExpiresAt    *string            `json:"expires_at,omitempty"`
	SessionToken *string            `json:"session_token,omitempty"`
	StateToken   *string            `json:"state_token,omitempty"`
	CallbackUrl  *string            `json:"callback_url,omitempty"`
	Devices      *[]ResponseDevices `json:"devices,omitempty"`
}

// NewResponse instantiates a new Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponse() *Response {
	this := Response{}
	return &this
}

// NewResponseWithDefaults instantiates a new Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWithDefaults() *Response {
	this := Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Response) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Response) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Response) SetStatus(v string) {
	o.Status = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Response) GetUser() ResponseUser {
	if o == nil || o.User == nil {
		var ret ResponseUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetUserOk() (*ResponseUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Response) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given ResponseUser and assigns it to the User field.
func (o *Response) SetUser(v ResponseUser) {
	o.User = &v
}

// GetReturnToUrl returns the ReturnToUrl field value if set, zero value otherwise.
func (o *Response) GetReturnToUrl() string {
	if o == nil || o.ReturnToUrl == nil {
		var ret string
		return ret
	}
	return *o.ReturnToUrl
}

// GetReturnToUrlOk returns a tuple with the ReturnToUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetReturnToUrlOk() (*string, bool) {
	if o == nil || o.ReturnToUrl == nil {
		return nil, false
	}
	return o.ReturnToUrl, true
}

// HasReturnToUrl returns a boolean if a field has been set.
func (o *Response) HasReturnToUrl() bool {
	if o != nil && o.ReturnToUrl != nil {
		return true
	}

	return false
}

// SetReturnToUrl gets a reference to the given string and assigns it to the ReturnToUrl field.
func (o *Response) SetReturnToUrl(v string) {
	o.ReturnToUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *Response) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *Response) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *Response) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *Response) GetSessionToken() string {
	if o == nil || o.SessionToken == nil {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetSessionTokenOk() (*string, bool) {
	if o == nil || o.SessionToken == nil {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *Response) HasSessionToken() bool {
	if o != nil && o.SessionToken != nil {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *Response) SetSessionToken(v string) {
	o.SessionToken = &v
}

// GetStateToken returns the StateToken field value if set, zero value otherwise.
func (o *Response) GetStateToken() string {
	if o == nil || o.StateToken == nil {
		var ret string
		return ret
	}
	return *o.StateToken
}

// GetStateTokenOk returns a tuple with the StateToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetStateTokenOk() (*string, bool) {
	if o == nil || o.StateToken == nil {
		return nil, false
	}
	return o.StateToken, true
}

// HasStateToken returns a boolean if a field has been set.
func (o *Response) HasStateToken() bool {
	if o != nil && o.StateToken != nil {
		return true
	}

	return false
}

// SetStateToken gets a reference to the given string and assigns it to the StateToken field.
func (o *Response) SetStateToken(v string) {
	o.StateToken = &v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *Response) GetCallbackUrl() string {
	if o == nil || o.CallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetCallbackUrlOk() (*string, bool) {
	if o == nil || o.CallbackUrl == nil {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *Response) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl != nil {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *Response) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *Response) GetDevices() []ResponseDevices {
	if o == nil || o.Devices == nil {
		var ret []ResponseDevices
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetDevicesOk() (*[]ResponseDevices, bool) {
	if o == nil || o.Devices == nil {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *Response) HasDevices() bool {
	if o != nil && o.Devices != nil {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []ResponseDevices and assigns it to the Devices field.
func (o *Response) SetDevices(v []ResponseDevices) {
	o.Devices = &v
}

func (o Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.ReturnToUrl != nil {
		toSerialize["return_to_url"] = o.ReturnToUrl
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.SessionToken != nil {
		toSerialize["session_token"] = o.SessionToken
	}
	if o.StateToken != nil {
		toSerialize["state_token"] = o.StateToken
	}
	if o.CallbackUrl != nil {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	if o.Devices != nil {
		toSerialize["devices"] = o.Devices
	}
	return json.Marshal(toSerialize)
}

type NullableResponse struct {
	value *Response
	isSet bool
}

func (v NullableResponse) Get() *Response {
	return v.value
}

func (v *NullableResponse) Set(val *Response) {
	v.value = val
	v.isSet = true
}

func (v NullableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponse(val *Response) *NullableResponse {
	return &NullableResponse{value: val, isSet: true}
}

func (v NullableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
