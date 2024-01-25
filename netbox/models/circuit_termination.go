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

// CircuitTermination circuit termination
//
// swagger:model CircuitTermination
type CircuitTermination struct {

	// occupied
	// Read Only: true
	Occupied *bool `json:"_occupied,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable end
	// Read Only: true
	// Min Length: 1
	CableEnd string `json:"cable_end,omitempty"`

	// circuit
	// Required: true
	Circuit *NestedCircuit `json:"circuit"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	//
	// Return the appropriate serializer for the link termination model.
	//
	// Read Only: true
	LinkPeers []interface{} `json:"link_peers"`

	// Link peers type
	// Read Only: true
	LinkPeersType string `json:"link_peers_type,omitempty"`

	// Mark connected
	//
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty"`

	// Port speed (Kbps)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	PortSpeed *int64 `json:"port_speed,omitempty"`

	// Patch panel/port(s)
	// Max Length: 100
	PpInfo string `json:"pp_info,omitempty"`

	// provider network
	ProviderNetwork *NestedProviderNetwork `json:"provider_network,omitempty"`

	// site
	Site *NestedSite `json:"site,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Termination
	// Required: true
	// Enum: [A Z]
	TermSide *string `json:"term_side"`

	// Upstream speed (Kbps)
	//
	// Upstream speed, if different from port speed
	// Maximum: 2.147483647e+09
	// Minimum: 0
	UpstreamSpeed *int64 `json:"upstream_speed,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Cross-connect ID
	// Max Length: 50
	XconnectID string `json:"xconnect_id,omitempty"`
}

// Validate validates this circuit termination
func (m *CircuitTermination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCableEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCircuit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePpInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTermSide(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpstreamSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXconnectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CircuitTermination) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *CircuitTermination) validateCableEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.CableEnd) { // not required
		return nil
	}

	if err := validate.MinLength("cable_end", "body", m.CableEnd, 1); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) validateCircuit(formats strfmt.Registry) error {

	if err := validate.Required("circuit", "body", m.Circuit); err != nil {
		return err
	}

	if m.Circuit != nil {
		if err := m.Circuit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("circuit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("circuit")
			}
			return err
		}
	}

	return nil
}

func (m *CircuitTermination) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) validatePortSpeed(formats strfmt.Registry) error {
	if swag.IsZero(m.PortSpeed) { // not required
		return nil
	}

	if err := validate.MinimumInt("port_speed", "body", *m.PortSpeed, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port_speed", "body", *m.PortSpeed, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) validatePpInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.PpInfo) { // not required
		return nil
	}

	if err := validate.MaxLength("pp_info", "body", m.PpInfo, 100); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) validateProviderNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.ProviderNetwork) { // not required
		return nil
	}

	if m.ProviderNetwork != nil {
		if err := m.ProviderNetwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider_network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider_network")
			}
			return err
		}
	}

	return nil
}

func (m *CircuitTermination) validateSite(formats strfmt.Registry) error {
	if swag.IsZero(m.Site) { // not required
		return nil
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *CircuitTermination) validateTags(formats strfmt.Registry) error {
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

var circuitTerminationTypeTermSidePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","Z"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		circuitTerminationTypeTermSidePropEnum = append(circuitTerminationTypeTermSidePropEnum, v)
	}
}

const (

	// CircuitTerminationTermSideA captures enum value "A"
	CircuitTerminationTermSideA string = "A"

	// CircuitTerminationTermSideZ captures enum value "Z"
	CircuitTerminationTermSideZ string = "Z"
)

// prop value enum
func (m *CircuitTermination) validateTermSideEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, circuitTerminationTypeTermSidePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CircuitTermination) validateTermSide(formats strfmt.Registry) error {

	if err := validate.Required("term_side", "body", m.TermSide); err != nil {
		return err
	}

	// value enum
	if err := m.validateTermSideEnum("term_side", "body", *m.TermSide); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) validateUpstreamSpeed(formats strfmt.Registry) error {
	if swag.IsZero(m.UpstreamSpeed) { // not required
		return nil
	}

	if err := validate.MinimumInt("upstream_speed", "body", *m.UpstreamSpeed, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("upstream_speed", "body", *m.UpstreamSpeed, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) validateXconnectID(formats strfmt.Registry) error {
	if swag.IsZero(m.XconnectID) { // not required
		return nil
	}

	if err := validate.MaxLength("xconnect_id", "body", m.XconnectID, 50); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this circuit termination based on the context it is used
func (m *CircuitTermination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOccupied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCableEnd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCircuit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkPeers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkPeersType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProviderNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
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

func (m *CircuitTermination) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {

		if swag.IsZero(m.Cable) { // not required
			return nil
		}

		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *CircuitTermination) contextValidateCableEnd(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_end", "body", string(m.CableEnd)); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) contextValidateCircuit(ctx context.Context, formats strfmt.Registry) error {

	if m.Circuit != nil {

		if err := m.Circuit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("circuit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("circuit")
			}
			return err
		}
	}

	return nil
}

func (m *CircuitTermination) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) contextValidateLinkPeers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peers", "body", []interface{}(m.LinkPeers)); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) contextValidateLinkPeersType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peers_type", "body", string(m.LinkPeersType)); err != nil {
		return err
	}

	return nil
}

func (m *CircuitTermination) contextValidateProviderNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.ProviderNetwork != nil {

		if swag.IsZero(m.ProviderNetwork) { // not required
			return nil
		}

		if err := m.ProviderNetwork.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider_network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provider_network")
			}
			return err
		}
	}

	return nil
}

func (m *CircuitTermination) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {

		if swag.IsZero(m.Site) { // not required
			return nil
		}

		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *CircuitTermination) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CircuitTermination) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CircuitTermination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CircuitTermination) UnmarshalBinary(b []byte) error {
	var res CircuitTermination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
