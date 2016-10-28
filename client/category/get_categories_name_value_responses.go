package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// GetCategoriesNameValueReader is a Reader for the GetCategoriesNameValue structure.
type GetCategoriesNameValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCategoriesNameValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCategoriesNameValueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCategoriesNameValueDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetCategoriesNameValueOK creates a GetCategoriesNameValueOK with default headers values
func NewGetCategoriesNameValueOK() *GetCategoriesNameValueOK {
	return &GetCategoriesNameValueOK{}
}

/*GetCategoriesNameValueOK handles this case with default header values.

Successful
*/
type GetCategoriesNameValueOK struct {
	Payload *models.CategoryValueIntentResponse
}

func (o *GetCategoriesNameValueOK) Error() string {
	return fmt.Sprintf("[GET /categories/{name}/{value}][%d] getCategoriesNameValueOK  %+v", 200, o.Payload)
}

func (o *GetCategoriesNameValueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryValueIntentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCategoriesNameValueDefault creates a GetCategoriesNameValueDefault with default headers values
func NewGetCategoriesNameValueDefault(code int) *GetCategoriesNameValueDefault {
	return &GetCategoriesNameValueDefault{
		_statusCode: code,
	}
}

/*GetCategoriesNameValueDefault handles this case with default header values.

Internal Error
*/
type GetCategoriesNameValueDefault struct {
	_statusCode int

	Payload *models.CategoryStatus
}

// Code gets the status code for the get categories name value default response
func (o *GetCategoriesNameValueDefault) Code() int {
	return o._statusCode
}

func (o *GetCategoriesNameValueDefault) Error() string {
	return fmt.Sprintf("[GET /categories/{name}/{value}][%d] GetCategoriesNameValue default  %+v", o._statusCode, o.Payload)
}

func (o *GetCategoriesNameValueDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
