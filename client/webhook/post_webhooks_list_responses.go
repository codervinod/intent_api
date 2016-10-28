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

// PostWebhooksListReader is a Reader for the PostWebhooksList structure.
type PostWebhooksListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWebhooksListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostWebhooksListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostWebhooksListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostWebhooksListOK creates a PostWebhooksListOK with default headers values
func NewPostWebhooksListOK() *PostWebhooksListOK {
	return &PostWebhooksListOK{}
}

/*PostWebhooksListOK handles this case with default header values.

Successful operation
*/
type PostWebhooksListOK struct {
	Payload *models.WebhookListIntentResponse
}

func (o *PostWebhooksListOK) Error() string {
	return fmt.Sprintf("[POST /webhooks/list][%d] postWebhooksListOK  %+v", 200, o.Payload)
}

func (o *PostWebhooksListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebhookListIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebhooksListDefault creates a PostWebhooksListDefault with default headers values
func NewPostWebhooksListDefault(code int) *PostWebhooksListDefault {
	return &PostWebhooksListDefault{
		_statusCode: code,
	}
}

/*PostWebhooksListDefault handles this case with default header values.

Error
*/
type PostWebhooksListDefault struct {
	_statusCode int

	Payload *models.WebhookStatus
}

// Code gets the status code for the post webhooks list default response
func (o *PostWebhooksListDefault) Code() int {
	return o._statusCode
}

func (o *PostWebhooksListDefault) Error() string {
	return fmt.Sprintf("[POST /webhooks/list][%d] PostWebhooksList default  %+v", o._statusCode, o.Payload)
}

func (o *PostWebhooksListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebhookStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
