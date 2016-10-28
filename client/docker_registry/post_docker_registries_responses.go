package docker_registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// PostDockerRegistriesReader is a Reader for the PostDockerRegistries structure.
type PostDockerRegistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDockerRegistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostDockerRegistriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostDockerRegistriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostDockerRegistriesOK creates a PostDockerRegistriesOK with default headers values
func NewPostDockerRegistriesOK() *PostDockerRegistriesOK {
	return &PostDockerRegistriesOK{}
}

/*PostDockerRegistriesOK handles this case with default header values.

Created docker registry
*/
type PostDockerRegistriesOK struct {
	Payload *models.DockerRegistryIntentResponse
}

func (o *PostDockerRegistriesOK) Error() string {
	return fmt.Sprintf("[POST /docker_registries][%d] postDockerRegistriesOK  %+v", 200, o.Payload)
}

func (o *PostDockerRegistriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DockerRegistryIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDockerRegistriesDefault creates a PostDockerRegistriesDefault with default headers values
func NewPostDockerRegistriesDefault(code int) *PostDockerRegistriesDefault {
	return &PostDockerRegistriesDefault{
		_statusCode: code,
	}
}

/*PostDockerRegistriesDefault handles this case with default header values.

Failed to create docker registry
*/
type PostDockerRegistriesDefault struct {
	_statusCode int

	Payload *models.DockerRegistryStatus
}

// Code gets the status code for the post docker registries default response
func (o *PostDockerRegistriesDefault) Code() int {
	return o._statusCode
}

func (o *PostDockerRegistriesDefault) Error() string {
	return fmt.Sprintf("[POST /docker_registries][%d] PostDockerRegistries default  %+v", o._statusCode, o.Payload)
}

func (o *PostDockerRegistriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DockerRegistryStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
