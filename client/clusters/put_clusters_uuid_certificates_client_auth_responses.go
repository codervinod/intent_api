package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutClustersUUIDCertificatesClientAuthReader is a Reader for the PutClustersUUIDCertificatesClientAuth structure.
type PutClustersUUIDCertificatesClientAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClustersUUIDCertificatesClientAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutClustersUUIDCertificatesClientAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutClustersUUIDCertificatesClientAuthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutClustersUUIDCertificatesClientAuthOK creates a PutClustersUUIDCertificatesClientAuthOK with default headers values
func NewPutClustersUUIDCertificatesClientAuthOK() *PutClustersUUIDCertificatesClientAuthOK {
	return &PutClustersUUIDCertificatesClientAuthOK{}
}

/*PutClustersUUIDCertificatesClientAuthOK handles this case with default header values.

Request Succeeded
*/
type PutClustersUUIDCertificatesClientAuthOK struct {
}

func (o *PutClustersUUIDCertificatesClientAuthOK) Error() string {
	return fmt.Sprintf("[PUT /clusters/{uuid}/certificates/client_auth][%d] putClustersUuidCertificatesClientAuthOK ", 200)
}

func (o *PutClustersUUIDCertificatesClientAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutClustersUUIDCertificatesClientAuthDefault creates a PutClustersUUIDCertificatesClientAuthDefault with default headers values
func NewPutClustersUUIDCertificatesClientAuthDefault(code int) *PutClustersUUIDCertificatesClientAuthDefault {
	return &PutClustersUUIDCertificatesClientAuthDefault{
		_statusCode: code,
	}
}

/*PutClustersUUIDCertificatesClientAuthDefault handles this case with default header values.

Internal Error
*/
type PutClustersUUIDCertificatesClientAuthDefault struct {
	_statusCode int
}

// Code gets the status code for the put clusters UUID certificates client auth default response
func (o *PutClustersUUIDCertificatesClientAuthDefault) Code() int {
	return o._statusCode
}

func (o *PutClustersUUIDCertificatesClientAuthDefault) Error() string {
	return fmt.Sprintf("[PUT /clusters/{uuid}/certificates/client_auth][%d] PutClustersUUIDCertificatesClientAuth default ", o._statusCode)
}

func (o *PutClustersUUIDCertificatesClientAuthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}