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

// DeleteVolumeGroupSnapshotsUUIDReader is a Reader for the DeleteVolumeGroupSnapshotsUUID structure.
type DeleteVolumeGroupSnapshotsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVolumeGroupSnapshotsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteVolumeGroupSnapshotsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteVolumeGroupSnapshotsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteVolumeGroupSnapshotsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteVolumeGroupSnapshotsUUIDOK creates a DeleteVolumeGroupSnapshotsUUIDOK with default headers values
func NewDeleteVolumeGroupSnapshotsUUIDOK() *DeleteVolumeGroupSnapshotsUUIDOK {
	return &DeleteVolumeGroupSnapshotsUUIDOK{}
}

/*DeleteVolumeGroupSnapshotsUUIDOK handles this case with default header values.

Request Succeeded
*/
type DeleteVolumeGroupSnapshotsUUIDOK struct {
}

func (o *DeleteVolumeGroupSnapshotsUUIDOK) Error() string {
	return fmt.Sprintf("[DELETE /volume_group_snapshots/{uuid}][%d] deleteVolumeGroupSnapshotsUuidOK ", 200)
}

func (o *DeleteVolumeGroupSnapshotsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVolumeGroupSnapshotsUUIDNotFound creates a DeleteVolumeGroupSnapshotsUUIDNotFound with default headers values
func NewDeleteVolumeGroupSnapshotsUUIDNotFound() *DeleteVolumeGroupSnapshotsUUIDNotFound {
	return &DeleteVolumeGroupSnapshotsUUIDNotFound{}
}

/*DeleteVolumeGroupSnapshotsUUIDNotFound handles this case with default header values.

Invalid UUID Provided
*/
type DeleteVolumeGroupSnapshotsUUIDNotFound struct {
	Payload *models.VolumeGroupSnapshotStatus
}

func (o *DeleteVolumeGroupSnapshotsUUIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /volume_group_snapshots/{uuid}][%d] deleteVolumeGroupSnapshotsUuidNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVolumeGroupSnapshotsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeGroupSnapshotsUUIDDefault creates a DeleteVolumeGroupSnapshotsUUIDDefault with default headers values
func NewDeleteVolumeGroupSnapshotsUUIDDefault(code int) *DeleteVolumeGroupSnapshotsUUIDDefault {
	return &DeleteVolumeGroupSnapshotsUUIDDefault{
		_statusCode: code,
	}
}

/*DeleteVolumeGroupSnapshotsUUIDDefault handles this case with default header values.

Internal Error
*/
type DeleteVolumeGroupSnapshotsUUIDDefault struct {
	_statusCode int

	Payload *models.VolumeGroupSnapshotStatus
}

// Code gets the status code for the delete volume group snapshots UUID default response
func (o *DeleteVolumeGroupSnapshotsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVolumeGroupSnapshotsUUIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /volume_group_snapshots/{uuid}][%d] DeleteVolumeGroupSnapshotsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVolumeGroupSnapshotsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
