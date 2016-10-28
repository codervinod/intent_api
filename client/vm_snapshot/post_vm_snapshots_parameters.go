package vm_snapshot

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

// NewPostVMSnapshotsParams creates a new PostVMSnapshotsParams object
// with the default values initialized.
func NewPostVMSnapshotsParams() *PostVMSnapshotsParams {
	var ()
	return &PostVMSnapshotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVMSnapshotsParamsWithTimeout creates a new PostVMSnapshotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVMSnapshotsParamsWithTimeout(timeout time.Duration) *PostVMSnapshotsParams {
	var ()
	return &PostVMSnapshotsParams{

		timeout: timeout,
	}
}

// NewPostVMSnapshotsParamsWithContext creates a new PostVMSnapshotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVMSnapshotsParamsWithContext(ctx context.Context) *PostVMSnapshotsParams {
	var ()
	return &PostVMSnapshotsParams{

		Context: ctx,
	}
}

/*PostVMSnapshotsParams contains all the parameters to send to the API endpoint
for the post VM snapshots operation typically these are written to a http.Request
*/
type PostVMSnapshotsParams struct {

	/*Body*/
	Body *models.VMSnapshotIntentInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post VM snapshots params
func (o *PostVMSnapshotsParams) WithTimeout(timeout time.Duration) *PostVMSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post VM snapshots params
func (o *PostVMSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post VM snapshots params
func (o *PostVMSnapshotsParams) WithContext(ctx context.Context) *PostVMSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post VM snapshots params
func (o *PostVMSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post VM snapshots params
func (o *PostVMSnapshotsParams) WithBody(body *models.VMSnapshotIntentInput) *PostVMSnapshotsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post VM snapshots params
func (o *PostVMSnapshotsParams) SetBody(body *models.VMSnapshotIntentInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVMSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.VMSnapshotIntentInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
