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

// PutCategoriesNameReader is a Reader for the PutCategoriesName structure.
type PutCategoriesNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCategoriesNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutCategoriesNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutCategoriesNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutCategoriesNameOK creates a PutCategoriesNameOK with default headers values
func NewPutCategoriesNameOK() *PutCategoriesNameOK {
	return &PutCategoriesNameOK{}
}

/*PutCategoriesNameOK handles this case with default header values.

Success
*/
type PutCategoriesNameOK struct {
	Payload *models.CategoryIntentResponse
}

func (o *PutCategoriesNameOK) Error() string {
	return fmt.Sprintf("[PUT /categories/{name}][%d] putCategoriesNameOK  %+v", 200, o.Payload)
}

func (o *PutCategoriesNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCategoriesNameDefault creates a PutCategoriesNameDefault with default headers values
func NewPutCategoriesNameDefault(code int) *PutCategoriesNameDefault {
	return &PutCategoriesNameDefault{
		_statusCode: code,
	}
}

/*PutCategoriesNameDefault handles this case with default header values.

Error
*/
type PutCategoriesNameDefault struct {
	_statusCode int

	Payload *models.CategoryStatus
}

// Code gets the status code for the put categories name default response
func (o *PutCategoriesNameDefault) Code() int {
	return o._statusCode
}

func (o *PutCategoriesNameDefault) Error() string {
	return fmt.Sprintf("[PUT /categories/{name}][%d] PutCategoriesName default  %+v", o._statusCode, o.Payload)
}

func (o *PutCategoriesNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
