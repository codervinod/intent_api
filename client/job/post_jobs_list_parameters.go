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

	"github.com/intent_api/models"
)

// NewPostJobsListParams creates a new PostJobsListParams object
// with the default values initialized.
func NewPostJobsListParams() *PostJobsListParams {
	var ()
	return &PostJobsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostJobsListParamsWithTimeout creates a new PostJobsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostJobsListParamsWithTimeout(timeout time.Duration) *PostJobsListParams {
	var ()
	return &PostJobsListParams{

		timeout: timeout,
	}
}

// NewPostJobsListParamsWithContext creates a new PostJobsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostJobsListParamsWithContext(ctx context.Context) *PostJobsListParams {
	var ()
	return &PostJobsListParams{

		Context: ctx,
	}
}

/*PostJobsListParams contains all the parameters to send to the API endpoint
for the post jobs list operation typically these are written to a http.Request
*/
type PostJobsListParams struct {

	/*GetEntitiesRequest*/
	GetEntitiesRequest *models.JobListMetadata

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post jobs list params
func (o *PostJobsListParams) WithTimeout(timeout time.Duration) *PostJobsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post jobs list params
func (o *PostJobsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post jobs list params
func (o *PostJobsListParams) WithContext(ctx context.Context) *PostJobsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post jobs list params
func (o *PostJobsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithGetEntitiesRequest adds the getEntitiesRequest to the post jobs list params
func (o *PostJobsListParams) WithGetEntitiesRequest(getEntitiesRequest *models.JobListMetadata) *PostJobsListParams {
	o.SetGetEntitiesRequest(getEntitiesRequest)
	return o
}

// SetGetEntitiesRequest adds the getEntitiesRequest to the post jobs list params
func (o *PostJobsListParams) SetGetEntitiesRequest(getEntitiesRequest *models.JobListMetadata) {
	o.GetEntitiesRequest = getEntitiesRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostJobsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.GetEntitiesRequest == nil {
		o.GetEntitiesRequest = new(models.JobListMetadata)
	}

	if err := r.SetBodyParam(o.GetEntitiesRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
