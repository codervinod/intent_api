package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostClustersUUIDCertificatesPemkeyImportReader is a Reader for the PostClustersUUIDCertificatesPemkeyImport structure.
type PostClustersUUIDCertificatesPemkeyImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClustersUUIDCertificatesPemkeyImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostClustersUUIDCertificatesPemkeyImportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostClustersUUIDCertificatesPemkeyImportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostClustersUUIDCertificatesPemkeyImportOK creates a PostClustersUUIDCertificatesPemkeyImportOK with default headers values
func NewPostClustersUUIDCertificatesPemkeyImportOK() *PostClustersUUIDCertificatesPemkeyImportOK {
	return &PostClustersUUIDCertificatesPemkeyImportOK{}
}

/*PostClustersUUIDCertificatesPemkeyImportOK handles this case with default header values.

Request Succeeded
*/
type PostClustersUUIDCertificatesPemkeyImportOK struct {
}

func (o *PostClustersUUIDCertificatesPemkeyImportOK) Error() string {
	return fmt.Sprintf("[POST /clusters/{uuid}/certificates/pemkey/import][%d] postClustersUuidCertificatesPemkeyImportOK ", 200)
}

func (o *PostClustersUUIDCertificatesPemkeyImportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostClustersUUIDCertificatesPemkeyImportDefault creates a PostClustersUUIDCertificatesPemkeyImportDefault with default headers values
func NewPostClustersUUIDCertificatesPemkeyImportDefault(code int) *PostClustersUUIDCertificatesPemkeyImportDefault {
	return &PostClustersUUIDCertificatesPemkeyImportDefault{
		_statusCode: code,
	}
}

/*PostClustersUUIDCertificatesPemkeyImportDefault handles this case with default header values.

Internal Error
*/
type PostClustersUUIDCertificatesPemkeyImportDefault struct {
	_statusCode int
}

// Code gets the status code for the post clusters UUID certificates pemkey import default response
func (o *PostClustersUUIDCertificatesPemkeyImportDefault) Code() int {
	return o._statusCode
}

func (o *PostClustersUUIDCertificatesPemkeyImportDefault) Error() string {
	return fmt.Sprintf("[POST /clusters/{uuid}/certificates/pemkey/import][%d] PostClustersUUIDCertificatesPemkeyImport default ", o._statusCode)
}

func (o *PostClustersUUIDCertificatesPemkeyImportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
