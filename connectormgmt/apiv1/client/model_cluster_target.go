/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
	"fmt"
)

// ClusterTarget - struct for ClusterTarget
type ClusterTarget struct {
	AddonClusterTarget *AddonClusterTarget
	CloudProviderClusterTarget *CloudProviderClusterTarget
}

// AddonClusterTargetAsClusterTarget is a convenience function that returns AddonClusterTarget wrapped in ClusterTarget
func AddonClusterTargetAsClusterTarget(v *AddonClusterTarget) ClusterTarget {
	return ClusterTarget{ AddonClusterTarget: v}
}

// CloudProviderClusterTargetAsClusterTarget is a convenience function that returns CloudProviderClusterTarget wrapped in ClusterTarget
func CloudProviderClusterTargetAsClusterTarget(v *CloudProviderClusterTarget) ClusterTarget {
	return ClusterTarget{ CloudProviderClusterTarget: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ClusterTarget) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddonClusterTarget
	err = json.Unmarshal(data, &dst.AddonClusterTarget)
	if err == nil {
		jsonAddonClusterTarget, _ := json.Marshal(dst.AddonClusterTarget)
		if string(jsonAddonClusterTarget) == "{}" { // empty struct
			dst.AddonClusterTarget = nil
		} else {
			match++
		}
	} else {
		dst.AddonClusterTarget = nil
	}

	// try to unmarshal data into CloudProviderClusterTarget
	err = json.Unmarshal(data, &dst.CloudProviderClusterTarget)
	if err == nil {
		jsonCloudProviderClusterTarget, _ := json.Marshal(dst.CloudProviderClusterTarget)
		if string(jsonCloudProviderClusterTarget) == "{}" { // empty struct
			dst.CloudProviderClusterTarget = nil
		} else {
			match++
		}
	} else {
		dst.CloudProviderClusterTarget = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddonClusterTarget = nil
		dst.CloudProviderClusterTarget = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ClusterTarget)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ClusterTarget)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ClusterTarget) MarshalJSON() ([]byte, error) {
	if src.AddonClusterTarget != nil {
		return json.Marshal(&src.AddonClusterTarget)
	}

	if src.CloudProviderClusterTarget != nil {
		return json.Marshal(&src.CloudProviderClusterTarget)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ClusterTarget) GetActualInstance() (interface{}) {
	if obj.AddonClusterTarget != nil {
		return obj.AddonClusterTarget
	}

	if obj.CloudProviderClusterTarget != nil {
		return obj.CloudProviderClusterTarget
	}

	// all schemas are nil
	return nil
}

type NullableClusterTarget struct {
	value *ClusterTarget
	isSet bool
}

func (v NullableClusterTarget) Get() *ClusterTarget {
	return v.value
}

func (v *NullableClusterTarget) Set(val *ClusterTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterTarget(val *ClusterTarget) *NullableClusterTarget {
	return &NullableClusterTarget{value: val, isSet: true}
}

func (v NullableClusterTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

