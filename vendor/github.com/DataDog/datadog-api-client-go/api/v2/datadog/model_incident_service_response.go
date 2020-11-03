/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// IncidentServiceResponse Response with an incident service payload.
type IncidentServiceResponse struct {
	Data IncidentServiceResponseData `json:"data"`
	// Included objects from relationships.
	Included *[]IncidentServiceIncludedItems `json:"included,omitempty"`
}

// NewIncidentServiceResponse instantiates a new IncidentServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentServiceResponse(data IncidentServiceResponseData) *IncidentServiceResponse {
	this := IncidentServiceResponse{}
	this.Data = data
	return &this
}

// NewIncidentServiceResponseWithDefaults instantiates a new IncidentServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentServiceResponseWithDefaults() *IncidentServiceResponse {
	this := IncidentServiceResponse{}
	return &this
}

// GetData returns the Data field value
func (o *IncidentServiceResponse) GetData() IncidentServiceResponseData {
	if o == nil {
		var ret IncidentServiceResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentServiceResponse) GetDataOk() (*IncidentServiceResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IncidentServiceResponse) SetData(v IncidentServiceResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentServiceResponse) GetIncluded() []IncidentServiceIncludedItems {
	if o == nil || o.Included == nil {
		var ret []IncidentServiceIncludedItems
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentServiceResponse) GetIncludedOk() (*[]IncidentServiceIncludedItems, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentServiceResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []IncidentServiceIncludedItems and assigns it to the Included field.
func (o *IncidentServiceResponse) SetIncluded(v []IncidentServiceIncludedItems) {
	o.Included = &v
}

func (o IncidentServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentServiceResponse struct {
	value *IncidentServiceResponse
	isSet bool
}

func (v NullableIncidentServiceResponse) Get() *IncidentServiceResponse {
	return v.value
}

func (v *NullableIncidentServiceResponse) Set(val *IncidentServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentServiceResponse(val *IncidentServiceResponse) *NullableIncidentServiceResponse {
	return &NullableIncidentServiceResponse{value: val, isSet: true}
}

func (v NullableIncidentServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}