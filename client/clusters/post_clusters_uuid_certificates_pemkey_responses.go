package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostClustersUUIDCertificatesPemkeyReader is a Reader for the PostClustersUUIDCertificatesPemkey structure.
type PostClustersUUIDCertificatesPemkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClustersUUIDCertificatesPemkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostClustersUUIDCertificatesPemkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostClustersUUIDCertificatesPemkeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostClustersUUIDCertificatesPemkeyOK creates a PostClustersUUIDCertificatesPemkeyOK with default headers values
func NewPostClustersUUIDCertificatesPemkeyOK() *PostClustersUUIDCertificatesPemkeyOK {
	return &PostClustersUUIDCertificatesPemkeyOK{}
}

/*PostClustersUUIDCertificatesPemkeyOK handles this case with default header values.

Request Succeeded
*/
type PostClustersUUIDCertificatesPemkeyOK struct {
}

func (o *PostClustersUUIDCertificatesPemkeyOK) Error() string {
	return fmt.Sprintf("[POST /clusters/{uuid}/certificates/pemkey][%d] postClustersUuidCertificatesPemkeyOK ", 200)
}

func (o *PostClustersUUIDCertificatesPemkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostClustersUUIDCertificatesPemkeyDefault creates a PostClustersUUIDCertificatesPemkeyDefault with default headers values
func NewPostClustersUUIDCertificatesPemkeyDefault(code int) *PostClustersUUIDCertificatesPemkeyDefault {
	return &PostClustersUUIDCertificatesPemkeyDefault{
		_statusCode: code,
	}
}

/*PostClustersUUIDCertificatesPemkeyDefault handles this case with default header values.

Internal Error
*/
type PostClustersUUIDCertificatesPemkeyDefault struct {
	_statusCode int
}

// Code gets the status code for the post clusters UUID certificates pemkey default response
func (o *PostClustersUUIDCertificatesPemkeyDefault) Code() int {
	return o._statusCode
}

func (o *PostClustersUUIDCertificatesPemkeyDefault) Error() string {
	return fmt.Sprintf("[POST /clusters/{uuid}/certificates/pemkey][%d] PostClustersUUIDCertificatesPemkey default ", o._statusCode)
}

func (o *PostClustersUUIDCertificatesPemkeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}