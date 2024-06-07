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

package virtualization

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

// VirtualizationInterfacesPartialUpdateReader is a Reader for the VirtualizationInterfacesPartialUpdate structure.
type VirtualizationInterfacesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationInterfacesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationInterfacesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationInterfacesPartialUpdateOK creates a VirtualizationInterfacesPartialUpdateOK with default headers values
func NewVirtualizationInterfacesPartialUpdateOK() *VirtualizationInterfacesPartialUpdateOK {
	return &VirtualizationInterfacesPartialUpdateOK{}
}

/*
VirtualizationInterfacesPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationInterfacesPartialUpdateOK virtualization interfaces partial update o k
*/
type VirtualizationInterfacesPartialUpdateOK struct {
	Payload *models.VMInterface
}

// IsSuccess returns true when this virtualization interfaces partial update o k response has a 2xx status code
func (o *VirtualizationInterfacesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization interfaces partial update o k response has a 3xx status code
func (o *VirtualizationInterfacesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization interfaces partial update o k response has a 4xx status code
func (o *VirtualizationInterfacesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization interfaces partial update o k response has a 5xx status code
func (o *VirtualizationInterfacesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization interfaces partial update o k response a status code equal to that given
func (o *VirtualizationInterfacesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization interfaces partial update o k response
func (o *VirtualizationInterfacesPartialUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationInterfacesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/interfaces/{id}/][%d] virtualizationInterfacesPartialUpdateOK %s", 200, payload)
}

func (o *VirtualizationInterfacesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/interfaces/{id}/][%d] virtualizationInterfacesPartialUpdateOK %s", 200, payload)
}

func (o *VirtualizationInterfacesPartialUpdateOK) GetPayload() *models.VMInterface {
	return o.Payload
}

func (o *VirtualizationInterfacesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationInterfacesPartialUpdateDefault creates a VirtualizationInterfacesPartialUpdateDefault with default headers values
func NewVirtualizationInterfacesPartialUpdateDefault(code int) *VirtualizationInterfacesPartialUpdateDefault {
	return &VirtualizationInterfacesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationInterfacesPartialUpdateDefault describes a response with status code -1, with default header values.

VirtualizationInterfacesPartialUpdateDefault virtualization interfaces partial update default
*/
type VirtualizationInterfacesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization interfaces partial update default response has a 2xx status code
func (o *VirtualizationInterfacesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization interfaces partial update default response has a 3xx status code
func (o *VirtualizationInterfacesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization interfaces partial update default response has a 4xx status code
func (o *VirtualizationInterfacesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization interfaces partial update default response has a 5xx status code
func (o *VirtualizationInterfacesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization interfaces partial update default response a status code equal to that given
func (o *VirtualizationInterfacesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization interfaces partial update default response
func (o *VirtualizationInterfacesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationInterfacesPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/interfaces/{id}/][%d] virtualization_interfaces_partial_update default %s", o._statusCode, payload)
}

func (o *VirtualizationInterfacesPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/interfaces/{id}/][%d] virtualization_interfaces_partial_update default %s", o._statusCode, payload)
}

func (o *VirtualizationInterfacesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationInterfacesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
