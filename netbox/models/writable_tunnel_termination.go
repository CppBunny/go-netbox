// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableTunnelTermination writable tunnel termination
//
// swagger:model WritableTunnelTermination
type WritableTunnelTermination struct {

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// outside ip
	OutsideIP *int64 `json:"outside_ip,omitempty"`

	// Role
	// Enum: ["peer","hub","spoke"]
	Role string `json:"role,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// termination
	// Read Only: true
	Termination interface{} `json:"termination,omitempty"`

	// termination id
	TerminationID int64 `json:"termination_id,omitempty"`

	// termination type
	// Required: true
	TerminationType *string `json:"termination_type"`

	// tunnel
	// Required: true
	Tunnel *int64 `json:"tunnel"`
}

// Validate validates this writable tunnel termination
func (m *WritableTunnelTermination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTunnel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var writableTunnelTerminationTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["peer","hub","spoke"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableTunnelTerminationTypeRolePropEnum = append(writableTunnelTerminationTypeRolePropEnum, v)
	}
}

const (

	// WritableTunnelTerminationRolePeer captures enum value "peer"
	WritableTunnelTerminationRolePeer string = "peer"

	// WritableTunnelTerminationRoleHub captures enum value "hub"
	WritableTunnelTerminationRoleHub string = "hub"

	// WritableTunnelTerminationRoleSpoke captures enum value "spoke"
	WritableTunnelTerminationRoleSpoke string = "spoke"
)

// prop value enum
func (m *WritableTunnelTermination) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableTunnelTerminationTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableTunnelTermination) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

func (m *WritableTunnelTermination) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableTunnelTermination) validateTerminationType(formats strfmt.Registry) error {

	if err := validate.Required("termination_type", "body", m.TerminationType); err != nil {
		return err
	}

	return nil
}

func (m *WritableTunnelTermination) validateTunnel(formats strfmt.Registry) error {

	if err := validate.Required("tunnel", "body", m.Tunnel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable tunnel termination based on the context it is used
func (m *WritableTunnelTermination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableTunnelTermination) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableTunnelTermination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableTunnelTermination) UnmarshalBinary(b []byte) error {
	var res WritableTunnelTermination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
