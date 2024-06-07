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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CppBunny/go-netbox/netbox/models"
)

// TenancyContactAssignmentsUpdateReader is a Reader for the TenancyContactAssignmentsUpdate structure.
type TenancyContactAssignmentsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactAssignmentsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactAssignmentsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactAssignmentsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactAssignmentsUpdateOK creates a TenancyContactAssignmentsUpdateOK with default headers values
func NewTenancyContactAssignmentsUpdateOK() *TenancyContactAssignmentsUpdateOK {
	return &TenancyContactAssignmentsUpdateOK{}
}

/*
TenancyContactAssignmentsUpdateOK describes a response with status code 200, with default header values.

TenancyContactAssignmentsUpdateOK tenancy contact assignments update o k
*/
type TenancyContactAssignmentsUpdateOK struct {
	Payload *models.ContactAssignment
}

// IsSuccess returns true when this tenancy contact assignments update o k response has a 2xx status code
func (o *TenancyContactAssignmentsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact assignments update o k response has a 3xx status code
func (o *TenancyContactAssignmentsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact assignments update o k response has a 4xx status code
func (o *TenancyContactAssignmentsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact assignments update o k response has a 5xx status code
func (o *TenancyContactAssignmentsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact assignments update o k response a status code equal to that given
func (o *TenancyContactAssignmentsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contact assignments update o k response
func (o *TenancyContactAssignmentsUpdateOK) Code() int {
	return 200
}

func (o *TenancyContactAssignmentsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/contact-assignments/{id}/][%d] tenancyContactAssignmentsUpdateOK %s", 200, payload)
}

func (o *TenancyContactAssignmentsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/contact-assignments/{id}/][%d] tenancyContactAssignmentsUpdateOK %s", 200, payload)
}

func (o *TenancyContactAssignmentsUpdateOK) GetPayload() *models.ContactAssignment {
	return o.Payload
}

func (o *TenancyContactAssignmentsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactAssignmentsUpdateDefault creates a TenancyContactAssignmentsUpdateDefault with default headers values
func NewTenancyContactAssignmentsUpdateDefault(code int) *TenancyContactAssignmentsUpdateDefault {
	return &TenancyContactAssignmentsUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactAssignmentsUpdateDefault describes a response with status code -1, with default header values.

TenancyContactAssignmentsUpdateDefault tenancy contact assignments update default
*/
type TenancyContactAssignmentsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contact assignments update default response has a 2xx status code
func (o *TenancyContactAssignmentsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact assignments update default response has a 3xx status code
func (o *TenancyContactAssignmentsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact assignments update default response has a 4xx status code
func (o *TenancyContactAssignmentsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact assignments update default response has a 5xx status code
func (o *TenancyContactAssignmentsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact assignments update default response a status code equal to that given
func (o *TenancyContactAssignmentsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contact assignments update default response
func (o *TenancyContactAssignmentsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactAssignmentsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/contact-assignments/{id}/][%d] tenancy_contact-assignments_update default %s", o._statusCode, payload)
}

func (o *TenancyContactAssignmentsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/contact-assignments/{id}/][%d] tenancy_contact-assignments_update default %s", o._statusCode, payload)
}

func (o *TenancyContactAssignmentsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactAssignmentsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
