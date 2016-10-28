package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// GetContainersUUIDReader is a Reader for the GetContainersUUID structure.
type GetContainersUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContainersUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetContainersUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetContainersUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetContainersUUIDOK creates a GetContainersUUIDOK with default headers values
func NewGetContainersUUIDOK() *GetContainersUUIDOK {
	return &GetContainersUUIDOK{}
}

/*GetContainersUUIDOK handles this case with default header values.

Success
*/
type GetContainersUUIDOK struct {
	Payload *models.ContainerIntentResponse
}

func (o *GetContainersUUIDOK) Error() string {
	return fmt.Sprintf("[GET /containers/{uuid}][%d] getContainersUuidOK  %+v", 200, o.Payload)
}

func (o *GetContainersUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContainersUUIDDefault creates a GetContainersUUIDDefault with default headers values
func NewGetContainersUUIDDefault(code int) *GetContainersUUIDDefault {
	return &GetContainersUUIDDefault{
		_statusCode: code,
	}
}

/*GetContainersUUIDDefault handles this case with default header values.

Server Error
*/
type GetContainersUUIDDefault struct {
	_statusCode int

	Payload *models.ContainerStatus
}

// Code gets the status code for the get containers UUID default response
func (o *GetContainersUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetContainersUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /containers/{uuid}][%d] GetContainersUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetContainersUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
