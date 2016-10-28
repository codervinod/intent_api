package project

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

// NewGetProjectsUUIDParams creates a new GetProjectsUUIDParams object
// with the default values initialized.
func NewGetProjectsUUIDParams() *GetProjectsUUIDParams {
	var ()
	return &GetProjectsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsUUIDParamsWithTimeout creates a new GetProjectsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectsUUIDParamsWithTimeout(timeout time.Duration) *GetProjectsUUIDParams {
	var ()
	return &GetProjectsUUIDParams{

		timeout: timeout,
	}
}

// NewGetProjectsUUIDParamsWithContext creates a new GetProjectsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectsUUIDParamsWithContext(ctx context.Context) *GetProjectsUUIDParams {
	var ()
	return &GetProjectsUUIDParams{

		Context: ctx,
	}
}

/*GetProjectsUUIDParams contains all the parameters to send to the API endpoint
for the get projects UUID operation typically these are written to a http.Request
*/
type GetProjectsUUIDParams struct {

	/*UUID
	  The UUID of the entity.

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get projects UUID params
func (o *GetProjectsUUIDParams) WithTimeout(timeout time.Duration) *GetProjectsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects UUID params
func (o *GetProjectsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects UUID params
func (o *GetProjectsUUIDParams) WithContext(ctx context.Context) *GetProjectsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects UUID params
func (o *GetProjectsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithUUID adds the uuid to the get projects UUID params
func (o *GetProjectsUUIDParams) WithUUID(uuid string) *GetProjectsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get projects UUID params
func (o *GetProjectsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
