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

// NestedModuleType nested module type
//
// swagger:model NestedModuleType
type NestedModuleType struct {

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// manufacturer
	Manufacturer *NestedManufacturer `json:"manufacturer,omitempty"`

	// Model
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Model *string `json:"model"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this nested module type
func (m *NestedModuleType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManufacturer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedModuleType) validateManufacturer(formats strfmt.Registry) error {
	if swag.IsZero(m.Manufacturer) { // not required
		return nil
	}

	if m.Manufacturer != nil {
		if err := m.Manufacturer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *NestedModuleType) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	if err := validate.MinLength("model", "body", *m.Model, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("model", "body", *m.Model, 100); err != nil {
		return err
	}

	return nil
}

func (m *NestedModuleType) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nested module type based on the context it is used
func (m *NestedModuleType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManufacturer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedModuleType) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *NestedModuleType) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NestedModuleType) contextValidateManufacturer(ctx context.Context, formats strfmt.Registry) error {

	if m.Manufacturer != nil {

		if swag.IsZero(m.Manufacturer) { // not required
			return nil
		}

		if err := m.Manufacturer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *NestedModuleType) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedModuleType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedModuleType) UnmarshalBinary(b []byte) error {
	var res NestedModuleType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
