/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// UnmanagedAttributePolicy the model 'UnmanagedAttributePolicy'
type UnmanagedAttributePolicy string

// List of UnmanagedAttributePolicy
const (
	ENABLED UnmanagedAttributePolicy = "ENABLED"
	ADMIN_VIEW UnmanagedAttributePolicy = "ADMIN_VIEW"
	ADMIN_EDIT UnmanagedAttributePolicy = "ADMIN_EDIT"
)

// All allowed values of UnmanagedAttributePolicy enum
var AllowedUnmanagedAttributePolicyEnumValues = []UnmanagedAttributePolicy{
	"ENABLED",
	"ADMIN_VIEW",
	"ADMIN_EDIT",
}

func (v *UnmanagedAttributePolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UnmanagedAttributePolicy(value)
	for _, existing := range AllowedUnmanagedAttributePolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UnmanagedAttributePolicy", value)
}

// NewUnmanagedAttributePolicyFromValue returns a pointer to a valid UnmanagedAttributePolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUnmanagedAttributePolicyFromValue(v string) (*UnmanagedAttributePolicy, error) {
	ev := UnmanagedAttributePolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UnmanagedAttributePolicy: valid values are %v", v, AllowedUnmanagedAttributePolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UnmanagedAttributePolicy) IsValid() bool {
	for _, existing := range AllowedUnmanagedAttributePolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnmanagedAttributePolicy value
func (v UnmanagedAttributePolicy) Ptr() *UnmanagedAttributePolicy {
	return &v
}

type NullableUnmanagedAttributePolicy struct {
	value *UnmanagedAttributePolicy
	isSet bool
}

func (v NullableUnmanagedAttributePolicy) Get() *UnmanagedAttributePolicy {
	return v.value
}

func (v *NullableUnmanagedAttributePolicy) Set(val *UnmanagedAttributePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableUnmanagedAttributePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableUnmanagedAttributePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnmanagedAttributePolicy(val *UnmanagedAttributePolicy) *NullableUnmanagedAttributePolicy {
	return &NullableUnmanagedAttributePolicy{value: val, isSet: true}
}

func (v NullableUnmanagedAttributePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnmanagedAttributePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

