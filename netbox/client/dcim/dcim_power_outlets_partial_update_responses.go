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

// DcimPowerOutletsPartialUpdateReader is a Reader for the DcimPowerOutletsPartialUpdate structure.
type DcimPowerOutletsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletsPartialUpdateOK creates a DcimPowerOutletsPartialUpdateOK with default headers values
func NewDcimPowerOutletsPartialUpdateOK() *DcimPowerOutletsPartialUpdateOK {
	return &DcimPowerOutletsPartialUpdateOK{}
}

/*
DcimPowerOutletsPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletsPartialUpdateOK dcim power outlets partial update o k
*/
type DcimPowerOutletsPartialUpdateOK struct {
	Payload *models.PowerOutlet
}

// IsSuccess returns true when this dcim power outlets partial update o k response has a 2xx status code
func (o *DcimPowerOutletsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets partial update o k response has a 3xx status code
func (o *DcimPowerOutletsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets partial update o k response has a 4xx status code
func (o *DcimPowerOutletsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets partial update o k response has a 5xx status code
func (o *DcimPowerOutletsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets partial update o k response a status code equal to that given
func (o *DcimPowerOutletsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power outlets partial update o k response
func (o *DcimPowerOutletsPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerOutletsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlets/{id}/][%d] dcimPowerOutletsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlets/{id}/][%d] dcimPowerOutletsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsPartialUpdateOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletsPartialUpdateDefault creates a DcimPowerOutletsPartialUpdateDefault with default headers values
func NewDcimPowerOutletsPartialUpdateDefault(code int) *DcimPowerOutletsPartialUpdateDefault {
	return &DcimPowerOutletsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerOutletsPartialUpdateDefault dcim power outlets partial update default
*/
type DcimPowerOutletsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power outlets partial update default response has a 2xx status code
func (o *DcimPowerOutletsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlets partial update default response has a 3xx status code
func (o *DcimPowerOutletsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlets partial update default response has a 4xx status code
func (o *DcimPowerOutletsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlets partial update default response has a 5xx status code
func (o *DcimPowerOutletsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlets partial update default response a status code equal to that given
func (o *DcimPowerOutletsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power outlets partial update default response
func (o *DcimPowerOutletsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerOutletsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlets/{id}/][%d] dcim_power-outlets_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlets/{id}/][%d] dcim_power-outlets_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
