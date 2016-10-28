package workflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// DeleteWorkflowsUUIDReader is a Reader for the DeleteWorkflowsUUID structure.
type DeleteWorkflowsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkflowsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteWorkflowsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteWorkflowsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteWorkflowsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteWorkflowsUUIDOK creates a DeleteWorkflowsUUIDOK with default headers values
func NewDeleteWorkflowsUUIDOK() *DeleteWorkflowsUUIDOK {
	return &DeleteWorkflowsUUIDOK{}
}

/*DeleteWorkflowsUUIDOK handles this case with default header values.

Request Succeeded
*/
type DeleteWorkflowsUUIDOK struct {
}

func (o *DeleteWorkflowsUUIDOK) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{uuid}][%d] deleteWorkflowsUuidOK ", 200)
}

func (o *DeleteWorkflowsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkflowsUUIDNotFound creates a DeleteWorkflowsUUIDNotFound with default headers values
func NewDeleteWorkflowsUUIDNotFound() *DeleteWorkflowsUUIDNotFound {
	return &DeleteWorkflowsUUIDNotFound{}
}

/*DeleteWorkflowsUUIDNotFound handles this case with default header values.

Invalid UUID Provided
*/
type DeleteWorkflowsUUIDNotFound struct {
	Payload *models.WorkflowStatus
}

func (o *DeleteWorkflowsUUIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{uuid}][%d] deleteWorkflowsUuidNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWorkflowsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkflowsUUIDDefault creates a DeleteWorkflowsUUIDDefault with default headers values
func NewDeleteWorkflowsUUIDDefault(code int) *DeleteWorkflowsUUIDDefault {
	return &DeleteWorkflowsUUIDDefault{
		_statusCode: code,
	}
}

/*DeleteWorkflowsUUIDDefault handles this case with default header values.

Internal Error
*/
type DeleteWorkflowsUUIDDefault struct {
	_statusCode int

	Payload *models.WorkflowStatus
}

// Code gets the status code for the delete workflows UUID default response
func (o *DeleteWorkflowsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteWorkflowsUUIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{uuid}][%d] DeleteWorkflowsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteWorkflowsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
