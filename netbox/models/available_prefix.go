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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AvailablePrefix available prefix
//
// swagger:model AvailablePrefix
type AvailablePrefix struct {

	// Family
	// Read Only: true
	Family int64 `json:"family,omitempty"`

	// Prefix
	// Read Only: true
	// Min Length: 1
	Prefix string `json:"prefix,omitempty"`

	// vrf
	Vrf *NestedVRF `json:"vrf,omitempty"`
}

// Validate validates this available prefix
func (m *AvailablePrefix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVrf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailablePrefix) validatePrefix(formats strfmt.Registry) error {
	if swag.IsZero(m.Prefix) { // not required
		return nil
	}

	if err := validate.MinLength("prefix", "body", m.Prefix, 1); err != nil {
		return err
	}

	return nil
}

func (m *AvailablePrefix) validateVrf(formats strfmt.Registry) error {
	if swag.IsZero(m.Vrf) { // not required
		return nil
	}

	if m.Vrf != nil {
		if err := m.Vrf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vrf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vrf")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this available prefix based on the context it is used
func (m *AvailablePrefix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFamily(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrefix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVrf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailablePrefix) contextValidateFamily(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "family", "body", int64(m.Family)); err != nil {
		return err
	}

	return nil
}

func (m *AvailablePrefix) contextValidatePrefix(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "prefix", "body", string(m.Prefix)); err != nil {
		return err
	}

	return nil
}

func (m *AvailablePrefix) contextValidateVrf(ctx context.Context, formats strfmt.Registry) error {

	if m.Vrf != nil {

		if swag.IsZero(m.Vrf) { // not required
			return nil
		}

		if err := m.Vrf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vrf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vrf")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AvailablePrefix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailablePrefix) UnmarshalBinary(b []byte) error {
	var res AvailablePrefix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
