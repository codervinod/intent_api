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

// GetClustersUUIDDirectoryServicesDirectoryServiceUUIDReader is a Reader for the GetClustersUUIDDirectoryServicesDirectoryServiceUUID structure.
type GetClustersUUIDDirectoryServicesDirectoryServiceUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClustersUUIDDirectoryServicesDirectoryServiceUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClustersUUIDDirectoryServicesDirectoryServiceUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetClustersUUIDDirectoryServicesDirectoryServiceUUIDOK creates a GetClustersUUIDDirectoryServicesDirectoryServiceUUIDOK with default headers values
func NewGetClustersUUIDDirectoryServicesDirectoryServiceUUIDOK() *GetClustersUUIDDirectoryServicesDirectoryServiceUUIDOK {
	return &GetClustersUUIDDirectoryServicesDirectoryServiceUUIDOK{}
}

/*GetClustersUUIDDirectoryServicesDirectoryServiceUUIDOK handles this case with default header values.

Success
*/
type GetClustersUUIDDirectoryServicesDirectoryServiceUUIDOK struct {
	Payload *models.DirectoryServiceIntentResponse
}

func (o *GetClustersUUIDDirectoryServicesDirectoryServiceUUIDOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{uuid}/directory_services/{directory_service_uuid}][%d] getClustersUuidDirectoryServicesDirectoryServiceUuidOK  %+v", 200, o.Payload)
}

func (o *GetClustersUUIDDirectoryServicesDirectoryServiceUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryServiceIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound creates a GetClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound with default headers values
func NewGetClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound() *GetClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound {
	return &GetClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound{}
}

/*GetClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound handles this case with default header values.

Invalid Directory Name Provided
*/
type GetClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound struct {
	Payload *models.DirectoryServiceStatus
}

func (o *GetClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{uuid}/directory_services/{directory_service_uuid}][%d] getClustersUuidDirectoryServicesDirectoryServiceUuidNotFound  %+v", 404, o.Payload)
}

func (o *GetClustersUUIDDirectoryServicesDirectoryServiceUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryServiceStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault creates a GetClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault with default headers values
func NewGetClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault(code int) *GetClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault {
	return &GetClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault{
		_statusCode: code,
	}
}

/*GetClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault handles this case with default header values.

Failed to retrieve the directory
*/
type GetClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault struct {
	_statusCode int

	Payload *models.DirectoryServiceStatus
}

// Code gets the status code for the get clusters UUID directory services directory service UUID default response
func (o *GetClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /clusters/{uuid}/directory_services/{directory_service_uuid}][%d] GetClustersUUIDDirectoryServicesDirectoryServiceUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetClustersUUIDDirectoryServicesDirectoryServiceUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryServiceStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
