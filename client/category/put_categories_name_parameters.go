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

// NewPutCategoriesNameParams creates a new PutCategoriesNameParams object
// with the default values initialized.
func NewPutCategoriesNameParams() *PutCategoriesNameParams {
	var ()
	return &PutCategoriesNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCategoriesNameParamsWithTimeout creates a new PutCategoriesNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCategoriesNameParamsWithTimeout(timeout time.Duration) *PutCategoriesNameParams {
	var ()
	return &PutCategoriesNameParams{

		timeout: timeout,
	}
}

// NewPutCategoriesNameParamsWithContext creates a new PutCategoriesNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCategoriesNameParamsWithContext(ctx context.Context) *PutCategoriesNameParams {
	var ()
	return &PutCategoriesNameParams{

		Context: ctx,
	}
}

/*PutCategoriesNameParams contains all the parameters to send to the API endpoint
for the put categories name operation typically these are written to a http.Request
*/
type PutCategoriesNameParams struct {

	/*Name*/
	Name string
	/*Request*/
	Request *models.CategoryIntentInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put categories name params
func (o *PutCategoriesNameParams) WithTimeout(timeout time.Duration) *PutCategoriesNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put categories name params
func (o *PutCategoriesNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put categories name params
func (o *PutCategoriesNameParams) WithContext(ctx context.Context) *PutCategoriesNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put categories name params
func (o *PutCategoriesNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithName adds the name to the put categories name params
func (o *PutCategoriesNameParams) WithName(name string) *PutCategoriesNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put categories name params
func (o *PutCategoriesNameParams) SetName(name string) {
	o.Name = name
}

// WithRequest adds the request to the put categories name params
func (o *PutCategoriesNameParams) WithRequest(request *models.CategoryIntentInput) *PutCategoriesNameParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the put categories name params
func (o *PutCategoriesNameParams) SetRequest(request *models.CategoryIntentInput) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PutCategoriesNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

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
