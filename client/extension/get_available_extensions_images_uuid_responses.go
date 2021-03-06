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

// GetAvailableExtensionsImagesUUIDReader is a Reader for the GetAvailableExtensionsImagesUUID structure.
type GetAvailableExtensionsImagesUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvailableExtensionsImagesUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetAvailableExtensionsImagesUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetAvailableExtensionsImagesUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetAvailableExtensionsImagesUUIDNotFound creates a GetAvailableExtensionsImagesUUIDNotFound with default headers values
func NewGetAvailableExtensionsImagesUUIDNotFound() *GetAvailableExtensionsImagesUUIDNotFound {
	return &GetAvailableExtensionsImagesUUIDNotFound{}
}

/*GetAvailableExtensionsImagesUUIDNotFound handles this case with default header values.

Extension does not exists
*/
type GetAvailableExtensionsImagesUUIDNotFound struct {
	Payload *models.AvailableExtensionStatus
}

func (o *GetAvailableExtensionsImagesUUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /available_extensions/images/{uuid}][%d] getAvailableExtensionsImagesUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetAvailableExtensionsImagesUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailableExtensionStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAvailableExtensionsImagesUUIDDefault creates a GetAvailableExtensionsImagesUUIDDefault with default headers values
func NewGetAvailableExtensionsImagesUUIDDefault(code int) *GetAvailableExtensionsImagesUUIDDefault {
	return &GetAvailableExtensionsImagesUUIDDefault{
		_statusCode: code,
	}
}

/*GetAvailableExtensionsImagesUUIDDefault handles this case with default header values.

Error
*/
type GetAvailableExtensionsImagesUUIDDefault struct {
	_statusCode int

	Payload *models.AvailableExtensionStatus
}

// Code gets the status code for the get available extensions images UUID default response
func (o *GetAvailableExtensionsImagesUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAvailableExtensionsImagesUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /available_extensions/images/{uuid}][%d] GetAvailableExtensionsImagesUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetAvailableExtensionsImagesUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailableExtensionStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
