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

// checks if the UPAttributePermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UPAttributePermissions{}

// UPAttributePermissions struct for UPAttributePermissions
type UPAttributePermissions struct {
	View []string `json:"view,omitempty"`
	Edit []string `json:"edit,omitempty"`
}

// NewUPAttributePermissions instantiates a new UPAttributePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUPAttributePermissions() *UPAttributePermissions {
	this := UPAttributePermissions{}
	return &this
}

// NewUPAttributePermissionsWithDefaults instantiates a new UPAttributePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUPAttributePermissionsWithDefaults() *UPAttributePermissions {
	this := UPAttributePermissions{}
	return &this
}

// GetView returns the View field value if set, zero value otherwise.
func (o *UPAttributePermissions) GetView() []string {
	if o == nil || IsNil(o.View) {
		var ret []string
		return ret
	}
	return o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UPAttributePermissions) GetViewOk() ([]string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *UPAttributePermissions) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given []string and assigns it to the View field.
func (o *UPAttributePermissions) SetView(v []string) {
	o.View = v
}

// GetEdit returns the Edit field value if set, zero value otherwise.
func (o *UPAttributePermissions) GetEdit() []string {
	if o == nil || IsNil(o.Edit) {
		var ret []string
		return ret
	}
	return o.Edit
}

// GetEditOk returns a tuple with the Edit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UPAttributePermissions) GetEditOk() ([]string, bool) {
	if o == nil || IsNil(o.Edit) {
		return nil, false
	}
	return o.Edit, true
}

// HasEdit returns a boolean if a field has been set.
func (o *UPAttributePermissions) HasEdit() bool {
	if o != nil && !IsNil(o.Edit) {
		return true
	}

	return false
}

// SetEdit gets a reference to the given []string and assigns it to the Edit field.
func (o *UPAttributePermissions) SetEdit(v []string) {
	o.Edit = v
}

func (o UPAttributePermissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UPAttributePermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	if !IsNil(o.Edit) {
		toSerialize["edit"] = o.Edit
	}
	return toSerialize, nil
}

type NullableUPAttributePermissions struct {
	value *UPAttributePermissions
	isSet bool
}

func (v NullableUPAttributePermissions) Get() *UPAttributePermissions {
	return v.value
}

func (v *NullableUPAttributePermissions) Set(val *UPAttributePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableUPAttributePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableUPAttributePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUPAttributePermissions(val *UPAttributePermissions) *NullableUPAttributePermissions {
	return &NullableUPAttributePermissions{value: val, isSet: true}
}

func (v NullableUPAttributePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUPAttributePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


