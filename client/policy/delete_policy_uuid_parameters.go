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
)

// NewDeletePolicyUUIDParams creates a new DeletePolicyUUIDParams object
// with the default values initialized.
func NewDeletePolicyUUIDParams() *DeletePolicyUUIDParams {
	var ()
	return &DeletePolicyUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePolicyUUIDParamsWithTimeout creates a new DeletePolicyUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePolicyUUIDParamsWithTimeout(timeout time.Duration) *DeletePolicyUUIDParams {
	var ()
	return &DeletePolicyUUIDParams{

		timeout: timeout,
	}
}

// NewDeletePolicyUUIDParamsWithContext creates a new DeletePolicyUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePolicyUUIDParamsWithContext(ctx context.Context) *DeletePolicyUUIDParams {
	var ()
	return &DeletePolicyUUIDParams{

		Context: ctx,
	}
}

/*DeletePolicyUUIDParams contains all the parameters to send to the API endpoint
for the delete policy UUID operation typically these are written to a http.Request
*/
type DeletePolicyUUIDParams struct {

	/*UUID
	  The UUID of the entity.

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the delete policy UUID params
func (o *DeletePolicyUUIDParams) WithTimeout(timeout time.Duration) *DeletePolicyUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete policy UUID params
func (o *DeletePolicyUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete policy UUID params
func (o *DeletePolicyUUIDParams) WithContext(ctx context.Context) *DeletePolicyUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete policy UUID params
func (o *DeletePolicyUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithUUID adds the uuid to the delete policy UUID params
func (o *DeletePolicyUUIDParams) WithUUID(uuid string) *DeletePolicyUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete policy UUID params
func (o *DeletePolicyUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePolicyUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
