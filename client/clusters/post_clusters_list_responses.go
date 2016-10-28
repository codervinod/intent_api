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

// PostClustersListReader is a Reader for the PostClustersList structure.
type PostClustersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClustersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostClustersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostClustersListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostClustersListOK creates a PostClustersListOK with default headers values
func NewPostClustersListOK() *PostClustersListOK {
	return &PostClustersListOK{}
}

/*PostClustersListOK handles this case with default header values.

Request Succeeded
*/
type PostClustersListOK struct {
	Payload *models.ClusterListIntentResponse
}

func (o *PostClustersListOK) Error() string {
	return fmt.Sprintf("[POST /clusters/list][%d] postClustersListOK  %+v", 200, o.Payload)
}

func (o *PostClustersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterListIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClustersListDefault creates a PostClustersListDefault with default headers values
func NewPostClustersListDefault(code int) *PostClustersListDefault {
	return &PostClustersListDefault{
		_statusCode: code,
	}
}

/*PostClustersListDefault handles this case with default header values.

Error
*/
type PostClustersListDefault struct {
	_statusCode int

	Payload *models.ClusterStatus
}

// Code gets the status code for the post clusters list default response
func (o *PostClustersListDefault) Code() int {
	return o._statusCode
}

func (o *PostClustersListDefault) Error() string {
	return fmt.Sprintf("[POST /clusters/list][%d] PostClustersList default  %+v", o._statusCode, o.Payload)
}

func (o *PostClustersListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
