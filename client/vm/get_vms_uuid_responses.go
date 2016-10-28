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

// GetVmsUUIDReader is a Reader for the GetVmsUUID structure.
type GetVmsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVmsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVmsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetVmsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetVmsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetVmsUUIDOK creates a GetVmsUUIDOK with default headers values
func NewGetVmsUUIDOK() *GetVmsUUIDOK {
	return &GetVmsUUIDOK{}
}

/*GetVmsUUIDOK handles this case with default header values.

Request succeeded
*/
type GetVmsUUIDOK struct {
	Payload *models.VMIntentResponse
}

func (o *GetVmsUUIDOK) Error() string {
	return fmt.Sprintf("[GET /vms/{uuid}][%d] getVmsUuidOK  %+v", 200, o.Payload)
}

func (o *GetVmsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmsUUIDNotFound creates a GetVmsUUIDNotFound with default headers values
func NewGetVmsUUIDNotFound() *GetVmsUUIDNotFound {
	return &GetVmsUUIDNotFound{}
}

/*GetVmsUUIDNotFound handles this case with default header values.

Invalid UUID Provided
*/
type GetVmsUUIDNotFound struct {
	Payload *models.VMStatus
}

func (o *GetVmsUUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /vms/{uuid}][%d] getVmsUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetVmsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmsUUIDDefault creates a GetVmsUUIDDefault with default headers values
func NewGetVmsUUIDDefault(code int) *GetVmsUUIDDefault {
	return &GetVmsUUIDDefault{
		_statusCode: code,
	}
}

/*GetVmsUUIDDefault handles this case with default header values.

Internal Error
*/
type GetVmsUUIDDefault struct {
	_statusCode int

	Payload *models.VMStatus
}

// Code gets the status code for the get vms UUID default response
func (o *GetVmsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetVmsUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /vms/{uuid}][%d] GetVmsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetVmsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
