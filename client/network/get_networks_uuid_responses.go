package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// GetNetworksUUIDReader is a Reader for the GetNetworksUUID structure.
type GetNetworksUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworksUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNetworksUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetNetworksUUIDOK creates a GetNetworksUUIDOK with default headers values
func NewGetNetworksUUIDOK() *GetNetworksUUIDOK {
	return &GetNetworksUUIDOK{}
}

/*GetNetworksUUIDOK handles this case with default header values.

Success
*/
type GetNetworksUUIDOK struct {
	Payload *models.Network
}

func (o *GetNetworksUUIDOK) Error() string {
	return fmt.Sprintf("[GET /networks/{uuid}][%d] getNetworksUuidOK  %+v", 200, o.Payload)
}

func (o *GetNetworksUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksUUIDDefault creates a GetNetworksUUIDDefault with default headers values
func NewGetNetworksUUIDDefault(code int) *GetNetworksUUIDDefault {
	return &GetNetworksUUIDDefault{
		_statusCode: code,
	}
}

/*GetNetworksUUIDDefault handles this case with default header values.

Error
*/
type GetNetworksUUIDDefault struct {
	_statusCode int

	Payload *models.NetworkStatus
}

// Code gets the status code for the get networks UUID default response
func (o *GetNetworksUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{uuid}][%d] GetNetworksUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
