package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PutUsersUUIDReader is a Reader for the PutUsersUUID structure.
type PutUsersUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUsersUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutUsersUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutUsersUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutUsersUUIDOK creates a PutUsersUUIDOK with default headers values
func NewPutUsersUUIDOK() *PutUsersUUIDOK {
	return &PutUsersUUIDOK{}
}

/*PutUsersUUIDOK handles this case with default header values.

Updated user object
*/
type PutUsersUUIDOK struct {
	Payload *models.User
}

func (o *PutUsersUUIDOK) Error() string {
	return fmt.Sprintf("[PUT /users/{uuid}][%d] putUsersUuidOK  %+v", 200, o.Payload)
}

func (o *PutUsersUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUsersUUIDDefault creates a PutUsersUUIDDefault with default headers values
func NewPutUsersUUIDDefault(code int) *PutUsersUUIDDefault {
	return &PutUsersUUIDDefault{
		_statusCode: code,
	}
}

/*PutUsersUUIDDefault handles this case with default header values.

Failed to updated user
*/
type PutUsersUUIDDefault struct {
	_statusCode int

	Payload *models.UserStatus
}

// Code gets the status code for the put users UUID default response
func (o *PutUsersUUIDDefault) Code() int {
	return o._statusCode
}

func (o *PutUsersUUIDDefault) Error() string {
	return fmt.Sprintf("[PUT /users/{uuid}][%d] PutUsersUUID default  %+v", o._statusCode, o.Payload)
}

func (o *PutUsersUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
