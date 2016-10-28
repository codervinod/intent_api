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
)

// NewDeletePermissionsUUIDParams creates a new DeletePermissionsUUIDParams object
// with the default values initialized.
func NewDeletePermissionsUUIDParams() *DeletePermissionsUUIDParams {
	var ()
	return &DeletePermissionsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePermissionsUUIDParamsWithTimeout creates a new DeletePermissionsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePermissionsUUIDParamsWithTimeout(timeout time.Duration) *DeletePermissionsUUIDParams {
	var ()
	return &DeletePermissionsUUIDParams{

		timeout: timeout,
	}
}

// NewDeletePermissionsUUIDParamsWithContext creates a new DeletePermissionsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePermissionsUUIDParamsWithContext(ctx context.Context) *DeletePermissionsUUIDParams {
	var ()
	return &DeletePermissionsUUIDParams{

		Context: ctx,
	}
}

/*DeletePermissionsUUIDParams contains all the parameters to send to the API endpoint
for the delete permissions UUID operation typically these are written to a http.Request
*/
type DeletePermissionsUUIDParams struct {

	/*UUID
	  UUID of permission

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the delete permissions UUID params
func (o *DeletePermissionsUUIDParams) WithTimeout(timeout time.Duration) *DeletePermissionsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete permissions UUID params
func (o *DeletePermissionsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete permissions UUID params
func (o *DeletePermissionsUUIDParams) WithContext(ctx context.Context) *DeletePermissionsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete permissions UUID params
func (o *DeletePermissionsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithUUID adds the uuid to the delete permissions UUID params
func (o *DeletePermissionsUUIDParams) WithUUID(uuid string) *DeletePermissionsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete permissions UUID params
func (o *DeletePermissionsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePermissionsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
