package policy

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

// NewPutPolicyUUIDParams creates a new PutPolicyUUIDParams object
// with the default values initialized.
func NewPutPolicyUUIDParams() *PutPolicyUUIDParams {
	var ()
	return &PutPolicyUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutPolicyUUIDParamsWithTimeout creates a new PutPolicyUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutPolicyUUIDParamsWithTimeout(timeout time.Duration) *PutPolicyUUIDParams {
	var ()
	return &PutPolicyUUIDParams{

		timeout: timeout,
	}
}

// NewPutPolicyUUIDParamsWithContext creates a new PutPolicyUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutPolicyUUIDParamsWithContext(ctx context.Context) *PutPolicyUUIDParams {
	var ()
	return &PutPolicyUUIDParams{

		Context: ctx,
	}
}

/*PutPolicyUUIDParams contains all the parameters to send to the API endpoint
for the put policy UUID operation typically these are written to a http.Request
*/
type PutPolicyUUIDParams struct {

	/*Request*/
	Request *models.Policy
	/*UUID
	  The UUID of the entity.

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put policy UUID params
func (o *PutPolicyUUIDParams) WithTimeout(timeout time.Duration) *PutPolicyUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put policy UUID params
func (o *PutPolicyUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put policy UUID params
func (o *PutPolicyUUIDParams) WithContext(ctx context.Context) *PutPolicyUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put policy UUID params
func (o *PutPolicyUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithRequest adds the request to the put policy UUID params
func (o *PutPolicyUUIDParams) WithRequest(request *models.Policy) *PutPolicyUUIDParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the put policy UUID params
func (o *PutPolicyUUIDParams) SetRequest(request *models.Policy) {
	o.Request = request
}

// WithUUID adds the uuid to the put policy UUID params
func (o *PutPolicyUUIDParams) WithUUID(uuid string) *PutPolicyUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the put policy UUID params
func (o *PutPolicyUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PutPolicyUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Request == nil {
		o.Request = new(models.Policy)
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