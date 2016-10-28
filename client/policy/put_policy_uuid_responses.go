package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// PutPolicyUUIDReader is a Reader for the PutPolicyUUID structure.
type PutPolicyUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPolicyUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutPolicyUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutPolicyUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutPolicyUUIDOK creates a PutPolicyUUIDOK with default headers values
func NewPutPolicyUUIDOK() *PutPolicyUUIDOK {
	return &PutPolicyUUIDOK{}
}

/*PutPolicyUUIDOK handles this case with default header values.

Success
*/
type PutPolicyUUIDOK struct {
}

func (o *PutPolicyUUIDOK) Error() string {
	return fmt.Sprintf("[PUT /policy/{uuid}][%d] putPolicyUuidOK ", 200)
}

func (o *PutPolicyUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutPolicyUUIDDefault creates a PutPolicyUUIDDefault with default headers values
func NewPutPolicyUUIDDefault(code int) *PutPolicyUUIDDefault {
	return &PutPolicyUUIDDefault{
		_statusCode: code,
	}
}

/*PutPolicyUUIDDefault handles this case with default header values.

Error
*/
type PutPolicyUUIDDefault struct {
	_statusCode int

	Payload *models.PolicyStatus
}

// Code gets the status code for the put policy UUID default response
func (o *PutPolicyUUIDDefault) Code() int {
	return o._statusCode
}

func (o *PutPolicyUUIDDefault) Error() string {
	return fmt.Sprintf("[PUT /policy/{uuid}][%d] PutPolicyUUID default  %+v", o._statusCode, o.Payload)
}

func (o *PutPolicyUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
