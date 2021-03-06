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

// AppSso struct for AppSso
type AppSso struct {
	ClientId     *string            `json:"client_id,omitempty"`
	ClientSecret *string            `json:"client_secret,omitempty"`
	MetadataUrl  *string            `json:"metadata_url,omitempty"`
	AcsUrl       *string            `json:"acs_url,omitempty"`
	SlsUrl       *string            `json:"sls_url,omitempty"`
	Issuer       *string            `json:"issuer,omitempty"`
	Certificate  *AppSsoCertificate `json:"certificate,omitempty"`
}

// NewAppSso instantiates a new AppSso object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSso() *AppSso {
	this := AppSso{}
	return &this
}

// NewAppSsoWithDefaults instantiates a new AppSso object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSsoWithDefaults() *AppSso {
	this := AppSso{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AppSso) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSso) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AppSso) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AppSso) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AppSso) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSso) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AppSso) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AppSso) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetMetadataUrl returns the MetadataUrl field value if set, zero value otherwise.
func (o *AppSso) GetMetadataUrl() string {
	if o == nil || o.MetadataUrl == nil {
		var ret string
		return ret
	}
	return *o.MetadataUrl
}

// GetMetadataUrlOk returns a tuple with the MetadataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSso) GetMetadataUrlOk() (*string, bool) {
	if o == nil || o.MetadataUrl == nil {
		return nil, false
	}
	return o.MetadataUrl, true
}

// HasMetadataUrl returns a boolean if a field has been set.
func (o *AppSso) HasMetadataUrl() bool {
	if o != nil && o.MetadataUrl != nil {
		return true
	}

	return false
}

// SetMetadataUrl gets a reference to the given string and assigns it to the MetadataUrl field.
func (o *AppSso) SetMetadataUrl(v string) {
	o.MetadataUrl = &v
}

// GetAcsUrl returns the AcsUrl field value if set, zero value otherwise.
func (o *AppSso) GetAcsUrl() string {
	if o == nil || o.AcsUrl == nil {
		var ret string
		return ret
	}
	return *o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSso) GetAcsUrlOk() (*string, bool) {
	if o == nil || o.AcsUrl == nil {
		return nil, false
	}
	return o.AcsUrl, true
}

// HasAcsUrl returns a boolean if a field has been set.
func (o *AppSso) HasAcsUrl() bool {
	if o != nil && o.AcsUrl != nil {
		return true
	}

	return false
}

// SetAcsUrl gets a reference to the given string and assigns it to the AcsUrl field.
func (o *AppSso) SetAcsUrl(v string) {
	o.AcsUrl = &v
}

// GetSlsUrl returns the SlsUrl field value if set, zero value otherwise.
func (o *AppSso) GetSlsUrl() string {
	if o == nil || o.SlsUrl == nil {
		var ret string
		return ret
	}
	return *o.SlsUrl
}

// GetSlsUrlOk returns a tuple with the SlsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSso) GetSlsUrlOk() (*string, bool) {
	if o == nil || o.SlsUrl == nil {
		return nil, false
	}
	return o.SlsUrl, true
}

// HasSlsUrl returns a boolean if a field has been set.
func (o *AppSso) HasSlsUrl() bool {
	if o != nil && o.SlsUrl != nil {
		return true
	}

	return false
}

// SetSlsUrl gets a reference to the given string and assigns it to the SlsUrl field.
func (o *AppSso) SetSlsUrl(v string) {
	o.SlsUrl = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *AppSso) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSso) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *AppSso) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *AppSso) SetIssuer(v string) {
	o.Issuer = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *AppSso) GetCertificate() AppSsoCertificate {
	if o == nil || o.Certificate == nil {
		var ret AppSsoCertificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSso) GetCertificateOk() (*AppSsoCertificate, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *AppSso) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given AppSsoCertificate and assigns it to the Certificate field.
func (o *AppSso) SetCertificate(v AppSsoCertificate) {
	o.Certificate = &v
}

func (o AppSso) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.MetadataUrl != nil {
		toSerialize["metadata_url"] = o.MetadataUrl
	}
	if o.AcsUrl != nil {
		toSerialize["acs_url"] = o.AcsUrl
	}
	if o.SlsUrl != nil {
		toSerialize["sls_url"] = o.SlsUrl
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	return json.Marshal(toSerialize)
}

type NullableAppSso struct {
	value *AppSso
	isSet bool
}

func (v NullableAppSso) Get() *AppSso {
	return v.value
}

func (v *NullableAppSso) Set(val *AppSso) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSso) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSso) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSso(val *AppSso) *NullableAppSso {
	return &NullableAppSso{value: val, isSet: true}
}

func (v NullableAppSso) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSso) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
