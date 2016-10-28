package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// DeleteClustersUUIDCloudCredentialsCloudTypeReader is a Reader for the DeleteClustersUUIDCloudCredentialsCloudType structure.
type DeleteClustersUUIDCloudCredentialsCloudTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClustersUUIDCloudCredentialsCloudTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClustersUUIDCloudCredentialsCloudTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteClustersUUIDCloudCredentialsCloudTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteClustersUUIDCloudCredentialsCloudTypeOK creates a DeleteClustersUUIDCloudCredentialsCloudTypeOK with default headers values
func NewDeleteClustersUUIDCloudCredentialsCloudTypeOK() *DeleteClustersUUIDCloudCredentialsCloudTypeOK {
	return &DeleteClustersUUIDCloudCredentialsCloudTypeOK{}
}

/*DeleteClustersUUIDCloudCredentialsCloudTypeOK handles this case with default header values.

Request succeeded
*/
type DeleteClustersUUIDCloudCredentialsCloudTypeOK struct {
}

func (o *DeleteClustersUUIDCloudCredentialsCloudTypeOK) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{uuid}/cloud_credentials/{cloud_type}][%d] deleteClustersUuidCloudCredentialsCloudTypeOK ", 200)
}

func (o *DeleteClustersUUIDCloudCredentialsCloudTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClustersUUIDCloudCredentialsCloudTypeDefault creates a DeleteClustersUUIDCloudCredentialsCloudTypeDefault with default headers values
func NewDeleteClustersUUIDCloudCredentialsCloudTypeDefault(code int) *DeleteClustersUUIDCloudCredentialsCloudTypeDefault {
	return &DeleteClustersUUIDCloudCredentialsCloudTypeDefault{
		_statusCode: code,
	}
}

/*DeleteClustersUUIDCloudCredentialsCloudTypeDefault handles this case with default header values.

Internal Error
*/
type DeleteClustersUUIDCloudCredentialsCloudTypeDefault struct {
	_statusCode int

	Payload *models.CloudCredentialsStatus
}

// Code gets the status code for the delete clusters UUID cloud credentials cloud type default response
func (o *DeleteClustersUUIDCloudCredentialsCloudTypeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteClustersUUIDCloudCredentialsCloudTypeDefault) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{uuid}/cloud_credentials/{cloud_type}][%d] DeleteClustersUUIDCloudCredentialsCloudType default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClustersUUIDCloudCredentialsCloudTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudCredentialsStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}