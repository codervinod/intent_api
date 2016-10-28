package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// GetUsersUUIDReader is a Reader for the GetUsersUUID structure.
type GetUsersUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetUsersUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetUsersUUIDOK creates a GetUsersUUIDOK with default headers values
func NewGetUsersUUIDOK() *GetUsersUUIDOK {
	return &GetUsersUUIDOK{}
}

/*GetUsersUUIDOK handles this case with default header values.

Retrieved user object
*/
type GetUsersUUIDOK struct {
	Payload *models.User
}

func (o *GetUsersUUIDOK) Error() string {
	return fmt.Sprintf("[GET /users/{uuid}][%d] getUsersUuidOK  %+v", 200, o.Payload)
}

func (o *GetUsersUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUUIDDefault creates a GetUsersUUIDDefault with default headers values
func NewGetUsersUUIDDefault(code int) *GetUsersUUIDDefault {
	return &GetUsersUUIDDefault{
		_statusCode: code,
	}
}

/*GetUsersUUIDDefault handles this case with default header values.

Failed to retrieve specified user
*/
type GetUsersUUIDDefault struct {
	_statusCode int

	Payload *models.UserStatus
}

// Code gets the status code for the get users UUID default response
func (o *GetUsersUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetUsersUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /users/{uuid}][%d] GetUsersUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetUsersUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
