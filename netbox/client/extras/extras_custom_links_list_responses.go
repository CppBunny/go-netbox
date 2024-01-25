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
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasCustomLinksListReader is a Reader for the ExtrasCustomLinksList structure.
type ExtrasCustomLinksListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomLinksListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomLinksListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomLinksListOK creates a ExtrasCustomLinksListOK with default headers values
func NewExtrasCustomLinksListOK() *ExtrasCustomLinksListOK {
	return &ExtrasCustomLinksListOK{}
}

/*
ExtrasCustomLinksListOK describes a response with status code 200, with default header values.

ExtrasCustomLinksListOK extras custom links list o k
*/
type ExtrasCustomLinksListOK struct {
	Payload *ExtrasCustomLinksListOKBody
}

// IsSuccess returns true when this extras custom links list o k response has a 2xx status code
func (o *ExtrasCustomLinksListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom links list o k response has a 3xx status code
func (o *ExtrasCustomLinksListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom links list o k response has a 4xx status code
func (o *ExtrasCustomLinksListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom links list o k response has a 5xx status code
func (o *ExtrasCustomLinksListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom links list o k response a status code equal to that given
func (o *ExtrasCustomLinksListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras custom links list o k response
func (o *ExtrasCustomLinksListOK) Code() int {
	return 200
}

func (o *ExtrasCustomLinksListOK) Error() string {
	return fmt.Sprintf("[GET /extras/custom-links/][%d] extrasCustomLinksListOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomLinksListOK) String() string {
	return fmt.Sprintf("[GET /extras/custom-links/][%d] extrasCustomLinksListOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomLinksListOK) GetPayload() *ExtrasCustomLinksListOKBody {
	return o.Payload
}

func (o *ExtrasCustomLinksListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExtrasCustomLinksListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomLinksListDefault creates a ExtrasCustomLinksListDefault with default headers values
func NewExtrasCustomLinksListDefault(code int) *ExtrasCustomLinksListDefault {
	return &ExtrasCustomLinksListDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomLinksListDefault describes a response with status code -1, with default header values.

ExtrasCustomLinksListDefault extras custom links list default
*/
type ExtrasCustomLinksListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras custom links list default response has a 2xx status code
func (o *ExtrasCustomLinksListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom links list default response has a 3xx status code
func (o *ExtrasCustomLinksListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom links list default response has a 4xx status code
func (o *ExtrasCustomLinksListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom links list default response has a 5xx status code
func (o *ExtrasCustomLinksListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom links list default response a status code equal to that given
func (o *ExtrasCustomLinksListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras custom links list default response
func (o *ExtrasCustomLinksListDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasCustomLinksListDefault) Error() string {
	return fmt.Sprintf("[GET /extras/custom-links/][%d] extras_custom-links_list default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomLinksListDefault) String() string {
	return fmt.Sprintf("[GET /extras/custom-links/][%d] extras_custom-links_list default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomLinksListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomLinksListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ExtrasCustomLinksListOKBody extras custom links list o k body
swagger:model ExtrasCustomLinksListOKBody
*/
type ExtrasCustomLinksListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.CustomLink `json:"results"`
}

// Validate validates this extras custom links list o k body
func (o *ExtrasCustomLinksListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExtrasCustomLinksListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("extrasCustomLinksListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasCustomLinksListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasCustomLinksListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasCustomLinksListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasCustomLinksListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasCustomLinksListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("extrasCustomLinksListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasCustomLinksListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasCustomLinksListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this extras custom links list o k body based on the context it is used
func (o *ExtrasCustomLinksListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExtrasCustomLinksListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasCustomLinksListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasCustomLinksListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExtrasCustomLinksListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExtrasCustomLinksListOKBody) UnmarshalBinary(b []byte) error {
	var res ExtrasCustomLinksListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
