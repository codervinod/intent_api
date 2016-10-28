package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PostGroupsReader is a Reader for the PostGroups structure.
type PostGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostGroupsOK creates a PostGroupsOK with default headers values
func NewPostGroupsOK() *PostGroupsOK {
	return &PostGroupsOK{}
}

/*PostGroupsOK handles this case with default header values.

Success
*/
type PostGroupsOK struct {
	Payload *models.GroupsGetEntitiesResponse
}

func (o *PostGroupsOK) Error() string {
	return fmt.Sprintf("[POST /groups][%d] postGroupsOK  %+v", 200, o.Payload)
}

func (o *PostGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupsGetEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsDefault creates a PostGroupsDefault with default headers values
func NewPostGroupsDefault(code int) *PostGroupsDefault {
	return &PostGroupsDefault{
		_statusCode: code,
	}
}

/*PostGroupsDefault handles this case with default header values.

Error
*/
type PostGroupsDefault struct {
	_statusCode int

	Payload *models.Status
}

// Code gets the status code for the post groups default response
func (o *PostGroupsDefault) Code() int {
	return o._statusCode
}

func (o *PostGroupsDefault) Error() string {
	return fmt.Sprintf("[POST /groups][%d] PostGroups default  %+v", o._statusCode, o.Payload)
}

func (o *PostGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
