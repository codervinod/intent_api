package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PutVmsUUIDReader is a Reader for the PutVmsUUID structure.
type PutVmsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutVmsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutVmsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutVmsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutVmsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutVmsUUIDOK creates a PutVmsUUIDOK with default headers values
func NewPutVmsUUIDOK() *PutVmsUUIDOK {
	return &PutVmsUUIDOK{}
}

/*PutVmsUUIDOK handles this case with default header values.

Request Succeeded
*/
type PutVmsUUIDOK struct {
	Payload *models.VMIntentResponse
}

func (o *PutVmsUUIDOK) Error() string {
	return fmt.Sprintf("[PUT /vms/{uuid}][%d] putVmsUuidOK  %+v", 200, o.Payload)
}

func (o *PutVmsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVmsUUIDNotFound creates a PutVmsUUIDNotFound with default headers values
func NewPutVmsUUIDNotFound() *PutVmsUUIDNotFound {
	return &PutVmsUUIDNotFound{}
}

/*PutVmsUUIDNotFound handles this case with default header values.

Invalid UUID Provided
*/
type PutVmsUUIDNotFound struct {
	Payload *models.VMStatus
}

func (o *PutVmsUUIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /vms/{uuid}][%d] putVmsUuidNotFound  %+v", 404, o.Payload)
}

func (o *PutVmsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVmsUUIDDefault creates a PutVmsUUIDDefault with default headers values
func NewPutVmsUUIDDefault(code int) *PutVmsUUIDDefault {
	return &PutVmsUUIDDefault{
		_statusCode: code,
	}
}

/*PutVmsUUIDDefault handles this case with default header values.

Internal Error
*/
type PutVmsUUIDDefault struct {
	_statusCode int

	Payload *models.VMStatus
}

// Code gets the status code for the put vms UUID default response
func (o *PutVmsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *PutVmsUUIDDefault) Error() string {
	return fmt.Sprintf("[PUT /vms/{uuid}][%d] PutVmsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *PutVmsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
