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

// GetClustersMulticlusterConfigUUIDReader is a Reader for the GetClustersMulticlusterConfigUUID structure.
type GetClustersMulticlusterConfigUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClustersMulticlusterConfigUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClustersMulticlusterConfigUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetClustersMulticlusterConfigUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetClustersMulticlusterConfigUUIDOK creates a GetClustersMulticlusterConfigUUIDOK with default headers values
func NewGetClustersMulticlusterConfigUUIDOK() *GetClustersMulticlusterConfigUUIDOK {
	return &GetClustersMulticlusterConfigUUIDOK{}
}

/*GetClustersMulticlusterConfigUUIDOK handles this case with default header values.

Success
*/
type GetClustersMulticlusterConfigUUIDOK struct {
	Payload *models.MulticlusterConfigStatus
}

func (o *GetClustersMulticlusterConfigUUIDOK) Error() string {
	return fmt.Sprintf("[GET /clusters/multicluster_config/{uuid}][%d] getClustersMulticlusterConfigUuidOK  %+v", 200, o.Payload)
}

func (o *GetClustersMulticlusterConfigUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MulticlusterConfigStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersMulticlusterConfigUUIDDefault creates a GetClustersMulticlusterConfigUUIDDefault with default headers values
func NewGetClustersMulticlusterConfigUUIDDefault(code int) *GetClustersMulticlusterConfigUUIDDefault {
	return &GetClustersMulticlusterConfigUUIDDefault{
		_statusCode: code,
	}
}

/*GetClustersMulticlusterConfigUUIDDefault handles this case with default header values.

Internal Error
*/
type GetClustersMulticlusterConfigUUIDDefault struct {
	_statusCode int
}

// Code gets the status code for the get clusters multicluster config UUID default response
func (o *GetClustersMulticlusterConfigUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetClustersMulticlusterConfigUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /clusters/multicluster_config/{uuid}][%d] GetClustersMulticlusterConfigUUID default ", o._statusCode)
}

func (o *GetClustersMulticlusterConfigUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
