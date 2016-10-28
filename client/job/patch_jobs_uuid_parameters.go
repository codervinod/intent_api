package job

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

// NewPatchJobsUUIDParams creates a new PatchJobsUUIDParams object
// with the default values initialized.
func NewPatchJobsUUIDParams() *PatchJobsUUIDParams {
	var ()
	return &PatchJobsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchJobsUUIDParamsWithTimeout creates a new PatchJobsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchJobsUUIDParamsWithTimeout(timeout time.Duration) *PatchJobsUUIDParams {
	var ()
	return &PatchJobsUUIDParams{

		timeout: timeout,
	}
}

// NewPatchJobsUUIDParamsWithContext creates a new PatchJobsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchJobsUUIDParamsWithContext(ctx context.Context) *PatchJobsUUIDParams {
	var ()
	return &PatchJobsUUIDParams{

		Context: ctx,
	}
}

/*PatchJobsUUIDParams contains all the parameters to send to the API endpoint
for the patch jobs UUID operation typically these are written to a http.Request
*/
type PatchJobsUUIDParams struct {

	/*Body
	  Request body

	*/
	Body *models.JobIntentInput
	/*UUID
	  Uuid of Job to be updated

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the patch jobs UUID params
func (o *PatchJobsUUIDParams) WithTimeout(timeout time.Duration) *PatchJobsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch jobs UUID params
func (o *PatchJobsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch jobs UUID params
func (o *PatchJobsUUIDParams) WithContext(ctx context.Context) *PatchJobsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch jobs UUID params
func (o *PatchJobsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the patch jobs UUID params
func (o *PatchJobsUUIDParams) WithBody(body *models.JobIntentInput) *PatchJobsUUIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch jobs UUID params
func (o *PatchJobsUUIDParams) SetBody(body *models.JobIntentInput) {
	o.Body = body
}

// WithUUID adds the uuid to the patch jobs UUID params
func (o *PatchJobsUUIDParams) WithUUID(uuid string) *PatchJobsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the patch jobs UUID params
func (o *PatchJobsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchJobsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.JobIntentInput)
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
