package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// GetCategoriesNameReader is a Reader for the GetCategoriesName structure.
type GetCategoriesNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCategoriesNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCategoriesNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCategoriesNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetCategoriesNameOK creates a GetCategoriesNameOK with default headers values
func NewGetCategoriesNameOK() *GetCategoriesNameOK {
	return &GetCategoriesNameOK{}
}

/*GetCategoriesNameOK handles this case with default header values.

Successful
*/
type GetCategoriesNameOK struct {
	Payload *models.CategoryIntentResponse
}

func (o *GetCategoriesNameOK) Error() string {
	return fmt.Sprintf("[GET /categories/{name}][%d] getCategoriesNameOK  %+v", 200, o.Payload)
}

func (o *GetCategoriesNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCategoriesNameDefault creates a GetCategoriesNameDefault with default headers values
func NewGetCategoriesNameDefault(code int) *GetCategoriesNameDefault {
	return &GetCategoriesNameDefault{
		_statusCode: code,
	}
}

/*GetCategoriesNameDefault handles this case with default header values.

Internal Error
*/
type GetCategoriesNameDefault struct {
	_statusCode int

	Payload *models.CategoryStatus
}

// Code gets the status code for the get categories name default response
func (o *GetCategoriesNameDefault) Code() int {
	return o._statusCode
}

func (o *GetCategoriesNameDefault) Error() string {
	return fmt.Sprintf("[GET /categories/{name}][%d] GetCategoriesName default  %+v", o._statusCode, o.Payload)
}

func (o *GetCategoriesNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
