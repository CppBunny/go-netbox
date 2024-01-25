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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimRackRolesReadReader is a Reader for the DcimRackRolesRead structure.
type DcimRackRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackRolesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackRolesReadOK creates a DcimRackRolesReadOK with default headers values
func NewDcimRackRolesReadOK() *DcimRackRolesReadOK {
	return &DcimRackRolesReadOK{}
}

/*
DcimRackRolesReadOK describes a response with status code 200, with default header values.

DcimRackRolesReadOK dcim rack roles read o k
*/
type DcimRackRolesReadOK struct {
	Payload *models.RackRole
}

// IsSuccess returns true when this dcim rack roles read o k response has a 2xx status code
func (o *DcimRackRolesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack roles read o k response has a 3xx status code
func (o *DcimRackRolesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack roles read o k response has a 4xx status code
func (o *DcimRackRolesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack roles read o k response has a 5xx status code
func (o *DcimRackRolesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack roles read o k response a status code equal to that given
func (o *DcimRackRolesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rack roles read o k response
func (o *DcimRackRolesReadOK) Code() int {
	return 200
}

func (o *DcimRackRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-roles/{id}/][%d] dcimRackRolesReadOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/rack-roles/{id}/][%d] dcimRackRolesReadOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesReadOK) GetPayload() *models.RackRole {
	return o.Payload
}

func (o *DcimRackRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackRolesReadDefault creates a DcimRackRolesReadDefault with default headers values
func NewDcimRackRolesReadDefault(code int) *DcimRackRolesReadDefault {
	return &DcimRackRolesReadDefault{
		_statusCode: code,
	}
}

/*
DcimRackRolesReadDefault describes a response with status code -1, with default header values.

DcimRackRolesReadDefault dcim rack roles read default
*/
type DcimRackRolesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim rack roles read default response has a 2xx status code
func (o *DcimRackRolesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rack roles read default response has a 3xx status code
func (o *DcimRackRolesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rack roles read default response has a 4xx status code
func (o *DcimRackRolesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rack roles read default response has a 5xx status code
func (o *DcimRackRolesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rack roles read default response a status code equal to that given
func (o *DcimRackRolesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim rack roles read default response
func (o *DcimRackRolesReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackRolesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-roles/{id}/][%d] dcim_rack-roles_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/rack-roles/{id}/][%d] dcim_rack-roles_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
