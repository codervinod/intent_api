package catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// GetCatalogItemsUUIDReader is a Reader for the GetCatalogItemsUUID structure.
type GetCatalogItemsUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogItemsUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCatalogItemsUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCatalogItemsUUIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetCatalogItemsUUIDOK creates a GetCatalogItemsUUIDOK with default headers values
func NewGetCatalogItemsUUIDOK() *GetCatalogItemsUUIDOK {
	return &GetCatalogItemsUUIDOK{}
}

/*GetCatalogItemsUUIDOK handles this case with default header values.

Success
*/
type GetCatalogItemsUUIDOK struct {
	Payload *models.Catalog
}

func (o *GetCatalogItemsUUIDOK) Error() string {
	return fmt.Sprintf("[GET /catalog_items/{uuid}][%d] getCatalogItemsUuidOK  %+v", 200, o.Payload)
}

func (o *GetCatalogItemsUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Catalog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogItemsUUIDDefault creates a GetCatalogItemsUUIDDefault with default headers values
func NewGetCatalogItemsUUIDDefault(code int) *GetCatalogItemsUUIDDefault {
	return &GetCatalogItemsUUIDDefault{
		_statusCode: code,
	}
}

/*GetCatalogItemsUUIDDefault handles this case with default header values.

Error
*/
type GetCatalogItemsUUIDDefault struct {
	_statusCode int

	Payload *models.CatalogStatus
}

// Code gets the status code for the get catalog items UUID default response
func (o *GetCatalogItemsUUIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCatalogItemsUUIDDefault) Error() string {
	return fmt.Sprintf("[GET /catalog_items/{uuid}][%d] GetCatalogItemsUUID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCatalogItemsUUIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
