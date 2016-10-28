package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// PostDataChangedRegionsReader is a Reader for the PostDataChangedRegions structure.
type PostDataChangedRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDataChangedRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostDataChangedRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostDataChangedRegionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostDataChangedRegionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPostDataChangedRegionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostDataChangedRegionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostDataChangedRegionsOK creates a PostDataChangedRegionsOK with default headers values
func NewPostDataChangedRegionsOK() *PostDataChangedRegionsOK {
	return &PostDataChangedRegionsOK{}
}

/*PostDataChangedRegionsOK handles this case with default header values.

Success
*/
type PostDataChangedRegionsOK struct {
	Payload *models.ChangedRegions
}

func (o *PostDataChangedRegionsOK) Error() string {
	return fmt.Sprintf("[POST /data/changed_regions][%d] postDataChangedRegionsOK  %+v", 200, o.Payload)
}

func (o *PostDataChangedRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangedRegions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDataChangedRegionsBadRequest creates a PostDataChangedRegionsBadRequest with default headers values
func NewPostDataChangedRegionsBadRequest() *PostDataChangedRegionsBadRequest {
	return &PostDataChangedRegionsBadRequest{}
}

/*PostDataChangedRegionsBadRequest handles this case with default header values.

Bad Request. Returned due to validation errors.
*/
type PostDataChangedRegionsBadRequest struct {
	Payload *models.ChangedRegionsStatus
}

func (o *PostDataChangedRegionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /data/changed_regions][%d] postDataChangedRegionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostDataChangedRegionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangedRegionsStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDataChangedRegionsNotFound creates a PostDataChangedRegionsNotFound with default headers values
func NewPostDataChangedRegionsNotFound() *PostDataChangedRegionsNotFound {
	return &PostDataChangedRegionsNotFound{}
}

/*PostDataChangedRegionsNotFound handles this case with default header values.

Not Found. Returned when the file(s) specified by the snapshot_file_path and/or the reference_snapshot_file_path does not exist.

*/
type PostDataChangedRegionsNotFound struct {
	Payload *models.ChangedRegionsStatus
}

func (o *PostDataChangedRegionsNotFound) Error() string {
	return fmt.Sprintf("[POST /data/changed_regions][%d] postDataChangedRegionsNotFound  %+v", 404, o.Payload)
}

func (o *PostDataChangedRegionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangedRegionsStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDataChangedRegionsServiceUnavailable creates a PostDataChangedRegionsServiceUnavailable with default headers values
func NewPostDataChangedRegionsServiceUnavailable() *PostDataChangedRegionsServiceUnavailable {
	return &PostDataChangedRegionsServiceUnavailable{}
}

/*PostDataChangedRegionsServiceUnavailable handles this case with default header values.

Service Unavailable. Returned when the system cannot currently handle the request possibly due to overloading.

*/
type PostDataChangedRegionsServiceUnavailable struct {
	Payload *models.ChangedRegionsStatus
}

func (o *PostDataChangedRegionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /data/changed_regions][%d] postDataChangedRegionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostDataChangedRegionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangedRegionsStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDataChangedRegionsDefault creates a PostDataChangedRegionsDefault with default headers values
func NewPostDataChangedRegionsDefault(code int) *PostDataChangedRegionsDefault {
	return &PostDataChangedRegionsDefault{
		_statusCode: code,
	}
}

/*PostDataChangedRegionsDefault handles this case with default header values.

Internal server error
*/
type PostDataChangedRegionsDefault struct {
	_statusCode int

	Payload *models.ChangedRegionsStatus
}

// Code gets the status code for the post data changed regions default response
func (o *PostDataChangedRegionsDefault) Code() int {
	return o._statusCode
}

func (o *PostDataChangedRegionsDefault) Error() string {
	return fmt.Sprintf("[POST /data/changed_regions][%d] PostDataChangedRegions default  %+v", o._statusCode, o.Payload)
}

func (o *PostDataChangedRegionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangedRegionsStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
