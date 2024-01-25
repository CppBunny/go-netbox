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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamServicesCreateReader is a Reader for the IpamServicesCreate structure.
type IpamServicesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamServicesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServicesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServicesCreateCreated creates a IpamServicesCreateCreated with default headers values
func NewIpamServicesCreateCreated() *IpamServicesCreateCreated {
	return &IpamServicesCreateCreated{}
}

/*
IpamServicesCreateCreated describes a response with status code 201, with default header values.

IpamServicesCreateCreated ipam services create created
*/
type IpamServicesCreateCreated struct {
	Payload *models.Service
}

// IsSuccess returns true when this ipam services create created response has a 2xx status code
func (o *IpamServicesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam services create created response has a 3xx status code
func (o *IpamServicesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam services create created response has a 4xx status code
func (o *IpamServicesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam services create created response has a 5xx status code
func (o *IpamServicesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam services create created response a status code equal to that given
func (o *IpamServicesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam services create created response
func (o *IpamServicesCreateCreated) Code() int {
	return 201
}

func (o *IpamServicesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/services/][%d] ipamServicesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamServicesCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/services/][%d] ipamServicesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamServicesCreateCreated) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServicesCreateDefault creates a IpamServicesCreateDefault with default headers values
func NewIpamServicesCreateDefault(code int) *IpamServicesCreateDefault {
	return &IpamServicesCreateDefault{
		_statusCode: code,
	}
}

/*
IpamServicesCreateDefault describes a response with status code -1, with default header values.

IpamServicesCreateDefault ipam services create default
*/
type IpamServicesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam services create default response has a 2xx status code
func (o *IpamServicesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam services create default response has a 3xx status code
func (o *IpamServicesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam services create default response has a 4xx status code
func (o *IpamServicesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam services create default response has a 5xx status code
func (o *IpamServicesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam services create default response a status code equal to that given
func (o *IpamServicesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam services create default response
func (o *IpamServicesCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpamServicesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/services/][%d] ipam_services_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServicesCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/services/][%d] ipam_services_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServicesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServicesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
