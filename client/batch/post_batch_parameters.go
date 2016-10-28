package batch

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

// NewPostBatchParams creates a new PostBatchParams object
// with the default values initialized.
func NewPostBatchParams() *PostBatchParams {
	var ()
	return &PostBatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostBatchParamsWithTimeout creates a new PostBatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostBatchParamsWithTimeout(timeout time.Duration) *PostBatchParams {
	var ()
	return &PostBatchParams{

		timeout: timeout,
	}
}

// NewPostBatchParamsWithContext creates a new PostBatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostBatchParamsWithContext(ctx context.Context) *PostBatchParams {
	var ()
	return &PostBatchParams{

		Context: ctx,
	}
}

/*PostBatchParams contains all the parameters to send to the API endpoint
for the post batch operation typically these are written to a http.Request
*/
type PostBatchParams struct {

	/*IntentList
	  List of intent APIs

	*/
	IntentList *models.BatchRequest

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post batch params
func (o *PostBatchParams) WithTimeout(timeout time.Duration) *PostBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post batch params
func (o *PostBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post batch params
func (o *PostBatchParams) WithContext(ctx context.Context) *PostBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post batch params
func (o *PostBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIntentList adds the intentList to the post batch params
func (o *PostBatchParams) WithIntentList(intentList *models.BatchRequest) *PostBatchParams {
	o.SetIntentList(intentList)
	return o
}

// SetIntentList adds the intentList to the post batch params
func (o *PostBatchParams) SetIntentList(intentList *models.BatchRequest) {
	o.IntentList = intentList
}

// WriteToRequest writes these params to a swagger request
func (o *PostBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.IntentList == nil {
		o.IntentList = new(models.BatchRequest)
	}

	if err := r.SetBodyParam(o.IntentList); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
