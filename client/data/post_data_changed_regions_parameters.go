package data

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

// NewPostDataChangedRegionsParams creates a new PostDataChangedRegionsParams object
// with the default values initialized.
func NewPostDataChangedRegionsParams() *PostDataChangedRegionsParams {
	var ()
	return &PostDataChangedRegionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDataChangedRegionsParamsWithTimeout creates a new PostDataChangedRegionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDataChangedRegionsParamsWithTimeout(timeout time.Duration) *PostDataChangedRegionsParams {
	var ()
	return &PostDataChangedRegionsParams{

		timeout: timeout,
	}
}

// NewPostDataChangedRegionsParamsWithContext creates a new PostDataChangedRegionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDataChangedRegionsParamsWithContext(ctx context.Context) *PostDataChangedRegionsParams {
	var ()
	return &PostDataChangedRegionsParams{

		Context: ctx,
	}
}

/*PostDataChangedRegionsParams contains all the parameters to send to the API endpoint
for the post data changed regions operation typically these are written to a http.Request
*/
type PostDataChangedRegionsParams struct {

	/*Body*/
	Body *models.ChangedRegionsQuery

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post data changed regions params
func (o *PostDataChangedRegionsParams) WithTimeout(timeout time.Duration) *PostDataChangedRegionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post data changed regions params
func (o *PostDataChangedRegionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post data changed regions params
func (o *PostDataChangedRegionsParams) WithContext(ctx context.Context) *PostDataChangedRegionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post data changed regions params
func (o *PostDataChangedRegionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post data changed regions params
func (o *PostDataChangedRegionsParams) WithBody(body *models.ChangedRegionsQuery) *PostDataChangedRegionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post data changed regions params
func (o *PostDataChangedRegionsParams) SetBody(body *models.ChangedRegionsQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDataChangedRegionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.ChangedRegionsQuery)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
