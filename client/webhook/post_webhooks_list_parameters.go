package webhook

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

// NewPostWebhooksListParams creates a new PostWebhooksListParams object
// with the default values initialized.
func NewPostWebhooksListParams() *PostWebhooksListParams {
	var ()
	return &PostWebhooksListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWebhooksListParamsWithTimeout creates a new PostWebhooksListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWebhooksListParamsWithTimeout(timeout time.Duration) *PostWebhooksListParams {
	var ()
	return &PostWebhooksListParams{

		timeout: timeout,
	}
}

// NewPostWebhooksListParamsWithContext creates a new PostWebhooksListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWebhooksListParamsWithContext(ctx context.Context) *PostWebhooksListParams {
	var ()
	return &PostWebhooksListParams{

		Context: ctx,
	}
}

/*PostWebhooksListParams contains all the parameters to send to the API endpoint
for the post webhooks list operation typically these are written to a http.Request
*/
type PostWebhooksListParams struct {

	/*GetEntitiesRequest*/
	GetEntitiesRequest *models.WebhookListMetadata

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post webhooks list params
func (o *PostWebhooksListParams) WithTimeout(timeout time.Duration) *PostWebhooksListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post webhooks list params
func (o *PostWebhooksListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post webhooks list params
func (o *PostWebhooksListParams) WithContext(ctx context.Context) *PostWebhooksListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post webhooks list params
func (o *PostWebhooksListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithGetEntitiesRequest adds the getEntitiesRequest to the post webhooks list params
func (o *PostWebhooksListParams) WithGetEntitiesRequest(getEntitiesRequest *models.WebhookListMetadata) *PostWebhooksListParams {
	o.SetGetEntitiesRequest(getEntitiesRequest)
	return o
}

// SetGetEntitiesRequest adds the getEntitiesRequest to the post webhooks list params
func (o *PostWebhooksListParams) SetGetEntitiesRequest(getEntitiesRequest *models.WebhookListMetadata) {
	o.GetEntitiesRequest = getEntitiesRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostWebhooksListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.GetEntitiesRequest == nil {
		o.GetEntitiesRequest = new(models.WebhookListMetadata)
	}

	if err := r.SetBodyParam(o.GetEntitiesRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
