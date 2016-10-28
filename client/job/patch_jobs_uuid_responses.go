package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PatchJobsUUIDReader is a Reader for the PatchJobsUUID structure.
type PatchJobsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchJobsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchJobsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPatchJobsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchJobsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPatchJobsUUIDOK creates a PatchJobsUUIDOK with default headers values
func NewPatchJobsUUIDOK() *PatchJobsUUIDOK {
	return &PatchJobsUUIDOK{}
}

/*PatchJobsUUIDOK handles this case with default header values.

Request Succeeded
*/
type PatchJobsUUIDOK struct {
	Payload *models.JobIntentResponse
}

func (o *PatchJobsUUIDOK) Error() string {
	return fmt.Sprintf("[PATCH /jobs/{uuid}][%d] patchJobsUuidOK  %+v", 200, o.Payload)
}

func (o *PatchJobsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJobsUUIDNotFound creates a PatchJobsUUIDNotFound with default headers values
func NewPatchJobsUUIDNotFound() *PatchJobsUUIDNotFound {
	return &PatchJobsUUIDNotFound{}
}

/*PatchJobsUUIDNotFound handles this case with default header values.

Invalid UUID Provided
*/
type PatchJobsUUIDNotFound struct {
	Payload *models.JobStatus
}

func (o *PatchJobsUUIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /jobs/{uuid}][%d] patchJobsUuidNotFound  %+v", 404, o.Payload)
}

func (o *PatchJobsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJobsUUIDDefault creates a PatchJobsUUIDDefault with default headers values
func NewPatchJobsUUIDDefault(code int) *PatchJobsUUIDDefault {
	return &PatchJobsUUIDDefault{
		_statusCode: code,
	}
}

/*PatchJobsUUIDDefault handles this case with default header values.

Internal Error
*/
type PatchJobsUUIDDefault struct {
	_statusCode int

	Payload *models.JobStatus
}

// Code gets the status code for the patch jobs UUID default response
func (o *PatchJobsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchJobsUUIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /jobs/{uuid}][%d] PatchJobsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchJobsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
