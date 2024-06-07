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

// DcimPowerPortsPartialUpdateReader is a Reader for the DcimPowerPortsPartialUpdate structure.
type DcimPowerPortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortsPartialUpdateOK creates a DcimPowerPortsPartialUpdateOK with default headers values
func NewDcimPowerPortsPartialUpdateOK() *DcimPowerPortsPartialUpdateOK {
	return &DcimPowerPortsPartialUpdateOK{}
}

/*
DcimPowerPortsPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortsPartialUpdateOK dcim power ports partial update o k
*/
type DcimPowerPortsPartialUpdateOK struct {
	Payload *models.PowerPort
}

// IsSuccess returns true when this dcim power ports partial update o k response has a 2xx status code
func (o *DcimPowerPortsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power ports partial update o k response has a 3xx status code
func (o *DcimPowerPortsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power ports partial update o k response has a 4xx status code
func (o *DcimPowerPortsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power ports partial update o k response has a 5xx status code
func (o *DcimPowerPortsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power ports partial update o k response a status code equal to that given
func (o *DcimPowerPortsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power ports partial update o k response
func (o *DcimPowerPortsPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerPortsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/power-ports/{id}/][%d] dcimPowerPortsPartialUpdateOK %s", 200, payload)
}

func (o *DcimPowerPortsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/power-ports/{id}/][%d] dcimPowerPortsPartialUpdateOK %s", 200, payload)
}

func (o *DcimPowerPortsPartialUpdateOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortsPartialUpdateDefault creates a DcimPowerPortsPartialUpdateDefault with default headers values
func NewDcimPowerPortsPartialUpdateDefault(code int) *DcimPowerPortsPartialUpdateDefault {
	return &DcimPowerPortsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerPortsPartialUpdateDefault dcim power ports partial update default
*/
type DcimPowerPortsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power ports partial update default response has a 2xx status code
func (o *DcimPowerPortsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power ports partial update default response has a 3xx status code
func (o *DcimPowerPortsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power ports partial update default response has a 4xx status code
func (o *DcimPowerPortsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power ports partial update default response has a 5xx status code
func (o *DcimPowerPortsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power ports partial update default response a status code equal to that given
func (o *DcimPowerPortsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power ports partial update default response
func (o *DcimPowerPortsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortsPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/power-ports/{id}/][%d] dcim_power-ports_partial_update default %s", o._statusCode, payload)
}

func (o *DcimPowerPortsPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/power-ports/{id}/][%d] dcim_power-ports_partial_update default %s", o._statusCode, payload)
}

func (o *DcimPowerPortsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
