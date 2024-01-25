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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimRacksCreateReader is a Reader for the DcimRacksCreate structure.
type DcimRacksCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRacksCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRacksCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRacksCreateCreated creates a DcimRacksCreateCreated with default headers values
func NewDcimRacksCreateCreated() *DcimRacksCreateCreated {
	return &DcimRacksCreateCreated{}
}

/*
DcimRacksCreateCreated describes a response with status code 201, with default header values.

DcimRacksCreateCreated dcim racks create created
*/
type DcimRacksCreateCreated struct {
	Payload *models.Rack
}

// IsSuccess returns true when this dcim racks create created response has a 2xx status code
func (o *DcimRacksCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim racks create created response has a 3xx status code
func (o *DcimRacksCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim racks create created response has a 4xx status code
func (o *DcimRacksCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim racks create created response has a 5xx status code
func (o *DcimRacksCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim racks create created response a status code equal to that given
func (o *DcimRacksCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim racks create created response
func (o *DcimRacksCreateCreated) Code() int {
	return 201
}

func (o *DcimRacksCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/racks/][%d] dcimRacksCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRacksCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/racks/][%d] dcimRacksCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRacksCreateCreated) GetPayload() *models.Rack {
	return o.Payload
}

func (o *DcimRacksCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRacksCreateDefault creates a DcimRacksCreateDefault with default headers values
func NewDcimRacksCreateDefault(code int) *DcimRacksCreateDefault {
	return &DcimRacksCreateDefault{
		_statusCode: code,
	}
}

/*
DcimRacksCreateDefault describes a response with status code -1, with default header values.

DcimRacksCreateDefault dcim racks create default
*/
type DcimRacksCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim racks create default response has a 2xx status code
func (o *DcimRacksCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim racks create default response has a 3xx status code
func (o *DcimRacksCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim racks create default response has a 4xx status code
func (o *DcimRacksCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim racks create default response has a 5xx status code
func (o *DcimRacksCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim racks create default response a status code equal to that given
func (o *DcimRacksCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim racks create default response
func (o *DcimRacksCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRacksCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/racks/][%d] dcim_racks_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/racks/][%d] dcim_racks_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
