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

// ExtrasObjectChangesListReader is a Reader for the ExtrasObjectChangesList structure.
type ExtrasObjectChangesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasObjectChangesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasObjectChangesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasObjectChangesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasObjectChangesListOK creates a ExtrasObjectChangesListOK with default headers values
func NewExtrasObjectChangesListOK() *ExtrasObjectChangesListOK {
	return &ExtrasObjectChangesListOK{}
}

/*
ExtrasObjectChangesListOK describes a response with status code 200, with default header values.

ExtrasObjectChangesListOK extras object changes list o k
*/
type ExtrasObjectChangesListOK struct {
	Payload *ExtrasObjectChangesListOKBody
}

// IsSuccess returns true when this extras object changes list o k response has a 2xx status code
func (o *ExtrasObjectChangesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras object changes list o k response has a 3xx status code
func (o *ExtrasObjectChangesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras object changes list o k response has a 4xx status code
func (o *ExtrasObjectChangesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras object changes list o k response has a 5xx status code
func (o *ExtrasObjectChangesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras object changes list o k response a status code equal to that given
func (o *ExtrasObjectChangesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras object changes list o k response
func (o *ExtrasObjectChangesListOK) Code() int {
	return 200
}

func (o *ExtrasObjectChangesListOK) Error() string {
	return fmt.Sprintf("[GET /extras/object-changes/][%d] extrasObjectChangesListOK  %+v", 200, o.Payload)
}

func (o *ExtrasObjectChangesListOK) String() string {
	return fmt.Sprintf("[GET /extras/object-changes/][%d] extrasObjectChangesListOK  %+v", 200, o.Payload)
}

func (o *ExtrasObjectChangesListOK) GetPayload() *ExtrasObjectChangesListOKBody {
	return o.Payload
}

func (o *ExtrasObjectChangesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExtrasObjectChangesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasObjectChangesListDefault creates a ExtrasObjectChangesListDefault with default headers values
func NewExtrasObjectChangesListDefault(code int) *ExtrasObjectChangesListDefault {
	return &ExtrasObjectChangesListDefault{
		_statusCode: code,
	}
}

/*
ExtrasObjectChangesListDefault describes a response with status code -1, with default header values.

ExtrasObjectChangesListDefault extras object changes list default
*/
type ExtrasObjectChangesListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras object changes list default response has a 2xx status code
func (o *ExtrasObjectChangesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras object changes list default response has a 3xx status code
func (o *ExtrasObjectChangesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras object changes list default response has a 4xx status code
func (o *ExtrasObjectChangesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras object changes list default response has a 5xx status code
func (o *ExtrasObjectChangesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras object changes list default response a status code equal to that given
func (o *ExtrasObjectChangesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras object changes list default response
func (o *ExtrasObjectChangesListDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasObjectChangesListDefault) Error() string {
	return fmt.Sprintf("[GET /extras/object-changes/][%d] extras_object-changes_list default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasObjectChangesListDefault) String() string {
	return fmt.Sprintf("[GET /extras/object-changes/][%d] extras_object-changes_list default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasObjectChangesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasObjectChangesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ExtrasObjectChangesListOKBody extras object changes list o k body
swagger:model ExtrasObjectChangesListOKBody
*/
type ExtrasObjectChangesListOKBody struct {

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
	Results []*models.ObjectChange `json:"results"`
}

// Validate validates this extras object changes list o k body
func (o *ExtrasObjectChangesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ExtrasObjectChangesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("extrasObjectChangesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasObjectChangesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasObjectChangesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasObjectChangesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasObjectChangesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasObjectChangesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("extrasObjectChangesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasObjectChangesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasObjectChangesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this extras object changes list o k body based on the context it is used
func (o *ExtrasObjectChangesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExtrasObjectChangesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasObjectChangesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasObjectChangesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExtrasObjectChangesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExtrasObjectChangesListOKBody) UnmarshalBinary(b []byte) error {
	var res ExtrasObjectChangesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
