package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// GetClustersUUIDReader is a Reader for the GetClustersUUID structure.
type GetClustersUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClustersUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClustersUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetClustersUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetClustersUUIDOK creates a GetClustersUUIDOK with default headers values
func NewGetClustersUUIDOK() *GetClustersUUIDOK {
	return &GetClustersUUIDOK{}
}

/*GetClustersUUIDOK handles this case with default header values.

Request succeeded
*/
type GetClustersUUIDOK struct {
	Payload *models.ClusterIntentResponse
}

func (o *GetClustersUUIDOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{uuid}][%d] getClustersUuidOK  %+v", 200, o.Payload)
}

func (o *GetClustersUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersUUIDDefault creates a GetClustersUUIDDefault with default headers values
func NewGetClustersUUIDDefault(code int) *GetClustersUUIDDefault {
	return &GetClustersUUIDDefault{
		_statusCode: code,
	}
}

/*GetClustersUUIDDefault handles this case with default header values.

Internal Error
*/
type GetClustersUUIDDefault struct {
	_statusCode int

	Payload *models.ClusterStatus
}

// Code gets the status code for the get clusters UUID default response
func (o *GetClustersUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetClustersUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /clusters/{uuid}][%d] GetClustersUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetClustersUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
