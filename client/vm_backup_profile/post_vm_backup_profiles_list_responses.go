package vm_backup_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PostVMBackupProfilesListReader is a Reader for the PostVMBackupProfilesList structure.
type PostVMBackupProfilesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVMBackupProfilesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostVMBackupProfilesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostVMBackupProfilesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostVMBackupProfilesListOK creates a PostVMBackupProfilesListOK with default headers values
func NewPostVMBackupProfilesListOK() *PostVMBackupProfilesListOK {
	return &PostVMBackupProfilesListOK{}
}

/*PostVMBackupProfilesListOK handles this case with default header values.

Successful operation
*/
type PostVMBackupProfilesListOK struct {
	Payload *models.VMBackupProfileList
}

func (o *PostVMBackupProfilesListOK) Error() string {
	return fmt.Sprintf("[POST /vm_backup_profiles/list][%d] postVmBackupProfilesListOK  %+v", 200, o.Payload)
}

func (o *PostVMBackupProfilesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMBackupProfileList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVMBackupProfilesListDefault creates a PostVMBackupProfilesListDefault with default headers values
func NewPostVMBackupProfilesListDefault(code int) *PostVMBackupProfilesListDefault {
	return &PostVMBackupProfilesListDefault{
		_statusCode: code,
	}
}

/*PostVMBackupProfilesListDefault handles this case with default header values.

Error
*/
type PostVMBackupProfilesListDefault struct {
	_statusCode int

	Payload *models.VMBackupProfileStatus
}

// Code gets the status code for the post VM backup profiles list default response
func (o *PostVMBackupProfilesListDefault) Code() int {
	return o._statusCode
}

func (o *PostVMBackupProfilesListDefault) Error() string {
	return fmt.Sprintf("[POST /vm_backup_profiles/list][%d] PostVMBackupProfilesList default  %+v", o._statusCode, o.Payload)
}

func (o *PostVMBackupProfilesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMBackupProfileStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}