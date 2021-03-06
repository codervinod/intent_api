package clusters

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

// NewGetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams creates a new GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams object
// with the default values initialized.
func NewGetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams() *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams {
	var ()
	return &GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParamsWithTimeout creates a new GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParamsWithTimeout(timeout time.Duration) *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams {
	var ()
	return &GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams{

		timeout: timeout,
	}
}

// NewGetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParamsWithContext creates a new GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParamsWithContext(ctx context.Context) *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams {
	var ()
	return &GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams{

		Context: ctx,
	}
}

/*GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams contains all the parameters to send to the API endpoint
for the get clusters UUID cloud credentials cloud type cloud credentials UUID operation typically these are written to a http.Request
*/
type GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams struct {

	/*CloudCredentialsUUID*/
	CloudCredentialsUUID int64
	/*CloudType*/
	CloudType string
	/*UUID*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get clusters UUID cloud credentials cloud type cloud credentials UUID params
func (o *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams) WithTimeout(timeout time.Duration) *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get clusters UUID cloud credentials cloud type cloud credentials UUID params
func (o *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get clusters UUID cloud credentials cloud type cloud credentials UUID params
func (o *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams) WithContext(ctx context.Context) *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get clusters UUID cloud credentials cloud type cloud credentials UUID params
func (o *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCloudCredentialsUUID adds the cloudCredentialsUUID to the get clusters UUID cloud credentials cloud type cloud credentials UUID params
func (o *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams) WithCloudCredentialsUUID(cloudCredentialsUUID int64) *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams {
	o.SetCloudCredentialsUUID(cloudCredentialsUUID)
	return o
}

// SetCloudCredentialsUUID adds the cloudCredentialsUuid to the get clusters UUID cloud credentials cloud type cloud credentials UUID params
func (o *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams) SetCloudCredentialsUUID(cloudCredentialsUUID int64) {
	o.CloudCredentialsUUID = cloudCredentialsUUID
}

// WithCloudType adds the cloudType to the get clusters UUID cloud credentials cloud type cloud credentials UUID params
func (o *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams) WithCloudType(cloudType string) *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the get clusters UUID cloud credentials cloud type cloud credentials UUID params
func (o *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WithUUID adds the uuid to the get clusters UUID cloud credentials cloud type cloud credentials UUID params
func (o *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams) WithUUID(uuid string) *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get clusters UUID cloud credentials cloud type cloud credentials UUID params
func (o *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetClustersUUIDCloudCredentialsCloudTypeCloudCredentialsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param cloud_credentials_uuid
	if err := r.SetPathParam("cloud_credentials_uuid", swag.FormatInt64(o.CloudCredentialsUUID)); err != nil {
		return err
	}

	// path param cloud_type
	if err := r.SetPathParam("cloud_type", o.CloudType); err != nil {
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
