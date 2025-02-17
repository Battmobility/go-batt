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

// checks if the RealmEventsConfigRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmEventsConfigRepresentation{}

// RealmEventsConfigRepresentation struct for RealmEventsConfigRepresentation
type RealmEventsConfigRepresentation struct {
	EventsEnabled *bool `json:"eventsEnabled,omitempty"`
	EventsExpiration *int64 `json:"eventsExpiration,omitempty"`
	EventsListeners []string `json:"eventsListeners,omitempty"`
	EnabledEventTypes []string `json:"enabledEventTypes,omitempty"`
	AdminEventsEnabled *bool `json:"adminEventsEnabled,omitempty"`
	AdminEventsDetailsEnabled *bool `json:"adminEventsDetailsEnabled,omitempty"`
}

// NewRealmEventsConfigRepresentation instantiates a new RealmEventsConfigRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmEventsConfigRepresentation() *RealmEventsConfigRepresentation {
	this := RealmEventsConfigRepresentation{}
	return &this
}

// NewRealmEventsConfigRepresentationWithDefaults instantiates a new RealmEventsConfigRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmEventsConfigRepresentationWithDefaults() *RealmEventsConfigRepresentation {
	this := RealmEventsConfigRepresentation{}
	return &this
}

// GetEventsEnabled returns the EventsEnabled field value if set, zero value otherwise.
func (o *RealmEventsConfigRepresentation) GetEventsEnabled() bool {
	if o == nil || IsNil(o.EventsEnabled) {
		var ret bool
		return ret
	}
	return *o.EventsEnabled
}

// GetEventsEnabledOk returns a tuple with the EventsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmEventsConfigRepresentation) GetEventsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EventsEnabled) {
		return nil, false
	}
	return o.EventsEnabled, true
}

// HasEventsEnabled returns a boolean if a field has been set.
func (o *RealmEventsConfigRepresentation) HasEventsEnabled() bool {
	if o != nil && !IsNil(o.EventsEnabled) {
		return true
	}

	return false
}

// SetEventsEnabled gets a reference to the given bool and assigns it to the EventsEnabled field.
func (o *RealmEventsConfigRepresentation) SetEventsEnabled(v bool) {
	o.EventsEnabled = &v
}

// GetEventsExpiration returns the EventsExpiration field value if set, zero value otherwise.
func (o *RealmEventsConfigRepresentation) GetEventsExpiration() int64 {
	if o == nil || IsNil(o.EventsExpiration) {
		var ret int64
		return ret
	}
	return *o.EventsExpiration
}

// GetEventsExpirationOk returns a tuple with the EventsExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmEventsConfigRepresentation) GetEventsExpirationOk() (*int64, bool) {
	if o == nil || IsNil(o.EventsExpiration) {
		return nil, false
	}
	return o.EventsExpiration, true
}

// HasEventsExpiration returns a boolean if a field has been set.
func (o *RealmEventsConfigRepresentation) HasEventsExpiration() bool {
	if o != nil && !IsNil(o.EventsExpiration) {
		return true
	}

	return false
}

// SetEventsExpiration gets a reference to the given int64 and assigns it to the EventsExpiration field.
func (o *RealmEventsConfigRepresentation) SetEventsExpiration(v int64) {
	o.EventsExpiration = &v
}

// GetEventsListeners returns the EventsListeners field value if set, zero value otherwise.
func (o *RealmEventsConfigRepresentation) GetEventsListeners() []string {
	if o == nil || IsNil(o.EventsListeners) {
		var ret []string
		return ret
	}
	return o.EventsListeners
}

// GetEventsListenersOk returns a tuple with the EventsListeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmEventsConfigRepresentation) GetEventsListenersOk() ([]string, bool) {
	if o == nil || IsNil(o.EventsListeners) {
		return nil, false
	}
	return o.EventsListeners, true
}

// HasEventsListeners returns a boolean if a field has been set.
func (o *RealmEventsConfigRepresentation) HasEventsListeners() bool {
	if o != nil && !IsNil(o.EventsListeners) {
		return true
	}

	return false
}

// SetEventsListeners gets a reference to the given []string and assigns it to the EventsListeners field.
func (o *RealmEventsConfigRepresentation) SetEventsListeners(v []string) {
	o.EventsListeners = v
}

// GetEnabledEventTypes returns the EnabledEventTypes field value if set, zero value otherwise.
func (o *RealmEventsConfigRepresentation) GetEnabledEventTypes() []string {
	if o == nil || IsNil(o.EnabledEventTypes) {
		var ret []string
		return ret
	}
	return o.EnabledEventTypes
}

