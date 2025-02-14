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

// checks if the IDToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IDToken{}

// IDToken struct for IDToken
type IDToken struct {
	Jti *string `json:"jti,omitempty"`
	Exp *int64 `json:"exp,omitempty"`
	Nbf *int64 `json:"nbf,omitempty"`
	Iat *int64 `json:"iat,omitempty"`
	Iss *string `json:"iss,omitempty"`
	Sub *string `json:"sub,omitempty"`
	Typ *string `json:"typ,omitempty"`
	Azp *string `json:"azp,omitempty"`
	OtherClaims map[string]interface{} `json:"otherClaims,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	AuthTime *int64 `json:"auth_time,omitempty"`
	Sid *string `json:"sid,omitempty"`
	AtHash *string `json:"at_hash,omitempty"`
	CHash *string `json:"c_hash,omitempty"`
	Name *string `json:"name,omitempty"`
	GivenName *string `json:"given_name,omitempty"`
	FamilyName *string `json:"family_name,omitempty"`
	MiddleName *string `json:"middle_name,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	PreferredUsername *string `json:"preferred_username,omitempty"`
	Profile *string `json:"profile,omitempty"`
	Picture *string `json:"picture,omitempty"`
	Website *string `json:"website,omitempty"`
	Email *string `json:"email,omitempty"`
	EmailVerified *bool `json:"email_verified,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Birthdate *string `json:"birthdate,omitempty"`
	Zoneinfo *string `json:"zoneinfo,omitempty"`
	Locale *string `json:"locale,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	PhoneNumberVerified *bool `json:"phone_number_verified,omitempty"`
	Address *AddressClaimSet `json:"address,omitempty"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	ClaimsLocales *string `json:"claims_locales,omitempty"`
	Acr *string `json:"acr,omitempty"`
	SHash *string `json:"s_hash,omitempty"`
}

// NewIDToken instantiates a new IDToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIDToken() *IDToken {
	this := IDToken{}
	return &this
}

// NewIDTokenWithDefaults instantiates a new IDToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIDTokenWithDefaults() *IDToken {
	this := IDToken{}
	return &this
}

// GetJti returns the Jti field value if set, zero value otherwise.
func (o *IDToken) GetJti() string {
	if o == nil || IsNil(o.Jti) {
		var ret string
		return ret
	}
	return *o.Jti
}

// GetJtiOk returns a tuple with the Jti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetJtiOk() (*string, bool) {
	if o == nil || IsNil(o.Jti) {
		return nil, false
	}
	return o.Jti, true
}

// HasJti returns a boolean if a field has been set.
func (o *IDToken) HasJti() bool {
	if o != nil && !IsNil(o.Jti) {
		return true
	}

	return false
}

// SetJti gets a reference to the given string and assigns it to the Jti field.
func (o *IDToken) SetJti(v string) {
	o.Jti = &v
}

// GetExp returns the Exp field value if set, zero value otherwise.
func (o *IDToken) GetExp() int64 {
	if o == nil || IsNil(o.Exp) {
		var ret int64
		return ret
	}
	return *o.Exp
}

// GetExpOk returns a tuple with the Exp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetExpOk() (*int64, bool) {
	if o == nil || IsNil(o.Exp) {
		return nil, false
	}
	return o.Exp, true
}

// HasExp returns a boolean if a field has been set.
func (o *IDToken) HasExp() bool {
	if o != nil && !IsNil(o.Exp) {
		return true
	}

	return false
}

// SetExp gets a reference to the given int64 and assigns it to the Exp field.
func (o *IDToken) SetExp(v int64) {
	o.Exp = &v
}

// GetNbf returns the Nbf field value if set, zero value otherwise.
func (o *IDToken) GetNbf() int64 {
	if o == nil || IsNil(o.Nbf) {
		var ret int64
		return ret
	}
	return *o.Nbf
}

// GetNbfOk returns a tuple with the Nbf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetNbfOk() (*int64, bool) {
	if o == nil || IsNil(o.Nbf) {
		return nil, false
	}
	return o.Nbf, true
}

// HasNbf returns a boolean if a field has been set.
func (o *IDToken) HasNbf() bool {
	if o != nil && !IsNil(o.Nbf) {
		return true
	}

	return false
}

// SetNbf gets a reference to the given int64 and assigns it to the Nbf field.
func (o *IDToken) SetNbf(v int64) {
	o.Nbf = &v
}

// GetIat returns the Iat field value if set, zero value otherwise.
func (o *IDToken) GetIat() int64 {
	if o == nil || IsNil(o.Iat) {
		var ret int64
		return ret
	}
	return *o.Iat
}

// GetIatOk returns a tuple with the Iat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetIatOk() (*int64, bool) {
	if o == nil || IsNil(o.Iat) {
		return nil, false
	}
	return o.Iat, true
}

// HasIat returns a boolean if a field has been set.
func (o *IDToken) HasIat() bool {
	if o != nil && !IsNil(o.Iat) {
		return true
	}

	return false
}

// SetIat gets a reference to the given int64 and assigns it to the Iat field.
func (o *IDToken) SetIat(v int64) {
	o.Iat = &v
}

// GetIss returns the Iss field value if set, zero value otherwise.
func (o *IDToken) GetIss() string {
	if o == nil || IsNil(o.Iss) {
		var ret string
		return ret
	}
	return *o.Iss
}

// GetIssOk returns a tuple with the Iss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetIssOk() (*string, bool) {
	if o == nil || IsNil(o.Iss) {
		return nil, false
	}
	return o.Iss, true
}

// HasIss returns a boolean if a field has been set.
func (o *IDToken) HasIss() bool {
	if o != nil && !IsNil(o.Iss) {
		return true
	}

	return false
}

// SetIss gets a reference to the given string and assigns it to the Iss field.
func (o *IDToken) SetIss(v string) {
	o.Iss = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *IDToken) GetSub() string {
	if o == nil || IsNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetSubOk() (*string, bool) {
	if o == nil || IsNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *IDToken) HasSub() bool {
	if o != nil && !IsNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *IDToken) SetSub(v string) {
	o.Sub = &v
}

// GetTyp returns the Typ field value if set, zero value otherwise.
func (o *IDToken) GetTyp() string {
	if o == nil || IsNil(o.Typ) {
		var ret string
		return ret
	}
	return *o.Typ
}

// GetTypOk returns a tuple with the Typ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetTypOk() (*string, bool) {
	if o == nil || IsNil(o.Typ) {
		return nil, false
	}
	return o.Typ, true
}

// HasTyp returns a boolean if a field has been set.
func (o *IDToken) HasTyp() bool {
	if o != nil && !IsNil(o.Typ) {
		return true
	}

	return false
}

// SetTyp gets a reference to the given string and assigns it to the Typ field.
func (o *IDToken) SetTyp(v string) {
	o.Typ = &v
}

// GetAzp returns the Azp field value if set, zero value otherwise.
func (o *IDToken) GetAzp() string {
	if o == nil || IsNil(o.Azp) {
		var ret string
		return ret
	}
	return *o.Azp
}

// GetAzpOk returns a tuple with the Azp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetAzpOk() (*string, bool) {
	if o == nil || IsNil(o.Azp) {
		return nil, false
	}
	return o.Azp, true
}

// HasAzp returns a boolean if a field has been set.
func (o *IDToken) HasAzp() bool {
	if o != nil && !IsNil(o.Azp) {
		return true
	}

	return false
}

// SetAzp gets a reference to the given string and assigns it to the Azp field.
func (o *IDToken) SetAzp(v string) {
	o.Azp = &v
}

// GetOtherClaims returns the OtherClaims field value if set, zero value otherwise.
func (o *IDToken) GetOtherClaims() map[string]interface{} {
	if o == nil || IsNil(o.OtherClaims) {
		var ret map[string]interface{}
		return ret
	}
	return o.OtherClaims
}

// GetOtherClaimsOk returns a tuple with the OtherClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetOtherClaimsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OtherClaims) {
		return map[string]interface{}{}, false
	}
	return o.OtherClaims, true
}

// HasOtherClaims returns a boolean if a field has been set.
func (o *IDToken) HasOtherClaims() bool {
	if o != nil && !IsNil(o.OtherClaims) {
		return true
	}

	return false
}

// SetOtherClaims gets a reference to the given map[string]interface{} and assigns it to the OtherClaims field.
func (o *IDToken) SetOtherClaims(v map[string]interface{}) {
	o.OtherClaims = v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *IDToken) GetNonce() string {
	if o == nil || IsNil(o.Nonce) {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetNonceOk() (*string, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *IDToken) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *IDToken) SetNonce(v string) {
	o.Nonce = &v
}

// GetAuthTime returns the AuthTime field value if set, zero value otherwise.
func (o *IDToken) GetAuthTime() int64 {
	if o == nil || IsNil(o.AuthTime) {
		var ret int64
		return ret
	}
	return *o.AuthTime
}

// GetAuthTimeOk returns a tuple with the AuthTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetAuthTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthTime) {
		return nil, false
	}
	return o.AuthTime, true
}

// HasAuthTime returns a boolean if a field has been set.
func (o *IDToken) HasAuthTime() bool {
	if o != nil && !IsNil(o.AuthTime) {
		return true
	}

	return false
}

// SetAuthTime gets a reference to the given int64 and assigns it to the AuthTime field.
func (o *IDToken) SetAuthTime(v int64) {
	o.AuthTime = &v
}

// GetSid returns the Sid field value if set, zero value otherwise.
func (o *IDToken) GetSid() string {
	if o == nil || IsNil(o.Sid) {
		var ret string
		return ret
	}
	return *o.Sid
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetSidOk() (*string, bool) {
	if o == nil || IsNil(o.Sid) {
		return nil, false
	}
	return o.Sid, true
}

// HasSid returns a boolean if a field has been set.
func (o *IDToken) HasSid() bool {
	if o != nil && !IsNil(o.Sid) {
		return true
	}

	return false
}

// SetSid gets a reference to the given string and assigns it to the Sid field.
func (o *IDToken) SetSid(v string) {
	o.Sid = &v
}

// GetAtHash returns the AtHash field value if set, zero value otherwise.
func (o *IDToken) GetAtHash() string {
	if o == nil || IsNil(o.AtHash) {
		var ret string
		return ret
	}
	return *o.AtHash
}

// GetAtHashOk returns a tuple with the AtHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetAtHashOk() (*string, bool) {
	if o == nil || IsNil(o.AtHash) {
		return nil, false
	}
	return o.AtHash, true
}

// HasAtHash returns a boolean if a field has been set.
func (o *IDToken) HasAtHash() bool {
	if o != nil && !IsNil(o.AtHash) {
		return true
	}

	return false
}

// SetAtHash gets a reference to the given string and assigns it to the AtHash field.
func (o *IDToken) SetAtHash(v string) {
	o.AtHash = &v
}

// GetCHash returns the CHash field value if set, zero value otherwise.
func (o *IDToken) GetCHash() string {
	if o == nil || IsNil(o.CHash) {
		var ret string
		return ret
	}
	return *o.CHash
}

// GetCHashOk returns a tuple with the CHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetCHashOk() (*string, bool) {
	if o == nil || IsNil(o.CHash) {
		return nil, false
	}
	return o.CHash, true
}

// HasCHash returns a boolean if a field has been set.
func (o *IDToken) HasCHash() bool {
	if o != nil && !IsNil(o.CHash) {
		return true
	}

	return false
}

// SetCHash gets a reference to the given string and assigns it to the CHash field.
func (o *IDToken) SetCHash(v string) {
	o.CHash = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IDToken) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IDToken) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IDToken) SetName(v string) {
	o.Name = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *IDToken) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *IDToken) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *IDToken) SetGivenName(v string) {
	o.GivenName = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *IDToken) GetFamilyName() string {
	if o == nil || IsNil(o.FamilyName) {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetFamilyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FamilyName) {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *IDToken) HasFamilyName() bool {
	if o != nil && !IsNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *IDToken) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *IDToken) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *IDToken) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *IDToken) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *IDToken) GetNickname() string {
	if o == nil || IsNil(o.Nickname) {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname) {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *IDToken) HasNickname() bool {
	if o != nil && !IsNil(o.Nickname) {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *IDToken) SetNickname(v string) {
	o.Nickname = &v
}

// GetPreferredUsername returns the PreferredUsername field value if set, zero value otherwise.
func (o *IDToken) GetPreferredUsername() string {
	if o == nil || IsNil(o.PreferredUsername) {
		var ret string
		return ret
	}
	return *o.PreferredUsername
}

// GetPreferredUsernameOk returns a tuple with the PreferredUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetPreferredUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredUsername) {
		return nil, false
	}
	return o.PreferredUsername, true
}

// HasPreferredUsername returns a boolean if a field has been set.
func (o *IDToken) HasPreferredUsername() bool {
	if o != nil && !IsNil(o.PreferredUsername) {
		return true
	}

	return false
}

// SetPreferredUsername gets a reference to the given string and assigns it to the PreferredUsername field.
func (o *IDToken) SetPreferredUsername(v string) {
	o.PreferredUsername = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *IDToken) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *IDToken) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *IDToken) SetProfile(v string) {
	o.Profile = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *IDToken) GetPicture() string {
	if o == nil || IsNil(o.Picture) {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetPictureOk() (*string, bool) {
	if o == nil || IsNil(o.Picture) {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *IDToken) HasPicture() bool {
	if o != nil && !IsNil(o.Picture) {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *IDToken) SetPicture(v string) {
	o.Picture = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *IDToken) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *IDToken) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *IDToken) SetWebsite(v string) {
	o.Website = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *IDToken) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *IDToken) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *IDToken) SetEmail(v string) {
	o.Email = &v
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise.
func (o *IDToken) GetEmailVerified() bool {
	if o == nil || IsNil(o.EmailVerified) {
		var ret bool
		return ret
	}
	return *o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailVerified) {
		return nil, false
	}
	return o.EmailVerified, true
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *IDToken) HasEmailVerified() bool {
	if o != nil && !IsNil(o.EmailVerified) {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given bool and assigns it to the EmailVerified field.
func (o *IDToken) SetEmailVerified(v bool) {
	o.EmailVerified = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *IDToken) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *IDToken) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *IDToken) SetGender(v string) {
	o.Gender = &v
}

// GetBirthdate returns the Birthdate field value if set, zero value otherwise.
func (o *IDToken) GetBirthdate() string {
	if o == nil || IsNil(o.Birthdate) {
		var ret string
		return ret
	}
	return *o.Birthdate
}

// GetBirthdateOk returns a tuple with the Birthdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetBirthdateOk() (*string, bool) {
	if o == nil || IsNil(o.Birthdate) {
		return nil, false
	}
	return o.Birthdate, true
}

// HasBirthdate returns a boolean if a field has been set.
func (o *IDToken) HasBirthdate() bool {
	if o != nil && !IsNil(o.Birthdate) {
		return true
	}

	return false
}

// SetBirthdate gets a reference to the given string and assigns it to the Birthdate field.
func (o *IDToken) SetBirthdate(v string) {
	o.Birthdate = &v
}

// GetZoneinfo returns the Zoneinfo field value if set, zero value otherwise.
func (o *IDToken) GetZoneinfo() string {
	if o == nil || IsNil(o.Zoneinfo) {
		var ret string
		return ret
	}
	return *o.Zoneinfo
}

// GetZoneinfoOk returns a tuple with the Zoneinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetZoneinfoOk() (*string, bool) {
	if o == nil || IsNil(o.Zoneinfo) {
		return nil, false
	}
	return o.Zoneinfo, true
}

// HasZoneinfo returns a boolean if a field has been set.
func (o *IDToken) HasZoneinfo() bool {
	if o != nil && !IsNil(o.Zoneinfo) {
		return true
	}

	return false
}

// SetZoneinfo gets a reference to the given string and assigns it to the Zoneinfo field.
func (o *IDToken) SetZoneinfo(v string) {
	o.Zoneinfo = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *IDToken) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *IDToken) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *IDToken) SetLocale(v string) {
	o.Locale = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *IDToken) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *IDToken) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *IDToken) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPhoneNumberVerified returns the PhoneNumberVerified field value if set, zero value otherwise.
func (o *IDToken) GetPhoneNumberVerified() bool {
	if o == nil || IsNil(o.PhoneNumberVerified) {
		var ret bool
		return ret
	}
	return *o.PhoneNumberVerified
}

// GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetPhoneNumberVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.PhoneNumberVerified) {
		return nil, false
	}
	return o.PhoneNumberVerified, true
}

// HasPhoneNumberVerified returns a boolean if a field has been set.
func (o *IDToken) HasPhoneNumberVerified() bool {
	if o != nil && !IsNil(o.PhoneNumberVerified) {
		return true
	}

	return false
}

// SetPhoneNumberVerified gets a reference to the given bool and assigns it to the PhoneNumberVerified field.
func (o *IDToken) SetPhoneNumberVerified(v bool) {
	o.PhoneNumberVerified = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *IDToken) GetAddress() AddressClaimSet {
	if o == nil || IsNil(o.Address) {
		var ret AddressClaimSet
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetAddressOk() (*AddressClaimSet, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *IDToken) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AddressClaimSet and assigns it to the Address field.
func (o *IDToken) SetAddress(v AddressClaimSet) {
	o.Address = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IDToken) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IDToken) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *IDToken) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetClaimsLocales returns the ClaimsLocales field value if set, zero value otherwise.
func (o *IDToken) GetClaimsLocales() string {
	if o == nil || IsNil(o.ClaimsLocales) {
		var ret string
		return ret
	}
	return *o.ClaimsLocales
}

// GetClaimsLocalesOk returns a tuple with the ClaimsLocales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetClaimsLocalesOk() (*string, bool) {
	if o == nil || IsNil(o.ClaimsLocales) {
		return nil, false
	}
	return o.ClaimsLocales, true
}

// HasClaimsLocales returns a boolean if a field has been set.
func (o *IDToken) HasClaimsLocales() bool {
	if o != nil && !IsNil(o.ClaimsLocales) {
		return true
	}

	return false
}

// SetClaimsLocales gets a reference to the given string and assigns it to the ClaimsLocales field.
func (o *IDToken) SetClaimsLocales(v string) {
	o.ClaimsLocales = &v
}

// GetAcr returns the Acr field value if set, zero value otherwise.
func (o *IDToken) GetAcr() string {
	if o == nil || IsNil(o.Acr) {
		var ret string
		return ret
	}
	return *o.Acr
}

// GetAcrOk returns a tuple with the Acr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetAcrOk() (*string, bool) {
	if o == nil || IsNil(o.Acr) {
		return nil, false
	}
	return o.Acr, true
}

// HasAcr returns a boolean if a field has been set.
func (o *IDToken) HasAcr() bool {
	if o != nil && !IsNil(o.Acr) {
		return true
	}

	return false
}

// SetAcr gets a reference to the given string and assigns it to the Acr field.
func (o *IDToken) SetAcr(v string) {
	o.Acr = &v
}

// GetSHash returns the SHash field value if set, zero value otherwise.
func (o *IDToken) GetSHash() string {
	if o == nil || IsNil(o.SHash) {
		var ret string
		return ret
	}
	return *o.SHash
}

// GetSHashOk returns a tuple with the SHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDToken) GetSHashOk() (*string, bool) {
	if o == nil || IsNil(o.SHash) {
		return nil, false
	}
	return o.SHash, true
}

// HasSHash returns a boolean if a field has been set.
func (o *IDToken) HasSHash() bool {
	if o != nil && !IsNil(o.SHash) {
		return true
	}

	return false
}

// SetSHash gets a reference to the given string and assigns it to the SHash field.
func (o *IDToken) SetSHash(v string) {
	o.SHash = &v
}

func (o IDToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IDToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Jti) {
		toSerialize["jti"] = o.Jti
	}
	if !IsNil(o.Exp) {
		toSerialize["exp"] = o.Exp
	}
	if !IsNil(o.Nbf) {
		toSerialize["nbf"] = o.Nbf
	}
	if !IsNil(o.Iat) {
		toSerialize["iat"] = o.Iat
	}
	if !IsNil(o.Iss) {
		toSerialize["iss"] = o.Iss
	}
	if !IsNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	if !IsNil(o.Typ) {
		toSerialize["typ"] = o.Typ
	}
	if !IsNil(o.Azp) {
		toSerialize["azp"] = o.Azp
	}
	if !IsNil(o.OtherClaims) {
		toSerialize["otherClaims"] = o.OtherClaims
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.AuthTime) {
		toSerialize["auth_time"] = o.AuthTime
	}
	if !IsNil(o.Sid) {
		toSerialize["sid"] = o.Sid
	}
	if !IsNil(o.AtHash) {
		toSerialize["at_hash"] = o.AtHash
	}
	if !IsNil(o.CHash) {
		toSerialize["c_hash"] = o.CHash
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.GivenName) {
		toSerialize["given_name"] = o.GivenName
	}
	if !IsNil(o.FamilyName) {
		toSerialize["family_name"] = o.FamilyName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middle_name"] = o.MiddleName
	}
	if !IsNil(o.Nickname) {
		toSerialize["nickname"] = o.Nickname
	}
	if !IsNil(o.PreferredUsername) {
		toSerialize["preferred_username"] = o.PreferredUsername
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Picture) {
		toSerialize["picture"] = o.Picture
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EmailVerified) {
		toSerialize["email_verified"] = o.EmailVerified
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.Birthdate) {
		toSerialize["birthdate"] = o.Birthdate
	}
	if !IsNil(o.Zoneinfo) {
		toSerialize["zoneinfo"] = o.Zoneinfo
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.PhoneNumberVerified) {
		toSerialize["phone_number_verified"] = o.PhoneNumberVerified
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ClaimsLocales) {
		toSerialize["claims_locales"] = o.ClaimsLocales
	}
	if !IsNil(o.Acr) {
		toSerialize["acr"] = o.Acr
	}
	if !IsNil(o.SHash) {
		toSerialize["s_hash"] = o.SHash
	}
	return toSerialize, nil
}

type NullableIDToken struct {
	value *IDToken
	isSet bool
}

func (v NullableIDToken) Get() *IDToken {
	return v.value
}

func (v *NullableIDToken) Set(val *IDToken) {
	v.value = val
	v.isSet = true
}

func (v NullableIDToken) IsSet() bool {
	return v.isSet
}

func (v *NullableIDToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIDToken(val *IDToken) *NullableIDToken {
	return &NullableIDToken{value: val, isSet: true}
}

func (v NullableIDToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIDToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


