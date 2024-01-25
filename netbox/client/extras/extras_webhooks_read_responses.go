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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasWebhooksReadReader is a Reader for the ExtrasWebhooksRead structure.
type ExtrasWebhooksReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasWebhooksReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasWebhooksReadOK creates a ExtrasWebhooksReadOK with default headers values
func NewExtrasWebhooksReadOK() *ExtrasWebhooksReadOK {
	return &ExtrasWebhooksReadOK{}
}

/*
ExtrasWebhooksReadOK describes a response with status code 200, with default header values.

ExtrasWebhooksReadOK extras webhooks read o k
*/
type ExtrasWebhooksReadOK struct {
	Payload *models.Webhook
}

// IsSuccess returns true when this extras webhooks read o k response has a 2xx status code
func (o *ExtrasWebhooksReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras webhooks read o k response has a 3xx status code
func (o *ExtrasWebhooksReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras webhooks read o k response has a 4xx status code
func (o *ExtrasWebhooksReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras webhooks read o k response has a 5xx status code
func (o *ExtrasWebhooksReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras webhooks read o k response a status code equal to that given
func (o *ExtrasWebhooksReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras webhooks read o k response
func (o *ExtrasWebhooksReadOK) Code() int {
	return 200
}

func (o *ExtrasWebhooksReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/webhooks/{id}/][%d] extrasWebhooksReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasWebhooksReadOK) String() string {
	return fmt.Sprintf("[GET /extras/webhooks/{id}/][%d] extrasWebhooksReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasWebhooksReadOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasWebhooksReadDefault creates a ExtrasWebhooksReadDefault with default headers values
func NewExtrasWebhooksReadDefault(code int) *ExtrasWebhooksReadDefault {
	return &ExtrasWebhooksReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasWebhooksReadDefault describes a response with status code -1, with default header values.

ExtrasWebhooksReadDefault extras webhooks read default
*/
type ExtrasWebhooksReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras webhooks read default response has a 2xx status code
func (o *ExtrasWebhooksReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras webhooks read default response has a 3xx status code
func (o *ExtrasWebhooksReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras webhooks read default response has a 4xx status code
func (o *ExtrasWebhooksReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras webhooks read default response has a 5xx status code
func (o *ExtrasWebhooksReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras webhooks read default response a status code equal to that given
func (o *ExtrasWebhooksReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras webhooks read default response
func (o *ExtrasWebhooksReadDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasWebhooksReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/webhooks/{id}/][%d] extras_webhooks_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasWebhooksReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/webhooks/{id}/][%d] extras_webhooks_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasWebhooksReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasWebhooksReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
