package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PostProjectsListReader is a Reader for the PostProjectsList structure.
type PostProjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostProjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostProjectsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostProjectsListOK creates a PostProjectsListOK with default headers values
func NewPostProjectsListOK() *PostProjectsListOK {
	return &PostProjectsListOK{}
}

/*PostProjectsListOK handles this case with default header values.

Retrieved all projects
*/
type PostProjectsListOK struct {
	Payload *models.ProjectList
}

func (o *PostProjectsListOK) Error() string {
	return fmt.Sprintf("[POST /projects/list][%d] postProjectsListOK  %+v", 200, o.Payload)
}

func (o *PostProjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProjectsListDefault creates a PostProjectsListDefault with default headers values
func NewPostProjectsListDefault(code int) *PostProjectsListDefault {
	return &PostProjectsListDefault{
		_statusCode: code,
	}
}

/*PostProjectsListDefault handles this case with default header values.

Failed to retrieve all projects
*/
type PostProjectsListDefault struct {
	_statusCode int

	Payload *models.ProjectStatus
}

// Code gets the status code for the post projects list default response
func (o *PostProjectsListDefault) Code() int {
	return o._statusCode
}

func (o *PostProjectsListDefault) Error() string {
	return fmt.Sprintf("[POST /projects/list][%d] PostProjectsList default  %+v", o._statusCode, o.Payload)
}

func (o *PostProjectsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
