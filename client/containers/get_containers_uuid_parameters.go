package containers

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

// NewGetContainersUUIDParams creates a new GetContainersUUIDParams object
// with the default values initialized.
func NewGetContainersUUIDParams() *GetContainersUUIDParams {
	var ()
	return &GetContainersUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContainersUUIDParamsWithTimeout creates a new GetContainersUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContainersUUIDParamsWithTimeout(timeout time.Duration) *GetContainersUUIDParams {
	var ()
	return &GetContainersUUIDParams{

		timeout: timeout,
	}
}

// NewGetContainersUUIDParamsWithContext creates a new GetContainersUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContainersUUIDParamsWithContext(ctx context.Context) *GetContainersUUIDParams {
	var ()
	return &GetContainersUUIDParams{

		Context: ctx,
	}
}

/*GetContainersUUIDParams contains all the parameters to send to the API endpoint
for the get containers UUID operation typically these are written to a http.Request
*/
type GetContainersUUIDParams struct {

	/*UUID
	  Uuid of the container

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get containers UUID params
func (o *GetContainersUUIDParams) WithTimeout(timeout time.Duration) *GetContainersUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get containers UUID params
func (o *GetContainersUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get containers UUID params
func (o *GetContainersUUIDParams) WithContext(ctx context.Context) *GetContainersUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get containers UUID params
func (o *GetContainersUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithUUID adds the uuid to the get containers UUID params
func (o *GetContainersUUIDParams) WithUUID(uuid string) *GetContainersUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get containers UUID params
func (o *GetContainersUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetContainersUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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