package extension

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PostExtensionsReader is a Reader for the PostExtensions structure.
type PostExtensionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExtensionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostExtensionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostExtensionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostExtensionsOK creates a PostExtensionsOK with default headers values
func NewPostExtensionsOK() *PostExtensionsOK {
	return &PostExtensionsOK{}
}

/*PostExtensionsOK handles this case with default header values.

Successful operation
*/
type PostExtensionsOK struct {
	Payload *models.ExtensionIntentResponse
}

func (o *PostExtensionsOK) Error() string {
	return fmt.Sprintf("[POST /extensions][%d] postExtensionsOK  %+v", 200, o.Payload)
}

func (o *PostExtensionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtensionIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExtensionsDefault creates a PostExtensionsDefault with default headers values
func NewPostExtensionsDefault(code int) *PostExtensionsDefault {
	return &PostExtensionsDefault{
		_statusCode: code,
	}
}

/*PostExtensionsDefault handles this case with default header values.

Error
*/
type PostExtensionsDefault struct {
	_statusCode int

	Payload *models.ExtensionStatus
}

// Code gets the status code for the post extensions default response
func (o *PostExtensionsDefault) Code() int {
	return o._statusCode
}

func (o *PostExtensionsDefault) Error() string {
	return fmt.Sprintf("[POST /extensions][%d] PostExtensions default  %+v", o._statusCode, o.Payload)
}

func (o *PostExtensionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtensionStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}