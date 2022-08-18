/*
 * Red Hat Openshift SmartEvents Fleet Manager
 *
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * API version: 0.0.1
 * Contact: openbridge-dev@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsmgmtclient

import (
	"encoding/json"
	"time"
)

// ProcessorResponse struct for ProcessorResponse
type ProcessorResponse struct {
	Kind string `json:"kind"`
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	Href string `json:"href"`
	SubmittedAt time.Time `json:"submitted_at"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	Status ManagedResourceStatus `json:"status"`
	Owner string `json:"owner"`
	Type ProcessorType `json:"type"`
	Filters *[]BaseFilter `json:"filters,omitempty"`
	TransformationTemplate *string `json:"transformationTemplate,omitempty"`
	Action *Action `json:"action,omitempty"`
	Source *Source `json:"source,omitempty"`
}

// NewProcessorResponse instantiates a new ProcessorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorResponse(kind string, id string, href string, submittedAt time.Time, status ManagedResourceStatus, owner string, type_ ProcessorType) *ProcessorResponse {
	this := ProcessorResponse{}
	this.Kind = kind
	this.Id = id
	this.Href = href
	this.SubmittedAt = submittedAt
	this.Status = status
	this.Owner = owner
	this.Type = type_
	return &this
}

// NewProcessorResponseWithDefaults instantiates a new ProcessorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorResponseWithDefaults() *ProcessorResponse {
	this := ProcessorResponse{}
	return &this
}

// GetKind returns the Kind field value
func (o *ProcessorResponse) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ProcessorResponse) SetKind(v string) {
	o.Kind = v
}

// GetId returns the Id field value
func (o *ProcessorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProcessorResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProcessorResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProcessorResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProcessorResponse) SetName(v string) {
	o.Name = &v
}

// GetHref returns the Href field value
func (o *ProcessorResponse) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetHrefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ProcessorResponse) SetHref(v string) {
	o.Href = v
}

// GetSubmittedAt returns the SubmittedAt field value
func (o *ProcessorResponse) GetSubmittedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetSubmittedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubmittedAt, true
}

// SetSubmittedAt sets field value
func (o *ProcessorResponse) SetSubmittedAt(v time.Time) {
	o.SubmittedAt = v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *ProcessorResponse) GetPublishedAt() time.Time {
	if o == nil || o.PublishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || o.PublishedAt == nil {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *ProcessorResponse) HasPublishedAt() bool {
	if o != nil && o.PublishedAt != nil {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *ProcessorResponse) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ProcessorResponse) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ProcessorResponse) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *ProcessorResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetStatus returns the Status field value
func (o *ProcessorResponse) GetStatus() ManagedResourceStatus {
	if o == nil {
		var ret ManagedResourceStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetStatusOk() (*ManagedResourceStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ProcessorResponse) SetStatus(v ManagedResourceStatus) {
	o.Status = v
}

// GetOwner returns the Owner field value
func (o *ProcessorResponse) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *ProcessorResponse) SetOwner(v string) {
	o.Owner = v
}

// GetType returns the Type field value
func (o *ProcessorResponse) GetType() ProcessorType {
	if o == nil {
		var ret ProcessorType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetTypeOk() (*ProcessorType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProcessorResponse) SetType(v ProcessorType) {
	o.Type = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ProcessorResponse) GetFilters() []BaseFilter {
	if o == nil || o.Filters == nil {
		var ret []BaseFilter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetFiltersOk() (*[]BaseFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ProcessorResponse) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []BaseFilter and assigns it to the Filters field.
func (o *ProcessorResponse) SetFilters(v []BaseFilter) {
	o.Filters = &v
}

// GetTransformationTemplate returns the TransformationTemplate field value if set, zero value otherwise.
func (o *ProcessorResponse) GetTransformationTemplate() string {
	if o == nil || o.TransformationTemplate == nil {
		var ret string
		return ret
	}
	return *o.TransformationTemplate
}

// GetTransformationTemplateOk returns a tuple with the TransformationTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetTransformationTemplateOk() (*string, bool) {
	if o == nil || o.TransformationTemplate == nil {
		return nil, false
	}
	return o.TransformationTemplate, true
}

// HasTransformationTemplate returns a boolean if a field has been set.
func (o *ProcessorResponse) HasTransformationTemplate() bool {
	if o != nil && o.TransformationTemplate != nil {
		return true
	}

	return false
}

// SetTransformationTemplate gets a reference to the given string and assigns it to the TransformationTemplate field.
func (o *ProcessorResponse) SetTransformationTemplate(v string) {
	o.TransformationTemplate = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ProcessorResponse) GetAction() Action {
	if o == nil || o.Action == nil {
		var ret Action
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetActionOk() (*Action, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ProcessorResponse) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given Action and assigns it to the Action field.
func (o *ProcessorResponse) SetAction(v Action) {
	o.Action = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProcessorResponse) GetSource() Source {
	if o == nil || o.Source == nil {
		var ret Source
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetSourceOk() (*Source, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProcessorResponse) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given Source and assigns it to the Source field.
func (o *ProcessorResponse) SetSource(v Source) {
	o.Source = &v
}

func (o ProcessorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["submitted_at"] = o.SubmittedAt
	}
	if o.PublishedAt != nil {
		toSerialize["published_at"] = o.PublishedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.TransformationTemplate != nil {
		toSerialize["transformationTemplate"] = o.TransformationTemplate
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorResponse struct {
	value *ProcessorResponse
	isSet bool
}

func (v NullableProcessorResponse) Get() *ProcessorResponse {
	return v.value
}

func (v *NullableProcessorResponse) Set(val *ProcessorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorResponse(val *ProcessorResponse) *NullableProcessorResponse {
	return &NullableProcessorResponse{value: val, isSet: true}
}

func (v NullableProcessorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

