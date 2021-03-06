package extension

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// PostExtensionsListReader is a Reader for the PostExtensionsList structure.
type PostExtensionsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExtensionsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostExtensionsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostExtensionsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostExtensionsListOK creates a PostExtensionsListOK with default headers values
func NewPostExtensionsListOK() *PostExtensionsListOK {
	return &PostExtensionsListOK{}
}

/*PostExtensionsListOK handles this case with default header values.

Retrieved all installed extensions
*/
type PostExtensionsListOK struct {
	Payload *models.ExtensionListIntentResponse
}

func (o *PostExtensionsListOK) Error() string {
	return fmt.Sprintf("[POST /extensions/list][%d] postExtensionsListOK  %+v", 200, o.Payload)
}

func (o *PostExtensionsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtensionListIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExtensionsListDefault creates a PostExtensionsListDefault with default headers values
func NewPostExtensionsListDefault(code int) *PostExtensionsListDefault {
	return &PostExtensionsListDefault{
		_statusCode: code,
	}
}

/*PostExtensionsListDefault handles this case with default header values.

Failed to retrieve all extensions
*/
type PostExtensionsListDefault struct {
	_statusCode int

	Payload *models.ExtensionStatus
}

// Code gets the status code for the post extensions list default response
func (o *PostExtensionsListDefault) Code() int {
	return o._statusCode
}

func (o *PostExtensionsListDefault) Error() string {
	return fmt.Sprintf("[POST /extensions/list][%d] PostExtensionsList default  %+v", o._statusCode, o.Payload)
}

func (o *PostExtensionsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtensionStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
