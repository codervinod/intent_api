package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// GetPermissionsUUIDReader is a Reader for the GetPermissionsUUID structure.
type GetPermissionsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPermissionsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetPermissionsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetPermissionsUUIDOK creates a GetPermissionsUUIDOK with default headers values
func NewGetPermissionsUUIDOK() *GetPermissionsUUIDOK {
	return &GetPermissionsUUIDOK{}
}

/*GetPermissionsUUIDOK handles this case with default header values.

Successful operation
*/
type GetPermissionsUUIDOK struct {
	Payload *models.Permission
}

func (o *GetPermissionsUUIDOK) Error() string {
	return fmt.Sprintf("[GET /permissions/{uuid}][%d] getPermissionsUuidOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPermissionsUUIDDefault creates a GetPermissionsUUIDDefault with default headers values
func NewGetPermissionsUUIDDefault(code int) *GetPermissionsUUIDDefault {
	return &GetPermissionsUUIDDefault{
		_statusCode: code,
	}
}

/*GetPermissionsUUIDDefault handles this case with default header values.

Error
*/
type GetPermissionsUUIDDefault struct {
	_statusCode int

	Payload *models.PermissionStatus
}

// Code gets the status code for the get permissions UUID default response
func (o *GetPermissionsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetPermissionsUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /permissions/{uuid}][%d] GetPermissionsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetPermissionsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PermissionStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
