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

// TenancyContactGroupsUpdateReader is a Reader for the TenancyContactGroupsUpdate structure.
type TenancyContactGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactGroupsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactGroupsUpdateOK creates a TenancyContactGroupsUpdateOK with default headers values
func NewTenancyContactGroupsUpdateOK() *TenancyContactGroupsUpdateOK {
	return &TenancyContactGroupsUpdateOK{}
}

/*
TenancyContactGroupsUpdateOK describes a response with status code 200, with default header values.

TenancyContactGroupsUpdateOK tenancy contact groups update o k
*/
type TenancyContactGroupsUpdateOK struct {
	Payload *models.ContactGroup
}

// IsSuccess returns true when this tenancy contact groups update o k response has a 2xx status code
func (o *TenancyContactGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact groups update o k response has a 3xx status code
func (o *TenancyContactGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups update o k response has a 4xx status code
func (o *TenancyContactGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact groups update o k response has a 5xx status code
func (o *TenancyContactGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups update o k response a status code equal to that given
func (o *TenancyContactGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contact groups update o k response
func (o *TenancyContactGroupsUpdateOK) Code() int {
	return 200
}

func (o *TenancyContactGroupsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsUpdateOK %s", 200, payload)
}

func (o *TenancyContactGroupsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsUpdateOK %s", 200, payload)
}

func (o *TenancyContactGroupsUpdateOK) GetPayload() *models.ContactGroup {
	return o.Payload
}

func (o *TenancyContactGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactGroupsUpdateDefault creates a TenancyContactGroupsUpdateDefault with default headers values
func NewTenancyContactGroupsUpdateDefault(code int) *TenancyContactGroupsUpdateDefault {
	return &TenancyContactGroupsUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactGroupsUpdateDefault describes a response with status code -1, with default header values.

TenancyContactGroupsUpdateDefault tenancy contact groups update default
*/
type TenancyContactGroupsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contact groups update default response has a 2xx status code
func (o *TenancyContactGroupsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact groups update default response has a 3xx status code
func (o *TenancyContactGroupsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact groups update default response has a 4xx status code
func (o *TenancyContactGroupsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact groups update default response has a 5xx status code
func (o *TenancyContactGroupsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact groups update default response a status code equal to that given
func (o *TenancyContactGroupsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contact groups update default response
func (o *TenancyContactGroupsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactGroupsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/contact-groups/{id}/][%d] tenancy_contact-groups_update default %s", o._statusCode, payload)
}

func (o *TenancyContactGroupsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/contact-groups/{id}/][%d] tenancy_contact-groups_update default %s", o._statusCode, payload)
}

func (o *TenancyContactGroupsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactGroupsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
