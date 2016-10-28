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

// PutImagesUUIDReader is a Reader for the PutImagesUUID structure.
type PutImagesUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutImagesUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutImagesUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutImagesUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutImagesUUIDOK creates a PutImagesUUIDOK with default headers values
func NewPutImagesUUIDOK() *PutImagesUUIDOK {
	return &PutImagesUUIDOK{}
}

/*PutImagesUUIDOK handles this case with default header values.

Success
*/
type PutImagesUUIDOK struct {
}

func (o *PutImagesUUIDOK) Error() string {
	return fmt.Sprintf("[PUT /images/{uuid}][%d] putImagesUuidOK ", 200)
}

func (o *PutImagesUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutImagesUUIDDefault creates a PutImagesUUIDDefault with default headers values
func NewPutImagesUUIDDefault(code int) *PutImagesUUIDDefault {
	return &PutImagesUUIDDefault{
		_statusCode: code,
	}
}

/*PutImagesUUIDDefault handles this case with default header values.

Error
*/
type PutImagesUUIDDefault struct {
	_statusCode int

	Payload *models.ImageStatus
}

// Code gets the status code for the put images UUID default response
func (o *PutImagesUUIDDefault) Code() int {
	return o._statusCode
}

func (o *PutImagesUUIDDefault) Error() string {
	return fmt.Sprintf("[PUT /images/{uuid}][%d] PutImagesUUID default  %+v", o._statusCode, o.Payload)
}

func (o *PutImagesUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
