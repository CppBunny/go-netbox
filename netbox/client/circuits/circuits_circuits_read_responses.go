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

// CircuitsCircuitsReadReader is a Reader for the CircuitsCircuitsRead structure.
type CircuitsCircuitsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitsReadOK creates a CircuitsCircuitsReadOK with default headers values
func NewCircuitsCircuitsReadOK() *CircuitsCircuitsReadOK {
	return &CircuitsCircuitsReadOK{}
}

/*
CircuitsCircuitsReadOK describes a response with status code 200, with default header values.

CircuitsCircuitsReadOK circuits circuits read o k
*/
type CircuitsCircuitsReadOK struct {
	Payload *models.Circuit
}

// IsSuccess returns true when this circuits circuits read o k response has a 2xx status code
func (o *CircuitsCircuitsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuits read o k response has a 3xx status code
func (o *CircuitsCircuitsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuits read o k response has a 4xx status code
func (o *CircuitsCircuitsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuits read o k response has a 5xx status code
func (o *CircuitsCircuitsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuits read o k response a status code equal to that given
func (o *CircuitsCircuitsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits circuits read o k response
func (o *CircuitsCircuitsReadOK) Code() int {
	return 200
}

func (o *CircuitsCircuitsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /circuits/circuits/{id}/][%d] circuitsCircuitsReadOK %s", 200, payload)
}

func (o *CircuitsCircuitsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /circuits/circuits/{id}/][%d] circuitsCircuitsReadOK %s", 200, payload)
}

func (o *CircuitsCircuitsReadOK) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitsReadDefault creates a CircuitsCircuitsReadDefault with default headers values
func NewCircuitsCircuitsReadDefault(code int) *CircuitsCircuitsReadDefault {
	return &CircuitsCircuitsReadDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitsReadDefault describes a response with status code -1, with default header values.

CircuitsCircuitsReadDefault circuits circuits read default
*/
type CircuitsCircuitsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this circuits circuits read default response has a 2xx status code
func (o *CircuitsCircuitsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuits read default response has a 3xx status code
func (o *CircuitsCircuitsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuits read default response has a 4xx status code
func (o *CircuitsCircuitsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuits read default response has a 5xx status code
func (o *CircuitsCircuitsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuits read default response a status code equal to that given
func (o *CircuitsCircuitsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the circuits circuits read default response
func (o *CircuitsCircuitsReadDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitsReadDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /circuits/circuits/{id}/][%d] circuits_circuits_read default %s", o._statusCode, payload)
}

func (o *CircuitsCircuitsReadDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /circuits/circuits/{id}/][%d] circuits_circuits_read default %s", o._statusCode, payload)
}

func (o *CircuitsCircuitsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
