package oauth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// DeleteOauthClientClientIDReader is a Reader for the DeleteOauthClientClientID structure.
type DeleteOauthClientClientIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOauthClientClientIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteOauthClientClientIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteOauthClientClientIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteOauthClientClientIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteOauthClientClientIDOK creates a DeleteOauthClientClientIDOK with default headers values
func NewDeleteOauthClientClientIDOK() *DeleteOauthClientClientIDOK {
	return &DeleteOauthClientClientIDOK{}
}

/*DeleteOauthClientClientIDOK handles this case with default header values.

Successful operation
*/
type DeleteOauthClientClientIDOK struct {
}

func (o *DeleteOauthClientClientIDOK) Error() string {
	return fmt.Sprintf("[DELETE /oauth/client/{client_id}][%d] deleteOauthClientClientIdOK ", 200)
}

func (o *DeleteOauthClientClientIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOauthClientClientIDNotFound creates a DeleteOauthClientClientIDNotFound with default headers values
func NewDeleteOauthClientClientIDNotFound() *DeleteOauthClientClientIDNotFound {
	return &DeleteOauthClientClientIDNotFound{}
}

/*DeleteOauthClientClientIDNotFound handles this case with default header values.

Client id does not exists
*/
type DeleteOauthClientClientIDNotFound struct {
	Payload *models.OauthClientStatus
}

func (o *DeleteOauthClientClientIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /oauth/client/{client_id}][%d] deleteOauthClientClientIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOauthClientClientIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OauthClientStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOauthClientClientIDDefault creates a DeleteOauthClientClientIDDefault with default headers values
func NewDeleteOauthClientClientIDDefault(code int) *DeleteOauthClientClientIDDefault {
	return &DeleteOauthClientClientIDDefault{
		_statusCode: code,
	}
}

/*DeleteOauthClientClientIDDefault handles this case with default header values.

Error
*/
type DeleteOauthClientClientIDDefault struct {
	_statusCode int

	Payload *models.OauthClientStatus
}

// Code gets the status code for the delete oauth client client ID default response
func (o *DeleteOauthClientClientIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOauthClientClientIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /oauth/client/{client_id}][%d] DeleteOauthClientClientID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOauthClientClientIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OauthClientStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
