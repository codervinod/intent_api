package vm_backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// GetVMBackupsUUIDReader is a Reader for the GetVMBackupsUUID structure.
type GetVMBackupsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMBackupsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVMBackupsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetVMBackupsUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetVMBackupsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetVMBackupsUUIDOK creates a GetVMBackupsUUIDOK with default headers values
func NewGetVMBackupsUUIDOK() *GetVMBackupsUUIDOK {
	return &GetVMBackupsUUIDOK{}
}

/*GetVMBackupsUUIDOK handles this case with default header values.

Request succeeded
*/
type GetVMBackupsUUIDOK struct {
	Payload *models.VMBackupIntentResponse
}

func (o *GetVMBackupsUUIDOK) Error() string {
	return fmt.Sprintf("[GET /vm_backups/{uuid}][%d] getVmBackupsUuidOK  %+v", 200, o.Payload)
}

func (o *GetVMBackupsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMBackupIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMBackupsUUIDNotFound creates a GetVMBackupsUUIDNotFound with default headers values
func NewGetVMBackupsUUIDNotFound() *GetVMBackupsUUIDNotFound {
	return &GetVMBackupsUUIDNotFound{}
}

/*GetVMBackupsUUIDNotFound handles this case with default header values.

Invalid UUID Provided
*/
type GetVMBackupsUUIDNotFound struct {
	Payload *models.VMBackupStatus
}

func (o *GetVMBackupsUUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /vm_backups/{uuid}][%d] getVmBackupsUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetVMBackupsUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMBackupStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMBackupsUUIDDefault creates a GetVMBackupsUUIDDefault with default headers values
func NewGetVMBackupsUUIDDefault(code int) *GetVMBackupsUUIDDefault {
	return &GetVMBackupsUUIDDefault{
		_statusCode: code,
	}
}

/*GetVMBackupsUUIDDefault handles this case with default header values.

Internal Error
*/
type GetVMBackupsUUIDDefault struct {
	_statusCode int

	Payload *models.VMBackupStatus
}

// Code gets the status code for the get VM backups UUID default response
func (o *GetVMBackupsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetVMBackupsUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /vm_backups/{uuid}][%d] GetVMBackupsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetVMBackupsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMBackupStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
