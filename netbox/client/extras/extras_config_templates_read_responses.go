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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasConfigTemplatesReadReader is a Reader for the ExtrasConfigTemplatesRead structure.
type ExtrasConfigTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasConfigTemplatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasConfigTemplatesReadOK creates a ExtrasConfigTemplatesReadOK with default headers values
func NewExtrasConfigTemplatesReadOK() *ExtrasConfigTemplatesReadOK {
	return &ExtrasConfigTemplatesReadOK{}
}

/*
ExtrasConfigTemplatesReadOK describes a response with status code 200, with default header values.

ExtrasConfigTemplatesReadOK extras config templates read o k
*/
type ExtrasConfigTemplatesReadOK struct {
	Payload *models.ConfigTemplate
}

// IsSuccess returns true when this extras config templates read o k response has a 2xx status code
func (o *ExtrasConfigTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config templates read o k response has a 3xx status code
func (o *ExtrasConfigTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config templates read o k response has a 4xx status code
func (o *ExtrasConfigTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config templates read o k response has a 5xx status code
func (o *ExtrasConfigTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config templates read o k response a status code equal to that given
func (o *ExtrasConfigTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras config templates read o k response
func (o *ExtrasConfigTemplatesReadOK) Code() int {
	return 200
}

func (o *ExtrasConfigTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/config-templates/{id}/][%d] extrasConfigTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasConfigTemplatesReadOK) String() string {
	return fmt.Sprintf("[GET /extras/config-templates/{id}/][%d] extrasConfigTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasConfigTemplatesReadOK) GetPayload() *models.ConfigTemplate {
	return o.Payload
}

func (o *ExtrasConfigTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasConfigTemplatesReadDefault creates a ExtrasConfigTemplatesReadDefault with default headers values
func NewExtrasConfigTemplatesReadDefault(code int) *ExtrasConfigTemplatesReadDefault {
	return &ExtrasConfigTemplatesReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasConfigTemplatesReadDefault describes a response with status code -1, with default header values.

ExtrasConfigTemplatesReadDefault extras config templates read default
*/
type ExtrasConfigTemplatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras config templates read default response has a 2xx status code
func (o *ExtrasConfigTemplatesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras config templates read default response has a 3xx status code
func (o *ExtrasConfigTemplatesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras config templates read default response has a 4xx status code
func (o *ExtrasConfigTemplatesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras config templates read default response has a 5xx status code
func (o *ExtrasConfigTemplatesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras config templates read default response a status code equal to that given
func (o *ExtrasConfigTemplatesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras config templates read default response
func (o *ExtrasConfigTemplatesReadDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasConfigTemplatesReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/config-templates/{id}/][%d] extras_config-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigTemplatesReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/config-templates/{id}/][%d] extras_config-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigTemplatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigTemplatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}