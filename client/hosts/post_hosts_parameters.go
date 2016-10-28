package hosts

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

// NewPostHostsParams creates a new PostHostsParams object
// with the default values initialized.
func NewPostHostsParams() *PostHostsParams {
	var ()
	return &PostHostsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHostsParamsWithTimeout creates a new PostHostsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHostsParamsWithTimeout(timeout time.Duration) *PostHostsParams {
	var ()
	return &PostHostsParams{

		timeout: timeout,
	}
}

// NewPostHostsParamsWithContext creates a new PostHostsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHostsParamsWithContext(ctx context.Context) *PostHostsParams {
	var ()
	return &PostHostsParams{

		Context: ctx,
	}
}

/*PostHostsParams contains all the parameters to send to the API endpoint
for the post hosts operation typically these are written to a http.Request
*/
type PostHostsParams struct {

	/*HostIntent
	  Intent Spec of Host

	*/
	HostIntent *models.HostIntentInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post hosts params
func (o *PostHostsParams) WithTimeout(timeout time.Duration) *PostHostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post hosts params
func (o *PostHostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post hosts params
func (o *PostHostsParams) WithContext(ctx context.Context) *PostHostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post hosts params
func (o *PostHostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHostIntent adds the hostIntent to the post hosts params
func (o *PostHostsParams) WithHostIntent(hostIntent *models.HostIntentInput) *PostHostsParams {
	o.SetHostIntent(hostIntent)
	return o
}

// SetHostIntent adds the hostIntent to the post hosts params
func (o *PostHostsParams) SetHostIntent(hostIntent *models.HostIntentInput) {
	o.HostIntent = hostIntent
}

// WriteToRequest writes these params to a swagger request
func (o *PostHostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.HostIntent == nil {
		o.HostIntent = new(models.HostIntentInput)
	}

	if err := r.SetBodyParam(o.HostIntent); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
