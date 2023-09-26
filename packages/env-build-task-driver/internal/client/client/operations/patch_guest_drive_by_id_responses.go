// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/e2b-dev/infra/packages/env-build-task-driver/internal/client/models"
)

// PatchGuestDriveByIDReader is a Reader for the PatchGuestDriveByID structure.
type PatchGuestDriveByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGuestDriveByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchGuestDriveByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchGuestDriveByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchGuestDriveByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchGuestDriveByIDNoContent creates a PatchGuestDriveByIDNoContent with default headers values
func NewPatchGuestDriveByIDNoContent() *PatchGuestDriveByIDNoContent {
	return &PatchGuestDriveByIDNoContent{}
}

/*
PatchGuestDriveByIDNoContent describes a response with status code 204, with default header values.

Drive updated
*/
type PatchGuestDriveByIDNoContent struct {
}

// IsSuccess returns true when this patch guest drive by Id no content response has a 2xx status code
func (o *PatchGuestDriveByIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch guest drive by Id no content response has a 3xx status code
func (o *PatchGuestDriveByIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch guest drive by Id no content response has a 4xx status code
func (o *PatchGuestDriveByIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch guest drive by Id no content response has a 5xx status code
func (o *PatchGuestDriveByIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch guest drive by Id no content response a status code equal to that given
func (o *PatchGuestDriveByIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the patch guest drive by Id no content response
func (o *PatchGuestDriveByIDNoContent) Code() int {
	return 204
}

func (o *PatchGuestDriveByIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /drives/{drive_id}][%d] patchGuestDriveByIdNoContent ", 204)
}

func (o *PatchGuestDriveByIDNoContent) String() string {
	return fmt.Sprintf("[PATCH /drives/{drive_id}][%d] patchGuestDriveByIdNoContent ", 204)
}

func (o *PatchGuestDriveByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGuestDriveByIDBadRequest creates a PatchGuestDriveByIDBadRequest with default headers values
func NewPatchGuestDriveByIDBadRequest() *PatchGuestDriveByIDBadRequest {
	return &PatchGuestDriveByIDBadRequest{}
}

/*
PatchGuestDriveByIDBadRequest describes a response with status code 400, with default header values.

Drive cannot be updated due to bad input
*/
type PatchGuestDriveByIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this patch guest drive by Id bad request response has a 2xx status code
func (o *PatchGuestDriveByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch guest drive by Id bad request response has a 3xx status code
func (o *PatchGuestDriveByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch guest drive by Id bad request response has a 4xx status code
func (o *PatchGuestDriveByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch guest drive by Id bad request response has a 5xx status code
func (o *PatchGuestDriveByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch guest drive by Id bad request response a status code equal to that given
func (o *PatchGuestDriveByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch guest drive by Id bad request response
func (o *PatchGuestDriveByIDBadRequest) Code() int {
	return 400
}

func (o *PatchGuestDriveByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /drives/{drive_id}][%d] patchGuestDriveByIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchGuestDriveByIDBadRequest) String() string {
	return fmt.Sprintf("[PATCH /drives/{drive_id}][%d] patchGuestDriveByIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchGuestDriveByIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchGuestDriveByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGuestDriveByIDDefault creates a PatchGuestDriveByIDDefault with default headers values
func NewPatchGuestDriveByIDDefault(code int) *PatchGuestDriveByIDDefault {
	return &PatchGuestDriveByIDDefault{
		_statusCode: code,
	}
}

/*
PatchGuestDriveByIDDefault describes a response with status code -1, with default header values.

Internal server error.
*/
type PatchGuestDriveByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this patch guest drive by ID default response has a 2xx status code
func (o *PatchGuestDriveByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch guest drive by ID default response has a 3xx status code
func (o *PatchGuestDriveByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch guest drive by ID default response has a 4xx status code
func (o *PatchGuestDriveByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch guest drive by ID default response has a 5xx status code
func (o *PatchGuestDriveByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch guest drive by ID default response a status code equal to that given
func (o *PatchGuestDriveByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch guest drive by ID default response
func (o *PatchGuestDriveByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchGuestDriveByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /drives/{drive_id}][%d] patchGuestDriveByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchGuestDriveByIDDefault) String() string {
	return fmt.Sprintf("[PATCH /drives/{drive_id}][%d] patchGuestDriveByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchGuestDriveByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchGuestDriveByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
