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

	"github.com/codervinod/intent_api/models"
)

// NewPostVmsParams creates a new PostVmsParams object
// with the default values initialized.
func NewPostVmsParams() *PostVmsParams {
	var ()
	return &PostVmsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVmsParamsWithTimeout creates a new PostVmsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVmsParamsWithTimeout(timeout time.Duration) *PostVmsParams {
	var ()
	return &PostVmsParams{

		timeout: timeout,
	}
}

// NewPostVmsParamsWithContext creates a new PostVmsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVmsParamsWithContext(ctx context.Context) *PostVmsParams {
	var ()
	return &PostVmsParams{

		Context: ctx,
	}
}

/*PostVmsParams contains all the parameters to send to the API endpoint
for the post vms operation typically these are written to a http.Request
*/
type PostVmsParams struct {

	/*Body*/
	Body *models.VMIntentInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post vms params
func (o *PostVmsParams) WithTimeout(timeout time.Duration) *PostVmsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vms params
func (o *PostVmsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vms params
func (o *PostVmsParams) WithContext(ctx context.Context) *PostVmsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vms params
func (o *PostVmsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post vms params
func (o *PostVmsParams) WithBody(body *models.VMIntentInput) *PostVmsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vms params
func (o *PostVmsParams) SetBody(body *models.VMIntentInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVmsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.VMIntentInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
