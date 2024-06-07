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

package tenancy

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

// TenancyTenantsUpdateReader is a Reader for the TenancyTenantsUpdate structure.
type TenancyTenantsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantsUpdateOK creates a TenancyTenantsUpdateOK with default headers values
func NewTenancyTenantsUpdateOK() *TenancyTenantsUpdateOK {
	return &TenancyTenantsUpdateOK{}
}

/*
TenancyTenantsUpdateOK describes a response with status code 200, with default header values.

TenancyTenantsUpdateOK tenancy tenants update o k
*/
type TenancyTenantsUpdateOK struct {
	Payload *models.Tenant
}

// IsSuccess returns true when this tenancy tenants update o k response has a 2xx status code
func (o *TenancyTenantsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy tenants update o k response has a 3xx status code
func (o *TenancyTenantsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenants update o k response has a 4xx status code
func (o *TenancyTenantsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy tenants update o k response has a 5xx status code
func (o *TenancyTenantsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenants update o k response a status code equal to that given
func (o *TenancyTenantsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy tenants update o k response
func (o *TenancyTenantsUpdateOK) Code() int {
	return 200
}

func (o *TenancyTenantsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/tenants/{id}/][%d] tenancyTenantsUpdateOK %s", 200, payload)
}

func (o *TenancyTenantsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/tenants/{id}/][%d] tenancyTenantsUpdateOK %s", 200, payload)
}

func (o *TenancyTenantsUpdateOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *TenancyTenantsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantsUpdateDefault creates a TenancyTenantsUpdateDefault with default headers values
func NewTenancyTenantsUpdateDefault(code int) *TenancyTenantsUpdateDefault {
	return &TenancyTenantsUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyTenantsUpdateDefault describes a response with status code -1, with default header values.

TenancyTenantsUpdateDefault tenancy tenants update default
*/
type TenancyTenantsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy tenants update default response has a 2xx status code
func (o *TenancyTenantsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy tenants update default response has a 3xx status code
func (o *TenancyTenantsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy tenants update default response has a 4xx status code
func (o *TenancyTenantsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy tenants update default response has a 5xx status code
func (o *TenancyTenantsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy tenants update default response a status code equal to that given
func (o *TenancyTenantsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy tenants update default response
func (o *TenancyTenantsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyTenantsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/tenants/{id}/][%d] tenancy_tenants_update default %s", o._statusCode, payload)
}

func (o *TenancyTenantsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/tenants/{id}/][%d] tenancy_tenants_update default %s", o._statusCode, payload)
}

func (o *TenancyTenantsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
