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

// GetUsersMeReader is a Reader for the GetUsersMe structure.
type GetUsersMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetUsersMeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetUsersMeOK creates a GetUsersMeOK with default headers values
func NewGetUsersMeOK() *GetUsersMeOK {
	return &GetUsersMeOK{}
}

/*GetUsersMeOK handles this case with default header values.

Retrieved logged in user object
*/
type GetUsersMeOK struct {
	Payload *models.User
}

func (o *GetUsersMeOK) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] getUsersMeOK  %+v", 200, o.Payload)
}

func (o *GetUsersMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersMeDefault creates a GetUsersMeDefault with default headers values
func NewGetUsersMeDefault(code int) *GetUsersMeDefault {
	return &GetUsersMeDefault{
		_statusCode: code,
	}
}

/*GetUsersMeDefault handles this case with default header values.

Failed to retrieve currently logged in user
*/
type GetUsersMeDefault struct {
	_statusCode int

	Payload *models.UserStatus
}

// Code gets the status code for the get users me default response
func (o *GetUsersMeDefault) Code() int {
	return o._statusCode
}

func (o *GetUsersMeDefault) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] GetUsersMe default  %+v", o._statusCode, o.Payload)
}

func (o *GetUsersMeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}