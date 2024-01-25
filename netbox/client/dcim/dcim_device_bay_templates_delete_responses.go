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
)

// DcimDeviceBayTemplatesDeleteReader is a Reader for the DcimDeviceBayTemplatesDelete structure.
type DcimDeviceBayTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceBayTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceBayTemplatesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBayTemplatesDeleteNoContent creates a DcimDeviceBayTemplatesDeleteNoContent with default headers values
func NewDcimDeviceBayTemplatesDeleteNoContent() *DcimDeviceBayTemplatesDeleteNoContent {
	return &DcimDeviceBayTemplatesDeleteNoContent{}
}

/*
DcimDeviceBayTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceBayTemplatesDeleteNoContent dcim device bay templates delete no content
*/
type DcimDeviceBayTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim device bay templates delete no content response has a 2xx status code
func (o *DcimDeviceBayTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bay templates delete no content response has a 3xx status code
func (o *DcimDeviceBayTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates delete no content response has a 4xx status code
func (o *DcimDeviceBayTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bay templates delete no content response has a 5xx status code
func (o *DcimDeviceBayTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates delete no content response a status code equal to that given
func (o *DcimDeviceBayTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim device bay templates delete no content response
func (o *DcimDeviceBayTemplatesDeleteNoContent) Code() int {
	return 204
}

func (o *DcimDeviceBayTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesDeleteNoContent ", 204)
}

func (o *DcimDeviceBayTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesDeleteNoContent ", 204)
}

func (o *DcimDeviceBayTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimDeviceBayTemplatesDeleteDefault creates a DcimDeviceBayTemplatesDeleteDefault with default headers values
func NewDcimDeviceBayTemplatesDeleteDefault(code int) *DcimDeviceBayTemplatesDeleteDefault {
	return &DcimDeviceBayTemplatesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceBayTemplatesDeleteDefault describes a response with status code -1, with default header values.

DcimDeviceBayTemplatesDeleteDefault dcim device bay templates delete default
*/
type DcimDeviceBayTemplatesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim device bay templates delete default response has a 2xx status code
func (o *DcimDeviceBayTemplatesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device bay templates delete default response has a 3xx status code
func (o *DcimDeviceBayTemplatesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device bay templates delete default response has a 4xx status code
func (o *DcimDeviceBayTemplatesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device bay templates delete default response has a 5xx status code
func (o *DcimDeviceBayTemplatesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device bay templates delete default response a status code equal to that given
func (o *DcimDeviceBayTemplatesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim device bay templates delete default response
func (o *DcimDeviceBayTemplatesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceBayTemplatesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-bay-templates/{id}/][%d] dcim_device-bay-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-bay-templates/{id}/][%d] dcim_device-bay-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
