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

	"github.com/CppBunny/go-netbox/netbox/models"
)

// NewVpnTunnelTerminationsCreateParams creates a new VpnTunnelTerminationsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVpnTunnelTerminationsCreateParams() *VpnTunnelTerminationsCreateParams {
	return &VpnTunnelTerminationsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVpnTunnelTerminationsCreateParamsWithTimeout creates a new VpnTunnelTerminationsCreateParams object
// with the ability to set a timeout on a request.
func NewVpnTunnelTerminationsCreateParamsWithTimeout(timeout time.Duration) *VpnTunnelTerminationsCreateParams {
	return &VpnTunnelTerminationsCreateParams{
		timeout: timeout,
	}
}

// NewVpnTunnelTerminationsCreateParamsWithContext creates a new VpnTunnelTerminationsCreateParams object
// with the ability to set a context for a request.
func NewVpnTunnelTerminationsCreateParamsWithContext(ctx context.Context) *VpnTunnelTerminationsCreateParams {
	return &VpnTunnelTerminationsCreateParams{
		Context: ctx,
	}
}

// NewVpnTunnelTerminationsCreateParamsWithHTTPClient creates a new VpnTunnelTerminationsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVpnTunnelTerminationsCreateParamsWithHTTPClient(client *http.Client) *VpnTunnelTerminationsCreateParams {
	return &VpnTunnelTerminationsCreateParams{
		HTTPClient: client,
	}
}

/*
VpnTunnelTerminationsCreateParams contains all the parameters to send to the API endpoint

	for the vpn tunnel terminations create operation.

	Typically these are written to a http.Request.
*/
type VpnTunnelTerminationsCreateParams struct {

	// Data.
	Data *models.WritableTunnelTermination

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vpn tunnel terminations create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VpnTunnelTerminationsCreateParams) WithDefaults() *VpnTunnelTerminationsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vpn tunnel terminations create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VpnTunnelTerminationsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vpn tunnel terminations create params
func (o *VpnTunnelTerminationsCreateParams) WithTimeout(timeout time.Duration) *VpnTunnelTerminationsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vpn tunnel terminations create params
func (o *VpnTunnelTerminationsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vpn tunnel terminations create params
func (o *VpnTunnelTerminationsCreateParams) WithContext(ctx context.Context) *VpnTunnelTerminationsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vpn tunnel terminations create params
func (o *VpnTunnelTerminationsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vpn tunnel terminations create params
func (o *VpnTunnelTerminationsCreateParams) WithHTTPClient(client *http.Client) *VpnTunnelTerminationsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vpn tunnel terminations create params
func (o *VpnTunnelTerminationsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the vpn tunnel terminations create params
func (o *VpnTunnelTerminationsCreateParams) WithData(data *models.WritableTunnelTermination) *VpnTunnelTerminationsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the vpn tunnel terminations create params
func (o *VpnTunnelTerminationsCreateParams) SetData(data *models.WritableTunnelTermination) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VpnTunnelTerminationsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
