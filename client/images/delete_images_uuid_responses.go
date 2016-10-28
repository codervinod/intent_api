package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// DeleteImagesUUIDReader is a Reader for the DeleteImagesUUID structure.
type DeleteImagesUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImagesUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteImagesUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteImagesUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteImagesUUIDOK creates a DeleteImagesUUIDOK with default headers values
func NewDeleteImagesUUIDOK() *DeleteImagesUUIDOK {
	return &DeleteImagesUUIDOK{}
}

/*DeleteImagesUUIDOK handles this case with default header values.

Success
*/
type DeleteImagesUUIDOK struct {
}

func (o *DeleteImagesUUIDOK) Error() string {
	return fmt.Sprintf("[DELETE /images/{uuid}][%d] deleteImagesUuidOK ", 200)
}

func (o *DeleteImagesUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteImagesUUIDDefault creates a DeleteImagesUUIDDefault with default headers values
func NewDeleteImagesUUIDDefault(code int) *DeleteImagesUUIDDefault {
	return &DeleteImagesUUIDDefault{
		_statusCode: code,
	}
}

/*DeleteImagesUUIDDefault handles this case with default header values.

Error
*/
type DeleteImagesUUIDDefault struct {
	_statusCode int

	Payload *models.ImageStatus
}

// Code gets the status code for the delete images UUID default response
func (o *DeleteImagesUUIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteImagesUUIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /images/{uuid}][%d] DeleteImagesUUID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteImagesUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}