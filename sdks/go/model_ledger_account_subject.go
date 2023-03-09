/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the LedgerAccountSubject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LedgerAccountSubject{}

// LedgerAccountSubject struct for LedgerAccountSubject
type LedgerAccountSubject struct {
	Type string `json:"type"`
	Identifier string `json:"identifier"`
}

// NewLedgerAccountSubject instantiates a new LedgerAccountSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerAccountSubject(type_ string, identifier string) *LedgerAccountSubject {
	this := LedgerAccountSubject{}
	this.Type = type_
	this.Identifier = identifier
	return &this
}

// NewLedgerAccountSubjectWithDefaults instantiates a new LedgerAccountSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerAccountSubjectWithDefaults() *LedgerAccountSubject {
	this := LedgerAccountSubject{}
	return &this
}

// GetType returns the Type field value
func (o *LedgerAccountSubject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LedgerAccountSubject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LedgerAccountSubject) SetType(v string) {
	o.Type = v
}

// GetIdentifier returns the Identifier field value
func (o *LedgerAccountSubject) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *LedgerAccountSubject) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *LedgerAccountSubject) SetIdentifier(v string) {
	o.Identifier = v
}

func (o LedgerAccountSubject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LedgerAccountSubject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["identifier"] = o.Identifier
	return toSerialize, nil
}

type NullableLedgerAccountSubject struct {
	value *LedgerAccountSubject
	isSet bool
}

func (v NullableLedgerAccountSubject) Get() *LedgerAccountSubject {
	return v.value
}

func (v *NullableLedgerAccountSubject) Set(val *LedgerAccountSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerAccountSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerAccountSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerAccountSubject(val *LedgerAccountSubject) *NullableLedgerAccountSubject {
	return &NullableLedgerAccountSubject{value: val, isSet: true}
}

func (v NullableLedgerAccountSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerAccountSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

