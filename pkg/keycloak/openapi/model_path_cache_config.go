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

// checks if the PathCacheConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PathCacheConfig{}

// PathCacheConfig struct for PathCacheConfig
type PathCacheConfig struct {
	MaxEntries *int32 `json:"max-entries,omitempty"`
	Lifespan *int64 `json:"lifespan,omitempty"`
}

// NewPathCacheConfig instantiates a new PathCacheConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathCacheConfig() *PathCacheConfig {
	this := PathCacheConfig{}
	return &this
}

// NewPathCacheConfigWithDefaults instantiates a new PathCacheConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathCacheConfigWithDefaults() *PathCacheConfig {
	this := PathCacheConfig{}
	return &this
}

// GetMaxEntries returns the MaxEntries field value if set, zero value otherwise.
func (o *PathCacheConfig) GetMaxEntries() int32 {
	if o == nil || IsNil(o.MaxEntries) {
		var ret int32
		return ret
	}
	return *o.MaxEntries
}

// GetMaxEntriesOk returns a tuple with the MaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathCacheConfig) GetMaxEntriesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxEntries) {
		return nil, false
	}
	return o.MaxEntries, true
}

// HasMaxEntries returns a boolean if a field has been set.
func (o *PathCacheConfig) HasMaxEntries() bool {
	if o != nil && !IsNil(o.MaxEntries) {
		return true
	}

	return false
}

// SetMaxEntries gets a reference to the given int32 and assigns it to the MaxEntries field.
func (o *PathCacheConfig) SetMaxEntries(v int32) {
	o.MaxEntries = &v
}

// GetLifespan returns the Lifespan field value if set, zero value otherwise.
func (o *PathCacheConfig) GetLifespan() int64 {
	if o == nil || IsNil(o.Lifespan) {
		var ret int64
		return ret
	}
	return *o.Lifespan
}

// GetLifespanOk returns a tuple with the Lifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathCacheConfig) GetLifespanOk() (*int64, bool) {
	if o == nil || IsNil(o.Lifespan) {
		return nil, false
	}
	return o.Lifespan, true
}

// HasLifespan returns a boolean if a field has been set.
func (o *PathCacheConfig) HasLifespan() bool {
	if o != nil && !IsNil(o.Lifespan) {
		return true
	}

	return false
}

// SetLifespan gets a reference to the given int64 and assigns it to the Lifespan field.
func (o *PathCacheConfig) SetLifespan(v int64) {
	o.Lifespan = &v
}

func (o PathCacheConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PathCacheConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxEntries) {
		toSerialize["max-entries"] = o.MaxEntries
	}
	if !IsNil(o.Lifespan) {
		toSerialize["lifespan"] = o.Lifespan
	}
	return toSerialize, nil
}

type NullablePathCacheConfig struct {
	value *PathCacheConfig
	isSet bool
}

func (v NullablePathCacheConfig) Get() *PathCacheConfig {
	return v.value
}

func (v *NullablePathCacheConfig) Set(val *PathCacheConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePathCacheConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePathCacheConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathCacheConfig(val *PathCacheConfig) *NullablePathCacheConfig {
	return &NullablePathCacheConfig{value: val, isSet: true}
}

func (v NullablePathCacheConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathCacheConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


