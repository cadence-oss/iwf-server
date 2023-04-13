/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// checks if the WorkflowStartOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStartOptions{}

// WorkflowStartOptions struct for WorkflowStartOptions
type WorkflowStartOptions struct {
	WorkflowIDReusePolicy  *WorkflowIDReusePolicy `json:"workflowIDReusePolicy,omitempty"`
	CronSchedule           *string                `json:"cronSchedule,omitempty"`
	RetryPolicy            *WorkflowRetryPolicy   `json:"retryPolicy,omitempty"`
	SearchAttributes       []SearchAttribute      `json:"searchAttributes,omitempty"`
	WorkflowConfigOverride *WorkflowConfig        `json:"workflowConfigOverride,omitempty"`
	IdReusePolicy          *IDReusePolicy         `json:"idReusePolicy,omitempty"`
}

// NewWorkflowStartOptions instantiates a new WorkflowStartOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStartOptions() *WorkflowStartOptions {
	this := WorkflowStartOptions{}
	return &this
}

// NewWorkflowStartOptionsWithDefaults instantiates a new WorkflowStartOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStartOptionsWithDefaults() *WorkflowStartOptions {
	this := WorkflowStartOptions{}
	return &this
}

// GetWorkflowIDReusePolicy returns the WorkflowIDReusePolicy field value if set, zero value otherwise.
func (o *WorkflowStartOptions) GetWorkflowIDReusePolicy() WorkflowIDReusePolicy {
	if o == nil || IsNil(o.WorkflowIDReusePolicy) {
		var ret WorkflowIDReusePolicy
		return ret
	}
	return *o.WorkflowIDReusePolicy
}

// GetWorkflowIDReusePolicyOk returns a tuple with the WorkflowIDReusePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartOptions) GetWorkflowIDReusePolicyOk() (*WorkflowIDReusePolicy, bool) {
	if o == nil || IsNil(o.WorkflowIDReusePolicy) {
		return nil, false
	}
	return o.WorkflowIDReusePolicy, true
}

// HasWorkflowIDReusePolicy returns a boolean if a field has been set.
func (o *WorkflowStartOptions) HasWorkflowIDReusePolicy() bool {
	if o != nil && !IsNil(o.WorkflowIDReusePolicy) {
		return true
	}

	return false
}

// SetWorkflowIDReusePolicy gets a reference to the given WorkflowIDReusePolicy and assigns it to the WorkflowIDReusePolicy field.
func (o *WorkflowStartOptions) SetWorkflowIDReusePolicy(v WorkflowIDReusePolicy) {
	o.WorkflowIDReusePolicy = &v
}

// GetCronSchedule returns the CronSchedule field value if set, zero value otherwise.
func (o *WorkflowStartOptions) GetCronSchedule() string {
	if o == nil || IsNil(o.CronSchedule) {
		var ret string
		return ret
	}
	return *o.CronSchedule
}

// GetCronScheduleOk returns a tuple with the CronSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartOptions) GetCronScheduleOk() (*string, bool) {
	if o == nil || IsNil(o.CronSchedule) {
		return nil, false
	}
	return o.CronSchedule, true
}

// HasCronSchedule returns a boolean if a field has been set.
func (o *WorkflowStartOptions) HasCronSchedule() bool {
	if o != nil && !IsNil(o.CronSchedule) {
		return true
	}

	return false
}

// SetCronSchedule gets a reference to the given string and assigns it to the CronSchedule field.
func (o *WorkflowStartOptions) SetCronSchedule(v string) {
	o.CronSchedule = &v
}

// GetRetryPolicy returns the RetryPolicy field value if set, zero value otherwise.
func (o *WorkflowStartOptions) GetRetryPolicy() WorkflowRetryPolicy {
	if o == nil || IsNil(o.RetryPolicy) {
		var ret WorkflowRetryPolicy
		return ret
	}
	return *o.RetryPolicy
}

// GetRetryPolicyOk returns a tuple with the RetryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartOptions) GetRetryPolicyOk() (*WorkflowRetryPolicy, bool) {
	if o == nil || IsNil(o.RetryPolicy) {
		return nil, false
	}
	return o.RetryPolicy, true
}

// HasRetryPolicy returns a boolean if a field has been set.
func (o *WorkflowStartOptions) HasRetryPolicy() bool {
	if o != nil && !IsNil(o.RetryPolicy) {
		return true
	}

	return false
}

// SetRetryPolicy gets a reference to the given WorkflowRetryPolicy and assigns it to the RetryPolicy field.
func (o *WorkflowStartOptions) SetRetryPolicy(v WorkflowRetryPolicy) {
	o.RetryPolicy = &v
}

