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

// DeleteVMSnapshotsUUIDReader is a Reader for the DeleteVMSnapshotsUUID structure.
type DeleteVMSnapshotsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMSnapshotsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteVMSnapshotsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteVMSnapshotsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteVMSnapshotsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteVMSnapshotsUUIDOK creates a DeleteVMSnapshotsUUIDOK with default headers values
func NewDeleteVMSnapshotsUUIDOK() *DeleteVMSnapshotsUUIDOK {
	return &DeleteVMSnapshotsUUIDOK{}
}

/*DeleteVMSnapshotsUUIDOK handles this case with default header values.

Request Succeeded
*/
type DeleteVMSnapshotsUUIDOK struct {
}

func (o *DeleteVMSnapshotsUUIDOK) Error() string {
	return fmt.Sprintf("[DELETE /vm_snapshots/{uuid}][%d] deleteVmSnapshotsUuidOK ", 200)
}

func (o *DeleteVMSnapshotsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVMSnapshotsUUIDNotFound creates a DeleteVMSnapshotsUUIDNotFound with default headers values
func NewDeleteVMSnapshotsUUIDNotFound() *DeleteVMSnapshotsUUIDNotFound {
	return &DeleteVMSnapshotsUUIDNotFound{}
}

/*DeleteVMSnapshotsUUIDNotFound handles this case with default header values.

Invalid UUID Provided
*/
type DeleteVMSnapshotsUUIDNotFound struct {
	Payload *models.VMSnapshotStatus
}

func (o *DeleteVMSnapshotsUUIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /vm_snapshots/{uuid}][%d] deleteVmSnapshotsUuidNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVMSnapshotsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMSnapshotsUUIDDefault creates a DeleteVMSnapshotsUUIDDefault with default headers values
func NewDeleteVMSnapshotsUUIDDefault(code int) *DeleteVMSnapshotsUUIDDefault {
	return &DeleteVMSnapshotsUUIDDefault{
		_statusCode: code,
	}
}

/*DeleteVMSnapshotsUUIDDefault handles this case with default header values.

Internal Error
*/
type DeleteVMSnapshotsUUIDDefault struct {
	_statusCode int

	Payload *models.VMSnapshotStatus
}

// Code gets the status code for the delete VM snapshots UUID default response
func (o *DeleteVMSnapshotsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVMSnapshotsUUIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /vm_snapshots/{uuid}][%d] DeleteVMSnapshotsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVMSnapshotsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
