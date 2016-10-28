package user

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

// NewGetUsersMeParams creates a new GetUsersMeParams object
// with the default values initialized.
func NewGetUsersMeParams() *GetUsersMeParams {

	return &GetUsersMeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersMeParamsWithTimeout creates a new GetUsersMeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersMeParamsWithTimeout(timeout time.Duration) *GetUsersMeParams {

	return &GetUsersMeParams{

		timeout: timeout,
	}
}

// NewGetUsersMeParamsWithContext creates a new GetUsersMeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersMeParamsWithContext(ctx context.Context) *GetUsersMeParams {

	return &GetUsersMeParams{

		Context: ctx,
	}
}

/*GetUsersMeParams contains all the parameters to send to the API endpoint
for the get users me operation typically these are written to a http.Request
*/
type GetUsersMeParams struct {
	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get users me params
func (o *GetUsersMeParams) WithTimeout(timeout time.Duration) *GetUsersMeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users me params
func (o *GetUsersMeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users me params
func (o *GetUsersMeParams) WithContext(ctx context.Context) *GetUsersMeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users me params
func (o *GetUsersMeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersMeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}