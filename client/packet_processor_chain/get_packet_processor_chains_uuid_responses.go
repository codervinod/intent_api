package packet_processor_chain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// GetPacketProcessorChainsUUIDReader is a Reader for the GetPacketProcessorChainsUUID structure.
type GetPacketProcessorChainsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPacketProcessorChainsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPacketProcessorChainsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetPacketProcessorChainsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetPacketProcessorChainsUUIDOK creates a GetPacketProcessorChainsUUIDOK with default headers values
func NewGetPacketProcessorChainsUUIDOK() *GetPacketProcessorChainsUUIDOK {
	return &GetPacketProcessorChainsUUIDOK{}
}

/*GetPacketProcessorChainsUUIDOK handles this case with default header values.

Request succeeded
*/
type GetPacketProcessorChainsUUIDOK struct {
	Payload *models.PacketProcessorChainIntentResponse
}

func (o *GetPacketProcessorChainsUUIDOK) Error() string {
	return fmt.Sprintf("[GET /packet_processor_chains/{uuid}][%d] getPacketProcessorChainsUuidOK  %+v", 200, o.Payload)
}

func (o *GetPacketProcessorChainsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PacketProcessorChainIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPacketProcessorChainsUUIDDefault creates a GetPacketProcessorChainsUUIDDefault with default headers values
func NewGetPacketProcessorChainsUUIDDefault(code int) *GetPacketProcessorChainsUUIDDefault {
	return &GetPacketProcessorChainsUUIDDefault{
		_statusCode: code,
	}
}

/*GetPacketProcessorChainsUUIDDefault handles this case with default header values.

Internal Error
*/
type GetPacketProcessorChainsUUIDDefault struct {
	_statusCode int

	Payload *models.PacketProcessorChainStatus
}

// Code gets the status code for the get packet processor chains UUID default response
func (o *GetPacketProcessorChainsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetPacketProcessorChainsUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /packet_processor_chains/{uuid}][%d] GetPacketProcessorChainsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetPacketProcessorChainsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PacketProcessorChainStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
