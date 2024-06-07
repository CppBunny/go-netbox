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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CppBunny/go-netbox/netbox/models"
)

// IpamServicesUpdateReader is a Reader for the IpamServicesUpdate structure.
type IpamServicesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServicesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServicesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServicesUpdateOK creates a IpamServicesUpdateOK with default headers values
func NewIpamServicesUpdateOK() *IpamServicesUpdateOK {
	return &IpamServicesUpdateOK{}
}

/*
IpamServicesUpdateOK describes a response with status code 200, with default header values.

IpamServicesUpdateOK ipam services update o k
*/
type IpamServicesUpdateOK struct {
	Payload *models.Service
}

// IsSuccess returns true when this ipam services update o k response has a 2xx status code
func (o *IpamServicesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam services update o k response has a 3xx status code
func (o *IpamServicesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam services update o k response has a 4xx status code
func (o *IpamServicesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam services update o k response has a 5xx status code
func (o *IpamServicesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam services update o k response a status code equal to that given
func (o *IpamServicesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam services update o k response
func (o *IpamServicesUpdateOK) Code() int {
	return 200
}

func (o *IpamServicesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/services/{id}/][%d] ipamServicesUpdateOK %s", 200, payload)
}

func (o *IpamServicesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/services/{id}/][%d] ipamServicesUpdateOK %s", 200, payload)
}

func (o *IpamServicesUpdateOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServicesUpdateDefault creates a IpamServicesUpdateDefault with default headers values
func NewIpamServicesUpdateDefault(code int) *IpamServicesUpdateDefault {
	return &IpamServicesUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamServicesUpdateDefault describes a response with status code -1, with default header values.

IpamServicesUpdateDefault ipam services update default
*/
type IpamServicesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam services update default response has a 2xx status code
func (o *IpamServicesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam services update default response has a 3xx status code
func (o *IpamServicesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam services update default response has a 4xx status code
func (o *IpamServicesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam services update default response has a 5xx status code
func (o *IpamServicesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam services update default response a status code equal to that given
func (o *IpamServicesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam services update default response
func (o *IpamServicesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamServicesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/services/{id}/][%d] ipam_services_update default %s", o._statusCode, payload)
}

func (o *IpamServicesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/services/{id}/][%d] ipam_services_update default %s", o._statusCode, payload)
}

func (o *IpamServicesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServicesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
