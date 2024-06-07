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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CppBunny/go-netbox/netbox/models"
)

// CircuitsCircuitTerminationsPathsReader is a Reader for the CircuitsCircuitTerminationsPaths structure.
type CircuitsCircuitTerminationsPathsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsPathsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsPathsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTerminationsPathsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTerminationsPathsOK creates a CircuitsCircuitTerminationsPathsOK with default headers values
func NewCircuitsCircuitTerminationsPathsOK() *CircuitsCircuitTerminationsPathsOK {
	return &CircuitsCircuitTerminationsPathsOK{}
}

/*
CircuitsCircuitTerminationsPathsOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsPathsOK circuits circuit terminations paths o k
*/
type CircuitsCircuitTerminationsPathsOK struct {
	Payload *models.CircuitTermination
}

// IsSuccess returns true when this circuits circuit terminations paths o k response has a 2xx status code
func (o *CircuitsCircuitTerminationsPathsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit terminations paths o k response has a 3xx status code
func (o *CircuitsCircuitTerminationsPathsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit terminations paths o k response has a 4xx status code
func (o *CircuitsCircuitTerminationsPathsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit terminations paths o k response has a 5xx status code
func (o *CircuitsCircuitTerminationsPathsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit terminations paths o k response a status code equal to that given
func (o *CircuitsCircuitTerminationsPathsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits circuit terminations paths o k response
func (o *CircuitsCircuitTerminationsPathsOK) Code() int {
	return 200
}

func (o *CircuitsCircuitTerminationsPathsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/paths/][%d] circuitsCircuitTerminationsPathsOK %s", 200, payload)
}

func (o *CircuitsCircuitTerminationsPathsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/paths/][%d] circuitsCircuitTerminationsPathsOK %s", 200, payload)
}

func (o *CircuitsCircuitTerminationsPathsOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsPathsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsPathsDefault creates a CircuitsCircuitTerminationsPathsDefault with default headers values
func NewCircuitsCircuitTerminationsPathsDefault(code int) *CircuitsCircuitTerminationsPathsDefault {
	return &CircuitsCircuitTerminationsPathsDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitTerminationsPathsDefault describes a response with status code -1, with default header values.

CircuitsCircuitTerminationsPathsDefault circuits circuit terminations paths default
*/
type CircuitsCircuitTerminationsPathsDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this circuits circuit terminations paths default response has a 2xx status code
func (o *CircuitsCircuitTerminationsPathsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuit terminations paths default response has a 3xx status code
func (o *CircuitsCircuitTerminationsPathsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuit terminations paths default response has a 4xx status code
func (o *CircuitsCircuitTerminationsPathsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuit terminations paths default response has a 5xx status code
func (o *CircuitsCircuitTerminationsPathsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuit terminations paths default response a status code equal to that given
func (o *CircuitsCircuitTerminationsPathsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the circuits circuit terminations paths default response
func (o *CircuitsCircuitTerminationsPathsDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTerminationsPathsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/paths/][%d] circuits_circuit-terminations_paths default %s", o._statusCode, payload)
}

func (o *CircuitsCircuitTerminationsPathsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/paths/][%d] circuits_circuit-terminations_paths default %s", o._statusCode, payload)
}

func (o *CircuitsCircuitTerminationsPathsDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsPathsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
