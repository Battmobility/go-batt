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

// checks if the PublishedRealmRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublishedRealmRepresentation{}

// PublishedRealmRepresentation struct for PublishedRealmRepresentation
type PublishedRealmRepresentation struct {
	Realm *string `json:"realm,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
	TokenService *string `json:"token-service,omitempty"`
	AccountService *string `json:"account-service,omitempty"`
	TokensNotBefore *int32 `json:"tokens-not-before,omitempty"`
}

// NewPublishedRealmRepresentation instantiates a new PublishedRealmRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishedRealmRepresentation() *PublishedRealmRepresentation {
	this := PublishedRealmRepresentation{}
	return &this
}

// NewPublishedRealmRepresentationWithDefaults instantiates a new PublishedRealmRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishedRealmRepresentationWithDefaults() *PublishedRealmRepresentation {
	this := PublishedRealmRepresentation{}
	return &this
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *PublishedRealmRepresentation) GetRealm() string {
	if o == nil || IsNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishedRealmRepresentation) GetRealmOk() (*string, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *PublishedRealmRepresentation) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *PublishedRealmRepresentation) SetRealm(v string) {
	o.Realm = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *PublishedRealmRepresentation) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishedRealmRepresentation) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PublishedRealmRepresentation) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *PublishedRealmRepresentation) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetTokenService returns the TokenService field value if set, zero value otherwise.
func (o *PublishedRealmRepresentation) GetTokenService() string {
	if o == nil || IsNil(o.TokenService) {
		var ret string
		return ret
	}
	return *o.TokenService
}

// GetTokenServiceOk returns a tuple with the TokenService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishedRealmRepresentation) GetTokenServiceOk() (*string, bool) {
	if o == nil || IsNil(o.TokenService) {
		return nil, false
	}
	return o.TokenService, true
}

// HasTokenService returns a boolean if a field has been set.
func (o *PublishedRealmRepresentation) HasTokenService() bool {
	if o != nil && !IsNil(o.TokenService) {
		return true
	}

	return false
}

// SetTokenService gets a reference to the given string and assigns it to the TokenService field.
func (o *PublishedRealmRepresentation) SetTokenService(v string) {
	o.TokenService = &v
}

// GetAccountService returns the AccountService field value if set, zero value otherwise.
func (o *PublishedRealmRepresentation) GetAccountService() string {
	if o == nil || IsNil(o.AccountService) {
		var ret string
		return ret
	}
	return *o.AccountService
}

// GetAccountServiceOk returns a tuple with the AccountService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishedRealmRepresentation) GetAccountServiceOk() (*string, bool) {
	if o == nil || IsNil(o.AccountService) {
		return nil, false
	}
	return o.AccountService, true
}

// HasAccountService returns a boolean if a field has been set.
func (o *PublishedRealmRepresentation) HasAccountService() bool {
	if o != nil && !IsNil(o.AccountService) {
		return true
	}

	return false
}

// SetAccountService gets a reference to the given string and assigns it to the AccountService field.
func (o *PublishedRealmRepresentation) SetAccountService(v string) {
	o.AccountService = &v
}

// GetTokensNotBefore returns the TokensNotBefore field value if set, zero value otherwise.
func (o *PublishedRealmRepresentation) GetTokensNotBefore() int32 {
	if o == nil || IsNil(o.TokensNotBefore) {
		var ret int32
		return ret
	}
	return *o.TokensNotBefore
}

// GetTokensNotBeforeOk returns a tuple with the TokensNotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishedRealmRepresentation) GetTokensNotBeforeOk() (*int32, bool) {
	if o == nil || IsNil(o.TokensNotBefore) {
		return nil, false
	}
	return o.TokensNotBefore, true
}

// HasTokensNotBefore returns a boolean if a field has been set.
func (o *PublishedRealmRepresentation) HasTokensNotBefore() bool {
	if o != nil && !IsNil(o.TokensNotBefore) {
		return true
	}

	return false
}

// SetTokensNotBefore gets a reference to the given int32 and assigns it to the TokensNotBefore field.
func (o *PublishedRealmRepresentation) SetTokensNotBefore(v int32) {
	o.TokensNotBefore = &v
}

func (o PublishedRealmRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublishedRealmRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	if !IsNil(o.TokenService) {
		toSerialize["token-service"] = o.TokenService
	}
	if !IsNil(o.AccountService) {
		toSerialize["account-service"] = o.AccountService
	}
	if !IsNil(o.TokensNotBefore) {
		toSerialize["tokens-not-before"] = o.TokensNotBefore
	}
	return toSerialize, nil
}

type NullablePublishedRealmRepresentation struct {
	value *PublishedRealmRepresentation
	isSet bool
}

func (v NullablePublishedRealmRepresentation) Get() *PublishedRealmRepresentation {
	return v.value
}

func (v *NullablePublishedRealmRepresentation) Set(val *PublishedRealmRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishedRealmRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishedRealmRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishedRealmRepresentation(val *PublishedRealmRepresentation) *NullablePublishedRealmRepresentation {
	return &NullablePublishedRealmRepresentation{value: val, isSet: true}
}

func (v NullablePublishedRealmRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublishedRealmRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


