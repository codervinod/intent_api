package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteClustersUUIDCertificatesCaCertsCaNameReader is a Reader for the DeleteClustersUUIDCertificatesCaCertsCaName structure.
type DeleteClustersUUIDCertificatesCaCertsCaNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClustersUUIDCertificatesCaCertsCaNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClustersUUIDCertificatesCaCertsCaNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteClustersUUIDCertificatesCaCertsCaNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteClustersUUIDCertificatesCaCertsCaNameOK creates a DeleteClustersUUIDCertificatesCaCertsCaNameOK with default headers values
func NewDeleteClustersUUIDCertificatesCaCertsCaNameOK() *DeleteClustersUUIDCertificatesCaCertsCaNameOK {
	return &DeleteClustersUUIDCertificatesCaCertsCaNameOK{}
}

/*DeleteClustersUUIDCertificatesCaCertsCaNameOK handles this case with default header values.

Request Succeeded
*/
type DeleteClustersUUIDCertificatesCaCertsCaNameOK struct {
}

func (o *DeleteClustersUUIDCertificatesCaCertsCaNameOK) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{uuid}/certificates/ca_certs/{ca_name}][%d] deleteClustersUuidCertificatesCaCertsCaNameOK ", 200)
}

func (o *DeleteClustersUUIDCertificatesCaCertsCaNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClustersUUIDCertificatesCaCertsCaNameDefault creates a DeleteClustersUUIDCertificatesCaCertsCaNameDefault with default headers values
func NewDeleteClustersUUIDCertificatesCaCertsCaNameDefault(code int) *DeleteClustersUUIDCertificatesCaCertsCaNameDefault {
	return &DeleteClustersUUIDCertificatesCaCertsCaNameDefault{
		_statusCode: code,
	}
}

/*DeleteClustersUUIDCertificatesCaCertsCaNameDefault handles this case with default header values.

Internal Error
*/
type DeleteClustersUUIDCertificatesCaCertsCaNameDefault struct {
	_statusCode int
}

// Code gets the status code for the delete clusters UUID certificates ca certs ca name default response
func (o *DeleteClustersUUIDCertificatesCaCertsCaNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteClustersUUIDCertificatesCaCertsCaNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{uuid}/certificates/ca_certs/{ca_name}][%d] DeleteClustersUUIDCertificatesCaCertsCaName default ", o._statusCode)
}

func (o *DeleteClustersUUIDCertificatesCaCertsCaNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
