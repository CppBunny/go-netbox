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

// DcimCableTerminationsPartialUpdateReader is a Reader for the DcimCableTerminationsPartialUpdate structure.
type DcimCableTerminationsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCableTerminationsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCableTerminationsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCableTerminationsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCableTerminationsPartialUpdateOK creates a DcimCableTerminationsPartialUpdateOK with default headers values
func NewDcimCableTerminationsPartialUpdateOK() *DcimCableTerminationsPartialUpdateOK {
	return &DcimCableTerminationsPartialUpdateOK{}
}

/*
DcimCableTerminationsPartialUpdateOK describes a response with status code 200, with default header values.

DcimCableTerminationsPartialUpdateOK dcim cable terminations partial update o k
*/
type DcimCableTerminationsPartialUpdateOK struct {
	Payload *models.CableTermination
}

// IsSuccess returns true when this dcim cable terminations partial update o k response has a 2xx status code
func (o *DcimCableTerminationsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cable terminations partial update o k response has a 3xx status code
func (o *DcimCableTerminationsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cable terminations partial update o k response has a 4xx status code
func (o *DcimCableTerminationsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cable terminations partial update o k response has a 5xx status code
func (o *DcimCableTerminationsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cable terminations partial update o k response a status code equal to that given
func (o *DcimCableTerminationsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim cable terminations partial update o k response
func (o *DcimCableTerminationsPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimCableTerminationsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsPartialUpdateOK %s", 200, payload)
}

func (o *DcimCableTerminationsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsPartialUpdateOK %s", 200, payload)
}

func (o *DcimCableTerminationsPartialUpdateOK) GetPayload() *models.CableTermination {
	return o.Payload
}

func (o *DcimCableTerminationsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CableTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCableTerminationsPartialUpdateDefault creates a DcimCableTerminationsPartialUpdateDefault with default headers values
func NewDcimCableTerminationsPartialUpdateDefault(code int) *DcimCableTerminationsPartialUpdateDefault {
	return &DcimCableTerminationsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimCableTerminationsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimCableTerminationsPartialUpdateDefault dcim cable terminations partial update default
*/
type DcimCableTerminationsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim cable terminations partial update default response has a 2xx status code
func (o *DcimCableTerminationsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim cable terminations partial update default response has a 3xx status code
func (o *DcimCableTerminationsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim cable terminations partial update default response has a 4xx status code
func (o *DcimCableTerminationsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim cable terminations partial update default response has a 5xx status code
func (o *DcimCableTerminationsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim cable terminations partial update default response a status code equal to that given
func (o *DcimCableTerminationsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim cable terminations partial update default response
func (o *DcimCableTerminationsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimCableTerminationsPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/cable-terminations/{id}/][%d] dcim_cable-terminations_partial_update default %s", o._statusCode, payload)
}

func (o *DcimCableTerminationsPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/cable-terminations/{id}/][%d] dcim_cable-terminations_partial_update default %s", o._statusCode, payload)
}

func (o *DcimCableTerminationsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCableTerminationsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
