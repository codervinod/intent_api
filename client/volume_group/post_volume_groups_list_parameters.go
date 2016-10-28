package volume_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostVolumeGroupsListParams creates a new PostVolumeGroupsListParams object
// with the default values initialized.
func NewPostVolumeGroupsListParams() *PostVolumeGroupsListParams {

	return &PostVolumeGroupsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVolumeGroupsListParamsWithTimeout creates a new PostVolumeGroupsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVolumeGroupsListParamsWithTimeout(timeout time.Duration) *PostVolumeGroupsListParams {

	return &PostVolumeGroupsListParams{

		timeout: timeout,
	}
}

// NewPostVolumeGroupsListParamsWithContext creates a new PostVolumeGroupsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVolumeGroupsListParamsWithContext(ctx context.Context) *PostVolumeGroupsListParams {

	return &PostVolumeGroupsListParams{

		Context: ctx,
	}
}

/*PostVolumeGroupsListParams contains all the parameters to send to the API endpoint
for the post volume groups list operation typically these are written to a http.Request
*/
type PostVolumeGroupsListParams struct {
	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post volume groups list params
func (o *PostVolumeGroupsListParams) WithTimeout(timeout time.Duration) *PostVolumeGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post volume groups list params
func (o *PostVolumeGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post volume groups list params
func (o *PostVolumeGroupsListParams) WithContext(ctx context.Context) *PostVolumeGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post volume groups list params
func (o *PostVolumeGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *PostVolumeGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}