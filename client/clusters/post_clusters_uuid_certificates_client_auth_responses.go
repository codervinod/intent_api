package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostClustersUUIDCertificatesClientAuthReader is a Reader for the PostClustersUUIDCertificatesClientAuth structure.
type PostClustersUUIDCertificatesClientAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClustersUUIDCertificatesClientAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostClustersUUIDCertificatesClientAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostClustersUUIDCertificatesClientAuthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostClustersUUIDCertificatesClientAuthOK creates a PostClustersUUIDCertificatesClientAuthOK with default headers values
func NewPostClustersUUIDCertificatesClientAuthOK() *PostClustersUUIDCertificatesClientAuthOK {
	return &PostClustersUUIDCertificatesClientAuthOK{}
}

/*PostClustersUUIDCertificatesClientAuthOK handles this case with default header values.

Request Succeeded
*/
type PostClustersUUIDCertificatesClientAuthOK struct {
}

func (o *PostClustersUUIDCertificatesClientAuthOK) Error() string {
	return fmt.Sprintf("[POST /clusters/{uuid}/certificates/client_auth][%d] postClustersUuidCertificatesClientAuthOK ", 200)
}

func (o *PostClustersUUIDCertificatesClientAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostClustersUUIDCertificatesClientAuthDefault creates a PostClustersUUIDCertificatesClientAuthDefault with default headers values
func NewPostClustersUUIDCertificatesClientAuthDefault(code int) *PostClustersUUIDCertificatesClientAuthDefault {
	return &PostClustersUUIDCertificatesClientAuthDefault{
		_statusCode: code,
	}
}

/*PostClustersUUIDCertificatesClientAuthDefault handles this case with default header values.

Internal Error
*/
type PostClustersUUIDCertificatesClientAuthDefault struct {
	_statusCode int
}

// Code gets the status code for the post clusters UUID certificates client auth default response
func (o *PostClustersUUIDCertificatesClientAuthDefault) Code() int {
	return o._statusCode
}

func (o *PostClustersUUIDCertificatesClientAuthDefault) Error() string {
	return fmt.Sprintf("[POST /clusters/{uuid}/certificates/client_auth][%d] PostClustersUUIDCertificatesClientAuth default ", o._statusCode)
}

func (o *PostClustersUUIDCertificatesClientAuthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
