package vm_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PostVMSnapshotsReader is a Reader for the PostVMSnapshots structure.
type PostVMSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVMSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostVMSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostVMSnapshotsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostVMSnapshotsOK creates a PostVMSnapshotsOK with default headers values
func NewPostVMSnapshotsOK() *PostVMSnapshotsOK {
	return &PostVMSnapshotsOK{}
}

/*PostVMSnapshotsOK handles this case with default header values.

Request Succeeded
*/
type PostVMSnapshotsOK struct {
	Payload *models.VMSnapshotIntentResponse
}

func (o *PostVMSnapshotsOK) Error() string {
	return fmt.Sprintf("[POST /vm_snapshots][%d] postVmSnapshotsOK  %+v", 200, o.Payload)
}

func (o *PostVMSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMSnapshotIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVMSnapshotsDefault creates a PostVMSnapshotsDefault with default headers values
func NewPostVMSnapshotsDefault(code int) *PostVMSnapshotsDefault {
	return &PostVMSnapshotsDefault{
		_statusCode: code,
	}
}

/*PostVMSnapshotsDefault handles this case with default header values.

Internal Error
*/
type PostVMSnapshotsDefault struct {
	_statusCode int

	Payload *models.VMSnapshotStatus
}

// Code gets the status code for the post VM snapshots default response
func (o *PostVMSnapshotsDefault) Code() int {
	return o._statusCode
}

func (o *PostVMSnapshotsDefault) Error() string {
	return fmt.Sprintf("[POST /vm_snapshots][%d] PostVMSnapshots default  %+v", o._statusCode, o.Payload)
}

func (o *PostVMSnapshotsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
