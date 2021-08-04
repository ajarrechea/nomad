/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.3
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testclient

import (
	"encoding/json"
)

// QuotaLimit struct for QuotaLimit
type QuotaLimit struct {
	Hash        *string    `json:"Hash,omitempty"`
	Region      *string    `json:"Region,omitempty"`
	RegionLimit *Resources `json:"RegionLimit,omitempty"`
}

// NewQuotaLimit instantiates a new QuotaLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaLimit() *QuotaLimit {
	this := QuotaLimit{}
	return &this
}

// NewQuotaLimitWithDefaults instantiates a new QuotaLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaLimitWithDefaults() *QuotaLimit {
	this := QuotaLimit{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *QuotaLimit) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaLimit) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *QuotaLimit) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *QuotaLimit) SetHash(v string) {
	o.Hash = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *QuotaLimit) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaLimit) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *QuotaLimit) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *QuotaLimit) SetRegion(v string) {
	o.Region = &v
}

// GetRegionLimit returns the RegionLimit field value if set, zero value otherwise.
func (o *QuotaLimit) GetRegionLimit() Resources {
	if o == nil || o.RegionLimit == nil {
		var ret Resources
		return ret
	}
	return *o.RegionLimit
}

// GetRegionLimitOk returns a tuple with the RegionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaLimit) GetRegionLimitOk() (*Resources, bool) {
	if o == nil || o.RegionLimit == nil {
		return nil, false
	}
	return o.RegionLimit, true
}

// HasRegionLimit returns a boolean if a field has been set.
func (o *QuotaLimit) HasRegionLimit() bool {
	if o != nil && o.RegionLimit != nil {
		return true
	}

	return false
}

// SetRegionLimit gets a reference to the given Resources and assigns it to the RegionLimit field.
func (o *QuotaLimit) SetRegionLimit(v Resources) {
	o.RegionLimit = &v
}

func (o QuotaLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hash != nil {
		toSerialize["Hash"] = o.Hash
	}
	if o.Region != nil {
		toSerialize["Region"] = o.Region
	}
	if o.RegionLimit != nil {
		toSerialize["RegionLimit"] = o.RegionLimit
	}
	return json.Marshal(toSerialize)
}

type NullableQuotaLimit struct {
	value *QuotaLimit
	isSet bool
}

func (v NullableQuotaLimit) Get() *QuotaLimit {
	return v.value
}

func (v *NullableQuotaLimit) Set(val *QuotaLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaLimit(val *QuotaLimit) *NullableQuotaLimit {
	return &NullableQuotaLimit{value: val, isSet: true}
}

func (v NullableQuotaLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
