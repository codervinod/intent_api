package permission

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

// NewPutPermissionsUUIDParams creates a new PutPermissionsUUIDParams object
// with the default values initialized.
func NewPutPermissionsUUIDParams() *PutPermissionsUUIDParams {
	var ()
	return &PutPermissionsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPermissionsUUIDParamsWithTimeout creates a new PutPermissionsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPermissionsUUIDParamsWithTimeout(timeout time.Duration) *PutPermissionsUUIDParams {
	var ()
	return &PutPermissionsUUIDParams{

		timeout: timeout,
	}
}

// NewPutPermissionsUUIDParamsWithContext creates a new PutPermissionsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPermissionsUUIDParamsWithContext(ctx context.Context) *PutPermissionsUUIDParams {
	var ()
	return &PutPermissionsUUIDParams{

		Context: ctx,
	}
}

/*PutPermissionsUUIDParams contains all the parameters to send to the API endpoint
for the put permissions UUID operation typically these are written to a http.Request
*/
type PutPermissionsUUIDParams struct {

	/*Body
	  A Permission object

	*/
	Body *models.Permission
	/*UUID
	  ID of permission that needs to be updated

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put permissions UUID params
func (o *PutPermissionsUUIDParams) WithTimeout(timeout time.Duration) *PutPermissionsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put permissions UUID params
func (o *PutPermissionsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put permissions UUID params
func (o *PutPermissionsUUIDParams) WithContext(ctx context.Context) *PutPermissionsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put permissions UUID params
func (o *PutPermissionsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the put permissions UUID params
func (o *PutPermissionsUUIDParams) WithBody(body *models.Permission) *PutPermissionsUUIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put permissions UUID params
func (o *PutPermissionsUUIDParams) SetBody(body *models.Permission) {
	o.Body = body
}

// WithUUID adds the uuid to the put permissions UUID params
func (o *PutPermissionsUUIDParams) WithUUID(uuid string) *PutPermissionsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the put permissions UUID params
func (o *PutPermissionsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PutPermissionsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.Permission)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
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
