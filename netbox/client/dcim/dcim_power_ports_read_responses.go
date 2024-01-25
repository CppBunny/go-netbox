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

// DcimPowerPortsReadReader is a Reader for the DcimPowerPortsRead structure.
type DcimPowerPortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortsReadOK creates a DcimPowerPortsReadOK with default headers values
func NewDcimPowerPortsReadOK() *DcimPowerPortsReadOK {
	return &DcimPowerPortsReadOK{}
}

/*
DcimPowerPortsReadOK describes a response with status code 200, with default header values.

DcimPowerPortsReadOK dcim power ports read o k
*/
type DcimPowerPortsReadOK struct {
	Payload *models.PowerPort
}

// IsSuccess returns true when this dcim power ports read o k response has a 2xx status code
func (o *DcimPowerPortsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power ports read o k response has a 3xx status code
func (o *DcimPowerPortsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power ports read o k response has a 4xx status code
func (o *DcimPowerPortsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power ports read o k response has a 5xx status code
func (o *DcimPowerPortsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power ports read o k response a status code equal to that given
func (o *DcimPowerPortsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power ports read o k response
func (o *DcimPowerPortsReadOK) Code() int {
	return 200
}

func (o *DcimPowerPortsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/][%d] dcimPowerPortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/][%d] dcimPowerPortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsReadOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortsReadDefault creates a DcimPowerPortsReadDefault with default headers values
func NewDcimPowerPortsReadDefault(code int) *DcimPowerPortsReadDefault {
	return &DcimPowerPortsReadDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortsReadDefault describes a response with status code -1, with default header values.

DcimPowerPortsReadDefault dcim power ports read default
*/
type DcimPowerPortsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power ports read default response has a 2xx status code
func (o *DcimPowerPortsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power ports read default response has a 3xx status code
func (o *DcimPowerPortsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power ports read default response has a 4xx status code
func (o *DcimPowerPortsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power ports read default response has a 5xx status code
func (o *DcimPowerPortsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power ports read default response a status code equal to that given
func (o *DcimPowerPortsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power ports read default response
func (o *DcimPowerPortsReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/][%d] dcim_power-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortsReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/][%d] dcim_power-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