// GetSearchAttributes returns the SearchAttributes field value if set, zero value otherwise.
func (o *WorkflowStartOptions) GetSearchAttributes() []SearchAttribute {
	if o == nil || IsNil(o.SearchAttributes) {
		var ret []SearchAttribute
		return ret
	}
	return o.SearchAttributes
}

// GetSearchAttributesOk returns a tuple with the SearchAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartOptions) GetSearchAttributesOk() ([]SearchAttribute, bool) {
	if o == nil || IsNil(o.SearchAttributes) {
		return nil, false
	}
	return o.SearchAttributes, true
}

// HasSearchAttributes returns a boolean if a field has been set.
func (o *WorkflowStartOptions) HasSearchAttributes() bool {
	if o != nil && !IsNil(o.SearchAttributes) {
		return true
	}

	return false
}

// SetSearchAttributes gets a reference to the given []SearchAttribute and assigns it to the SearchAttributes field.
func (o *WorkflowStartOptions) SetSearchAttributes(v []SearchAttribute) {
	o.SearchAttributes = v
}

// GetWorkflowConfigOverride returns the WorkflowConfigOverride field value if set, zero value otherwise.
func (o *WorkflowStartOptions) GetWorkflowConfigOverride() WorkflowConfig {
	if o == nil || IsNil(o.WorkflowConfigOverride) {
		var ret WorkflowConfig
		return ret
	}
	return *o.WorkflowConfigOverride
}

// GetWorkflowConfigOverrideOk returns a tuple with the WorkflowConfigOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartOptions) GetWorkflowConfigOverrideOk() (*WorkflowConfig, bool) {
	if o == nil || IsNil(o.WorkflowConfigOverride) {
		return nil, false
	}
	return o.WorkflowConfigOverride, true
}

// HasWorkflowConfigOverride returns a boolean if a field has been set.
func (o *WorkflowStartOptions) HasWorkflowConfigOverride() bool {
	if o != nil && !IsNil(o.WorkflowConfigOverride) {
		return true
	}

	return false
}

// SetWorkflowConfigOverride gets a reference to the given WorkflowConfig and assigns it to the WorkflowConfigOverride field.
func (o *WorkflowStartOptions) SetWorkflowConfigOverride(v WorkflowConfig) {
	o.WorkflowConfigOverride = &v
}

// GetIdReusePolicy returns the IdReusePolicy field value if set, zero value otherwise.
func (o *WorkflowStartOptions) GetIdReusePolicy() IDReusePolicy {
	if o == nil || IsNil(o.IdReusePolicy) {
		var ret IDReusePolicy
		return ret
	}
	return *o.IdReusePolicy
}

// GetIdReusePolicyOk returns a tuple with the IdReusePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartOptions) GetIdReusePolicyOk() (*IDReusePolicy, bool) {
	if o == nil || IsNil(o.IdReusePolicy) {
		return nil, false
	}
	return o.IdReusePolicy, true
}

// HasIdReusePolicy returns a boolean if a field has been set.
func (o *WorkflowStartOptions) HasIdReusePolicy() bool {
	if o != nil && !IsNil(o.IdReusePolicy) {
		return true
	}

	return false
}

// SetIdReusePolicy gets a reference to the given IDReusePolicy and assigns it to the IdReusePolicy field.
func (o *WorkflowStartOptions) SetIdReusePolicy(v IDReusePolicy) {
	o.IdReusePolicy = &v
}

func (o WorkflowStartOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStartOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WorkflowIDReusePolicy) {
		toSerialize["workflowIDReusePolicy"] = o.WorkflowIDReusePolicy
	}
	if !IsNil(o.CronSchedule) {
		toSerialize["cronSchedule"] = o.CronSchedule
	}
	if !IsNil(o.RetryPolicy) {
		toSerialize["retryPolicy"] = o.RetryPolicy
	}
	if !IsNil(o.SearchAttributes) {
		toSerialize["searchAttributes"] = o.SearchAttributes
	}
	if !IsNil(o.WorkflowConfigOverride) {
		toSerialize["workflowConfigOverride"] = o.WorkflowConfigOverride
	}
	if !IsNil(o.IdReusePolicy) {
		toSerialize["idReusePolicy"] = o.IdReusePolicy
	}
	return toSerialize, nil
}

type NullableWorkflowStartOptions struct {
	value *WorkflowStartOptions
	isSet bool
}

func (v NullableWorkflowStartOptions) Get() *WorkflowStartOptions {
	return v.value
}

func (v *NullableWorkflowStartOptions) Set(val *WorkflowStartOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStartOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStartOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStartOptions(val *WorkflowStartOptions) *NullableWorkflowStartOptions {
	return &NullableWorkflowStartOptions{value: val, isSet: true}
}

func (v NullableWorkflowStartOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStartOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
