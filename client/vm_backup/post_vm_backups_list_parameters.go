package vm_backup

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

// NewPostVMBackupsListParams creates a new PostVMBackupsListParams object
// with the default values initialized.
func NewPostVMBackupsListParams() *PostVMBackupsListParams {
	var ()
	return &PostVMBackupsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVMBackupsListParamsWithTimeout creates a new PostVMBackupsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVMBackupsListParamsWithTimeout(timeout time.Duration) *PostVMBackupsListParams {
	var ()
	return &PostVMBackupsListParams{

		timeout: timeout,
	}
}

// NewPostVMBackupsListParamsWithContext creates a new PostVMBackupsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVMBackupsListParamsWithContext(ctx context.Context) *PostVMBackupsListParams {
	var ()
	return &PostVMBackupsListParams{

		Context: ctx,
	}
}

/*PostVMBackupsListParams contains all the parameters to send to the API endpoint
for the post VM backups list operation typically these are written to a http.Request
*/
type PostVMBackupsListParams struct {

	/*GetEntitiesRequest*/
	GetEntitiesRequest *models.VMBackupListMetadata

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post VM backups list params
func (o *PostVMBackupsListParams) WithTimeout(timeout time.Duration) *PostVMBackupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post VM backups list params
func (o *PostVMBackupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post VM backups list params
func (o *PostVMBackupsListParams) WithContext(ctx context.Context) *PostVMBackupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post VM backups list params
func (o *PostVMBackupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithGetEntitiesRequest adds the getEntitiesRequest to the post VM backups list params
func (o *PostVMBackupsListParams) WithGetEntitiesRequest(getEntitiesRequest *models.VMBackupListMetadata) *PostVMBackupsListParams {
	o.SetGetEntitiesRequest(getEntitiesRequest)
	return o
}

// SetGetEntitiesRequest adds the getEntitiesRequest to the post VM backups list params
func (o *PostVMBackupsListParams) SetGetEntitiesRequest(getEntitiesRequest *models.VMBackupListMetadata) {
	o.GetEntitiesRequest = getEntitiesRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostVMBackupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.GetEntitiesRequest == nil {
		o.GetEntitiesRequest = new(models.VMBackupListMetadata)
	}

	if err := r.SetBodyParam(o.GetEntitiesRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
