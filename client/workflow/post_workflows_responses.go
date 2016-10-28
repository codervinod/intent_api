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

// PostWorkflowsReader is a Reader for the PostWorkflows structure.
type PostWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostWorkflowsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostWorkflowsOK creates a PostWorkflowsOK with default headers values
func NewPostWorkflowsOK() *PostWorkflowsOK {
	return &PostWorkflowsOK{}
}

/*PostWorkflowsOK handles this case with default header values.

Request Succeeded
*/
type PostWorkflowsOK struct {
	Payload *models.WorkflowIntentResponse
}

func (o *PostWorkflowsOK) Error() string {
	return fmt.Sprintf("[POST /workflows][%d] postWorkflowsOK  %+v", 200, o.Payload)
}

func (o *PostWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkflowsDefault creates a PostWorkflowsDefault with default headers values
func NewPostWorkflowsDefault(code int) *PostWorkflowsDefault {
	return &PostWorkflowsDefault{
		_statusCode: code,
	}
}

/*PostWorkflowsDefault handles this case with default header values.

Internal Error
*/
type PostWorkflowsDefault struct {
	_statusCode int

	Payload *models.WorkflowStatus
}

// Code gets the status code for the post workflows default response
func (o *PostWorkflowsDefault) Code() int {
	return o._statusCode
}

func (o *PostWorkflowsDefault) Error() string {
	return fmt.Sprintf("[POST /workflows][%d] PostWorkflows default  %+v", o._statusCode, o.Payload)
}

func (o *PostWorkflowsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}