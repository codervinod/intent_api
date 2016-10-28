package role

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

// NewPostRolesListParams creates a new PostRolesListParams object
// with the default values initialized.
func NewPostRolesListParams() *PostRolesListParams {
	var ()
	return &PostRolesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRolesListParamsWithTimeout creates a new PostRolesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRolesListParamsWithTimeout(timeout time.Duration) *PostRolesListParams {
	var ()
	return &PostRolesListParams{

		timeout: timeout,
	}
}

// NewPostRolesListParamsWithContext creates a new PostRolesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRolesListParamsWithContext(ctx context.Context) *PostRolesListParams {
	var ()
	return &PostRolesListParams{

		Context: ctx,
	}
}

/*PostRolesListParams contains all the parameters to send to the API endpoint
for the post roles list operation typically these are written to a http.Request
*/
type PostRolesListParams struct {

	/*GetEntitiesRequest*/
	GetEntitiesRequest *models.RoleListMetadata

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post roles list params
func (o *PostRolesListParams) WithTimeout(timeout time.Duration) *PostRolesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post roles list params
func (o *PostRolesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post roles list params
func (o *PostRolesListParams) WithContext(ctx context.Context) *PostRolesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post roles list params
func (o *PostRolesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithGetEntitiesRequest adds the getEntitiesRequest to the post roles list params
func (o *PostRolesListParams) WithGetEntitiesRequest(getEntitiesRequest *models.RoleListMetadata) *PostRolesListParams {
	o.SetGetEntitiesRequest(getEntitiesRequest)
	return o
}

// SetGetEntitiesRequest adds the getEntitiesRequest to the post roles list params
func (o *PostRolesListParams) SetGetEntitiesRequest(getEntitiesRequest *models.RoleListMetadata) {
	o.GetEntitiesRequest = getEntitiesRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostRolesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.GetEntitiesRequest == nil {
		o.GetEntitiesRequest = new(models.RoleListMetadata)
	}

	if err := r.SetBodyParam(o.GetEntitiesRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}