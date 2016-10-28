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

	"github.com/intent_api/models"
)

// NewPutCategoriesNameValueParams creates a new PutCategoriesNameValueParams object
// with the default values initialized.
func NewPutCategoriesNameValueParams() *PutCategoriesNameValueParams {
	var ()
	return &PutCategoriesNameValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCategoriesNameValueParamsWithTimeout creates a new PutCategoriesNameValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCategoriesNameValueParamsWithTimeout(timeout time.Duration) *PutCategoriesNameValueParams {
	var ()
	return &PutCategoriesNameValueParams{

		timeout: timeout,
	}
}

// NewPutCategoriesNameValueParamsWithContext creates a new PutCategoriesNameValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCategoriesNameValueParamsWithContext(ctx context.Context) *PutCategoriesNameValueParams {
	var ()
	return &PutCategoriesNameValueParams{

		Context: ctx,
	}
}

/*PutCategoriesNameValueParams contains all the parameters to send to the API endpoint
for the put categories name value operation typically these are written to a http.Request
*/
type PutCategoriesNameValueParams struct {

	/*Name*/
	Name string
	/*Request*/
	Request *models.CategoryValueIntentInput
	/*Value*/
	Value string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put categories name value params
func (o *PutCategoriesNameValueParams) WithTimeout(timeout time.Duration) *PutCategoriesNameValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put categories name value params
func (o *PutCategoriesNameValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put categories name value params
func (o *PutCategoriesNameValueParams) WithContext(ctx context.Context) *PutCategoriesNameValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put categories name value params
func (o *PutCategoriesNameValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithName adds the name to the put categories name value params
func (o *PutCategoriesNameValueParams) WithName(name string) *PutCategoriesNameValueParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put categories name value params
func (o *PutCategoriesNameValueParams) SetName(name string) {
	o.Name = name
}

// WithRequest adds the request to the put categories name value params
func (o *PutCategoriesNameValueParams) WithRequest(request *models.CategoryValueIntentInput) *PutCategoriesNameValueParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the put categories name value params
func (o *PutCategoriesNameValueParams) SetRequest(request *models.CategoryValueIntentInput) {
	o.Request = request
}

// WithValue adds the value to the put categories name value params
func (o *PutCategoriesNameValueParams) WithValue(value string) *PutCategoriesNameValueParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the put categories name value params
func (o *PutCategoriesNameValueParams) SetValue(value string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *PutCategoriesNameValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Request == nil {
		o.Request = new(models.CategoryValueIntentInput)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	// path param value
	if err := r.SetPathParam("value", o.Value); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}