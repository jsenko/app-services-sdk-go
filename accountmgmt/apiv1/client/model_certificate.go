/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
)

// Certificate struct for Certificate
type Certificate struct {
	Cert string `json:"cert"`
	Id string `json:"id"`
	Key string `json:"key"`
	Metadata map[string]string `json:"metadata"`
	OrganizationId string `json:"organization_id"`
	Serial CertificateSerial `json:"serial"`
}

// NewCertificate instantiates a new Certificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificate(cert string, id string, key string, metadata map[string]string, organizationId string, serial CertificateSerial) *Certificate {
	this := Certificate{}
	this.Cert = cert
	this.Id = id
	this.Key = key
	this.Metadata = metadata
	this.OrganizationId = organizationId
	this.Serial = serial
	return &this
}

// NewCertificateWithDefaults instantiates a new Certificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateWithDefaults() *Certificate {
	this := Certificate{}
	return &this
}

// GetCert returns the Cert field value
func (o *Certificate) GetCert() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cert
}

// GetCertOk returns a tuple with the Cert field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cert, true
}

// SetCert sets field value
func (o *Certificate) SetCert(v string) {
	o.Cert = v
}

// GetId returns the Id field value
func (o *Certificate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Certificate) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *Certificate) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Certificate) SetKey(v string) {
	o.Key = v
}

// GetMetadata returns the Metadata field value
func (o *Certificate) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetMetadataOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Certificate) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *Certificate) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetOrganizationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *Certificate) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetSerial returns the Serial field value
func (o *Certificate) GetSerial() CertificateSerial {
	if o == nil {
		var ret CertificateSerial
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetSerialOk() (*CertificateSerial, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *Certificate) SetSerial(v CertificateSerial) {
	o.Serial = v
}

func (o Certificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cert"] = o.Cert
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableCertificate struct {
	value *Certificate
	isSet bool
}

func (v NullableCertificate) Get() *Certificate {
	return v.value
}

func (v *NullableCertificate) Set(val *Certificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificate(val *Certificate) *NullableCertificate {
	return &NullableCertificate{value: val, isSet: true}
}

func (v NullableCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


