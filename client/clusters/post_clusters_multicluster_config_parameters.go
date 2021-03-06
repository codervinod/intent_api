package clusters

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

// NewPostClustersMulticlusterConfigParams creates a new PostClustersMulticlusterConfigParams object
// with the default values initialized.
func NewPostClustersMulticlusterConfigParams() *PostClustersMulticlusterConfigParams {
	var ()
	return &PostClustersMulticlusterConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClustersMulticlusterConfigParamsWithTimeout creates a new PostClustersMulticlusterConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClustersMulticlusterConfigParamsWithTimeout(timeout time.Duration) *PostClustersMulticlusterConfigParams {
	var ()
	return &PostClustersMulticlusterConfigParams{

		timeout: timeout,
	}
}

// NewPostClustersMulticlusterConfigParamsWithContext creates a new PostClustersMulticlusterConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClustersMulticlusterConfigParamsWithContext(ctx context.Context) *PostClustersMulticlusterConfigParams {
	var ()
	return &PostClustersMulticlusterConfigParams{

		Context: ctx,
	}
}

/*PostClustersMulticlusterConfigParams contains all the parameters to send to the API endpoint
for the post clusters multicluster config operation typically these are written to a http.Request
*/
type PostClustersMulticlusterConfigParams struct {

	/*Spec*/
	Spec *models.MulticlusterConfigSpec

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post clusters multicluster config params
func (o *PostClustersMulticlusterConfigParams) WithTimeout(timeout time.Duration) *PostClustersMulticlusterConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post clusters multicluster config params
func (o *PostClustersMulticlusterConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post clusters multicluster config params
func (o *PostClustersMulticlusterConfigParams) WithContext(ctx context.Context) *PostClustersMulticlusterConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post clusters multicluster config params
func (o *PostClustersMulticlusterConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithSpec adds the spec to the post clusters multicluster config params
func (o *PostClustersMulticlusterConfigParams) WithSpec(spec *models.MulticlusterConfigSpec) *PostClustersMulticlusterConfigParams {
	o.SetSpec(spec)
	return o
}

// SetSpec adds the spec to the post clusters multicluster config params
func (o *PostClustersMulticlusterConfigParams) SetSpec(spec *models.MulticlusterConfigSpec) {
	o.Spec = spec
}

// WriteToRequest writes these params to a swagger request
func (o *PostClustersMulticlusterConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Spec == nil {
		o.Spec = new(models.MulticlusterConfigSpec)
	}

	if err := r.SetBodyParam(o.Spec); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
