package idempotence_identifiers

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

// NewPostIdempotenceIdentifiersParams creates a new PostIdempotenceIdentifiersParams object
// with the default values initialized.
func NewPostIdempotenceIdentifiersParams() *PostIdempotenceIdentifiersParams {
	var ()
	return &PostIdempotenceIdentifiersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIdempotenceIdentifiersParamsWithTimeout creates a new PostIdempotenceIdentifiersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIdempotenceIdentifiersParamsWithTimeout(timeout time.Duration) *PostIdempotenceIdentifiersParams {
	var ()
	return &PostIdempotenceIdentifiersParams{

		timeout: timeout,
	}
}

// NewPostIdempotenceIdentifiersParamsWithContext creates a new PostIdempotenceIdentifiersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIdempotenceIdentifiersParamsWithContext(ctx context.Context) *PostIdempotenceIdentifiersParams {
	var ()
	return &PostIdempotenceIdentifiersParams{

		Context: ctx,
	}
}

/*PostIdempotenceIdentifiersParams contains all the parameters to send to the API endpoint
for the post idempotence identifiers operation typically these are written to a http.Request
*/
type PostIdempotenceIdentifiersParams struct {

	/*Body
	  An idempotence identifier object

	*/
	Body *models.IdempotenceIdentifiersInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post idempotence identifiers params
func (o *PostIdempotenceIdentifiersParams) WithTimeout(timeout time.Duration) *PostIdempotenceIdentifiersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post idempotence identifiers params
func (o *PostIdempotenceIdentifiersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post idempotence identifiers params
func (o *PostIdempotenceIdentifiersParams) WithContext(ctx context.Context) *PostIdempotenceIdentifiersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post idempotence identifiers params
func (o *PostIdempotenceIdentifiersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post idempotence identifiers params
func (o *PostIdempotenceIdentifiersParams) WithBody(body *models.IdempotenceIdentifiersInput) *PostIdempotenceIdentifiersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post idempotence identifiers params
func (o *PostIdempotenceIdentifiersParams) SetBody(body *models.IdempotenceIdentifiersInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostIdempotenceIdentifiersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.IdempotenceIdentifiersInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