// GetEnabledEventTypesOk returns a tuple with the EnabledEventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmEventsConfigRepresentation) GetEnabledEventTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.EnabledEventTypes) {
		return nil, false
	}
	return o.EnabledEventTypes, true
}

// HasEnabledEventTypes returns a boolean if a field has been set.
func (o *RealmEventsConfigRepresentation) HasEnabledEventTypes() bool {
	if o != nil && !IsNil(o.EnabledEventTypes) {
		return true
	}

	return false
}

// SetEnabledEventTypes gets a reference to the given []string and assigns it to the EnabledEventTypes field.
func (o *RealmEventsConfigRepresentation) SetEnabledEventTypes(v []string) {
	o.EnabledEventTypes = v
}

// GetAdminEventsEnabled returns the AdminEventsEnabled field value if set, zero value otherwise.
func (o *RealmEventsConfigRepresentation) GetAdminEventsEnabled() bool {
	if o == nil || IsNil(o.AdminEventsEnabled) {
		var ret bool
		return ret
	}
	return *o.AdminEventsEnabled
}

// GetAdminEventsEnabledOk returns a tuple with the AdminEventsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmEventsConfigRepresentation) GetAdminEventsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AdminEventsEnabled) {
		return nil, false
	}
	return o.AdminEventsEnabled, true
}

// HasAdminEventsEnabled returns a boolean if a field has been set.
func (o *RealmEventsConfigRepresentation) HasAdminEventsEnabled() bool {
	if o != nil && !IsNil(o.AdminEventsEnabled) {
		return true
	}

	return false
}

// SetAdminEventsEnabled gets a reference to the given bool and assigns it to the AdminEventsEnabled field.
func (o *RealmEventsConfigRepresentation) SetAdminEventsEnabled(v bool) {
	o.AdminEventsEnabled = &v
}

// GetAdminEventsDetailsEnabled returns the AdminEventsDetailsEnabled field value if set, zero value otherwise.
func (o *RealmEventsConfigRepresentation) GetAdminEventsDetailsEnabled() bool {
	if o == nil || IsNil(o.AdminEventsDetailsEnabled) {
		var ret bool
		return ret
	}
	return *o.AdminEventsDetailsEnabled
}

// GetAdminEventsDetailsEnabledOk returns a tuple with the AdminEventsDetailsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmEventsConfigRepresentation) GetAdminEventsDetailsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AdminEventsDetailsEnabled) {
		return nil, false
	}
	return o.AdminEventsDetailsEnabled, true
}

// HasAdminEventsDetailsEnabled returns a boolean if a field has been set.
func (o *RealmEventsConfigRepresentation) HasAdminEventsDetailsEnabled() bool {
	if o != nil && !IsNil(o.AdminEventsDetailsEnabled) {
		return true
	}

	return false
}

// SetAdminEventsDetailsEnabled gets a reference to the given bool and assigns it to the AdminEventsDetailsEnabled field.
func (o *RealmEventsConfigRepresentation) SetAdminEventsDetailsEnabled(v bool) {
	o.AdminEventsDetailsEnabled = &v
}

func (o RealmEventsConfigRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmEventsConfigRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventsEnabled) {
		toSerialize["eventsEnabled"] = o.EventsEnabled
	}
	if !IsNil(o.EventsExpiration) {
		toSerialize["eventsExpiration"] = o.EventsExpiration
	}
	if !IsNil(o.EventsListeners) {
		toSerialize["eventsListeners"] = o.EventsListeners
	}
	if !IsNil(o.EnabledEventTypes) {
		toSerialize["enabledEventTypes"] = o.EnabledEventTypes
	}
	if !IsNil(o.AdminEventsEnabled) {
		toSerialize["adminEventsEnabled"] = o.AdminEventsEnabled
	}
	if !IsNil(o.AdminEventsDetailsEnabled) {
		toSerialize["adminEventsDetailsEnabled"] = o.AdminEventsDetailsEnabled
	}
	return toSerialize, nil
}

type NullableRealmEventsConfigRepresentation struct {
	value *RealmEventsConfigRepresentation
	isSet bool
}

func (v NullableRealmEventsConfigRepresentation) Get() *RealmEventsConfigRepresentation {
	return v.value
}

func (v *NullableRealmEventsConfigRepresentation) Set(val *RealmEventsConfigRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmEventsConfigRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmEventsConfigRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmEventsConfigRepresentation(val *RealmEventsConfigRepresentation) *NullableRealmEventsConfigRepresentation {
	return &NullableRealmEventsConfigRepresentation{value: val, isSet: true}
}

func (v NullableRealmEventsConfigRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmEventsConfigRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


