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

// PostAvailabilityZonesListReader is a Reader for the PostAvailabilityZonesList structure.
type PostAvailabilityZonesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAvailabilityZonesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAvailabilityZonesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostAvailabilityZonesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostAvailabilityZonesListOK creates a PostAvailabilityZonesListOK with default headers values
func NewPostAvailabilityZonesListOK() *PostAvailabilityZonesListOK {
	return &PostAvailabilityZonesListOK{}
}

/*PostAvailabilityZonesListOK handles this case with default header values.

Successful operation
*/
type PostAvailabilityZonesListOK struct {
	Payload *models.AvailabilityZoneList
}

func (o *PostAvailabilityZonesListOK) Error() string {
	return fmt.Sprintf("[POST /availability_zones/list][%d] postAvailabilityZonesListOK  %+v", 200, o.Payload)
}

func (o *PostAvailabilityZonesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailabilityZoneList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAvailabilityZonesListDefault creates a PostAvailabilityZonesListDefault with default headers values
func NewPostAvailabilityZonesListDefault(code int) *PostAvailabilityZonesListDefault {
	return &PostAvailabilityZonesListDefault{
		_statusCode: code,
	}
}

/*PostAvailabilityZonesListDefault handles this case with default header values.

Error
*/
type PostAvailabilityZonesListDefault struct {
	_statusCode int

	Payload *models.AvailabilityZoneStatus
}

// Code gets the status code for the post availability zones list default response
func (o *PostAvailabilityZonesListDefault) Code() int {
	return o._statusCode
}

func (o *PostAvailabilityZonesListDefault) Error() string {
	return fmt.Sprintf("[POST /availability_zones/list][%d] PostAvailabilityZonesList default  %+v", o._statusCode, o.Payload)
}

func (o *PostAvailabilityZonesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailabilityZoneStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
