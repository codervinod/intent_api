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

// PostVolumeGroupSnapshotsListReader is a Reader for the PostVolumeGroupSnapshotsList structure.
type PostVolumeGroupSnapshotsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVolumeGroupSnapshotsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostVolumeGroupSnapshotsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostVolumeGroupSnapshotsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostVolumeGroupSnapshotsListOK creates a PostVolumeGroupSnapshotsListOK with default headers values
func NewPostVolumeGroupSnapshotsListOK() *PostVolumeGroupSnapshotsListOK {
	return &PostVolumeGroupSnapshotsListOK{}
}

/*PostVolumeGroupSnapshotsListOK handles this case with default header values.

Successful operation
*/
type PostVolumeGroupSnapshotsListOK struct {
	Payload *models.VolumeGroupSnapshotListIntentResponse
}

func (o *PostVolumeGroupSnapshotsListOK) Error() string {
	return fmt.Sprintf("[POST /volume_group_snapshots/list][%d] postVolumeGroupSnapshotsListOK  %+v", 200, o.Payload)
}

func (o *PostVolumeGroupSnapshotsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupSnapshotListIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVolumeGroupSnapshotsListDefault creates a PostVolumeGroupSnapshotsListDefault with default headers values
func NewPostVolumeGroupSnapshotsListDefault(code int) *PostVolumeGroupSnapshotsListDefault {
	return &PostVolumeGroupSnapshotsListDefault{
		_statusCode: code,
	}
}

/*PostVolumeGroupSnapshotsListDefault handles this case with default header values.

Error
*/
type PostVolumeGroupSnapshotsListDefault struct {
	_statusCode int

	Payload *models.VolumeGroupSnapshotStatus
}

// Code gets the status code for the post volume group snapshots list default response
func (o *PostVolumeGroupSnapshotsListDefault) Code() int {
	return o._statusCode
}

func (o *PostVolumeGroupSnapshotsListDefault) Error() string {
	return fmt.Sprintf("[POST /volume_group_snapshots/list][%d] PostVolumeGroupSnapshotsList default  %+v", o._statusCode, o.Payload)
}

func (o *PostVolumeGroupSnapshotsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
