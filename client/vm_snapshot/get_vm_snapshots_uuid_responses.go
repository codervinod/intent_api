package vm_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// GetVMSnapshotsUUIDReader is a Reader for the GetVMSnapshotsUUID structure.
type GetVMSnapshotsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMSnapshotsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVMSnapshotsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetVMSnapshotsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetVMSnapshotsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetVMSnapshotsUUIDOK creates a GetVMSnapshotsUUIDOK with default headers values
func NewGetVMSnapshotsUUIDOK() *GetVMSnapshotsUUIDOK {
	return &GetVMSnapshotsUUIDOK{}
}

/*GetVMSnapshotsUUIDOK handles this case with default header values.

Request succeeded
*/
type GetVMSnapshotsUUIDOK struct {
	Payload *models.VMSnapshotIntentResponse
}

func (o *GetVMSnapshotsUUIDOK) Error() string {
	return fmt.Sprintf("[GET /vm_snapshots/{uuid}][%d] getVmSnapshotsUuidOK  %+v", 200, o.Payload)
}

func (o *GetVMSnapshotsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMSnapshotIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMSnapshotsUUIDNotFound creates a GetVMSnapshotsUUIDNotFound with default headers values
func NewGetVMSnapshotsUUIDNotFound() *GetVMSnapshotsUUIDNotFound {
	return &GetVMSnapshotsUUIDNotFound{}
}

/*GetVMSnapshotsUUIDNotFound handles this case with default header values.

Invalid UUID Provided
*/
type GetVMSnapshotsUUIDNotFound struct {
	Payload *models.VMSnapshotStatus
}

func (o *GetVMSnapshotsUUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /vm_snapshots/{uuid}][%d] getVmSnapshotsUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetVMSnapshotsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMSnapshotsUUIDDefault creates a GetVMSnapshotsUUIDDefault with default headers values
func NewGetVMSnapshotsUUIDDefault(code int) *GetVMSnapshotsUUIDDefault {
	return &GetVMSnapshotsUUIDDefault{
		_statusCode: code,
	}
}

/*GetVMSnapshotsUUIDDefault handles this case with default header values.

Internal Error
*/
type GetVMSnapshotsUUIDDefault struct {
	_statusCode int

	Payload *models.VMSnapshotStatus
}

// Code gets the status code for the get VM snapshots UUID default response
func (o *GetVMSnapshotsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetVMSnapshotsUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /vm_snapshots/{uuid}][%d] GetVMSnapshotsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetVMSnapshotsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMSnapshotStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
