package docker_registry

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

// NewPostDockerRegistriesParams creates a new PostDockerRegistriesParams object
// with the default values initialized.
func NewPostDockerRegistriesParams() *PostDockerRegistriesParams {
	var ()
	return &PostDockerRegistriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDockerRegistriesParamsWithTimeout creates a new PostDockerRegistriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDockerRegistriesParamsWithTimeout(timeout time.Duration) *PostDockerRegistriesParams {
	var ()
	return &PostDockerRegistriesParams{

		timeout: timeout,
	}
}

// NewPostDockerRegistriesParamsWithContext creates a new PostDockerRegistriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDockerRegistriesParamsWithContext(ctx context.Context) *PostDockerRegistriesParams {
	var ()
	return &PostDockerRegistriesParams{

		Context: ctx,
	}
}

/*PostDockerRegistriesParams contains all the parameters to send to the API endpoint
for the post docker registries operation typically these are written to a http.Request
*/
type PostDockerRegistriesParams struct {

	/*Spec
	  Docker registry spec

	*/
	Spec *models.DockerRegistryIntentInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post docker registries params
func (o *PostDockerRegistriesParams) WithTimeout(timeout time.Duration) *PostDockerRegistriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post docker registries params
func (o *PostDockerRegistriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post docker registries params
func (o *PostDockerRegistriesParams) WithContext(ctx context.Context) *PostDockerRegistriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post docker registries params
func (o *PostDockerRegistriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithSpec adds the spec to the post docker registries params
func (o *PostDockerRegistriesParams) WithSpec(spec *models.DockerRegistryIntentInput) *PostDockerRegistriesParams {
	o.SetSpec(spec)
	return o
}

// SetSpec adds the spec to the post docker registries params
func (o *PostDockerRegistriesParams) SetSpec(spec *models.DockerRegistryIntentInput) {
	o.Spec = spec
}

// WriteToRequest writes these params to a swagger request
func (o *PostDockerRegistriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Spec == nil {
		o.Spec = new(models.DockerRegistryIntentInput)
	}

	if err := r.SetBodyParam(o.Spec); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
