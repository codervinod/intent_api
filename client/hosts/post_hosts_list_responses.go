package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PostHostsListReader is a Reader for the PostHostsList structure.
type PostHostsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHostsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostHostsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostHostsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostHostsListOK creates a PostHostsListOK with default headers values
func NewPostHostsListOK() *PostHostsListOK {
	return &PostHostsListOK{}
}

/*PostHostsListOK handles this case with default header values.

Success
*/
type PostHostsListOK struct {
	Payload []*models.HostListIntentResponse
}

func (o *PostHostsListOK) Error() string {
	return fmt.Sprintf("[POST /hosts/list][%d] postHostsListOK  %+v", 200, o.Payload)
}

func (o *PostHostsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostHostsListDefault creates a PostHostsListDefault with default headers values
func NewPostHostsListDefault(code int) *PostHostsListDefault {
	return &PostHostsListDefault{
		_statusCode: code,
	}
}

/*PostHostsListDefault handles this case with default header values.

Server Error
*/
type PostHostsListDefault struct {
	_statusCode int

	Payload *models.HostStatus
}

// Code gets the status code for the post hosts list default response
func (o *PostHostsListDefault) Code() int {
	return o._statusCode
}

func (o *PostHostsListDefault) Error() string {
	return fmt.Sprintf("[POST /hosts/list][%d] PostHostsList default  %+v", o._statusCode, o.Payload)
}

func (o *PostHostsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}