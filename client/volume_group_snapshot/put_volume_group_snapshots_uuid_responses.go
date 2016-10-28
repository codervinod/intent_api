package volume_group_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// PutVolumeGroupSnapshotsUUIDReader is a Reader for the PutVolumeGroupSnapshotsUUID structure.
type PutVolumeGroupSnapshotsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutVolumeGroupSnapshotsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutVolumeGroupSnapshotsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutVolumeGroupSnapshotsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutVolumeGroupSnapshotsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutVolumeGroupSnapshotsUUIDOK creates a PutVolumeGroupSnapshotsUUIDOK with default headers values
func NewPutVolumeGroupSnapshotsUUIDOK() *PutVolumeGroupSnapshotsUUIDOK {
	return &PutVolumeGroupSnapshotsUUIDOK{}
}

/*PutVolumeGroupSnapshotsUUIDOK handles this case with default header values.

Request succeeded
*/
type PutVolumeGroupSnapshotsUUIDOK struct {
	Payload *models.VolumeGroupSnapshotIntentResponse
}

func (o *PutVolumeGroupSnapshotsUUIDOK) Error() string {
	return fmt.Sprintf("[PUT /volume_group_snapshots/{uuid}][%d] putVolumeGroupSnapshotsUuidOK  %+v", 200, o.Payload)
}

func (o *PutVolumeGroupSnapshotsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupSnapshotIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVolumeGroupSnapshotsUUIDNotFound creates a PutVolumeGroupSnapshotsUUIDNotFound with default headers values
func NewPutVolumeGroupSnapshotsUUIDNotFound() *PutVolumeGroupSnapshotsUUIDNotFound {
	return &PutVolumeGroupSnapshotsUUIDNotFound{}
}

/*PutVolumeGroupSnapshotsUUIDNotFound handles this case with default header values.

Invalid UUID Provided
*/
type PutVolumeGroupSnapshotsUUIDNotFound struct {
	Payload *models.VolumeGroupSnapshotStatus
}

func (o *PutVolumeGroupSnapshotsUUIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /volume_group_snapshots/{uuid}][%d] putVolumeGroupSnapshotsUuidNotFound  %+v", 404, o.Payload)
}

func (o *PutVolumeGroupSnapshotsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVolumeGroupSnapshotsUUIDDefault creates a PutVolumeGroupSnapshotsUUIDDefault with default headers values
func NewPutVolumeGroupSnapshotsUUIDDefault(code int) *PutVolumeGroupSnapshotsUUIDDefault {
	return &PutVolumeGroupSnapshotsUUIDDefault{
		_statusCode: code,
	}
}

/*PutVolumeGroupSnapshotsUUIDDefault handles this case with default header values.

Internal Error
*/
type PutVolumeGroupSnapshotsUUIDDefault struct {
	_statusCode int

	Payload *models.VolumeGroupSnapshotStatus
}

// Code gets the status code for the put volume group snapshots UUID default response
func (o *PutVolumeGroupSnapshotsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *PutVolumeGroupSnapshotsUUIDDefault) Error() string {
	return fmt.Sprintf("[PUT /volume_group_snapshots/{uuid}][%d] PutVolumeGroupSnapshotsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *PutVolumeGroupSnapshotsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
