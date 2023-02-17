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

// WorkflowGetResponse struct for WorkflowGetResponse
type WorkflowGetResponse struct {
	WorkflowRunId  string                  `json:"workflowRunId"`
	WorkflowStatus WorkflowStatus          `json:"workflowStatus"`
	Results        []StateCompletionOutput `json:"results,omitempty"`
	ErrorType      *WorkflowErrorType      `json:"errorType,omitempty"`
	ErrorMessage   *string                 `json:"errorMessage,omitempty"`
}

// NewWorkflowGetResponse instantiates a new WorkflowGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowGetResponse(workflowRunId string, workflowStatus WorkflowStatus) *WorkflowGetResponse {
	this := WorkflowGetResponse{}
	this.WorkflowRunId = workflowRunId
	this.WorkflowStatus = workflowStatus
	return &this
}

// NewWorkflowGetResponseWithDefaults instantiates a new WorkflowGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowGetResponseWithDefaults() *WorkflowGetResponse {
	this := WorkflowGetResponse{}
	return &this
}

// GetWorkflowRunId returns the WorkflowRunId field value
func (o *WorkflowGetResponse) GetWorkflowRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value
// and a boolean to check if the value has been set.
func (o *WorkflowGetResponse) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowRunId, true
}

// SetWorkflowRunId sets field value
func (o *WorkflowGetResponse) SetWorkflowRunId(v string) {
	o.WorkflowRunId = v
}

// GetWorkflowStatus returns the WorkflowStatus field value
func (o *WorkflowGetResponse) GetWorkflowStatus() WorkflowStatus {
	if o == nil {
		var ret WorkflowStatus
		return ret
	}

	return o.WorkflowStatus
}

// GetWorkflowStatusOk returns a tuple with the WorkflowStatus field value
// and a boolean to check if the value has been set.
func (o *WorkflowGetResponse) GetWorkflowStatusOk() (*WorkflowStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowStatus, true
}

// SetWorkflowStatus sets field value
func (o *WorkflowGetResponse) SetWorkflowStatus(v WorkflowStatus) {
	o.WorkflowStatus = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *WorkflowGetResponse) GetResults() []StateCompletionOutput {
	if o == nil || o.Results == nil {
		var ret []StateCompletionOutput
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetResponse) GetResultsOk() ([]StateCompletionOutput, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *WorkflowGetResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []StateCompletionOutput and assigns it to the Results field.
func (o *WorkflowGetResponse) SetResults(v []StateCompletionOutput) {
	o.Results = v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *WorkflowGetResponse) GetErrorType() WorkflowErrorType {
	if o == nil || o.ErrorType == nil {
		var ret WorkflowErrorType
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetResponse) GetErrorTypeOk() (*WorkflowErrorType, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *WorkflowGetResponse) HasErrorType() bool {
	if o != nil && o.ErrorType != nil {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given WorkflowErrorType and assigns it to the ErrorType field.
func (o *WorkflowGetResponse) SetErrorType(v WorkflowErrorType) {
	o.ErrorType = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *WorkflowGetResponse) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *WorkflowGetResponse) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *WorkflowGetResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o WorkflowGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	if true {
		toSerialize["workflowStatus"] = o.WorkflowStatus
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.ErrorType != nil {
		toSerialize["errorType"] = o.ErrorType
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowGetResponse struct {
	value *WorkflowGetResponse
	isSet bool
}

func (v NullableWorkflowGetResponse) Get() *WorkflowGetResponse {
	return v.value
}

func (v *NullableWorkflowGetResponse) Set(val *WorkflowGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowGetResponse(val *WorkflowGetResponse) *NullableWorkflowGetResponse {
	return &NullableWorkflowGetResponse{value: val, isSet: true}
}

func (v NullableWorkflowGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
