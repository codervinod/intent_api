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

// NewPostClustersUUIDCertificatesSvmCertsKmsUUIDParams creates a new PostClustersUUIDCertificatesSvmCertsKmsUUIDParams object
// with the default values initialized.
func NewPostClustersUUIDCertificatesSvmCertsKmsUUIDParams() *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams {
	var ()
	return &PostClustersUUIDCertificatesSvmCertsKmsUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClustersUUIDCertificatesSvmCertsKmsUUIDParamsWithTimeout creates a new PostClustersUUIDCertificatesSvmCertsKmsUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClustersUUIDCertificatesSvmCertsKmsUUIDParamsWithTimeout(timeout time.Duration) *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams {
	var ()
	return &PostClustersUUIDCertificatesSvmCertsKmsUUIDParams{

		timeout: timeout,
	}
}

// NewPostClustersUUIDCertificatesSvmCertsKmsUUIDParamsWithContext creates a new PostClustersUUIDCertificatesSvmCertsKmsUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClustersUUIDCertificatesSvmCertsKmsUUIDParamsWithContext(ctx context.Context) *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams {
	var ()
	return &PostClustersUUIDCertificatesSvmCertsKmsUUIDParams{

		Context: ctx,
	}
}

/*PostClustersUUIDCertificatesSvmCertsKmsUUIDParams contains all the parameters to send to the API endpoint
for the post clusters UUID certificates svm certs kms UUID operation typically these are written to a http.Request
*/
type PostClustersUUIDCertificatesSvmCertsKmsUUIDParams struct {

	/*KmsUUID*/
	KmsUUID string
	/*Spec*/
	Spec *models.CertificateSpecList
	/*UUID*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post clusters UUID certificates svm certs kms UUID params
func (o *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams) WithTimeout(timeout time.Duration) *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post clusters UUID certificates svm certs kms UUID params
func (o *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post clusters UUID certificates svm certs kms UUID params
func (o *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams) WithContext(ctx context.Context) *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post clusters UUID certificates svm certs kms UUID params
func (o *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithKmsUUID adds the kmsUUID to the post clusters UUID certificates svm certs kms UUID params
func (o *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams) WithKmsUUID(kmsUUID string) *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams {
	o.SetKmsUUID(kmsUUID)
	return o
}

// SetKmsUUID adds the kmsUuid to the post clusters UUID certificates svm certs kms UUID params
func (o *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams) SetKmsUUID(kmsUUID string) {
	o.KmsUUID = kmsUUID
}

// WithSpec adds the spec to the post clusters UUID certificates svm certs kms UUID params
func (o *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams) WithSpec(spec *models.CertificateSpecList) *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams {
	o.SetSpec(spec)
	return o
}

// SetSpec adds the spec to the post clusters UUID certificates svm certs kms UUID params
func (o *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams) SetSpec(spec *models.CertificateSpecList) {
	o.Spec = spec
}

// WithUUID adds the uuid to the post clusters UUID certificates svm certs kms UUID params
func (o *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams) WithUUID(uuid string) *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the post clusters UUID certificates svm certs kms UUID params
func (o *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PostClustersUUIDCertificatesSvmCertsKmsUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param kms_uuid
	if err := r.SetPathParam("kms_uuid", o.KmsUUID); err != nil {
		return err
	}

	if o.Spec == nil {
		o.Spec = new(models.CertificateSpecList)
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
