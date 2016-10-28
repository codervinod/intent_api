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

// PostClustersUUIDDirectoryServicesListReader is a Reader for the PostClustersUUIDDirectoryServicesList structure.
type PostClustersUUIDDirectoryServicesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClustersUUIDDirectoryServicesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostClustersUUIDDirectoryServicesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostClustersUUIDDirectoryServicesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostClustersUUIDDirectoryServicesListOK creates a PostClustersUUIDDirectoryServicesListOK with default headers values
func NewPostClustersUUIDDirectoryServicesListOK() *PostClustersUUIDDirectoryServicesListOK {
	return &PostClustersUUIDDirectoryServicesListOK{}
}

/*PostClustersUUIDDirectoryServicesListOK handles this case with default header values.

Success
*/
type PostClustersUUIDDirectoryServicesListOK struct {
	Payload *models.DirectoryServiceListIntentResponse
}

func (o *PostClustersUUIDDirectoryServicesListOK) Error() string {
	return fmt.Sprintf("[POST /clusters/{uuid}/directory_services/list][%d] postClustersUuidDirectoryServicesListOK  %+v", 200, o.Payload)
}

func (o *PostClustersUUIDDirectoryServicesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryServiceListIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClustersUUIDDirectoryServicesListDefault creates a PostClustersUUIDDirectoryServicesListDefault with default headers values
func NewPostClustersUUIDDirectoryServicesListDefault(code int) *PostClustersUUIDDirectoryServicesListDefault {
	return &PostClustersUUIDDirectoryServicesListDefault{
		_statusCode: code,
	}
}

/*PostClustersUUIDDirectoryServicesListDefault handles this case with default header values.

Failed to retrieve all directories
*/
type PostClustersUUIDDirectoryServicesListDefault struct {
	_statusCode int

	Payload *models.DirectoryServiceStatus
}

// Code gets the status code for the post clusters UUID directory services list default response
func (o *PostClustersUUIDDirectoryServicesListDefault) Code() int {
	return o._statusCode
}

func (o *PostClustersUUIDDirectoryServicesListDefault) Error() string {
	return fmt.Sprintf("[POST /clusters/{uuid}/directory_services/list][%d] PostClustersUUIDDirectoryServicesList default  %+v", o._statusCode, o.Payload)
}

func (o *PostClustersUUIDDirectoryServicesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryServiceStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
