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

// AvailableIP available IP
//
// swagger:model AvailableIP
type AvailableIP struct {

	// Address
	// Read Only: true
	// Min Length: 1
	Address string `json:"address,omitempty"`

	// Family
	// Read Only: true
	Family int64 `json:"family,omitempty"`

	// vrf
	Vrf *NestedVRF `json:"vrf,omitempty"`
}

// Validate validates this available IP
func (m *AvailableIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
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

func (m *AvailableIP) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := validate.MinLength("address", "body", m.Address, 1); err != nil {
		return err
	}

	return nil
}

func (m *AvailableIP) validateVrf(formats strfmt.Registry) error {
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

// ContextValidate validate this available IP based on the context it is used
func (m *AvailableIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFamily(ctx, formats); err != nil {
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

func (m *AvailableIP) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "address", "body", string(m.Address)); err != nil {
		return err
	}

	return nil
}

func (m *AvailableIP) contextValidateFamily(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "family", "body", int64(m.Family)); err != nil {
		return err
	}

	return nil
}

func (m *AvailableIP) contextValidateVrf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AvailableIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailableIP) UnmarshalBinary(b []byte) error {
	var res AvailableIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
