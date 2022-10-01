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

// WorkflowSignalResponse struct for WorkflowSignalResponse
type WorkflowSignalResponse struct {
	WorkflowRunId *string `json:"workflowRunId,omitempty"`
}

// NewWorkflowSignalResponse instantiates a new WorkflowSignalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSignalResponse() *WorkflowSignalResponse {
	this := WorkflowSignalResponse{}
	return &this
}

// NewWorkflowSignalResponseWithDefaults instantiates a new WorkflowSignalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSignalResponseWithDefaults() *WorkflowSignalResponse {
	this := WorkflowSignalResponse{}
	return &this
}

// GetWorkflowRunId returns the WorkflowRunId field value if set, zero value otherwise.
func (o *WorkflowSignalResponse) GetWorkflowRunId() string {
	if o == nil || o.WorkflowRunId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSignalResponse) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil || o.WorkflowRunId == nil {
		return nil, false
	}
	return o.WorkflowRunId, true
}

// HasWorkflowRunId returns a boolean if a field has been set.
func (o *WorkflowSignalResponse) HasWorkflowRunId() bool {
	if o != nil && o.WorkflowRunId != nil {
		return true
	}

	return false
}

// SetWorkflowRunId gets a reference to the given string and assigns it to the WorkflowRunId field.
func (o *WorkflowSignalResponse) SetWorkflowRunId(v string) {
	o.WorkflowRunId = &v
}

func (o WorkflowSignalResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WorkflowRunId != nil {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowSignalResponse struct {
	value *WorkflowSignalResponse
	isSet bool
}

func (v NullableWorkflowSignalResponse) Get() *WorkflowSignalResponse {
	return v.value
}

func (v *NullableWorkflowSignalResponse) Set(val *WorkflowSignalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSignalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSignalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSignalResponse(val *WorkflowSignalResponse) *NullableWorkflowSignalResponse {
	return &NullableWorkflowSignalResponse{value: val, isSet: true}
}

func (v NullableWorkflowSignalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSignalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

