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

// NewPostClustersUUIDDirectoryServicesParams creates a new PostClustersUUIDDirectoryServicesParams object
// with the default values initialized.
func NewPostClustersUUIDDirectoryServicesParams() *PostClustersUUIDDirectoryServicesParams {
	var ()
	return &PostClustersUUIDDirectoryServicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClustersUUIDDirectoryServicesParamsWithTimeout creates a new PostClustersUUIDDirectoryServicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClustersUUIDDirectoryServicesParamsWithTimeout(timeout time.Duration) *PostClustersUUIDDirectoryServicesParams {
	var ()
	return &PostClustersUUIDDirectoryServicesParams{

		timeout: timeout,
	}
}

// NewPostClustersUUIDDirectoryServicesParamsWithContext creates a new PostClustersUUIDDirectoryServicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClustersUUIDDirectoryServicesParamsWithContext(ctx context.Context) *PostClustersUUIDDirectoryServicesParams {
	var ()
	return &PostClustersUUIDDirectoryServicesParams{

		Context: ctx,
	}
}

/*PostClustersUUIDDirectoryServicesParams contains all the parameters to send to the API endpoint
for the post clusters UUID directory services operation typically these are written to a http.Request
*/
type PostClustersUUIDDirectoryServicesParams struct {

	/*Body
	  Directory Service object

	*/
	Body *models.DirectoryServiceIntentInput
	/*UUID
	  The UUID of the entity.

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post clusters UUID directory services params
func (o *PostClustersUUIDDirectoryServicesParams) WithTimeout(timeout time.Duration) *PostClustersUUIDDirectoryServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post clusters UUID directory services params
func (o *PostClustersUUIDDirectoryServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post clusters UUID directory services params
func (o *PostClustersUUIDDirectoryServicesParams) WithContext(ctx context.Context) *PostClustersUUIDDirectoryServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post clusters UUID directory services params
func (o *PostClustersUUIDDirectoryServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post clusters UUID directory services params
func (o *PostClustersUUIDDirectoryServicesParams) WithBody(body *models.DirectoryServiceIntentInput) *PostClustersUUIDDirectoryServicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post clusters UUID directory services params
func (o *PostClustersUUIDDirectoryServicesParams) SetBody(body *models.DirectoryServiceIntentInput) {
	o.Body = body
}

// WithUUID adds the uuid to the post clusters UUID directory services params
func (o *PostClustersUUIDDirectoryServicesParams) WithUUID(uuid string) *PostClustersUUIDDirectoryServicesParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the post clusters UUID directory services params
func (o *PostClustersUUIDDirectoryServicesParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PostClustersUUIDDirectoryServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DirectoryServiceIntentInput)
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
