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

// PutVMSnapshotsUUIDReader is a Reader for the PutVMSnapshotsUUID structure.
type PutVMSnapshotsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutVMSnapshotsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutVMSnapshotsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutVMSnapshotsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutVMSnapshotsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutVMSnapshotsUUIDOK creates a PutVMSnapshotsUUIDOK with default headers values
func NewPutVMSnapshotsUUIDOK() *PutVMSnapshotsUUIDOK {
	return &PutVMSnapshotsUUIDOK{}
}

/*PutVMSnapshotsUUIDOK handles this case with default header values.

Request succeeded
*/
type PutVMSnapshotsUUIDOK struct {
	Payload *models.VMSnapshotIntentResponse
}

func (o *PutVMSnapshotsUUIDOK) Error() string {
	return fmt.Sprintf("[PUT /vm_snapshots/{uuid}][%d] putVmSnapshotsUuidOK  %+v", 200, o.Payload)
}

func (o *PutVMSnapshotsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMSnapshotIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVMSnapshotsUUIDNotFound creates a PutVMSnapshotsUUIDNotFound with default headers values
func NewPutVMSnapshotsUUIDNotFound() *PutVMSnapshotsUUIDNotFound {
	return &PutVMSnapshotsUUIDNotFound{}
}

/*PutVMSnapshotsUUIDNotFound handles this case with default header values.

Invalid UUID Provided
*/
type PutVMSnapshotsUUIDNotFound struct {
	Payload *models.VMSnapshotStatus
}

func (o *PutVMSnapshotsUUIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /vm_snapshots/{uuid}][%d] putVmSnapshotsUuidNotFound  %+v", 404, o.Payload)
}

func (o *PutVMSnapshotsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVMSnapshotsUUIDDefault creates a PutVMSnapshotsUUIDDefault with default headers values
func NewPutVMSnapshotsUUIDDefault(code int) *PutVMSnapshotsUUIDDefault {
	return &PutVMSnapshotsUUIDDefault{
		_statusCode: code,
	}
}

/*PutVMSnapshotsUUIDDefault handles this case with default header values.

Internal Error
*/
type PutVMSnapshotsUUIDDefault struct {
	_statusCode int

	Payload *models.VMSnapshotStatus
}

// Code gets the status code for the put VM snapshots UUID default response
func (o *PutVMSnapshotsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *PutVMSnapshotsUUIDDefault) Error() string {
	return fmt.Sprintf("[PUT /vm_snapshots/{uuid}][%d] PutVMSnapshotsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *PutVMSnapshotsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}