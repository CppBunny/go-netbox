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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// CircuitsCircuitTerminationsPartialUpdateReader is a Reader for the CircuitsCircuitTerminationsPartialUpdate structure.
type CircuitsCircuitTerminationsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTerminationsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTerminationsPartialUpdateOK creates a CircuitsCircuitTerminationsPartialUpdateOK with default headers values
func NewCircuitsCircuitTerminationsPartialUpdateOK() *CircuitsCircuitTerminationsPartialUpdateOK {
	return &CircuitsCircuitTerminationsPartialUpdateOK{}
}

/*
CircuitsCircuitTerminationsPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsPartialUpdateOK circuits circuit terminations partial update o k
*/
type CircuitsCircuitTerminationsPartialUpdateOK struct {
	Payload *models.CircuitTermination
}

// IsSuccess returns true when this circuits circuit terminations partial update o k response has a 2xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit terminations partial update o k response has a 3xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit terminations partial update o k response has a 4xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit terminations partial update o k response has a 5xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit terminations partial update o k response a status code equal to that given
func (o *CircuitsCircuitTerminationsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits circuit terminations partial update o k response
func (o *CircuitsCircuitTerminationsPartialUpdateOK) Code() int {
	return 200
}

func (o *CircuitsCircuitTerminationsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTerminationsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTerminationsPartialUpdateOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsPartialUpdateDefault creates a CircuitsCircuitTerminationsPartialUpdateDefault with default headers values
func NewCircuitsCircuitTerminationsPartialUpdateDefault(code int) *CircuitsCircuitTerminationsPartialUpdateDefault {
	return &CircuitsCircuitTerminationsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitTerminationsPartialUpdateDefault describes a response with status code -1, with default header values.

CircuitsCircuitTerminationsPartialUpdateDefault circuits circuit terminations partial update default
*/
type CircuitsCircuitTerminationsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this circuits circuit terminations partial update default response has a 2xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuit terminations partial update default response has a 3xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuit terminations partial update default response has a 4xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuit terminations partial update default response has a 5xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuit terminations partial update default response a status code equal to that given
func (o *CircuitsCircuitTerminationsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the circuits circuit terminations partial update default response
func (o *CircuitsCircuitTerminationsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTerminationsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/{id}/][%d] circuits_circuit-terminations_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTerminationsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/{id}/][%d] circuits_circuit-terminations_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTerminationsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
