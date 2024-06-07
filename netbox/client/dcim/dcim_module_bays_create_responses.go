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

// DcimModuleBaysCreateReader is a Reader for the DcimModuleBaysCreate structure.
type DcimModuleBaysCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBaysCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimModuleBaysCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBaysCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBaysCreateCreated creates a DcimModuleBaysCreateCreated with default headers values
func NewDcimModuleBaysCreateCreated() *DcimModuleBaysCreateCreated {
	return &DcimModuleBaysCreateCreated{}
}

/*
DcimModuleBaysCreateCreated describes a response with status code 201, with default header values.

DcimModuleBaysCreateCreated dcim module bays create created
*/
type DcimModuleBaysCreateCreated struct {
	Payload *models.ModuleBay
}

// IsSuccess returns true when this dcim module bays create created response has a 2xx status code
func (o *DcimModuleBaysCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bays create created response has a 3xx status code
func (o *DcimModuleBaysCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bays create created response has a 4xx status code
func (o *DcimModuleBaysCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bays create created response has a 5xx status code
func (o *DcimModuleBaysCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bays create created response a status code equal to that given
func (o *DcimModuleBaysCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim module bays create created response
func (o *DcimModuleBaysCreateCreated) Code() int {
	return 201
}

func (o *DcimModuleBaysCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/module-bays/][%d] dcimModuleBaysCreateCreated %s", 201, payload)
}

func (o *DcimModuleBaysCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/module-bays/][%d] dcimModuleBaysCreateCreated %s", 201, payload)
}

func (o *DcimModuleBaysCreateCreated) GetPayload() *models.ModuleBay {
	return o.Payload
}

func (o *DcimModuleBaysCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBaysCreateDefault creates a DcimModuleBaysCreateDefault with default headers values
func NewDcimModuleBaysCreateDefault(code int) *DcimModuleBaysCreateDefault {
	return &DcimModuleBaysCreateDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBaysCreateDefault describes a response with status code -1, with default header values.

DcimModuleBaysCreateDefault dcim module bays create default
*/
type DcimModuleBaysCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim module bays create default response has a 2xx status code
func (o *DcimModuleBaysCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bays create default response has a 3xx status code
func (o *DcimModuleBaysCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bays create default response has a 4xx status code
func (o *DcimModuleBaysCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bays create default response has a 5xx status code
func (o *DcimModuleBaysCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bays create default response a status code equal to that given
func (o *DcimModuleBaysCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim module bays create default response
func (o *DcimModuleBaysCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimModuleBaysCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/module-bays/][%d] dcim_module-bays_create default %s", o._statusCode, payload)
}

func (o *DcimModuleBaysCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/module-bays/][%d] dcim_module-bays_create default %s", o._statusCode, payload)
}

func (o *DcimModuleBaysCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBaysCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
