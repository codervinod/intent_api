package category

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

// NewPostCategoriesParams creates a new PostCategoriesParams object
// with the default values initialized.
func NewPostCategoriesParams() *PostCategoriesParams {
	var ()
	return &PostCategoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCategoriesParamsWithTimeout creates a new PostCategoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCategoriesParamsWithTimeout(timeout time.Duration) *PostCategoriesParams {
	var ()
	return &PostCategoriesParams{

		timeout: timeout,
	}
}

// NewPostCategoriesParamsWithContext creates a new PostCategoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCategoriesParamsWithContext(ctx context.Context) *PostCategoriesParams {
	var ()
	return &PostCategoriesParams{

		Context: ctx,
	}
}

/*PostCategoriesParams contains all the parameters to send to the API endpoint
for the post categories operation typically these are written to a http.Request
*/
type PostCategoriesParams struct {

	/*Request*/
	Request *models.CategoryIntentInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post categories params
func (o *PostCategoriesParams) WithTimeout(timeout time.Duration) *PostCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post categories params
func (o *PostCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post categories params
func (o *PostCategoriesParams) WithContext(ctx context.Context) *PostCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post categories params
func (o *PostCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithRequest adds the request to the post categories params
func (o *PostCategoriesParams) WithRequest(request *models.CategoryIntentInput) *PostCategoriesParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post categories params
func (o *PostCategoriesParams) SetRequest(request *models.CategoryIntentInput) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Request == nil {
		o.Request = new(models.CategoryIntentInput)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
