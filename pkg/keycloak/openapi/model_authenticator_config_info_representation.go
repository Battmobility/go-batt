/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AuthenticatorConfigInfoRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorConfigInfoRepresentation{}

// AuthenticatorConfigInfoRepresentation struct for AuthenticatorConfigInfoRepresentation
type AuthenticatorConfigInfoRepresentation struct {
	Name *string `json:"name,omitempty"`
	ProviderId *string `json:"providerId,omitempty"`
	HelpText *string `json:"helpText,omitempty"`
	Properties []ConfigPropertyRepresentation `json:"properties,omitempty"`
}

// NewAuthenticatorConfigInfoRepresentation instantiates a new AuthenticatorConfigInfoRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorConfigInfoRepresentation() *AuthenticatorConfigInfoRepresentation {
	this := AuthenticatorConfigInfoRepresentation{}
	return &this
}

// NewAuthenticatorConfigInfoRepresentationWithDefaults instantiates a new AuthenticatorConfigInfoRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorConfigInfoRepresentationWithDefaults() *AuthenticatorConfigInfoRepresentation {
	this := AuthenticatorConfigInfoRepresentation{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthenticatorConfigInfoRepresentation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorConfigInfoRepresentation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthenticatorConfigInfoRepresentation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthenticatorConfigInfoRepresentation) SetName(v string) {
	o.Name = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *AuthenticatorConfigInfoRepresentation) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId) {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorConfigInfoRepresentation) GetProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *AuthenticatorConfigInfoRepresentation) HasProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *AuthenticatorConfigInfoRepresentation) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetHelpText returns the HelpText field value if set, zero value otherwise.
func (o *AuthenticatorConfigInfoRepresentation) GetHelpText() string {
	if o == nil || IsNil(o.HelpText) {
		var ret string
		return ret
	}
	return *o.HelpText
}

// GetHelpTextOk returns a tuple with the HelpText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorConfigInfoRepresentation) GetHelpTextOk() (*string, bool) {
	if o == nil || IsNil(o.HelpText) {
		return nil, false
	}
	return o.HelpText, true
}

// HasHelpText returns a boolean if a field has been set.
func (o *AuthenticatorConfigInfoRepresentation) HasHelpText() bool {
	if o != nil && !IsNil(o.HelpText) {
		return true
	}

	return false
}

// SetHelpText gets a reference to the given string and assigns it to the HelpText field.
func (o *AuthenticatorConfigInfoRepresentation) SetHelpText(v string) {
	o.HelpText = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AuthenticatorConfigInfoRepresentation) GetProperties() []ConfigPropertyRepresentation {
	if o == nil || IsNil(o.Properties) {
		var ret []ConfigPropertyRepresentation
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorConfigInfoRepresentation) GetPropertiesOk() ([]ConfigPropertyRepresentation, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AuthenticatorConfigInfoRepresentation) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []ConfigPropertyRepresentation and assigns it to the Properties field.
func (o *AuthenticatorConfigInfoRepresentation) SetProperties(v []ConfigPropertyRepresentation) {
	o.Properties = v
}

func (o AuthenticatorConfigInfoRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorConfigInfoRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	if !IsNil(o.HelpText) {
		toSerialize["helpText"] = o.HelpText
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableAuthenticatorConfigInfoRepresentation struct {
	value *AuthenticatorConfigInfoRepresentation
	isSet bool
}

func (v NullableAuthenticatorConfigInfoRepresentation) Get() *AuthenticatorConfigInfoRepresentation {
	return v.value
}

func (v *NullableAuthenticatorConfigInfoRepresentation) Set(val *AuthenticatorConfigInfoRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorConfigInfoRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorConfigInfoRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorConfigInfoRepresentation(val *AuthenticatorConfigInfoRepresentation) *NullableAuthenticatorConfigInfoRepresentation {
	return &NullableAuthenticatorConfigInfoRepresentation{value: val, isSet: true}
}

func (v NullableAuthenticatorConfigInfoRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorConfigInfoRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


