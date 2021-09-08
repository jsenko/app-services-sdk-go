/*
 * Service Registry Fleet Manager
 *
 * Managed Service Registry cloud.redhat.com API Management API that lets you create new registry instances. Registry is a datastore for standard event schemas and API designs. Service Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Registry is an Managed version of upstream project called Apicurio Registry. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.
 *
 * API version: 0.0.6
 * Contact: rhosak-eval-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registrymgmtclient

import (
	"encoding/json"
	"fmt"
)

// RegistryInstanceTypeValue \"standard\": Standard, full-featured Registry instance  \"eval\": Evaluation (Trial) instance, provided for a limited time 
type RegistryInstanceTypeValue string

// List of RegistryInstanceTypeValue
const (
	REGISTRYINSTANCETYPEVALUE_STANDARD RegistryInstanceTypeValue = "standard"
	REGISTRYINSTANCETYPEVALUE_EVAL RegistryInstanceTypeValue = "eval"
)

var allowedRegistryInstanceTypeValueEnumValues = []RegistryInstanceTypeValue{
	"standard",
	"eval",
}

func (v *RegistryInstanceTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegistryInstanceTypeValue(value)
	for _, existing := range allowedRegistryInstanceTypeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegistryInstanceTypeValue", value)
}

// NewRegistryInstanceTypeValueFromValue returns a pointer to a valid RegistryInstanceTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegistryInstanceTypeValueFromValue(v string) (*RegistryInstanceTypeValue, error) {
	ev := RegistryInstanceTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegistryInstanceTypeValue: valid values are %v", v, allowedRegistryInstanceTypeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegistryInstanceTypeValue) IsValid() bool {
	for _, existing := range allowedRegistryInstanceTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RegistryInstanceTypeValue value
func (v RegistryInstanceTypeValue) Ptr() *RegistryInstanceTypeValue {
	return &v
}

type NullableRegistryInstanceTypeValue struct {
	value *RegistryInstanceTypeValue
	isSet bool
}

func (v NullableRegistryInstanceTypeValue) Get() *RegistryInstanceTypeValue {
	return v.value
}

func (v *NullableRegistryInstanceTypeValue) Set(val *RegistryInstanceTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryInstanceTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryInstanceTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryInstanceTypeValue(val *RegistryInstanceTypeValue) *NullableRegistryInstanceTypeValue {
	return &NullableRegistryInstanceTypeValue{value: val, isSet: true}
}

func (v NullableRegistryInstanceTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryInstanceTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
