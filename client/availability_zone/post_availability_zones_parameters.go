package availability_zone

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

// NewPostAvailabilityZonesParams creates a new PostAvailabilityZonesParams object
// with the default values initialized.
func NewPostAvailabilityZonesParams() *PostAvailabilityZonesParams {
	var ()
	return &PostAvailabilityZonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAvailabilityZonesParamsWithTimeout creates a new PostAvailabilityZonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAvailabilityZonesParamsWithTimeout(timeout time.Duration) *PostAvailabilityZonesParams {
	var ()
	return &PostAvailabilityZonesParams{

		timeout: timeout,
	}
}

// NewPostAvailabilityZonesParamsWithContext creates a new PostAvailabilityZonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAvailabilityZonesParamsWithContext(ctx context.Context) *PostAvailabilityZonesParams {
	var ()
	return &PostAvailabilityZonesParams{

		Context: ctx,
	}
}

/*PostAvailabilityZonesParams contains all the parameters to send to the API endpoint
for the post availability zones operation typically these are written to a http.Request
*/
type PostAvailabilityZonesParams struct {

	/*Request*/
	Request *models.AvailabilityZone

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post availability zones params
func (o *PostAvailabilityZonesParams) WithTimeout(timeout time.Duration) *PostAvailabilityZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post availability zones params
func (o *PostAvailabilityZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post availability zones params
func (o *PostAvailabilityZonesParams) WithContext(ctx context.Context) *PostAvailabilityZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post availability zones params
func (o *PostAvailabilityZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithRequest adds the request to the post availability zones params
func (o *PostAvailabilityZonesParams) WithRequest(request *models.AvailabilityZone) *PostAvailabilityZonesParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post availability zones params
func (o *PostAvailabilityZonesParams) SetRequest(request *models.AvailabilityZone) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostAvailabilityZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Request == nil {
		o.Request = new(models.AvailabilityZone)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
