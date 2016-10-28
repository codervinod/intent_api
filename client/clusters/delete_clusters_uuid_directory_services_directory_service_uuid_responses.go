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

// DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDReader is a Reader for the DeleteClustersUUIDDirectoryServicesDirectoryServiceUUID structure.
type DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDOK creates a DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDOK with default headers values
func NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDOK() *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDOK {
	return &DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDOK{}
}

/*DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDOK handles this case with default header values.

Success
*/
type DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDOK struct {
}

func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDOK) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{uuid}/directory_services/{directory_service_uuid}][%d] deleteClustersUuidDirectoryServicesDirectoryServiceUuidOK ", 200)
}

func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound creates a DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound with default headers values
func NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound() *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound {
	return &DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound{}
}

/*DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound handles this case with default header values.

Invalid Directory Uuid Provided
*/
type DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound struct {
	Payload *models.DirectoryServiceStatus
}

func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{uuid}/directory_services/{directory_service_uuid}][%d] deleteClustersUuidDirectoryServicesDirectoryServiceUuidNotFound  %+v", 404, o.Payload)
}

func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryServiceStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault creates a DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault with default headers values
func NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault(code int) *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault {
	return &DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault{
		_statusCode: code,
	}
}

/*DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault handles this case with default header values.

Failed to Delete the directory
*/
type DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault struct {
	_statusCode int

	Payload *models.DirectoryServiceStatus
}

// Code gets the status code for the delete clusters UUID directory services directory service UUID default response
func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{uuid}/directory_services/{directory_service_uuid}][%d] DeleteClustersUUIDDirectoryServicesDirectoryServiceUUID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryServiceStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
