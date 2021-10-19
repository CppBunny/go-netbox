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

// DcimLocationsReadReader is a Reader for the DcimLocationsRead structure.
type DcimLocationsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimLocationsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimLocationsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimLocationsReadOK creates a DcimLocationsReadOK with default headers values
func NewDcimLocationsReadOK() *DcimLocationsReadOK {
	return &DcimLocationsReadOK{}
}

/* DcimLocationsReadOK describes a response with status code 200, with default header values.

DcimLocationsReadOK dcim locations read o k
*/
type DcimLocationsReadOK struct {
	Payload *models.Location
}

func (o *DcimLocationsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/locations/{id}/][%d] dcimLocationsReadOK  %+v", 200, o.Payload)
}
func (o *DcimLocationsReadOK) GetPayload() *models.Location {
	return o.Payload
}

func (o *DcimLocationsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimLocationsReadDefault creates a DcimLocationsReadDefault with default headers values
func NewDcimLocationsReadDefault(code int) *DcimLocationsReadDefault {
	return &DcimLocationsReadDefault{
		_statusCode: code,
	}
}

/* DcimLocationsReadDefault describes a response with status code -1, with default header values.

DcimLocationsReadDefault dcim locations read default
*/
type DcimLocationsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim locations read default response
func (o *DcimLocationsReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimLocationsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/locations/{id}/][%d] dcim_locations_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimLocationsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}