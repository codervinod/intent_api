package groups

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

// NewPostGroupsParams creates a new PostGroupsParams object
// with the default values initialized.
func NewPostGroupsParams() *PostGroupsParams {
	var ()
	return &PostGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostGroupsParamsWithTimeout creates a new PostGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostGroupsParamsWithTimeout(timeout time.Duration) *PostGroupsParams {
	var ()
	return &PostGroupsParams{

		timeout: timeout,
	}
}

// NewPostGroupsParamsWithContext creates a new PostGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostGroupsParamsWithContext(ctx context.Context) *PostGroupsParams {
	var ()
	return &PostGroupsParams{

		Context: ctx,
	}
}

/*PostGroupsParams contains all the parameters to send to the API endpoint
for the post groups operation typically these are written to a http.Request
*/
type PostGroupsParams struct {

	/*GetEntitiesRequest*/
	GetEntitiesRequest *models.GroupsGetEntitiesRequest

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post groups params
func (o *PostGroupsParams) WithTimeout(timeout time.Duration) *PostGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post groups params
func (o *PostGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post groups params
func (o *PostGroupsParams) WithContext(ctx context.Context) *PostGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post groups params
func (o *PostGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithGetEntitiesRequest adds the getEntitiesRequest to the post groups params
func (o *PostGroupsParams) WithGetEntitiesRequest(getEntitiesRequest *models.GroupsGetEntitiesRequest) *PostGroupsParams {
	o.SetGetEntitiesRequest(getEntitiesRequest)
	return o
}

// SetGetEntitiesRequest adds the getEntitiesRequest to the post groups params
func (o *PostGroupsParams) SetGetEntitiesRequest(getEntitiesRequest *models.GroupsGetEntitiesRequest) {
	o.GetEntitiesRequest = getEntitiesRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.GetEntitiesRequest == nil {
		o.GetEntitiesRequest = new(models.GroupsGetEntitiesRequest)
	}

	if err := r.SetBodyParam(o.GetEntitiesRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
