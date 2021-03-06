package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// PutWebhooksUUIDReader is a Reader for the PutWebhooksUUID structure.
type PutWebhooksUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWebhooksUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutWebhooksUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutWebhooksUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutWebhooksUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutWebhooksUUIDOK creates a PutWebhooksUUIDOK with default headers values
func NewPutWebhooksUUIDOK() *PutWebhooksUUIDOK {
	return &PutWebhooksUUIDOK{}
}

/*PutWebhooksUUIDOK handles this case with default header values.

Request Succeeded
*/
type PutWebhooksUUIDOK struct {
	Payload *models.WebhookIntentResponse
}

func (o *PutWebhooksUUIDOK) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{uuid}][%d] putWebhooksUuidOK  %+v", 200, o.Payload)
}

func (o *PutWebhooksUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebhookIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebhooksUUIDNotFound creates a PutWebhooksUUIDNotFound with default headers values
func NewPutWebhooksUUIDNotFound() *PutWebhooksUUIDNotFound {
	return &PutWebhooksUUIDNotFound{}
}

/*PutWebhooksUUIDNotFound handles this case with default header values.

Invalid UUID Provided
*/
type PutWebhooksUUIDNotFound struct {
	Payload *models.WebhookStatus
}

func (o *PutWebhooksUUIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{uuid}][%d] putWebhooksUuidNotFound  %+v", 404, o.Payload)
}

func (o *PutWebhooksUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebhookStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWebhooksUUIDDefault creates a PutWebhooksUUIDDefault with default headers values
func NewPutWebhooksUUIDDefault(code int) *PutWebhooksUUIDDefault {
	return &PutWebhooksUUIDDefault{
		_statusCode: code,
	}
}

/*PutWebhooksUUIDDefault handles this case with default header values.

Internal error
*/
type PutWebhooksUUIDDefault struct {
	_statusCode int

	Payload *models.WebhookStatus
}

// Code gets the status code for the put webhooks UUID default response
func (o *PutWebhooksUUIDDefault) Code() int {
	return o._statusCode
}

func (o *PutWebhooksUUIDDefault) Error() string {
	return fmt.Sprintf("[PUT /webhooks/{uuid}][%d] PutWebhooksUUID default  %+v", o._statusCode, o.Payload)
}

func (o *PutWebhooksUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebhookStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
