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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// TenancyContactRolesUpdateReader is a Reader for the TenancyContactRolesUpdate structure.
type TenancyContactRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactRolesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactRolesUpdateOK creates a TenancyContactRolesUpdateOK with default headers values
func NewTenancyContactRolesUpdateOK() *TenancyContactRolesUpdateOK {
	return &TenancyContactRolesUpdateOK{}
}

/*
TenancyContactRolesUpdateOK describes a response with status code 200, with default header values.

TenancyContactRolesUpdateOK tenancy contact roles update o k
*/
type TenancyContactRolesUpdateOK struct {
	Payload *models.ContactRole
}

// IsSuccess returns true when this tenancy contact roles update o k response has a 2xx status code
func (o *TenancyContactRolesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact roles update o k response has a 3xx status code
func (o *TenancyContactRolesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact roles update o k response has a 4xx status code
func (o *TenancyContactRolesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact roles update o k response has a 5xx status code
func (o *TenancyContactRolesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact roles update o k response a status code equal to that given
func (o *TenancyContactRolesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contact roles update o k response
func (o *TenancyContactRolesUpdateOK) Code() int {
	return 200
}

func (o *TenancyContactRolesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/{id}/][%d] tenancyContactRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactRolesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/{id}/][%d] tenancyContactRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactRolesUpdateOK) GetPayload() *models.ContactRole {
	return o.Payload
}

func (o *TenancyContactRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactRolesUpdateDefault creates a TenancyContactRolesUpdateDefault with default headers values
func NewTenancyContactRolesUpdateDefault(code int) *TenancyContactRolesUpdateDefault {
	return &TenancyContactRolesUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactRolesUpdateDefault describes a response with status code -1, with default header values.

TenancyContactRolesUpdateDefault tenancy contact roles update default
*/
type TenancyContactRolesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contact roles update default response has a 2xx status code
func (o *TenancyContactRolesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact roles update default response has a 3xx status code
func (o *TenancyContactRolesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact roles update default response has a 4xx status code
func (o *TenancyContactRolesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact roles update default response has a 5xx status code
func (o *TenancyContactRolesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact roles update default response a status code equal to that given
func (o *TenancyContactRolesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contact roles update default response
func (o *TenancyContactRolesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactRolesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/{id}/][%d] tenancy_contact-roles_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/{id}/][%d] tenancy_contact-roles_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactRolesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
