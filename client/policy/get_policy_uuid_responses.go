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

// GetPolicyUUIDReader is a Reader for the GetPolicyUUID structure.
type GetPolicyUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPolicyUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetPolicyUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetPolicyUUIDOK creates a GetPolicyUUIDOK with default headers values
func NewGetPolicyUUIDOK() *GetPolicyUUIDOK {
	return &GetPolicyUUIDOK{}
}

/*GetPolicyUUIDOK handles this case with default header values.

Success
*/
type GetPolicyUUIDOK struct {
	Payload *models.Policy
}

func (o *GetPolicyUUIDOK) Error() string {
	return fmt.Sprintf("[GET /policy/{uuid}][%d] getPolicyUuidOK  %+v", 200, o.Payload)
}

func (o *GetPolicyUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyUUIDDefault creates a GetPolicyUUIDDefault with default headers values
func NewGetPolicyUUIDDefault(code int) *GetPolicyUUIDDefault {
	return &GetPolicyUUIDDefault{
		_statusCode: code,
	}
}

/*GetPolicyUUIDDefault handles this case with default header values.

Error
*/
type GetPolicyUUIDDefault struct {
	_statusCode int

	Payload *models.PolicyStatus
}

// Code gets the status code for the get policy UUID default response
func (o *GetPolicyUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetPolicyUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /policy/{uuid}][%d] GetPolicyUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetPolicyUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
