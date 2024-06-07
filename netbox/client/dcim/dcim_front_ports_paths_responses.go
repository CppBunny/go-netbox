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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CppBunny/go-netbox/netbox/models"
)

// DcimFrontPortsPathsReader is a Reader for the DcimFrontPortsPaths structure.
type DcimFrontPortsPathsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsPathsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsPathsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortsPathsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortsPathsOK creates a DcimFrontPortsPathsOK with default headers values
func NewDcimFrontPortsPathsOK() *DcimFrontPortsPathsOK {
	return &DcimFrontPortsPathsOK{}
}

/*
DcimFrontPortsPathsOK describes a response with status code 200, with default header values.

DcimFrontPortsPathsOK dcim front ports paths o k
*/
type DcimFrontPortsPathsOK struct {
	Payload *models.FrontPort
}

// IsSuccess returns true when this dcim front ports paths o k response has a 2xx status code
func (o *DcimFrontPortsPathsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front ports paths o k response has a 3xx status code
func (o *DcimFrontPortsPathsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front ports paths o k response has a 4xx status code
func (o *DcimFrontPortsPathsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front ports paths o k response has a 5xx status code
func (o *DcimFrontPortsPathsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front ports paths o k response a status code equal to that given
func (o *DcimFrontPortsPathsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim front ports paths o k response
func (o *DcimFrontPortsPathsOK) Code() int {
	return 200
}

func (o *DcimFrontPortsPathsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/paths/][%d] dcimFrontPortsPathsOK %s", 200, payload)
}

func (o *DcimFrontPortsPathsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/paths/][%d] dcimFrontPortsPathsOK %s", 200, payload)
}

func (o *DcimFrontPortsPathsOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsPathsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortsPathsDefault creates a DcimFrontPortsPathsDefault with default headers values
func NewDcimFrontPortsPathsDefault(code int) *DcimFrontPortsPathsDefault {
	return &DcimFrontPortsPathsDefault{
		_statusCode: code,
	}
}

/*
DcimFrontPortsPathsDefault describes a response with status code -1, with default header values.

DcimFrontPortsPathsDefault dcim front ports paths default
*/
type DcimFrontPortsPathsDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim front ports paths default response has a 2xx status code
func (o *DcimFrontPortsPathsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim front ports paths default response has a 3xx status code
func (o *DcimFrontPortsPathsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim front ports paths default response has a 4xx status code
func (o *DcimFrontPortsPathsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim front ports paths default response has a 5xx status code
func (o *DcimFrontPortsPathsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim front ports paths default response a status code equal to that given
func (o *DcimFrontPortsPathsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim front ports paths default response
func (o *DcimFrontPortsPathsDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortsPathsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/paths/][%d] dcim_front-ports_paths default %s", o._statusCode, payload)
}

func (o *DcimFrontPortsPathsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/paths/][%d] dcim_front-ports_paths default %s", o._statusCode, payload)
}

func (o *DcimFrontPortsPathsDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortsPathsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
