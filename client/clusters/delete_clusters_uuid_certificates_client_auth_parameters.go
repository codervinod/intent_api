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

// NewDeleteClustersUUIDCertificatesClientAuthParams creates a new DeleteClustersUUIDCertificatesClientAuthParams object
// with the default values initialized.
func NewDeleteClustersUUIDCertificatesClientAuthParams() *DeleteClustersUUIDCertificatesClientAuthParams {
	var ()
	return &DeleteClustersUUIDCertificatesClientAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClustersUUIDCertificatesClientAuthParamsWithTimeout creates a new DeleteClustersUUIDCertificatesClientAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClustersUUIDCertificatesClientAuthParamsWithTimeout(timeout time.Duration) *DeleteClustersUUIDCertificatesClientAuthParams {
	var ()
	return &DeleteClustersUUIDCertificatesClientAuthParams{

		timeout: timeout,
	}
}

// NewDeleteClustersUUIDCertificatesClientAuthParamsWithContext creates a new DeleteClustersUUIDCertificatesClientAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClustersUUIDCertificatesClientAuthParamsWithContext(ctx context.Context) *DeleteClustersUUIDCertificatesClientAuthParams {
	var ()
	return &DeleteClustersUUIDCertificatesClientAuthParams{

		Context: ctx,
	}
}

/*DeleteClustersUUIDCertificatesClientAuthParams contains all the parameters to send to the API endpoint
for the delete clusters UUID certificates client auth operation typically these are written to a http.Request
*/
type DeleteClustersUUIDCertificatesClientAuthParams struct {

	/*UUID*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the delete clusters UUID certificates client auth params
func (o *DeleteClustersUUIDCertificatesClientAuthParams) WithTimeout(timeout time.Duration) *DeleteClustersUUIDCertificatesClientAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete clusters UUID certificates client auth params
func (o *DeleteClustersUUIDCertificatesClientAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete clusters UUID certificates client auth params
func (o *DeleteClustersUUIDCertificatesClientAuthParams) WithContext(ctx context.Context) *DeleteClustersUUIDCertificatesClientAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete clusters UUID certificates client auth params
func (o *DeleteClustersUUIDCertificatesClientAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithUUID adds the uuid to the delete clusters UUID certificates client auth params
func (o *DeleteClustersUUIDCertificatesClientAuthParams) WithUUID(uuid string) *DeleteClustersUUIDCertificatesClientAuthParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete clusters UUID certificates client auth params
func (o *DeleteClustersUUIDCertificatesClientAuthParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClustersUUIDCertificatesClientAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
