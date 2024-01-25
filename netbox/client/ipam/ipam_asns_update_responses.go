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

// IpamAsnsUpdateReader is a Reader for the IpamAsnsUpdate structure.
type IpamAsnsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAsnsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAsnsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamAsnsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamAsnsUpdateOK creates a IpamAsnsUpdateOK with default headers values
func NewIpamAsnsUpdateOK() *IpamAsnsUpdateOK {
	return &IpamAsnsUpdateOK{}
}

/*
IpamAsnsUpdateOK describes a response with status code 200, with default header values.

IpamAsnsUpdateOK ipam asns update o k
*/
type IpamAsnsUpdateOK struct {
	Payload *models.ASN
}

// IsSuccess returns true when this ipam asns update o k response has a 2xx status code
func (o *IpamAsnsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam asns update o k response has a 3xx status code
func (o *IpamAsnsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam asns update o k response has a 4xx status code
func (o *IpamAsnsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam asns update o k response has a 5xx status code
func (o *IpamAsnsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam asns update o k response a status code equal to that given
func (o *IpamAsnsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam asns update o k response
func (o *IpamAsnsUpdateOK) Code() int {
	return 200
}

func (o *IpamAsnsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/asns/{id}/][%d] ipamAsnsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAsnsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/asns/{id}/][%d] ipamAsnsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAsnsUpdateOK) GetPayload() *models.ASN {
	return o.Payload
}

func (o *IpamAsnsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ASN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamAsnsUpdateDefault creates a IpamAsnsUpdateDefault with default headers values
func NewIpamAsnsUpdateDefault(code int) *IpamAsnsUpdateDefault {
	return &IpamAsnsUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamAsnsUpdateDefault describes a response with status code -1, with default header values.

IpamAsnsUpdateDefault ipam asns update default
*/
type IpamAsnsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam asns update default response has a 2xx status code
func (o *IpamAsnsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam asns update default response has a 3xx status code
func (o *IpamAsnsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam asns update default response has a 4xx status code
func (o *IpamAsnsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam asns update default response has a 5xx status code
func (o *IpamAsnsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam asns update default response a status code equal to that given
func (o *IpamAsnsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam asns update default response
func (o *IpamAsnsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamAsnsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/asns/{id}/][%d] ipam_asns_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAsnsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/asns/{id}/][%d] ipam_asns_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAsnsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAsnsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
