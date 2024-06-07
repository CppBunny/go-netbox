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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CppBunny/go-netbox/netbox/models"
)

// ExtrasConfigTemplatesUpdateReader is a Reader for the ExtrasConfigTemplatesUpdate structure.
type ExtrasConfigTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasConfigTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasConfigTemplatesUpdateOK creates a ExtrasConfigTemplatesUpdateOK with default headers values
func NewExtrasConfigTemplatesUpdateOK() *ExtrasConfigTemplatesUpdateOK {
	return &ExtrasConfigTemplatesUpdateOK{}
}

/*
ExtrasConfigTemplatesUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigTemplatesUpdateOK extras config templates update o k
*/
type ExtrasConfigTemplatesUpdateOK struct {
	Payload *models.ConfigTemplate
}

// IsSuccess returns true when this extras config templates update o k response has a 2xx status code
func (o *ExtrasConfigTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config templates update o k response has a 3xx status code
func (o *ExtrasConfigTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config templates update o k response has a 4xx status code
func (o *ExtrasConfigTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config templates update o k response has a 5xx status code
func (o *ExtrasConfigTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config templates update o k response a status code equal to that given
func (o *ExtrasConfigTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras config templates update o k response
func (o *ExtrasConfigTemplatesUpdateOK) Code() int {
	return 200
}

func (o *ExtrasConfigTemplatesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/config-templates/{id}/][%d] extrasConfigTemplatesUpdateOK %s", 200, payload)
}

func (o *ExtrasConfigTemplatesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/config-templates/{id}/][%d] extrasConfigTemplatesUpdateOK %s", 200, payload)
}

func (o *ExtrasConfigTemplatesUpdateOK) GetPayload() *models.ConfigTemplate {
	return o.Payload
}

func (o *ExtrasConfigTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasConfigTemplatesUpdateDefault creates a ExtrasConfigTemplatesUpdateDefault with default headers values
func NewExtrasConfigTemplatesUpdateDefault(code int) *ExtrasConfigTemplatesUpdateDefault {
	return &ExtrasConfigTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasConfigTemplatesUpdateDefault describes a response with status code -1, with default header values.

ExtrasConfigTemplatesUpdateDefault extras config templates update default
*/
type ExtrasConfigTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras config templates update default response has a 2xx status code
func (o *ExtrasConfigTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras config templates update default response has a 3xx status code
func (o *ExtrasConfigTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras config templates update default response has a 4xx status code
func (o *ExtrasConfigTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras config templates update default response has a 5xx status code
func (o *ExtrasConfigTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras config templates update default response a status code equal to that given
func (o *ExtrasConfigTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras config templates update default response
func (o *ExtrasConfigTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasConfigTemplatesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/config-templates/{id}/][%d] extras_config-templates_update default %s", o._statusCode, payload)
}

func (o *ExtrasConfigTemplatesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/config-templates/{id}/][%d] extras_config-templates_update default %s", o._statusCode, payload)
}

func (o *ExtrasConfigTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
