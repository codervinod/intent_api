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

// GetImagesUUIDReader is a Reader for the GetImagesUUID structure.
type GetImagesUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImagesUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetImagesUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetImagesUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetImagesUUIDOK creates a GetImagesUUIDOK with default headers values
func NewGetImagesUUIDOK() *GetImagesUUIDOK {
	return &GetImagesUUIDOK{}
}

/*GetImagesUUIDOK handles this case with default header values.

Success
*/
type GetImagesUUIDOK struct {
	Payload *models.Image
}

func (o *GetImagesUUIDOK) Error() string {
	return fmt.Sprintf("[GET /images/{uuid}][%d] getImagesUuidOK  %+v", 200, o.Payload)
}

func (o *GetImagesUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagesUUIDDefault creates a GetImagesUUIDDefault with default headers values
func NewGetImagesUUIDDefault(code int) *GetImagesUUIDDefault {
	return &GetImagesUUIDDefault{
		_statusCode: code,
	}
}

/*GetImagesUUIDDefault handles this case with default header values.

Error
*/
type GetImagesUUIDDefault struct {
	_statusCode int

	Payload *models.ImageStatus
}

// Code gets the status code for the get images UUID default response
func (o *GetImagesUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetImagesUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /images/{uuid}][%d] GetImagesUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetImagesUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
