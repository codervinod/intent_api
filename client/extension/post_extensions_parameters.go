package extension

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

// NewPostExtensionsParams creates a new PostExtensionsParams object
// with the default values initialized.
func NewPostExtensionsParams() *PostExtensionsParams {
	var ()
	return &PostExtensionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostExtensionsParamsWithTimeout creates a new PostExtensionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostExtensionsParamsWithTimeout(timeout time.Duration) *PostExtensionsParams {
	var ()
	return &PostExtensionsParams{

		timeout: timeout,
	}
}

// NewPostExtensionsParamsWithContext creates a new PostExtensionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostExtensionsParamsWithContext(ctx context.Context) *PostExtensionsParams {
	var ()
	return &PostExtensionsParams{

		Context: ctx,
	}
}

/*PostExtensionsParams contains all the parameters to send to the API endpoint
for the post extensions operation typically these are written to a http.Request
*/
type PostExtensionsParams struct {

	/*ExtensionManifest
	  Extension details

	*/
	ExtensionManifest *models.ExtensionIntentInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post extensions params
func (o *PostExtensionsParams) WithTimeout(timeout time.Duration) *PostExtensionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post extensions params
func (o *PostExtensionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post extensions params
func (o *PostExtensionsParams) WithContext(ctx context.Context) *PostExtensionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post extensions params
func (o *PostExtensionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithExtensionManifest adds the extensionManifest to the post extensions params
func (o *PostExtensionsParams) WithExtensionManifest(extensionManifest *models.ExtensionIntentInput) *PostExtensionsParams {
	o.SetExtensionManifest(extensionManifest)
	return o
}

// SetExtensionManifest adds the extensionManifest to the post extensions params
func (o *PostExtensionsParams) SetExtensionManifest(extensionManifest *models.ExtensionIntentInput) {
	o.ExtensionManifest = extensionManifest
}

// WriteToRequest writes these params to a swagger request
func (o *PostExtensionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.ExtensionManifest == nil {
		o.ExtensionManifest = new(models.ExtensionIntentInput)
	}

	if err := r.SetBodyParam(o.ExtensionManifest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}