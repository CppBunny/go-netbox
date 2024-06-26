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

// VirtualizationVirtualDisksUpdateReader is a Reader for the VirtualizationVirtualDisksUpdate structure.
type VirtualizationVirtualDisksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualDisksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualDisksUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationVirtualDisksUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationVirtualDisksUpdateOK creates a VirtualizationVirtualDisksUpdateOK with default headers values
func NewVirtualizationVirtualDisksUpdateOK() *VirtualizationVirtualDisksUpdateOK {
	return &VirtualizationVirtualDisksUpdateOK{}
}

/*
VirtualizationVirtualDisksUpdateOK describes a response with status code 200, with default header values.

VirtualizationVirtualDisksUpdateOK virtualization virtual disks update o k
*/
type VirtualizationVirtualDisksUpdateOK struct {
	Payload *models.VirtualDisk
}

// IsSuccess returns true when this virtualization virtual disks update o k response has a 2xx status code
func (o *VirtualizationVirtualDisksUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization virtual disks update o k response has a 3xx status code
func (o *VirtualizationVirtualDisksUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual disks update o k response has a 4xx status code
func (o *VirtualizationVirtualDisksUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization virtual disks update o k response has a 5xx status code
func (o *VirtualizationVirtualDisksUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual disks update o k response a status code equal to that given
func (o *VirtualizationVirtualDisksUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization virtual disks update o k response
func (o *VirtualizationVirtualDisksUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationVirtualDisksUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /virtualization/virtual-disks/{id}/][%d] virtualizationVirtualDisksUpdateOK %s", 200, payload)
}

func (o *VirtualizationVirtualDisksUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /virtualization/virtual-disks/{id}/][%d] virtualizationVirtualDisksUpdateOK %s", 200, payload)
}

func (o *VirtualizationVirtualDisksUpdateOK) GetPayload() *models.VirtualDisk {
	return o.Payload
}

func (o *VirtualizationVirtualDisksUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualDisk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationVirtualDisksUpdateDefault creates a VirtualizationVirtualDisksUpdateDefault with default headers values
func NewVirtualizationVirtualDisksUpdateDefault(code int) *VirtualizationVirtualDisksUpdateDefault {
	return &VirtualizationVirtualDisksUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationVirtualDisksUpdateDefault describes a response with status code -1, with default header values.

VirtualizationVirtualDisksUpdateDefault virtualization virtual disks update default
*/
type VirtualizationVirtualDisksUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization virtual disks update default response has a 2xx status code
func (o *VirtualizationVirtualDisksUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization virtual disks update default response has a 3xx status code
func (o *VirtualizationVirtualDisksUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization virtual disks update default response has a 4xx status code
func (o *VirtualizationVirtualDisksUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization virtual disks update default response has a 5xx status code
func (o *VirtualizationVirtualDisksUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization virtual disks update default response a status code equal to that given
func (o *VirtualizationVirtualDisksUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization virtual disks update default response
func (o *VirtualizationVirtualDisksUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationVirtualDisksUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /virtualization/virtual-disks/{id}/][%d] virtualization_virtual-disks_update default %s", o._statusCode, payload)
}

func (o *VirtualizationVirtualDisksUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /virtualization/virtual-disks/{id}/][%d] virtualization_virtual-disks_update default %s", o._statusCode, payload)
}

func (o *VirtualizationVirtualDisksUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualDisksUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
