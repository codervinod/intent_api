package availability_zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// DeleteAvailabilityZonesUUIDReader is a Reader for the DeleteAvailabilityZonesUUID structure.
type DeleteAvailabilityZonesUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAvailabilityZonesUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAvailabilityZonesUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteAvailabilityZonesUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteAvailabilityZonesUUIDOK creates a DeleteAvailabilityZonesUUIDOK with default headers values
func NewDeleteAvailabilityZonesUUIDOK() *DeleteAvailabilityZonesUUIDOK {
	return &DeleteAvailabilityZonesUUIDOK{}
}

/*DeleteAvailabilityZonesUUIDOK handles this case with default header values.

Success
*/
type DeleteAvailabilityZonesUUIDOK struct {
}

func (o *DeleteAvailabilityZonesUUIDOK) Error() string {
	return fmt.Sprintf("[DELETE /availability_zones/{uuid}][%d] deleteAvailabilityZonesUuidOK ", 200)
}

func (o *DeleteAvailabilityZonesUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAvailabilityZonesUUIDDefault creates a DeleteAvailabilityZonesUUIDDefault with default headers values
func NewDeleteAvailabilityZonesUUIDDefault(code int) *DeleteAvailabilityZonesUUIDDefault {
	return &DeleteAvailabilityZonesUUIDDefault{
		_statusCode: code,
	}
}

/*DeleteAvailabilityZonesUUIDDefault handles this case with default header values.

Error
*/
type DeleteAvailabilityZonesUUIDDefault struct {
	_statusCode int

	Payload *models.AvailabilityZoneStatus
}

// Code gets the status code for the delete availability zones UUID default response
func (o *DeleteAvailabilityZonesUUIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAvailabilityZonesUUIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /availability_zones/{uuid}][%d] DeleteAvailabilityZonesUUID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAvailabilityZonesUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailabilityZoneStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
