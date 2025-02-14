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

// checks if the RolesRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolesRepresentation{}

// RolesRepresentation struct for RolesRepresentation
type RolesRepresentation struct {
	Realm []RoleRepresentation `json:"realm,omitempty"`
	Client *map[string][]RoleRepresentation `json:"client,omitempty"`
	// Deprecated
	Application *map[string][]RoleRepresentation `json:"application,omitempty"`
}

// NewRolesRepresentation instantiates a new RolesRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolesRepresentation() *RolesRepresentation {
	this := RolesRepresentation{}
	return &this
}

// NewRolesRepresentationWithDefaults instantiates a new RolesRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolesRepresentationWithDefaults() *RolesRepresentation {
	this := RolesRepresentation{}
	return &this
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *RolesRepresentation) GetRealm() []RoleRepresentation {
	if o == nil || IsNil(o.Realm) {
		var ret []RoleRepresentation
		return ret
	}
	return o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesRepresentation) GetRealmOk() ([]RoleRepresentation, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *RolesRepresentation) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given []RoleRepresentation and assigns it to the Realm field.
func (o *RolesRepresentation) SetRealm(v []RoleRepresentation) {
	o.Realm = v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *RolesRepresentation) GetClient() map[string][]RoleRepresentation {
	if o == nil || IsNil(o.Client) {
		var ret map[string][]RoleRepresentation
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesRepresentation) GetClientOk() (*map[string][]RoleRepresentation, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *RolesRepresentation) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given map[string][]RoleRepresentation and assigns it to the Client field.
func (o *RolesRepresentation) SetClient(v map[string][]RoleRepresentation) {
	o.Client = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
// Deprecated
func (o *RolesRepresentation) GetApplication() map[string][]RoleRepresentation {
	if o == nil || IsNil(o.Application) {
		var ret map[string][]RoleRepresentation
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *RolesRepresentation) GetApplicationOk() (*map[string][]RoleRepresentation, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *RolesRepresentation) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given map[string][]RoleRepresentation and assigns it to the Application field.
// Deprecated
func (o *RolesRepresentation) SetApplication(v map[string][]RoleRepresentation) {
	o.Application = &v
}

func (o RolesRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolesRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	return toSerialize, nil
}

type NullableRolesRepresentation struct {
	value *RolesRepresentation
	isSet bool
}

func (v NullableRolesRepresentation) Get() *RolesRepresentation {
	return v.value
}

func (v *NullableRolesRepresentation) Set(val *RolesRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesRepresentation(val *RolesRepresentation) *NullableRolesRepresentation {
	return &NullableRolesRepresentation{value: val, isSet: true}
}

func (v NullableRolesRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


