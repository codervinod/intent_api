package oauth

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

// NewGetOauthClientClientIDParams creates a new GetOauthClientClientIDParams object
// with the default values initialized.
func NewGetOauthClientClientIDParams() *GetOauthClientClientIDParams {
	var ()
	return &GetOauthClientClientIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOauthClientClientIDParamsWithTimeout creates a new GetOauthClientClientIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOauthClientClientIDParamsWithTimeout(timeout time.Duration) *GetOauthClientClientIDParams {
	var ()
	return &GetOauthClientClientIDParams{

		timeout: timeout,
	}
}

// NewGetOauthClientClientIDParamsWithContext creates a new GetOauthClientClientIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOauthClientClientIDParamsWithContext(ctx context.Context) *GetOauthClientClientIDParams {
	var ()
	return &GetOauthClientClientIDParams{

		Context: ctx,
	}
}

/*GetOauthClientClientIDParams contains all the parameters to send to the API endpoint
for the get oauth client client ID operation typically these are written to a http.Request
*/
type GetOauthClientClientIDParams struct {

	/*ClientID*/
	ClientID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get oauth client client ID params
func (o *GetOauthClientClientIDParams) WithTimeout(timeout time.Duration) *GetOauthClientClientIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oauth client client ID params
func (o *GetOauthClientClientIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oauth client client ID params
func (o *GetOauthClientClientIDParams) WithContext(ctx context.Context) *GetOauthClientClientIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oauth client client ID params
func (o *GetOauthClientClientIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithClientID adds the clientID to the get oauth client client ID params
func (o *GetOauthClientClientIDParams) WithClientID(clientID string) *GetOauthClientClientIDParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the get oauth client client ID params
func (o *GetOauthClientClientIDParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOauthClientClientIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param client_id
	if err := r.SetPathParam("client_id", o.ClientID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
