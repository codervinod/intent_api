package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// PutJobsUUIDReader is a Reader for the PutJobsUUID structure.
type PutJobsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutJobsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutJobsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutJobsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutJobsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutJobsUUIDOK creates a PutJobsUUIDOK with default headers values
func NewPutJobsUUIDOK() *PutJobsUUIDOK {
	return &PutJobsUUIDOK{}
}

/*PutJobsUUIDOK handles this case with default header values.

Request Succeeded
*/
type PutJobsUUIDOK struct {
	Payload *models.JobIntentResponse
}

func (o *PutJobsUUIDOK) Error() string {
	return fmt.Sprintf("[PUT /jobs/{uuid}][%d] putJobsUuidOK  %+v", 200, o.Payload)
}

func (o *PutJobsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutJobsUUIDNotFound creates a PutJobsUUIDNotFound with default headers values
func NewPutJobsUUIDNotFound() *PutJobsUUIDNotFound {
	return &PutJobsUUIDNotFound{}
}

/*PutJobsUUIDNotFound handles this case with default header values.

Entity not found
*/
type PutJobsUUIDNotFound struct {
	Payload *models.JobStatus
}

func (o *PutJobsUUIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /jobs/{uuid}][%d] putJobsUuidNotFound  %+v", 404, o.Payload)
}

func (o *PutJobsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutJobsUUIDDefault creates a PutJobsUUIDDefault with default headers values
func NewPutJobsUUIDDefault(code int) *PutJobsUUIDDefault {
	return &PutJobsUUIDDefault{
		_statusCode: code,
	}
}

/*PutJobsUUIDDefault handles this case with default header values.

Internal Error
*/
type PutJobsUUIDDefault struct {
	_statusCode int

	Payload *models.JobStatus
}

// Code gets the status code for the put jobs UUID default response
func (o *PutJobsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *PutJobsUUIDDefault) Error() string {
	return fmt.Sprintf("[PUT /jobs/{uuid}][%d] PutJobsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *PutJobsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
