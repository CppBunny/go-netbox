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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CppBunny/go-netbox/netbox/models"
)

// DcimPowerFeedsTraceReader is a Reader for the DcimPowerFeedsTrace structure.
type DcimPowerFeedsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerFeedsTraceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsTraceOK creates a DcimPowerFeedsTraceOK with default headers values
func NewDcimPowerFeedsTraceOK() *DcimPowerFeedsTraceOK {
	return &DcimPowerFeedsTraceOK{}
}

/*
DcimPowerFeedsTraceOK describes a response with status code 200, with default header values.

DcimPowerFeedsTraceOK dcim power feeds trace o k
*/
type DcimPowerFeedsTraceOK struct {
	Payload *models.PowerFeed
}

// IsSuccess returns true when this dcim power feeds trace o k response has a 2xx status code
func (o *DcimPowerFeedsTraceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power feeds trace o k response has a 3xx status code
func (o *DcimPowerFeedsTraceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power feeds trace o k response has a 4xx status code
func (o *DcimPowerFeedsTraceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power feeds trace o k response has a 5xx status code
func (o *DcimPowerFeedsTraceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power feeds trace o k response a status code equal to that given
func (o *DcimPowerFeedsTraceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power feeds trace o k response
func (o *DcimPowerFeedsTraceOK) Code() int {
	return 200
}

func (o *DcimPowerFeedsTraceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/trace/][%d] dcimPowerFeedsTraceOK %s", 200, payload)
}

func (o *DcimPowerFeedsTraceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/trace/][%d] dcimPowerFeedsTraceOK %s", 200, payload)
}

func (o *DcimPowerFeedsTraceOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsTraceDefault creates a DcimPowerFeedsTraceDefault with default headers values
func NewDcimPowerFeedsTraceDefault(code int) *DcimPowerFeedsTraceDefault {
	return &DcimPowerFeedsTraceDefault{
		_statusCode: code,
	}
}

/*
DcimPowerFeedsTraceDefault describes a response with status code -1, with default header values.

DcimPowerFeedsTraceDefault dcim power feeds trace default
*/
type DcimPowerFeedsTraceDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power feeds trace default response has a 2xx status code
func (o *DcimPowerFeedsTraceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power feeds trace default response has a 3xx status code
func (o *DcimPowerFeedsTraceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power feeds trace default response has a 4xx status code
func (o *DcimPowerFeedsTraceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power feeds trace default response has a 5xx status code
func (o *DcimPowerFeedsTraceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power feeds trace default response a status code equal to that given
func (o *DcimPowerFeedsTraceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power feeds trace default response
func (o *DcimPowerFeedsTraceDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerFeedsTraceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/trace/][%d] dcim_power-feeds_trace default %s", o._statusCode, payload)
}

func (o *DcimPowerFeedsTraceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/trace/][%d] dcim_power-feeds_trace default %s", o._statusCode, payload)
}

func (o *DcimPowerFeedsTraceDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsTraceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
