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

package vpn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VpnTunnelsDeleteReader is a Reader for the VpnTunnelsDelete structure.
type VpnTunnelsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VpnTunnelsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVpnTunnelsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVpnTunnelsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVpnTunnelsDeleteNoContent creates a VpnTunnelsDeleteNoContent with default headers values
func NewVpnTunnelsDeleteNoContent() *VpnTunnelsDeleteNoContent {
	return &VpnTunnelsDeleteNoContent{}
}

/*
VpnTunnelsDeleteNoContent describes a response with status code 204, with default header values.

VpnTunnelsDeleteNoContent vpn tunnels delete no content
*/
type VpnTunnelsDeleteNoContent struct {
}

// IsSuccess returns true when this vpn tunnels delete no content response has a 2xx status code
func (o *VpnTunnelsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vpn tunnels delete no content response has a 3xx status code
func (o *VpnTunnelsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vpn tunnels delete no content response has a 4xx status code
func (o *VpnTunnelsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this vpn tunnels delete no content response has a 5xx status code
func (o *VpnTunnelsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this vpn tunnels delete no content response a status code equal to that given
func (o *VpnTunnelsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the vpn tunnels delete no content response
func (o *VpnTunnelsDeleteNoContent) Code() int {
	return 204
}

func (o *VpnTunnelsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /vpn/tunnels/{id}/][%d] vpnTunnelsDeleteNoContent", 204)
}

func (o *VpnTunnelsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /vpn/tunnels/{id}/][%d] vpnTunnelsDeleteNoContent", 204)
}

func (o *VpnTunnelsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVpnTunnelsDeleteDefault creates a VpnTunnelsDeleteDefault with default headers values
func NewVpnTunnelsDeleteDefault(code int) *VpnTunnelsDeleteDefault {
	return &VpnTunnelsDeleteDefault{
		_statusCode: code,
	}
}

/*
VpnTunnelsDeleteDefault describes a response with status code -1, with default header values.

VpnTunnelsDeleteDefault vpn tunnels delete default
*/
type VpnTunnelsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this vpn tunnels delete default response has a 2xx status code
func (o *VpnTunnelsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vpn tunnels delete default response has a 3xx status code
func (o *VpnTunnelsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vpn tunnels delete default response has a 4xx status code
func (o *VpnTunnelsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vpn tunnels delete default response has a 5xx status code
func (o *VpnTunnelsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vpn tunnels delete default response a status code equal to that given
func (o *VpnTunnelsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vpn tunnels delete default response
func (o *VpnTunnelsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VpnTunnelsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /vpn/tunnels/{id}/][%d] vpn_tunnels_delete default %s", o._statusCode, payload)
}

func (o *VpnTunnelsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /vpn/tunnels/{id}/][%d] vpn_tunnels_delete default %s", o._statusCode, payload)
}

func (o *VpnTunnelsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VpnTunnelsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
