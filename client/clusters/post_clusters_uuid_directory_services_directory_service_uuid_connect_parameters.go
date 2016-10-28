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

// NewPostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams creates a new PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams object
// with the default values initialized.
func NewPostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams() *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams {
	var ()
	return &PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParamsWithTimeout creates a new PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParamsWithTimeout(timeout time.Duration) *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams {
	var ()
	return &PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams{

		timeout: timeout,
	}
}

// NewPostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParamsWithContext creates a new PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParamsWithContext(ctx context.Context) *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams {
	var ()
	return &PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams{

		Context: ctx,
	}
}

/*PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams contains all the parameters to send to the API endpoint
for the post clusters UUID directory services directory service UUID connect operation typically these are written to a http.Request
*/
type PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams struct {

	/*Body
	  Credentials to connect the directory service

	*/
	Body *models.ConnectionInput
	/*DirectoryServiceUUID*/
	DirectoryServiceUUID string
	/*UUID
	  The UUID of the entity.

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post clusters UUID directory services directory service UUID connect params
func (o *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams) WithTimeout(timeout time.Duration) *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post clusters UUID directory services directory service UUID connect params
func (o *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post clusters UUID directory services directory service UUID connect params
func (o *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams) WithContext(ctx context.Context) *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post clusters UUID directory services directory service UUID connect params
func (o *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post clusters UUID directory services directory service UUID connect params
func (o *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams) WithBody(body *models.ConnectionInput) *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post clusters UUID directory services directory service UUID connect params
func (o *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams) SetBody(body *models.ConnectionInput) {
	o.Body = body
}

// WithDirectoryServiceUUID adds the directoryServiceUUID to the post clusters UUID directory services directory service UUID connect params
func (o *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams) WithDirectoryServiceUUID(directoryServiceUUID string) *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams {
	o.SetDirectoryServiceUUID(directoryServiceUUID)
	return o
}

// SetDirectoryServiceUUID adds the directoryServiceUuid to the post clusters UUID directory services directory service UUID connect params
func (o *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams) SetDirectoryServiceUUID(directoryServiceUUID string) {
	o.DirectoryServiceUUID = directoryServiceUUID
}

// WithUUID adds the uuid to the post clusters UUID directory services directory service UUID connect params
func (o *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams) WithUUID(uuid string) *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the post clusters UUID directory services directory service UUID connect params
func (o *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PostClustersUUIDDirectoryServicesDirectoryServiceUUIDConnectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.ConnectionInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param directory_service_uuid
	if err := r.SetPathParam("directory_service_uuid", o.DirectoryServiceUUID); err != nil {
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
