/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// EducationSchool Represents a school
type EducationSchool struct {
	// The unique idenfier for an entity. Read-only.
	Id *string `json:"id,omitempty"`
	// The organization name
	DisplayName *string `json:"displayName,omitempty"`
	// School number
	SchoolNumber *string `json:"schoolNumber,omitempty"`
}

// NewEducationSchool instantiates a new EducationSchool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEducationSchool() *EducationSchool {
	this := EducationSchool{}
	return &this
}

// NewEducationSchoolWithDefaults instantiates a new EducationSchool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEducationSchoolWithDefaults() *EducationSchool {
	this := EducationSchool{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EducationSchool) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EducationSchool) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EducationSchool) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EducationSchool) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EducationSchool) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EducationSchool) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSchoolNumber returns the SchoolNumber field value if set, zero value otherwise.
func (o *EducationSchool) GetSchoolNumber() string {
	if o == nil || o.SchoolNumber == nil {
		var ret string
		return ret
	}
	return *o.SchoolNumber
}

// GetSchoolNumberOk returns a tuple with the SchoolNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetSchoolNumberOk() (*string, bool) {
	if o == nil || o.SchoolNumber == nil {
		return nil, false
	}
	return o.SchoolNumber, true
}

// HasSchoolNumber returns a boolean if a field has been set.
func (o *EducationSchool) HasSchoolNumber() bool {
	if o != nil && o.SchoolNumber != nil {
		return true
	}

	return false
}

// SetSchoolNumber gets a reference to the given string and assigns it to the SchoolNumber field.
func (o *EducationSchool) SetSchoolNumber(v string) {
	o.SchoolNumber = &v
}

func (o EducationSchool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.SchoolNumber != nil {
		toSerialize["schoolNumber"] = o.SchoolNumber
	}
	return json.Marshal(toSerialize)
}

type NullableEducationSchool struct {
	value *EducationSchool
	isSet bool
}

func (v NullableEducationSchool) Get() *EducationSchool {
	return v.value
}

func (v *NullableEducationSchool) Set(val *EducationSchool) {
	v.value = val
	v.isSet = true
}

func (v NullableEducationSchool) IsSet() bool {
	return v.isSet
}

func (v *NullableEducationSchool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEducationSchool(val *EducationSchool) *NullableEducationSchool {
	return &NullableEducationSchool{value: val, isSet: true}
}

func (v NullableEducationSchool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEducationSchool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
