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

	"github.com/intent_api/models"
)

// NewPutClustersUUIDParams creates a new PutClustersUUIDParams object
// with the default values initialized.
func NewPutClustersUUIDParams() *PutClustersUUIDParams {
	var ()
	return &PutClustersUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClustersUUIDParamsWithTimeout creates a new PutClustersUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClustersUUIDParamsWithTimeout(timeout time.Duration) *PutClustersUUIDParams {
	var ()
	return &PutClustersUUIDParams{

		timeout: timeout,
	}
}

// NewPutClustersUUIDParamsWithContext creates a new PutClustersUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutClustersUUIDParamsWithContext(ctx context.Context) *PutClustersUUIDParams {
	var ()
	return &PutClustersUUIDParams{

		Context: ctx,
	}
}

/*PutClustersUUIDParams contains all the parameters to send to the API endpoint
for the put clusters UUID operation typically these are written to a http.Request
*/
type PutClustersUUIDParams struct {

	/*Spec*/
	Spec *models.ClusterIntentInput
	/*UUID*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put clusters UUID params
func (o *PutClustersUUIDParams) WithTimeout(timeout time.Duration) *PutClustersUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put clusters UUID params
func (o *PutClustersUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put clusters UUID params
func (o *PutClustersUUIDParams) WithContext(ctx context.Context) *PutClustersUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put clusters UUID params
func (o *PutClustersUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithSpec adds the spec to the put clusters UUID params
func (o *PutClustersUUIDParams) WithSpec(spec *models.ClusterIntentInput) *PutClustersUUIDParams {
	o.SetSpec(spec)
	return o
}

// SetSpec adds the spec to the put clusters UUID params
func (o *PutClustersUUIDParams) SetSpec(spec *models.ClusterIntentInput) {
	o.Spec = spec
}

// WithUUID adds the uuid to the put clusters UUID params
func (o *PutClustersUUIDParams) WithUUID(uuid string) *PutClustersUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the put clusters UUID params
func (o *PutClustersUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PutClustersUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Spec == nil {
		o.Spec = new(models.ClusterIntentInput)
	}

	if err := r.SetBodyParam(o.Spec); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
