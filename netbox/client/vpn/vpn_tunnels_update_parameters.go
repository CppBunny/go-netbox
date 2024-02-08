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

package vpn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewVpnTunnelsUpdateParams creates a new VpnTunnelsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVpnTunnelsUpdateParams() *VpnTunnelsUpdateParams {
	return &VpnTunnelsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVpnTunnelsUpdateParamsWithTimeout creates a new VpnTunnelsUpdateParams object
// with the ability to set a timeout on a request.
func NewVpnTunnelsUpdateParamsWithTimeout(timeout time.Duration) *VpnTunnelsUpdateParams {
	return &VpnTunnelsUpdateParams{
		timeout: timeout,
	}
}

// NewVpnTunnelsUpdateParamsWithContext creates a new VpnTunnelsUpdateParams object
// with the ability to set a context for a request.
func NewVpnTunnelsUpdateParamsWithContext(ctx context.Context) *VpnTunnelsUpdateParams {
	return &VpnTunnelsUpdateParams{
		Context: ctx,
	}
}

// NewVpnTunnelsUpdateParamsWithHTTPClient creates a new VpnTunnelsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVpnTunnelsUpdateParamsWithHTTPClient(client *http.Client) *VpnTunnelsUpdateParams {
	return &VpnTunnelsUpdateParams{
		HTTPClient: client,
	}
}

/*
VpnTunnelsUpdateParams contains all the parameters to send to the API endpoint

	for the vpn tunnels update operation.

	Typically these are written to a http.Request.
*/
type VpnTunnelsUpdateParams struct {

	// Data.
	Data *models.WritableTunnel

	/* ID.

	   A unique integer value identifying this tunnel.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vpn tunnels update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VpnTunnelsUpdateParams) WithDefaults() *VpnTunnelsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vpn tunnels update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VpnTunnelsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vpn tunnels update params
func (o *VpnTunnelsUpdateParams) WithTimeout(timeout time.Duration) *VpnTunnelsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vpn tunnels update params
func (o *VpnTunnelsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vpn tunnels update params
func (o *VpnTunnelsUpdateParams) WithContext(ctx context.Context) *VpnTunnelsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vpn tunnels update params
func (o *VpnTunnelsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vpn tunnels update params
func (o *VpnTunnelsUpdateParams) WithHTTPClient(client *http.Client) *VpnTunnelsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vpn tunnels update params
func (o *VpnTunnelsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the vpn tunnels update params
func (o *VpnTunnelsUpdateParams) WithData(data *models.WritableTunnel) *VpnTunnelsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the vpn tunnels update params
func (o *VpnTunnelsUpdateParams) SetData(data *models.WritableTunnel) {
	o.Data = data
}

// WithID adds the id to the vpn tunnels update params
func (o *VpnTunnelsUpdateParams) WithID(id int64) *VpnTunnelsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the vpn tunnels update params
func (o *VpnTunnelsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VpnTunnelsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}