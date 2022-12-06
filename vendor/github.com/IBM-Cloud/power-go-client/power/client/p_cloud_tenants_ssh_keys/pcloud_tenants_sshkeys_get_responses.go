// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants_ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudTenantsSshkeysGetReader is a Reader for the PcloudTenantsSshkeysGet structure.
type PcloudTenantsSshkeysGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTenantsSshkeysGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudTenantsSshkeysGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudTenantsSshkeysGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudTenantsSshkeysGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudTenantsSshkeysGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudTenantsSshkeysGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudTenantsSshkeysGetOK creates a PcloudTenantsSshkeysGetOK with default headers values
func NewPcloudTenantsSshkeysGetOK() *PcloudTenantsSshkeysGetOK {
	return &PcloudTenantsSshkeysGetOK{}
}

/* PcloudTenantsSshkeysGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudTenantsSshkeysGetOK struct {
	Payload *models.SSHKey
}

func (o *PcloudTenantsSshkeysGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysGetOK  %+v", 200, o.Payload)
}
func (o *PcloudTenantsSshkeysGetOK) GetPayload() *models.SSHKey {
	return o.Payload
}

func (o *PcloudTenantsSshkeysGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysGetBadRequest creates a PcloudTenantsSshkeysGetBadRequest with default headers values
func NewPcloudTenantsSshkeysGetBadRequest() *PcloudTenantsSshkeysGetBadRequest {
	return &PcloudTenantsSshkeysGetBadRequest{}
}

/* PcloudTenantsSshkeysGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudTenantsSshkeysGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysGetBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudTenantsSshkeysGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysGetUnauthorized creates a PcloudTenantsSshkeysGetUnauthorized with default headers values
func NewPcloudTenantsSshkeysGetUnauthorized() *PcloudTenantsSshkeysGetUnauthorized {
	return &PcloudTenantsSshkeysGetUnauthorized{}
}

/* PcloudTenantsSshkeysGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudTenantsSshkeysGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysGetUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudTenantsSshkeysGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysGetNotFound creates a PcloudTenantsSshkeysGetNotFound with default headers values
func NewPcloudTenantsSshkeysGetNotFound() *PcloudTenantsSshkeysGetNotFound {
	return &PcloudTenantsSshkeysGetNotFound{}
}

/* PcloudTenantsSshkeysGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudTenantsSshkeysGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysGetNotFound  %+v", 404, o.Payload)
}
func (o *PcloudTenantsSshkeysGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysGetInternalServerError creates a PcloudTenantsSshkeysGetInternalServerError with default headers values
func NewPcloudTenantsSshkeysGetInternalServerError() *PcloudTenantsSshkeysGetInternalServerError {
	return &PcloudTenantsSshkeysGetInternalServerError{}
}

/* PcloudTenantsSshkeysGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudTenantsSshkeysGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysGetInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudTenantsSshkeysGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
