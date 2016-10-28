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

// NewPostClustersUUIDCertificatesCaCertsParams creates a new PostClustersUUIDCertificatesCaCertsParams object
// with the default values initialized.
func NewPostClustersUUIDCertificatesCaCertsParams() *PostClustersUUIDCertificatesCaCertsParams {
	var ()
	return &PostClustersUUIDCertificatesCaCertsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClustersUUIDCertificatesCaCertsParamsWithTimeout creates a new PostClustersUUIDCertificatesCaCertsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClustersUUIDCertificatesCaCertsParamsWithTimeout(timeout time.Duration) *PostClustersUUIDCertificatesCaCertsParams {
	var ()
	return &PostClustersUUIDCertificatesCaCertsParams{

		timeout: timeout,
	}
}

// NewPostClustersUUIDCertificatesCaCertsParamsWithContext creates a new PostClustersUUIDCertificatesCaCertsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClustersUUIDCertificatesCaCertsParamsWithContext(ctx context.Context) *PostClustersUUIDCertificatesCaCertsParams {
	var ()
	return &PostClustersUUIDCertificatesCaCertsParams{

		Context: ctx,
	}
}

/*PostClustersUUIDCertificatesCaCertsParams contains all the parameters to send to the API endpoint
for the post clusters UUID certificates ca certs operation typically these are written to a http.Request
*/
type PostClustersUUIDCertificatesCaCertsParams struct {

	/*Spec*/
	Spec *models.CaCert
	/*UUID*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post clusters UUID certificates ca certs params
func (o *PostClustersUUIDCertificatesCaCertsParams) WithTimeout(timeout time.Duration) *PostClustersUUIDCertificatesCaCertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post clusters UUID certificates ca certs params
func (o *PostClustersUUIDCertificatesCaCertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post clusters UUID certificates ca certs params
func (o *PostClustersUUIDCertificatesCaCertsParams) WithContext(ctx context.Context) *PostClustersUUIDCertificatesCaCertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post clusters UUID certificates ca certs params
func (o *PostClustersUUIDCertificatesCaCertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithSpec adds the spec to the post clusters UUID certificates ca certs params
func (o *PostClustersUUIDCertificatesCaCertsParams) WithSpec(spec *models.CaCert) *PostClustersUUIDCertificatesCaCertsParams {
	o.SetSpec(spec)
	return o
}

// SetSpec adds the spec to the post clusters UUID certificates ca certs params
func (o *PostClustersUUIDCertificatesCaCertsParams) SetSpec(spec *models.CaCert) {
	o.Spec = spec
}

// WithUUID adds the uuid to the post clusters UUID certificates ca certs params
func (o *PostClustersUUIDCertificatesCaCertsParams) WithUUID(uuid string) *PostClustersUUIDCertificatesCaCertsParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the post clusters UUID certificates ca certs params
func (o *PostClustersUUIDCertificatesCaCertsParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PostClustersUUIDCertificatesCaCertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Spec == nil {
		o.Spec = new(models.CaCert)
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
