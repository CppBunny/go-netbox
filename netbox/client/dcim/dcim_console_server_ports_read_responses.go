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

// DcimConsoleServerPortsReadReader is a Reader for the DcimConsoleServerPortsRead structure.
type DcimConsoleServerPortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortsReadOK creates a DcimConsoleServerPortsReadOK with default headers values
func NewDcimConsoleServerPortsReadOK() *DcimConsoleServerPortsReadOK {
	return &DcimConsoleServerPortsReadOK{}
}

/*
DcimConsoleServerPortsReadOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsReadOK dcim console server ports read o k
*/
type DcimConsoleServerPortsReadOK struct {
	Payload *models.ConsoleServerPort
}

// IsSuccess returns true when this dcim console server ports read o k response has a 2xx status code
func (o *DcimConsoleServerPortsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server ports read o k response has a 3xx status code
func (o *DcimConsoleServerPortsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server ports read o k response has a 4xx status code
func (o *DcimConsoleServerPortsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server ports read o k response has a 5xx status code
func (o *DcimConsoleServerPortsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server ports read o k response a status code equal to that given
func (o *DcimConsoleServerPortsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console server ports read o k response
func (o *DcimConsoleServerPortsReadOK) Code() int {
	return 200
}

func (o *DcimConsoleServerPortsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortsReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortsReadOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortsReadDefault creates a DcimConsoleServerPortsReadDefault with default headers values
func NewDcimConsoleServerPortsReadDefault(code int) *DcimConsoleServerPortsReadDefault {
	return &DcimConsoleServerPortsReadDefault{
		_statusCode: code,
	}
}

/*
DcimConsoleServerPortsReadDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortsReadDefault dcim console server ports read default
*/
type DcimConsoleServerPortsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim console server ports read default response has a 2xx status code
func (o *DcimConsoleServerPortsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console server ports read default response has a 3xx status code
func (o *DcimConsoleServerPortsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console server ports read default response has a 4xx status code
func (o *DcimConsoleServerPortsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console server ports read default response has a 5xx status code
func (o *DcimConsoleServerPortsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console server ports read default response a status code equal to that given
func (o *DcimConsoleServerPortsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim console server ports read default response
func (o *DcimConsoleServerPortsReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/console-server-ports/{id}/][%d] dcim_console-server-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/console-server-ports/{id}/][%d] dcim_console-server-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
