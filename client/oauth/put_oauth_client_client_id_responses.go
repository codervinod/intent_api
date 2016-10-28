package oauth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// PutOauthClientClientIDReader is a Reader for the PutOauthClientClientID structure.
type PutOauthClientClientIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOauthClientClientIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutOauthClientClientIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutOauthClientClientIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutOauthClientClientIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutOauthClientClientIDOK creates a PutOauthClientClientIDOK with default headers values
func NewPutOauthClientClientIDOK() *PutOauthClientClientIDOK {
	return &PutOauthClientClientIDOK{}
}

/*PutOauthClientClientIDOK handles this case with default header values.

Successful operation
*/
type PutOauthClientClientIDOK struct {
	Payload *models.OauthClientResponse
}

func (o *PutOauthClientClientIDOK) Error() string {
	return fmt.Sprintf("[PUT /oauth/client/{client_id}][%d] putOauthClientClientIdOK  %+v", 200, o.Payload)
}

func (o *PutOauthClientClientIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OauthClientResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientClientIDNotFound creates a PutOauthClientClientIDNotFound with default headers values
func NewPutOauthClientClientIDNotFound() *PutOauthClientClientIDNotFound {
	return &PutOauthClientClientIDNotFound{}
}

/*PutOauthClientClientIDNotFound handles this case with default header values.

Client identifier does not exists
*/
type PutOauthClientClientIDNotFound struct {
	Payload *models.OauthClientStatus
}

func (o *PutOauthClientClientIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /oauth/client/{client_id}][%d] putOauthClientClientIdNotFound  %+v", 404, o.Payload)
}

func (o *PutOauthClientClientIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OauthClientStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientClientIDDefault creates a PutOauthClientClientIDDefault with default headers values
func NewPutOauthClientClientIDDefault(code int) *PutOauthClientClientIDDefault {
	return &PutOauthClientClientIDDefault{
		_statusCode: code,
	}
}

/*PutOauthClientClientIDDefault handles this case with default header values.

Error
*/
type PutOauthClientClientIDDefault struct {
	_statusCode int

	Payload *models.OauthClientStatus
}

// Code gets the status code for the put oauth client client ID default response
func (o *PutOauthClientClientIDDefault) Code() int {
	return o._statusCode
}

func (o *PutOauthClientClientIDDefault) Error() string {
	return fmt.Sprintf("[PUT /oauth/client/{client_id}][%d] PutOauthClientClientID default  %+v", o._statusCode, o.Payload)
}

func (o *PutOauthClientClientIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OauthClientStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
