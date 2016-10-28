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

// NewPutClustersUUIDCertificatesClientAuthParams creates a new PutClustersUUIDCertificatesClientAuthParams object
// with the default values initialized.
func NewPutClustersUUIDCertificatesClientAuthParams() *PutClustersUUIDCertificatesClientAuthParams {
	var ()
	return &PutClustersUUIDCertificatesClientAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClustersUUIDCertificatesClientAuthParamsWithTimeout creates a new PutClustersUUIDCertificatesClientAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClustersUUIDCertificatesClientAuthParamsWithTimeout(timeout time.Duration) *PutClustersUUIDCertificatesClientAuthParams {
	var ()
	return &PutClustersUUIDCertificatesClientAuthParams{

		timeout: timeout,
	}
}

// NewPutClustersUUIDCertificatesClientAuthParamsWithContext creates a new PutClustersUUIDCertificatesClientAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutClustersUUIDCertificatesClientAuthParamsWithContext(ctx context.Context) *PutClustersUUIDCertificatesClientAuthParams {
	var ()
	return &PutClustersUUIDCertificatesClientAuthParams{

		Context: ctx,
	}
}

/*PutClustersUUIDCertificatesClientAuthParams contains all the parameters to send to the API endpoint
for the put clusters UUID certificates client auth operation typically these are written to a http.Request
*/
type PutClustersUUIDCertificatesClientAuthParams struct {

	/*Spec*/
	Spec *models.CaChainSpec
	/*UUID*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put clusters UUID certificates client auth params
func (o *PutClustersUUIDCertificatesClientAuthParams) WithTimeout(timeout time.Duration) *PutClustersUUIDCertificatesClientAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put clusters UUID certificates client auth params
func (o *PutClustersUUIDCertificatesClientAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put clusters UUID certificates client auth params
func (o *PutClustersUUIDCertificatesClientAuthParams) WithContext(ctx context.Context) *PutClustersUUIDCertificatesClientAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put clusters UUID certificates client auth params
func (o *PutClustersUUIDCertificatesClientAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithSpec adds the spec to the put clusters UUID certificates client auth params
func (o *PutClustersUUIDCertificatesClientAuthParams) WithSpec(spec *models.CaChainSpec) *PutClustersUUIDCertificatesClientAuthParams {
	o.SetSpec(spec)
	return o
}

// SetSpec adds the spec to the put clusters UUID certificates client auth params
func (o *PutClustersUUIDCertificatesClientAuthParams) SetSpec(spec *models.CaChainSpec) {
	o.Spec = spec
}

// WithUUID adds the uuid to the put clusters UUID certificates client auth params
func (o *PutClustersUUIDCertificatesClientAuthParams) WithUUID(uuid string) *PutClustersUUIDCertificatesClientAuthParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the put clusters UUID certificates client auth params
func (o *PutClustersUUIDCertificatesClientAuthParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PutClustersUUIDCertificatesClientAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Spec == nil {
		o.Spec = new(models.CaChainSpec)
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
