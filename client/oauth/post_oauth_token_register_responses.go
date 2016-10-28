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

// PostOauthTokenRegisterReader is a Reader for the PostOauthTokenRegister structure.
type PostOauthTokenRegisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOauthTokenRegisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostOauthTokenRegisterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostOauthTokenRegisterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostOauthTokenRegisterOK creates a PostOauthTokenRegisterOK with default headers values
func NewPostOauthTokenRegisterOK() *PostOauthTokenRegisterOK {
	return &PostOauthTokenRegisterOK{}
}

/*PostOauthTokenRegisterOK handles this case with default header values.

Oauth token registration success
*/
type PostOauthTokenRegisterOK struct {
	Payload *models.OauthRegisterTokenResponse
}

func (o *PostOauthTokenRegisterOK) Error() string {
	return fmt.Sprintf("[POST /oauth/token/register][%d] postOauthTokenRegisterOK  %+v", 200, o.Payload)
}

func (o *PostOauthTokenRegisterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OauthRegisterTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthTokenRegisterDefault creates a PostOauthTokenRegisterDefault with default headers values
func NewPostOauthTokenRegisterDefault(code int) *PostOauthTokenRegisterDefault {
	return &PostOauthTokenRegisterDefault{
		_statusCode: code,
	}
}

/*PostOauthTokenRegisterDefault handles this case with default header values.

Failed to register oauth token
*/
type PostOauthTokenRegisterDefault struct {
	_statusCode int

	Payload *models.OauthStatus
}

// Code gets the status code for the post oauth token register default response
func (o *PostOauthTokenRegisterDefault) Code() int {
	return o._statusCode
}

func (o *PostOauthTokenRegisterDefault) Error() string {
	return fmt.Sprintf("[POST /oauth/token/register][%d] PostOauthTokenRegister default  %+v", o._statusCode, o.Payload)
}

func (o *PostOauthTokenRegisterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OauthStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}