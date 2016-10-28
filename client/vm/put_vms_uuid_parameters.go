package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intent_api/models"
)

// NewPutVmsUUIDParams creates a new PutVmsUUIDParams object
// with the default values initialized.
func NewPutVmsUUIDParams() *PutVmsUUIDParams {
	var ()
	return &PutVmsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutVmsUUIDParamsWithTimeout creates a new PutVmsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutVmsUUIDParamsWithTimeout(timeout time.Duration) *PutVmsUUIDParams {
	var ()
	return &PutVmsUUIDParams{

		timeout: timeout,
	}
}

// NewPutVmsUUIDParamsWithContext creates a new PutVmsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutVmsUUIDParamsWithContext(ctx context.Context) *PutVmsUUIDParams {
	var ()
	return &PutVmsUUIDParams{

		Context: ctx,
	}
}

/*PutVmsUUIDParams contains all the parameters to send to the API endpoint
for the put vms UUID operation typically these are written to a http.Request
*/
type PutVmsUUIDParams struct {

	/*Body*/
	Body *models.VMIntentInput
	/*UUID*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put vms UUID params
func (o *PutVmsUUIDParams) WithTimeout(timeout time.Duration) *PutVmsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put vms UUID params
func (o *PutVmsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put vms UUID params
func (o *PutVmsUUIDParams) WithContext(ctx context.Context) *PutVmsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put vms UUID params
func (o *PutVmsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the put vms UUID params
func (o *PutVmsUUIDParams) WithBody(body *models.VMIntentInput) *PutVmsUUIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put vms UUID params
func (o *PutVmsUUIDParams) SetBody(body *models.VMIntentInput) {
	o.Body = body
}

// WithUUID adds the uuid to the put vms UUID params
func (o *PutVmsUUIDParams) WithUUID(uuid string) *PutVmsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the put vms UUID params
func (o *PutVmsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PutVmsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.VMIntentInput)
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