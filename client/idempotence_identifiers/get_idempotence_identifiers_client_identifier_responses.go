package idempotence_identifiers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// GetIdempotenceIdentifiersClientIdentifierReader is a Reader for the GetIdempotenceIdentifiersClientIdentifier structure.
type GetIdempotenceIdentifiersClientIdentifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdempotenceIdentifiersClientIdentifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIdempotenceIdentifiersClientIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetIdempotenceIdentifiersClientIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetIdempotenceIdentifiersClientIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetIdempotenceIdentifiersClientIdentifierOK creates a GetIdempotenceIdentifiersClientIdentifierOK with default headers values
func NewGetIdempotenceIdentifiersClientIdentifierOK() *GetIdempotenceIdentifiersClientIdentifierOK {
	return &GetIdempotenceIdentifiersClientIdentifierOK{}
}

/*GetIdempotenceIdentifiersClientIdentifierOK handles this case with default header values.

Successful operation
*/
type GetIdempotenceIdentifiersClientIdentifierOK struct {
	Payload *models.IdempotenceIdentifiersResponse
}

func (o *GetIdempotenceIdentifiersClientIdentifierOK) Error() string {
	return fmt.Sprintf("[GET /idempotence_identifiers/{client_identifier}][%d] getIdempotenceIdentifiersClientIdentifierOK  %+v", 200, o.Payload)
}

func (o *GetIdempotenceIdentifiersClientIdentifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdempotenceIdentifiersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdempotenceIdentifiersClientIdentifierNotFound creates a GetIdempotenceIdentifiersClientIdentifierNotFound with default headers values
func NewGetIdempotenceIdentifiersClientIdentifierNotFound() *GetIdempotenceIdentifiersClientIdentifierNotFound {
	return &GetIdempotenceIdentifiersClientIdentifierNotFound{}
}

/*GetIdempotenceIdentifiersClientIdentifierNotFound handles this case with default header values.

Client identifier does not exists
*/
type GetIdempotenceIdentifiersClientIdentifierNotFound struct {
	Payload *models.IdempotenceIdentifiersStatus
}

func (o *GetIdempotenceIdentifiersClientIdentifierNotFound) Error() string {
	return fmt.Sprintf("[GET /idempotence_identifiers/{client_identifier}][%d] getIdempotenceIdentifiersClientIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *GetIdempotenceIdentifiersClientIdentifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdempotenceIdentifiersStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdempotenceIdentifiersClientIdentifierDefault creates a GetIdempotenceIdentifiersClientIdentifierDefault with default headers values
func NewGetIdempotenceIdentifiersClientIdentifierDefault(code int) *GetIdempotenceIdentifiersClientIdentifierDefault {
	return &GetIdempotenceIdentifiersClientIdentifierDefault{
		_statusCode: code,
	}
}

/*GetIdempotenceIdentifiersClientIdentifierDefault handles this case with default header values.

Error
*/
type GetIdempotenceIdentifiersClientIdentifierDefault struct {
	_statusCode int

	Payload *models.IdempotenceIdentifiersStatus
}

// Code gets the status code for the get idempotence identifiers client identifier default response
func (o *GetIdempotenceIdentifiersClientIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *GetIdempotenceIdentifiersClientIdentifierDefault) Error() string {
	return fmt.Sprintf("[GET /idempotence_identifiers/{client_identifier}][%d] GetIdempotenceIdentifiersClientIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *GetIdempotenceIdentifiersClientIdentifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdempotenceIdentifiersStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
