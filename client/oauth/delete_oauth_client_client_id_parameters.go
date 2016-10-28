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

// NewDeleteOauthClientClientIDParams creates a new DeleteOauthClientClientIDParams object
// with the default values initialized.
func NewDeleteOauthClientClientIDParams() *DeleteOauthClientClientIDParams {
	var ()
	return &DeleteOauthClientClientIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOauthClientClientIDParamsWithTimeout creates a new DeleteOauthClientClientIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOauthClientClientIDParamsWithTimeout(timeout time.Duration) *DeleteOauthClientClientIDParams {
	var ()
	return &DeleteOauthClientClientIDParams{

		timeout: timeout,
	}
}

// NewDeleteOauthClientClientIDParamsWithContext creates a new DeleteOauthClientClientIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOauthClientClientIDParamsWithContext(ctx context.Context) *DeleteOauthClientClientIDParams {
	var ()
	return &DeleteOauthClientClientIDParams{

		Context: ctx,
	}
}

/*DeleteOauthClientClientIDParams contains all the parameters to send to the API endpoint
for the delete oauth client client ID operation typically these are written to a http.Request
*/
type DeleteOauthClientClientIDParams struct {

	/*ClientID*/
	ClientID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the delete oauth client client ID params
func (o *DeleteOauthClientClientIDParams) WithTimeout(timeout time.Duration) *DeleteOauthClientClientIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete oauth client client ID params
func (o *DeleteOauthClientClientIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete oauth client client ID params
func (o *DeleteOauthClientClientIDParams) WithContext(ctx context.Context) *DeleteOauthClientClientIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete oauth client client ID params
func (o *DeleteOauthClientClientIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithClientID adds the clientID to the delete oauth client client ID params
func (o *DeleteOauthClientClientIDParams) WithClientID(clientID string) *DeleteOauthClientClientIDParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the delete oauth client client ID params
func (o *DeleteOauthClientClientIDParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOauthClientClientIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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