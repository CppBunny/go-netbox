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

package virtualization

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

	"github.com/CppBunny/go-netbox/netbox/models"
)

// NewVirtualizationVirtualDisksUpdateParams creates a new VirtualizationVirtualDisksUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationVirtualDisksUpdateParams() *VirtualizationVirtualDisksUpdateParams {
	return &VirtualizationVirtualDisksUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualDisksUpdateParamsWithTimeout creates a new VirtualizationVirtualDisksUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationVirtualDisksUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualDisksUpdateParams {
	return &VirtualizationVirtualDisksUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationVirtualDisksUpdateParamsWithContext creates a new VirtualizationVirtualDisksUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationVirtualDisksUpdateParamsWithContext(ctx context.Context) *VirtualizationVirtualDisksUpdateParams {
	return &VirtualizationVirtualDisksUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationVirtualDisksUpdateParamsWithHTTPClient creates a new VirtualizationVirtualDisksUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationVirtualDisksUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualDisksUpdateParams {
	return &VirtualizationVirtualDisksUpdateParams{
		HTTPClient: client,
	}
}

/*
VirtualizationVirtualDisksUpdateParams contains all the parameters to send to the API endpoint

	for the virtualization virtual disks update operation.

	Typically these are written to a http.Request.
*/
type VirtualizationVirtualDisksUpdateParams struct {

	// Data.
	Data *models.WritableVirtualDisk

	/* ID.

	   A unique integer value identifying this virtual chassis.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization virtual disks update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualDisksUpdateParams) WithDefaults() *VirtualizationVirtualDisksUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization virtual disks update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualDisksUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization virtual disks update params
func (o *VirtualizationVirtualDisksUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualDisksUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual disks update params
func (o *VirtualizationVirtualDisksUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual disks update params
func (o *VirtualizationVirtualDisksUpdateParams) WithContext(ctx context.Context) *VirtualizationVirtualDisksUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual disks update params
func (o *VirtualizationVirtualDisksUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual disks update params
func (o *VirtualizationVirtualDisksUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualDisksUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual disks update params
func (o *VirtualizationVirtualDisksUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization virtual disks update params
func (o *VirtualizationVirtualDisksUpdateParams) WithData(data *models.WritableVirtualDisk) *VirtualizationVirtualDisksUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization virtual disks update params
func (o *VirtualizationVirtualDisksUpdateParams) SetData(data *models.WritableVirtualDisk) {
	o.Data = data
}

// WithID adds the id to the virtualization virtual disks update params
func (o *VirtualizationVirtualDisksUpdateParams) WithID(id int64) *VirtualizationVirtualDisksUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization virtual disks update params
func (o *VirtualizationVirtualDisksUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualDisksUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
