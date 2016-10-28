package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// NewPutImagesUUIDParams creates a new PutImagesUUIDParams object
// with the default values initialized.
func NewPutImagesUUIDParams() *PutImagesUUIDParams {
	var ()
	return &PutImagesUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutImagesUUIDParamsWithTimeout creates a new PutImagesUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutImagesUUIDParamsWithTimeout(timeout time.Duration) *PutImagesUUIDParams {
	var ()
	return &PutImagesUUIDParams{

		timeout: timeout,
	}
}

// NewPutImagesUUIDParamsWithContext creates a new PutImagesUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutImagesUUIDParamsWithContext(ctx context.Context) *PutImagesUUIDParams {
	var ()
	return &PutImagesUUIDParams{

		Context: ctx,
	}
}

/*PutImagesUUIDParams contains all the parameters to send to the API endpoint
for the put images UUID operation typically these are written to a http.Request
*/
type PutImagesUUIDParams struct {

	/*Request*/
	Request *models.Image
	/*UUID
	  The UUID of the entity.

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put images UUID params
func (o *PutImagesUUIDParams) WithTimeout(timeout time.Duration) *PutImagesUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put images UUID params
func (o *PutImagesUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put images UUID params
func (o *PutImagesUUIDParams) WithContext(ctx context.Context) *PutImagesUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put images UUID params
func (o *PutImagesUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithRequest adds the request to the put images UUID params
func (o *PutImagesUUIDParams) WithRequest(request *models.Image) *PutImagesUUIDParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the put images UUID params
func (o *PutImagesUUIDParams) SetRequest(request *models.Image) {
	o.Request = request
}

// WithUUID adds the uuid to the put images UUID params
func (o *PutImagesUUIDParams) WithUUID(uuid string) *PutImagesUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the put images UUID params
func (o *PutImagesUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PutImagesUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Request == nil {
		o.Request = new(models.Image)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
