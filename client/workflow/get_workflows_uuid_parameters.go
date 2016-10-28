package workflow

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

// NewGetWorkflowsUUIDParams creates a new GetWorkflowsUUIDParams object
// with the default values initialized.
func NewGetWorkflowsUUIDParams() *GetWorkflowsUUIDParams {
	var ()
	return &GetWorkflowsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowsUUIDParamsWithTimeout creates a new GetWorkflowsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkflowsUUIDParamsWithTimeout(timeout time.Duration) *GetWorkflowsUUIDParams {
	var ()
	return &GetWorkflowsUUIDParams{

		timeout: timeout,
	}
}

// NewGetWorkflowsUUIDParamsWithContext creates a new GetWorkflowsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkflowsUUIDParamsWithContext(ctx context.Context) *GetWorkflowsUUIDParams {
	var ()
	return &GetWorkflowsUUIDParams{

		Context: ctx,
	}
}

/*GetWorkflowsUUIDParams contains all the parameters to send to the API endpoint
for the get workflows UUID operation typically these are written to a http.Request
*/
type GetWorkflowsUUIDParams struct {

	/*UUID
	  Uuid of workflow to get

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get workflows UUID params
func (o *GetWorkflowsUUIDParams) WithTimeout(timeout time.Duration) *GetWorkflowsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflows UUID params
func (o *GetWorkflowsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflows UUID params
func (o *GetWorkflowsUUIDParams) WithContext(ctx context.Context) *GetWorkflowsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflows UUID params
func (o *GetWorkflowsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithUUID adds the uuid to the get workflows UUID params
func (o *GetWorkflowsUUIDParams) WithUUID(uuid string) *GetWorkflowsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get workflows UUID params
func (o *GetWorkflowsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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