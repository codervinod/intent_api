package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// PostPermissionsReader is a Reader for the PostPermissions structure.
type PostPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostPermissionsOK creates a PostPermissionsOK with default headers values
func NewPostPermissionsOK() *PostPermissionsOK {
	return &PostPermissionsOK{}
}

/*PostPermissionsOK handles this case with default header values.

Successful operation
*/
type PostPermissionsOK struct {
	Payload *models.Permission
}

func (o *PostPermissionsOK) Error() string {
	return fmt.Sprintf("[POST /permissions][%d] postPermissionsOK  %+v", 200, o.Payload)
}

func (o *PostPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPermissionsDefault creates a PostPermissionsDefault with default headers values
func NewPostPermissionsDefault(code int) *PostPermissionsDefault {
	return &PostPermissionsDefault{
		_statusCode: code,
	}
}

/*PostPermissionsDefault handles this case with default header values.

Error
*/
type PostPermissionsDefault struct {
	_statusCode int

	Payload *models.PermissionStatus
}

// Code gets the status code for the post permissions default response
func (o *PostPermissionsDefault) Code() int {
	return o._statusCode
}

func (o *PostPermissionsDefault) Error() string {
	return fmt.Sprintf("[POST /permissions][%d] PostPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *PostPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PermissionStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
