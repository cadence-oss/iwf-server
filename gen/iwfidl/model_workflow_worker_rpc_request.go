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

// checks if the WorkflowWorkerRpcRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowWorkerRpcRequest{}

// WorkflowWorkerRpcRequest struct for WorkflowWorkerRpcRequest
type WorkflowWorkerRpcRequest struct {
	Context          Context           `json:"context"`
	WorkflowType     string            `json:"workflowType"`
	RpcName          string            `json:"rpcName"`
	Input            *EncodedObject    `json:"input,omitempty"`
	SearchAttributes []SearchAttribute `json:"searchAttributes,omitempty"`
	DataAttributes   []KeyValue        `json:"dataAttributes,omitempty"`
}

// NewWorkflowWorkerRpcRequest instantiates a new WorkflowWorkerRpcRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkerRpcRequest(context Context, workflowType string, rpcName string) *WorkflowWorkerRpcRequest {
	this := WorkflowWorkerRpcRequest{}
	this.Context = context
	this.WorkflowType = workflowType
	this.RpcName = rpcName
	return &this
}

// NewWorkflowWorkerRpcRequestWithDefaults instantiates a new WorkflowWorkerRpcRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkerRpcRequestWithDefaults() *WorkflowWorkerRpcRequest {
	this := WorkflowWorkerRpcRequest{}
	return &this
}

// GetContext returns the Context field value
func (o *WorkflowWorkerRpcRequest) GetContext() Context {
	if o == nil {
		var ret Context
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerRpcRequest) GetContextOk() (*Context, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *WorkflowWorkerRpcRequest) SetContext(v Context) {
	o.Context = v
}

// GetWorkflowType returns the WorkflowType field value
func (o *WorkflowWorkerRpcRequest) GetWorkflowType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerRpcRequest) GetWorkflowTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowType, true
}

// SetWorkflowType sets field value
func (o *WorkflowWorkerRpcRequest) SetWorkflowType(v string) {
	o.WorkflowType = v
}

// GetRpcName returns the RpcName field value
func (o *WorkflowWorkerRpcRequest) GetRpcName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RpcName
}

// GetRpcNameOk returns a tuple with the RpcName field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerRpcRequest) GetRpcNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RpcName, true
}

// SetRpcName sets field value
func (o *WorkflowWorkerRpcRequest) SetRpcName(v string) {
	o.RpcName = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *WorkflowWorkerRpcRequest) GetInput() EncodedObject {
	if o == nil || IsNil(o.Input) {
		var ret EncodedObject
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerRpcRequest) GetInputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowWorkerRpcRequest) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given EncodedObject and assigns it to the Input field.
func (o *WorkflowWorkerRpcRequest) SetInput(v EncodedObject) {
	o.Input = &v
}

// GetSearchAttributes returns the SearchAttributes field value if set, zero value otherwise.
func (o *WorkflowWorkerRpcRequest) GetSearchAttributes() []SearchAttribute {
	if o == nil || IsNil(o.SearchAttributes) {
		var ret []SearchAttribute
		return ret
	}
	return o.SearchAttributes
}

// GetSearchAttributesOk returns a tuple with the SearchAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerRpcRequest) GetSearchAttributesOk() ([]SearchAttribute, bool) {
	if o == nil || IsNil(o.SearchAttributes) {
		return nil, false
	}
	return o.SearchAttributes, true
}

// HasSearchAttributes returns a boolean if a field has been set.
func (o *WorkflowWorkerRpcRequest) HasSearchAttributes() bool {
	if o != nil && !IsNil(o.SearchAttributes) {
		return true
	}

	return false
}

// SetSearchAttributes gets a reference to the given []SearchAttribute and assigns it to the SearchAttributes field.
func (o *WorkflowWorkerRpcRequest) SetSearchAttributes(v []SearchAttribute) {
	o.SearchAttributes = v
}

// GetDataAttributes returns the DataAttributes field value if set, zero value otherwise.
func (o *WorkflowWorkerRpcRequest) GetDataAttributes() []KeyValue {
	if o == nil || IsNil(o.DataAttributes) {
		var ret []KeyValue
		return ret
	}
	return o.DataAttributes
}

// GetDataAttributesOk returns a tuple with the DataAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerRpcRequest) GetDataAttributesOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.DataAttributes) {
		return nil, false
	}
	return o.DataAttributes, true
}

// HasDataAttributes returns a boolean if a field has been set.
func (o *WorkflowWorkerRpcRequest) HasDataAttributes() bool {
	if o != nil && !IsNil(o.DataAttributes) {
		return true
	}

	return false
}

// SetDataAttributes gets a reference to the given []KeyValue and assigns it to the DataAttributes field.
func (o *WorkflowWorkerRpcRequest) SetDataAttributes(v []KeyValue) {
	o.DataAttributes = v
}

func (o WorkflowWorkerRpcRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowWorkerRpcRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	toSerialize["workflowType"] = o.WorkflowType
	toSerialize["rpcName"] = o.RpcName
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.SearchAttributes) {
		toSerialize["searchAttributes"] = o.SearchAttributes
	}
	if !IsNil(o.DataAttributes) {
		toSerialize["dataAttributes"] = o.DataAttributes
	}
	return toSerialize, nil
}

type NullableWorkflowWorkerRpcRequest struct {
	value *WorkflowWorkerRpcRequest
	isSet bool
}

func (v NullableWorkflowWorkerRpcRequest) Get() *WorkflowWorkerRpcRequest {
	return v.value
}

func (v *NullableWorkflowWorkerRpcRequest) Set(val *WorkflowWorkerRpcRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkerRpcRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkerRpcRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkerRpcRequest(val *WorkflowWorkerRpcRequest) *NullableWorkflowWorkerRpcRequest {
	return &NullableWorkflowWorkerRpcRequest{value: val, isSet: true}
}

func (v NullableWorkflowWorkerRpcRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkerRpcRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
