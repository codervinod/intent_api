package volume_group_backup

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

// NewPostVolumeGroupBackupsListParams creates a new PostVolumeGroupBackupsListParams object
// with the default values initialized.
func NewPostVolumeGroupBackupsListParams() *PostVolumeGroupBackupsListParams {
	var ()
	return &PostVolumeGroupBackupsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVolumeGroupBackupsListParamsWithTimeout creates a new PostVolumeGroupBackupsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVolumeGroupBackupsListParamsWithTimeout(timeout time.Duration) *PostVolumeGroupBackupsListParams {
	var ()
	return &PostVolumeGroupBackupsListParams{

		timeout: timeout,
	}
}

// NewPostVolumeGroupBackupsListParamsWithContext creates a new PostVolumeGroupBackupsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVolumeGroupBackupsListParamsWithContext(ctx context.Context) *PostVolumeGroupBackupsListParams {
	var ()
	return &PostVolumeGroupBackupsListParams{

		Context: ctx,
	}
}

/*PostVolumeGroupBackupsListParams contains all the parameters to send to the API endpoint
for the post volume group backups list operation typically these are written to a http.Request
*/
type PostVolumeGroupBackupsListParams struct {

	/*GetEntitiesRequest*/
	GetEntitiesRequest *models.VolumeGroupBackupListMetadata

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post volume group backups list params
func (o *PostVolumeGroupBackupsListParams) WithTimeout(timeout time.Duration) *PostVolumeGroupBackupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post volume group backups list params
func (o *PostVolumeGroupBackupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post volume group backups list params
func (o *PostVolumeGroupBackupsListParams) WithContext(ctx context.Context) *PostVolumeGroupBackupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post volume group backups list params
func (o *PostVolumeGroupBackupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithGetEntitiesRequest adds the getEntitiesRequest to the post volume group backups list params
func (o *PostVolumeGroupBackupsListParams) WithGetEntitiesRequest(getEntitiesRequest *models.VolumeGroupBackupListMetadata) *PostVolumeGroupBackupsListParams {
	o.SetGetEntitiesRequest(getEntitiesRequest)
	return o
}

// SetGetEntitiesRequest adds the getEntitiesRequest to the post volume group backups list params
func (o *PostVolumeGroupBackupsListParams) SetGetEntitiesRequest(getEntitiesRequest *models.VolumeGroupBackupListMetadata) {
	o.GetEntitiesRequest = getEntitiesRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostVolumeGroupBackupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.GetEntitiesRequest == nil {
		o.GetEntitiesRequest = new(models.VolumeGroupBackupListMetadata)
	}

	if err := r.SetBodyParam(o.GetEntitiesRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
