package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetImagesUUIDParams creates a new GetImagesUUIDParams object
// with the default values initialized.
func NewGetImagesUUIDParams() *GetImagesUUIDParams {
	var ()
	return &GetImagesUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImagesUUIDParamsWithTimeout creates a new GetImagesUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImagesUUIDParamsWithTimeout(timeout time.Duration) *GetImagesUUIDParams {
	var ()
	return &GetImagesUUIDParams{

		timeout: timeout,
	}
}

// NewGetImagesUUIDParamsWithContext creates a new GetImagesUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImagesUUIDParamsWithContext(ctx context.Context) *GetImagesUUIDParams {
	var ()
	return &GetImagesUUIDParams{

		Context: ctx,
	}
}

/*GetImagesUUIDParams contains all the parameters to send to the API endpoint
for the get images UUID operation typically these are written to a http.Request
*/
type GetImagesUUIDParams struct {

	/*IncludeVmdiskSizes*/
	IncludeVmdiskSizes *bool
	/*UUID
	  The UUID of the entity.

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get images UUID params
func (o *GetImagesUUIDParams) WithTimeout(timeout time.Duration) *GetImagesUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get images UUID params
func (o *GetImagesUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get images UUID params
func (o *GetImagesUUIDParams) WithContext(ctx context.Context) *GetImagesUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get images UUID params
func (o *GetImagesUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIncludeVmdiskSizes adds the includeVmdiskSizes to the get images UUID params
func (o *GetImagesUUIDParams) WithIncludeVmdiskSizes(includeVmdiskSizes *bool) *GetImagesUUIDParams {
	o.SetIncludeVmdiskSizes(includeVmdiskSizes)
	return o
}

// SetIncludeVmdiskSizes adds the includeVmdiskSizes to the get images UUID params
func (o *GetImagesUUIDParams) SetIncludeVmdiskSizes(includeVmdiskSizes *bool) {
	o.IncludeVmdiskSizes = includeVmdiskSizes
}

// WithUUID adds the uuid to the get images UUID params
func (o *GetImagesUUIDParams) WithUUID(uuid string) *GetImagesUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get images UUID params
func (o *GetImagesUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetImagesUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.IncludeVmdiskSizes != nil {

		// query param include_vmdisk_sizes
		var qrIncludeVmdiskSizes bool
		if o.IncludeVmdiskSizes != nil {
			qrIncludeVmdiskSizes = *o.IncludeVmdiskSizes
		}
		qIncludeVmdiskSizes := swag.FormatBool(qrIncludeVmdiskSizes)
		if qIncludeVmdiskSizes != "" {
			if err := r.SetQueryParam("include_vmdisk_sizes", qIncludeVmdiskSizes); err != nil {
				return err
			}
		}

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
