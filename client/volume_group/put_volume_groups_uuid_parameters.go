package volume_group

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

// NewPutVolumeGroupsUUIDParams creates a new PutVolumeGroupsUUIDParams object
// with the default values initialized.
func NewPutVolumeGroupsUUIDParams() *PutVolumeGroupsUUIDParams {
	var ()
	return &PutVolumeGroupsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutVolumeGroupsUUIDParamsWithTimeout creates a new PutVolumeGroupsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutVolumeGroupsUUIDParamsWithTimeout(timeout time.Duration) *PutVolumeGroupsUUIDParams {
	var ()
	return &PutVolumeGroupsUUIDParams{

		timeout: timeout,
	}
}

// NewPutVolumeGroupsUUIDParamsWithContext creates a new PutVolumeGroupsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutVolumeGroupsUUIDParamsWithContext(ctx context.Context) *PutVolumeGroupsUUIDParams {
	var ()
	return &PutVolumeGroupsUUIDParams{

		Context: ctx,
	}
}

/*PutVolumeGroupsUUIDParams contains all the parameters to send to the API endpoint
for the put volume groups UUID operation typically these are written to a http.Request
*/
type PutVolumeGroupsUUIDParams struct {

	/*Body
	  Volume group object

	*/
	Body *models.VolumeGroup
	/*UUID
	  The UUID of the entity.

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put volume groups UUID params
func (o *PutVolumeGroupsUUIDParams) WithTimeout(timeout time.Duration) *PutVolumeGroupsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put volume groups UUID params
func (o *PutVolumeGroupsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put volume groups UUID params
func (o *PutVolumeGroupsUUIDParams) WithContext(ctx context.Context) *PutVolumeGroupsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put volume groups UUID params
func (o *PutVolumeGroupsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the put volume groups UUID params
func (o *PutVolumeGroupsUUIDParams) WithBody(body *models.VolumeGroup) *PutVolumeGroupsUUIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put volume groups UUID params
func (o *PutVolumeGroupsUUIDParams) SetBody(body *models.VolumeGroup) {
	o.Body = body
}

// WithUUID adds the uuid to the put volume groups UUID params
func (o *PutVolumeGroupsUUIDParams) WithUUID(uuid string) *PutVolumeGroupsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the put volume groups UUID params
func (o *PutVolumeGroupsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PutVolumeGroupsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.VolumeGroup)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
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
