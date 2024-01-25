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

// IpamFhrpGroupsReadReader is a Reader for the IpamFhrpGroupsRead structure.
type IpamFhrpGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamFhrpGroupsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupsReadOK creates a IpamFhrpGroupsReadOK with default headers values
func NewIpamFhrpGroupsReadOK() *IpamFhrpGroupsReadOK {
	return &IpamFhrpGroupsReadOK{}
}

/*
IpamFhrpGroupsReadOK describes a response with status code 200, with default header values.

IpamFhrpGroupsReadOK ipam fhrp groups read o k
*/
type IpamFhrpGroupsReadOK struct {
	Payload *models.FHRPGroup
}

// IsSuccess returns true when this ipam fhrp groups read o k response has a 2xx status code
func (o *IpamFhrpGroupsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp groups read o k response has a 3xx status code
func (o *IpamFhrpGroupsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp groups read o k response has a 4xx status code
func (o *IpamFhrpGroupsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp groups read o k response has a 5xx status code
func (o *IpamFhrpGroupsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp groups read o k response a status code equal to that given
func (o *IpamFhrpGroupsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam fhrp groups read o k response
func (o *IpamFhrpGroupsReadOK) Code() int {
	return 200
}

func (o *IpamFhrpGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsReadOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupsReadOK) String() string {
	return fmt.Sprintf("[GET /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsReadOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupsReadOK) GetPayload() *models.FHRPGroup {
	return o.Payload
}

func (o *IpamFhrpGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamFhrpGroupsReadDefault creates a IpamFhrpGroupsReadDefault with default headers values
func NewIpamFhrpGroupsReadDefault(code int) *IpamFhrpGroupsReadDefault {
	return &IpamFhrpGroupsReadDefault{
		_statusCode: code,
	}
}

/*
IpamFhrpGroupsReadDefault describes a response with status code -1, with default header values.

IpamFhrpGroupsReadDefault ipam fhrp groups read default
*/
type IpamFhrpGroupsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam fhrp groups read default response has a 2xx status code
func (o *IpamFhrpGroupsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam fhrp groups read default response has a 3xx status code
func (o *IpamFhrpGroupsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam fhrp groups read default response has a 4xx status code
func (o *IpamFhrpGroupsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam fhrp groups read default response has a 5xx status code
func (o *IpamFhrpGroupsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam fhrp groups read default response a status code equal to that given
func (o *IpamFhrpGroupsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam fhrp groups read default response
func (o *IpamFhrpGroupsReadDefault) Code() int {
	return o._statusCode
}

func (o *IpamFhrpGroupsReadDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/fhrp-groups/{id}/][%d] ipam_fhrp-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupsReadDefault) String() string {
	return fmt.Sprintf("[GET /ipam/fhrp-groups/{id}/][%d] ipam_fhrp-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
