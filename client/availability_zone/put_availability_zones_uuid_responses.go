package availability_zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PutAvailabilityZonesUUIDReader is a Reader for the PutAvailabilityZonesUUID structure.
type PutAvailabilityZonesUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAvailabilityZonesUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAvailabilityZonesUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutAvailabilityZonesUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutAvailabilityZonesUUIDOK creates a PutAvailabilityZonesUUIDOK with default headers values
func NewPutAvailabilityZonesUUIDOK() *PutAvailabilityZonesUUIDOK {
	return &PutAvailabilityZonesUUIDOK{}
}

/*PutAvailabilityZonesUUIDOK handles this case with default header values.

Success
*/
type PutAvailabilityZonesUUIDOK struct {
}

func (o *PutAvailabilityZonesUUIDOK) Error() string {
	return fmt.Sprintf("[PUT /availability_zones/{uuid}][%d] putAvailabilityZonesUuidOK ", 200)
}

func (o *PutAvailabilityZonesUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAvailabilityZonesUUIDDefault creates a PutAvailabilityZonesUUIDDefault with default headers values
func NewPutAvailabilityZonesUUIDDefault(code int) *PutAvailabilityZonesUUIDDefault {
	return &PutAvailabilityZonesUUIDDefault{
		_statusCode: code,
	}
}

/*PutAvailabilityZonesUUIDDefault handles this case with default header values.

Error
*/
type PutAvailabilityZonesUUIDDefault struct {
	_statusCode int

	Payload *models.AvailabilityZoneStatus
}

// Code gets the status code for the put availability zones UUID default response
func (o *PutAvailabilityZonesUUIDDefault) Code() int {
	return o._statusCode
}

func (o *PutAvailabilityZonesUUIDDefault) Error() string {
	return fmt.Sprintf("[PUT /availability_zones/{uuid}][%d] PutAvailabilityZonesUUID default  %+v", o._statusCode, o.Payload)
}

func (o *PutAvailabilityZonesUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailabilityZoneStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}