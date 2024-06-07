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

// DcimInventoryItemsUpdateReader is a Reader for the DcimInventoryItemsUpdate structure.
type DcimInventoryItemsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemsUpdateOK creates a DcimInventoryItemsUpdateOK with default headers values
func NewDcimInventoryItemsUpdateOK() *DcimInventoryItemsUpdateOK {
	return &DcimInventoryItemsUpdateOK{}
}

/*
DcimInventoryItemsUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemsUpdateOK dcim inventory items update o k
*/
type DcimInventoryItemsUpdateOK struct {
	Payload *models.InventoryItem
}

// IsSuccess returns true when this dcim inventory items update o k response has a 2xx status code
func (o *DcimInventoryItemsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory items update o k response has a 3xx status code
func (o *DcimInventoryItemsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory items update o k response has a 4xx status code
func (o *DcimInventoryItemsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory items update o k response has a 5xx status code
func (o *DcimInventoryItemsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory items update o k response a status code equal to that given
func (o *DcimInventoryItemsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim inventory items update o k response
func (o *DcimInventoryItemsUpdateOK) Code() int {
	return 200
}

func (o *DcimInventoryItemsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/inventory-items/{id}/][%d] dcimInventoryItemsUpdateOK %s", 200, payload)
}

func (o *DcimInventoryItemsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/inventory-items/{id}/][%d] dcimInventoryItemsUpdateOK %s", 200, payload)
}

func (o *DcimInventoryItemsUpdateOK) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemsUpdateDefault creates a DcimInventoryItemsUpdateDefault with default headers values
func NewDcimInventoryItemsUpdateDefault(code int) *DcimInventoryItemsUpdateDefault {
	return &DcimInventoryItemsUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemsUpdateDefault describes a response with status code -1, with default header values.

DcimInventoryItemsUpdateDefault dcim inventory items update default
*/
type DcimInventoryItemsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim inventory items update default response has a 2xx status code
func (o *DcimInventoryItemsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory items update default response has a 3xx status code
func (o *DcimInventoryItemsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory items update default response has a 4xx status code
func (o *DcimInventoryItemsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory items update default response has a 5xx status code
func (o *DcimInventoryItemsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory items update default response a status code equal to that given
func (o *DcimInventoryItemsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim inventory items update default response
func (o *DcimInventoryItemsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/inventory-items/{id}/][%d] dcim_inventory-items_update default %s", o._statusCode, payload)
}

func (o *DcimInventoryItemsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/inventory-items/{id}/][%d] dcim_inventory-items_update default %s", o._statusCode, payload)
}

func (o *DcimInventoryItemsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
