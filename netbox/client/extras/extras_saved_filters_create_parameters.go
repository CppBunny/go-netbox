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

package extras

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

// NewExtrasSavedFiltersCreateParams creates a new ExtrasSavedFiltersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasSavedFiltersCreateParams() *ExtrasSavedFiltersCreateParams {
	return &ExtrasSavedFiltersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasSavedFiltersCreateParamsWithTimeout creates a new ExtrasSavedFiltersCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasSavedFiltersCreateParamsWithTimeout(timeout time.Duration) *ExtrasSavedFiltersCreateParams {
	return &ExtrasSavedFiltersCreateParams{
		timeout: timeout,
	}
}

// NewExtrasSavedFiltersCreateParamsWithContext creates a new ExtrasSavedFiltersCreateParams object
// with the ability to set a context for a request.
func NewExtrasSavedFiltersCreateParamsWithContext(ctx context.Context) *ExtrasSavedFiltersCreateParams {
	return &ExtrasSavedFiltersCreateParams{
		Context: ctx,
	}
}

// NewExtrasSavedFiltersCreateParamsWithHTTPClient creates a new ExtrasSavedFiltersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasSavedFiltersCreateParamsWithHTTPClient(client *http.Client) *ExtrasSavedFiltersCreateParams {
	return &ExtrasSavedFiltersCreateParams{
		HTTPClient: client,
	}
}

/*
ExtrasSavedFiltersCreateParams contains all the parameters to send to the API endpoint

	for the extras saved filters create operation.

	Typically these are written to a http.Request.
*/
type ExtrasSavedFiltersCreateParams struct {

	// Data.
	Data *models.SavedFilter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras saved filters create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasSavedFiltersCreateParams) WithDefaults() *ExtrasSavedFiltersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras saved filters create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasSavedFiltersCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras saved filters create params
func (o *ExtrasSavedFiltersCreateParams) WithTimeout(timeout time.Duration) *ExtrasSavedFiltersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras saved filters create params
func (o *ExtrasSavedFiltersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras saved filters create params
func (o *ExtrasSavedFiltersCreateParams) WithContext(ctx context.Context) *ExtrasSavedFiltersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras saved filters create params
func (o *ExtrasSavedFiltersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras saved filters create params
func (o *ExtrasSavedFiltersCreateParams) WithHTTPClient(client *http.Client) *ExtrasSavedFiltersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras saved filters create params
func (o *ExtrasSavedFiltersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras saved filters create params
func (o *ExtrasSavedFiltersCreateParams) WithData(data *models.SavedFilter) *ExtrasSavedFiltersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras saved filters create params
func (o *ExtrasSavedFiltersCreateParams) SetData(data *models.SavedFilter) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasSavedFiltersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
