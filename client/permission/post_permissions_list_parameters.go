package permission

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

// NewPostPermissionsListParams creates a new PostPermissionsListParams object
// with the default values initialized.
func NewPostPermissionsListParams() *PostPermissionsListParams {
	var ()
	return &PostPermissionsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPermissionsListParamsWithTimeout creates a new PostPermissionsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPermissionsListParamsWithTimeout(timeout time.Duration) *PostPermissionsListParams {
	var ()
	return &PostPermissionsListParams{

		timeout: timeout,
	}
}

// NewPostPermissionsListParamsWithContext creates a new PostPermissionsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPermissionsListParamsWithContext(ctx context.Context) *PostPermissionsListParams {
	var ()
	return &PostPermissionsListParams{

		Context: ctx,
	}
}

/*PostPermissionsListParams contains all the parameters to send to the API endpoint
for the post permissions list operation typically these are written to a http.Request
*/
type PostPermissionsListParams struct {

	/*GetEntitiesRequest*/
	GetEntitiesRequest *models.PermissionListMetadata

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post permissions list params
func (o *PostPermissionsListParams) WithTimeout(timeout time.Duration) *PostPermissionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post permissions list params
func (o *PostPermissionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post permissions list params
func (o *PostPermissionsListParams) WithContext(ctx context.Context) *PostPermissionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post permissions list params
func (o *PostPermissionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithGetEntitiesRequest adds the getEntitiesRequest to the post permissions list params
func (o *PostPermissionsListParams) WithGetEntitiesRequest(getEntitiesRequest *models.PermissionListMetadata) *PostPermissionsListParams {
	o.SetGetEntitiesRequest(getEntitiesRequest)
	return o
}

// SetGetEntitiesRequest adds the getEntitiesRequest to the post permissions list params
func (o *PostPermissionsListParams) SetGetEntitiesRequest(getEntitiesRequest *models.PermissionListMetadata) {
	o.GetEntitiesRequest = getEntitiesRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostPermissionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.GetEntitiesRequest == nil {
		o.GetEntitiesRequest = new(models.PermissionListMetadata)
	}

	if err := r.SetBodyParam(o.GetEntitiesRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
