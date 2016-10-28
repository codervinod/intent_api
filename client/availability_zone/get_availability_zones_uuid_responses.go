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

// GetAvailabilityZonesUUIDReader is a Reader for the GetAvailabilityZonesUUID structure.
type GetAvailabilityZonesUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvailabilityZonesUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAvailabilityZonesUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAvailabilityZonesUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetAvailabilityZonesUUIDOK creates a GetAvailabilityZonesUUIDOK with default headers values
func NewGetAvailabilityZonesUUIDOK() *GetAvailabilityZonesUUIDOK {
	return &GetAvailabilityZonesUUIDOK{}
}

/*GetAvailabilityZonesUUIDOK handles this case with default header values.

Success
*/
type GetAvailabilityZonesUUIDOK struct {
	Payload *models.AvailabilityZone
}

func (o *GetAvailabilityZonesUUIDOK) Error() string {
	return fmt.Sprintf("[GET /availability_zones/{uuid}][%d] getAvailabilityZonesUuidOK  %+v", 200, o.Payload)
}

func (o *GetAvailabilityZonesUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailabilityZone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAvailabilityZonesUUIDDefault creates a GetAvailabilityZonesUUIDDefault with default headers values
func NewGetAvailabilityZonesUUIDDefault(code int) *GetAvailabilityZonesUUIDDefault {
	return &GetAvailabilityZonesUUIDDefault{
		_statusCode: code,
	}
}

/*GetAvailabilityZonesUUIDDefault handles this case with default header values.

Error
*/
type GetAvailabilityZonesUUIDDefault struct {
	_statusCode int

	Payload *models.AvailabilityZoneStatus
}

// Code gets the status code for the get availability zones UUID default response
func (o *GetAvailabilityZonesUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAvailabilityZonesUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /availability_zones/{uuid}][%d] GetAvailabilityZonesUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetAvailabilityZonesUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailabilityZoneStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
