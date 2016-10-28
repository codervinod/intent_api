package volume_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PutVolumeGroupsUUIDReader is a Reader for the PutVolumeGroupsUUID structure.
type PutVolumeGroupsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutVolumeGroupsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutVolumeGroupsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutVolumeGroupsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutVolumeGroupsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutVolumeGroupsUUIDOK creates a PutVolumeGroupsUUIDOK with default headers values
func NewPutVolumeGroupsUUIDOK() *PutVolumeGroupsUUIDOK {
	return &PutVolumeGroupsUUIDOK{}
}

/*PutVolumeGroupsUUIDOK handles this case with default header values.

Successfully updated specified volume group
*/
type PutVolumeGroupsUUIDOK struct {
	Payload *models.VolumeGroup
}

func (o *PutVolumeGroupsUUIDOK) Error() string {
	return fmt.Sprintf("[PUT /volume_groups/{uuid}][%d] putVolumeGroupsUuidOK  %+v", 200, o.Payload)
}

func (o *PutVolumeGroupsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVolumeGroupsUUIDNotFound creates a PutVolumeGroupsUUIDNotFound with default headers values
func NewPutVolumeGroupsUUIDNotFound() *PutVolumeGroupsUUIDNotFound {
	return &PutVolumeGroupsUUIDNotFound{}
}

/*PutVolumeGroupsUUIDNotFound handles this case with default header values.

Specified volume group does not exist
*/
type PutVolumeGroupsUUIDNotFound struct {
	Payload *models.VolumeGroupStatus
}

func (o *PutVolumeGroupsUUIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /volume_groups/{uuid}][%d] putVolumeGroupsUuidNotFound  %+v", 404, o.Payload)
}

func (o *PutVolumeGroupsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVolumeGroupsUUIDDefault creates a PutVolumeGroupsUUIDDefault with default headers values
func NewPutVolumeGroupsUUIDDefault(code int) *PutVolumeGroupsUUIDDefault {
	return &PutVolumeGroupsUUIDDefault{
		_statusCode: code,
	}
}

/*PutVolumeGroupsUUIDDefault handles this case with default header values.

Failed to updated specified volume group
*/
type PutVolumeGroupsUUIDDefault struct {
	_statusCode int

	Payload *models.VolumeGroupStatus
}

// Code gets the status code for the put volume groups UUID default response
func (o *PutVolumeGroupsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *PutVolumeGroupsUUIDDefault) Error() string {
	return fmt.Sprintf("[PUT /volume_groups/{uuid}][%d] PutVolumeGroupsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *PutVolumeGroupsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}