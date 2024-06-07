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

// DcimPowerPortTemplatesCreateReader is a Reader for the DcimPowerPortTemplatesCreate structure.
type DcimPowerPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortTemplatesCreateCreated creates a DcimPowerPortTemplatesCreateCreated with default headers values
func NewDcimPowerPortTemplatesCreateCreated() *DcimPowerPortTemplatesCreateCreated {
	return &DcimPowerPortTemplatesCreateCreated{}
}

/*
DcimPowerPortTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimPowerPortTemplatesCreateCreated dcim power port templates create created
*/
type DcimPowerPortTemplatesCreateCreated struct {
	Payload *models.PowerPortTemplate
}

// IsSuccess returns true when this dcim power port templates create created response has a 2xx status code
func (o *DcimPowerPortTemplatesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power port templates create created response has a 3xx status code
func (o *DcimPowerPortTemplatesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power port templates create created response has a 4xx status code
func (o *DcimPowerPortTemplatesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power port templates create created response has a 5xx status code
func (o *DcimPowerPortTemplatesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power port templates create created response a status code equal to that given
func (o *DcimPowerPortTemplatesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim power port templates create created response
func (o *DcimPowerPortTemplatesCreateCreated) Code() int {
	return 201
}

func (o *DcimPowerPortTemplatesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/power-port-templates/][%d] dcimPowerPortTemplatesCreateCreated %s", 201, payload)
}

func (o *DcimPowerPortTemplatesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/power-port-templates/][%d] dcimPowerPortTemplatesCreateCreated %s", 201, payload)
}

func (o *DcimPowerPortTemplatesCreateCreated) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortTemplatesCreateDefault creates a DcimPowerPortTemplatesCreateDefault with default headers values
func NewDcimPowerPortTemplatesCreateDefault(code int) *DcimPowerPortTemplatesCreateDefault {
	return &DcimPowerPortTemplatesCreateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortTemplatesCreateDefault describes a response with status code -1, with default header values.

DcimPowerPortTemplatesCreateDefault dcim power port templates create default
*/
type DcimPowerPortTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power port templates create default response has a 2xx status code
func (o *DcimPowerPortTemplatesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power port templates create default response has a 3xx status code
func (o *DcimPowerPortTemplatesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power port templates create default response has a 4xx status code
func (o *DcimPowerPortTemplatesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power port templates create default response has a 5xx status code
func (o *DcimPowerPortTemplatesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power port templates create default response a status code equal to that given
func (o *DcimPowerPortTemplatesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power port templates create default response
func (o *DcimPowerPortTemplatesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortTemplatesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/power-port-templates/][%d] dcim_power-port-templates_create default %s", o._statusCode, payload)
}

func (o *DcimPowerPortTemplatesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/power-port-templates/][%d] dcim_power-port-templates_create default %s", o._statusCode, payload)
}

func (o *DcimPowerPortTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
