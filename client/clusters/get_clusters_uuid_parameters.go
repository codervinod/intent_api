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
)

// NewGetClustersUUIDParams creates a new GetClustersUUIDParams object
// with the default values initialized.
func NewGetClustersUUIDParams() *GetClustersUUIDParams {
	var ()
	return &GetClustersUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClustersUUIDParamsWithTimeout creates a new GetClustersUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClustersUUIDParamsWithTimeout(timeout time.Duration) *GetClustersUUIDParams {
	var ()
	return &GetClustersUUIDParams{

		timeout: timeout,
	}
}

// NewGetClustersUUIDParamsWithContext creates a new GetClustersUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClustersUUIDParamsWithContext(ctx context.Context) *GetClustersUUIDParams {
	var ()
	return &GetClustersUUIDParams{

		Context: ctx,
	}
}

/*GetClustersUUIDParams contains all the parameters to send to the API endpoint
for the get clusters UUID operation typically these are written to a http.Request
*/
type GetClustersUUIDParams struct {

	/*UUID*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get clusters UUID params
func (o *GetClustersUUIDParams) WithTimeout(timeout time.Duration) *GetClustersUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get clusters UUID params
func (o *GetClustersUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get clusters UUID params
func (o *GetClustersUUIDParams) WithContext(ctx context.Context) *GetClustersUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get clusters UUID params
func (o *GetClustersUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithUUID adds the uuid to the get clusters UUID params
func (o *GetClustersUUIDParams) WithUUID(uuid string) *GetClustersUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get clusters UUID params
func (o *GetClustersUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetClustersUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
