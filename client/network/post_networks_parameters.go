package network

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

// NewPostNetworksParams creates a new PostNetworksParams object
// with the default values initialized.
func NewPostNetworksParams() *PostNetworksParams {
	var ()
	return &PostNetworksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNetworksParamsWithTimeout creates a new PostNetworksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNetworksParamsWithTimeout(timeout time.Duration) *PostNetworksParams {
	var ()
	return &PostNetworksParams{

		timeout: timeout,
	}
}

// NewPostNetworksParamsWithContext creates a new PostNetworksParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNetworksParamsWithContext(ctx context.Context) *PostNetworksParams {
	var ()
	return &PostNetworksParams{

		Context: ctx,
	}
}

/*PostNetworksParams contains all the parameters to send to the API endpoint
for the post networks operation typically these are written to a http.Request
*/
type PostNetworksParams struct {

	/*Request*/
	Request *models.Network

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post networks params
func (o *PostNetworksParams) WithTimeout(timeout time.Duration) *PostNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post networks params
func (o *PostNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post networks params
func (o *PostNetworksParams) WithContext(ctx context.Context) *PostNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post networks params
func (o *PostNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithRequest adds the request to the post networks params
func (o *PostNetworksParams) WithRequest(request *models.Network) *PostNetworksParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post networks params
func (o *PostNetworksParams) SetRequest(request *models.Network) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Request == nil {
		o.Request = new(models.Network)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
