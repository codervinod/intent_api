package volume_group_backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PostVolumeGroupBackupsListReader is a Reader for the PostVolumeGroupBackupsList structure.
type PostVolumeGroupBackupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVolumeGroupBackupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostVolumeGroupBackupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostVolumeGroupBackupsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostVolumeGroupBackupsListOK creates a PostVolumeGroupBackupsListOK with default headers values
func NewPostVolumeGroupBackupsListOK() *PostVolumeGroupBackupsListOK {
	return &PostVolumeGroupBackupsListOK{}
}

/*PostVolumeGroupBackupsListOK handles this case with default header values.

Successful operation
*/
type PostVolumeGroupBackupsListOK struct {
	Payload *models.VolumeGroupBackupListIntentResponse
}

func (o *PostVolumeGroupBackupsListOK) Error() string {
	return fmt.Sprintf("[POST /volume_group_backups/list][%d] postVolumeGroupBackupsListOK  %+v", 200, o.Payload)
}

func (o *PostVolumeGroupBackupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupBackupListIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVolumeGroupBackupsListDefault creates a PostVolumeGroupBackupsListDefault with default headers values
func NewPostVolumeGroupBackupsListDefault(code int) *PostVolumeGroupBackupsListDefault {
	return &PostVolumeGroupBackupsListDefault{
		_statusCode: code,
	}
}

/*PostVolumeGroupBackupsListDefault handles this case with default header values.

Error
*/
type PostVolumeGroupBackupsListDefault struct {
	_statusCode int

	Payload *models.VolumeGroupBackupStatus
}

// Code gets the status code for the post volume group backups list default response
func (o *PostVolumeGroupBackupsListDefault) Code() int {
	return o._statusCode
}

func (o *PostVolumeGroupBackupsListDefault) Error() string {
	return fmt.Sprintf("[POST /volume_group_backups/list][%d] PostVolumeGroupBackupsList default  %+v", o._statusCode, o.Payload)
}

func (o *PostVolumeGroupBackupsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupBackupStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}