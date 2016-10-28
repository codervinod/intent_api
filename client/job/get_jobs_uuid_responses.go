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

// GetJobsUUIDReader is a Reader for the GetJobsUUID structure.
type GetJobsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetJobsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetJobsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetJobsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetJobsUUIDOK creates a GetJobsUUIDOK with default headers values
func NewGetJobsUUIDOK() *GetJobsUUIDOK {
	return &GetJobsUUIDOK{}
}

/*GetJobsUUIDOK handles this case with default header values.

Success
*/
type GetJobsUUIDOK struct {
	Payload *models.JobIntentResponse
}

func (o *GetJobsUUIDOK) Error() string {
	return fmt.Sprintf("[GET /jobs/{uuid}][%d] getJobsUuidOK  %+v", 200, o.Payload)
}

func (o *GetJobsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobsUUIDNotFound creates a GetJobsUUIDNotFound with default headers values
func NewGetJobsUUIDNotFound() *GetJobsUUIDNotFound {
	return &GetJobsUUIDNotFound{}
}

/*GetJobsUUIDNotFound handles this case with default header values.

Entity not found
*/
type GetJobsUUIDNotFound struct {
	Payload *models.JobStatus
}

func (o *GetJobsUUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /jobs/{uuid}][%d] getJobsUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetJobsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobsUUIDDefault creates a GetJobsUUIDDefault with default headers values
func NewGetJobsUUIDDefault(code int) *GetJobsUUIDDefault {
	return &GetJobsUUIDDefault{
		_statusCode: code,
	}
}

/*GetJobsUUIDDefault handles this case with default header values.

Internal Error
*/
type GetJobsUUIDDefault struct {
	_statusCode int

	Payload *models.JobStatus
}

// Code gets the status code for the get jobs UUID default response
func (o *GetJobsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetJobsUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /jobs/{uuid}][%d] GetJobsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetJobsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}