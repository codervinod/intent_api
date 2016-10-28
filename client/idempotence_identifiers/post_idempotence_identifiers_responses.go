package idempotence_identifiers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PostIdempotenceIdentifiersReader is a Reader for the PostIdempotenceIdentifiers structure.
type PostIdempotenceIdentifiersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIdempotenceIdentifiersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostIdempotenceIdentifiersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostIdempotenceIdentifiersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostIdempotenceIdentifiersOK creates a PostIdempotenceIdentifiersOK with default headers values
func NewPostIdempotenceIdentifiersOK() *PostIdempotenceIdentifiersOK {
	return &PostIdempotenceIdentifiersOK{}
}

/*PostIdempotenceIdentifiersOK handles this case with default header values.

Successful operation
*/
type PostIdempotenceIdentifiersOK struct {
	Payload *models.IdempotenceIdentifiersResponse
}

func (o *PostIdempotenceIdentifiersOK) Error() string {
	return fmt.Sprintf("[POST /idempotence_identifiers][%d] postIdempotenceIdentifiersOK  %+v", 200, o.Payload)
}

func (o *PostIdempotenceIdentifiersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdempotenceIdentifiersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIdempotenceIdentifiersDefault creates a PostIdempotenceIdentifiersDefault with default headers values
func NewPostIdempotenceIdentifiersDefault(code int) *PostIdempotenceIdentifiersDefault {
	return &PostIdempotenceIdentifiersDefault{
		_statusCode: code,
	}
}

/*PostIdempotenceIdentifiersDefault handles this case with default header values.

Error
*/
type PostIdempotenceIdentifiersDefault struct {
	_statusCode int

	Payload *models.IdempotenceIdentifiersStatus
}

// Code gets the status code for the post idempotence identifiers default response
func (o *PostIdempotenceIdentifiersDefault) Code() int {
	return o._statusCode
}

func (o *PostIdempotenceIdentifiersDefault) Error() string {
	return fmt.Sprintf("[POST /idempotence_identifiers][%d] PostIdempotenceIdentifiers default  %+v", o._statusCode, o.Payload)
}

func (o *PostIdempotenceIdentifiersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdempotenceIdentifiersStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}