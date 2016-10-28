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

// PostOauthGenerateTokenReader is a Reader for the PostOauthGenerateToken structure.
type PostOauthGenerateTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOauthGenerateTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostOauthGenerateTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostOauthGenerateTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostOauthGenerateTokenOK creates a PostOauthGenerateTokenOK with default headers values
func NewPostOauthGenerateTokenOK() *PostOauthGenerateTokenOK {
	return &PostOauthGenerateTokenOK{}
}

/*PostOauthGenerateTokenOK handles this case with default header values.

Oauth token generation success
*/
type PostOauthGenerateTokenOK struct {
	Payload *models.OauthGenerateTokenResponse
}

func (o *PostOauthGenerateTokenOK) Error() string {
	return fmt.Sprintf("[POST /oauth/generate_token][%d] postOauthGenerateTokenOK  %+v", 200, o.Payload)
}

func (o *PostOauthGenerateTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OauthGenerateTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthGenerateTokenDefault creates a PostOauthGenerateTokenDefault with default headers values
func NewPostOauthGenerateTokenDefault(code int) *PostOauthGenerateTokenDefault {
	return &PostOauthGenerateTokenDefault{
		_statusCode: code,
	}
}

/*PostOauthGenerateTokenDefault handles this case with default header values.

Failed to generate oauth token
*/
type PostOauthGenerateTokenDefault struct {
	_statusCode int

	Payload *models.OauthStatus
}

// Code gets the status code for the post oauth generate token default response
func (o *PostOauthGenerateTokenDefault) Code() int {
	return o._statusCode
}

func (o *PostOauthGenerateTokenDefault) Error() string {
	return fmt.Sprintf("[POST /oauth/generate_token][%d] PostOauthGenerateToken default  %+v", o._statusCode, o.Payload)
}

func (o *PostOauthGenerateTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OauthStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
