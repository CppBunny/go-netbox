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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CppBunny/go-netbox/netbox/models"
)

// ExtrasJournalEntriesUpdateReader is a Reader for the ExtrasJournalEntriesUpdate structure.
type ExtrasJournalEntriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJournalEntriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasJournalEntriesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasJournalEntriesUpdateOK creates a ExtrasJournalEntriesUpdateOK with default headers values
func NewExtrasJournalEntriesUpdateOK() *ExtrasJournalEntriesUpdateOK {
	return &ExtrasJournalEntriesUpdateOK{}
}

/*
ExtrasJournalEntriesUpdateOK describes a response with status code 200, with default header values.

ExtrasJournalEntriesUpdateOK extras journal entries update o k
*/
type ExtrasJournalEntriesUpdateOK struct {
	Payload *models.JournalEntry
}

// IsSuccess returns true when this extras journal entries update o k response has a 2xx status code
func (o *ExtrasJournalEntriesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras journal entries update o k response has a 3xx status code
func (o *ExtrasJournalEntriesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras journal entries update o k response has a 4xx status code
func (o *ExtrasJournalEntriesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras journal entries update o k response has a 5xx status code
func (o *ExtrasJournalEntriesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras journal entries update o k response a status code equal to that given
func (o *ExtrasJournalEntriesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras journal entries update o k response
func (o *ExtrasJournalEntriesUpdateOK) Code() int {
	return 200
}

func (o *ExtrasJournalEntriesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/journal-entries/{id}/][%d] extrasJournalEntriesUpdateOK %s", 200, payload)
}

func (o *ExtrasJournalEntriesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/journal-entries/{id}/][%d] extrasJournalEntriesUpdateOK %s", 200, payload)
}

func (o *ExtrasJournalEntriesUpdateOK) GetPayload() *models.JournalEntry {
	return o.Payload
}

func (o *ExtrasJournalEntriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JournalEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasJournalEntriesUpdateDefault creates a ExtrasJournalEntriesUpdateDefault with default headers values
func NewExtrasJournalEntriesUpdateDefault(code int) *ExtrasJournalEntriesUpdateDefault {
	return &ExtrasJournalEntriesUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasJournalEntriesUpdateDefault describes a response with status code -1, with default header values.

ExtrasJournalEntriesUpdateDefault extras journal entries update default
*/
type ExtrasJournalEntriesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras journal entries update default response has a 2xx status code
func (o *ExtrasJournalEntriesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras journal entries update default response has a 3xx status code
func (o *ExtrasJournalEntriesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras journal entries update default response has a 4xx status code
func (o *ExtrasJournalEntriesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras journal entries update default response has a 5xx status code
func (o *ExtrasJournalEntriesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras journal entries update default response a status code equal to that given
func (o *ExtrasJournalEntriesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras journal entries update default response
func (o *ExtrasJournalEntriesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasJournalEntriesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/journal-entries/{id}/][%d] extras_journal-entries_update default %s", o._statusCode, payload)
}

func (o *ExtrasJournalEntriesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/journal-entries/{id}/][%d] extras_journal-entries_update default %s", o._statusCode, payload)
}

func (o *ExtrasJournalEntriesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasJournalEntriesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
