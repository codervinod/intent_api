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

// NewPostOauthTokenParams creates a new PostOauthTokenParams object
// with the default values initialized.
func NewPostOauthTokenParams() *PostOauthTokenParams {
	var ()
	return &PostOauthTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOauthTokenParamsWithTimeout creates a new PostOauthTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOauthTokenParamsWithTimeout(timeout time.Duration) *PostOauthTokenParams {
	var ()
	return &PostOauthTokenParams{

		timeout: timeout,
	}
}

// NewPostOauthTokenParamsWithContext creates a new PostOauthTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOauthTokenParamsWithContext(ctx context.Context) *PostOauthTokenParams {
	var ()
	return &PostOauthTokenParams{

		Context: ctx,
	}
}

/*PostOauthTokenParams contains all the parameters to send to the API endpoint
for the post oauth token operation typically these are written to a http.Request
*/
type PostOauthTokenParams struct {

	/*ClientID
	  Your client ID

	*/
	ClientID string
	/*ClientSecret
	  Your client secret

	*/
	ClientSecret string
	/*Code
	  authorization_code received in redirect_uri
	must for grant_type = authorization_code


	*/
	Code *string
	/*GrantType
	  Must be authorization_code or refresh_token

	*/
	GrantType string
	/*RedirectURI
	  redirect uri used to get authorization
	must for grant_type = authorization_code


	*/
	RedirectURI *string
	/*RefreshToken
	  refresh_token received with previous token
	must for grant_type = refresh_token


	*/
	RefreshToken *string
	/*State
	  state parameter to prevent cross site origin attacks
	use only if passed in getting authorization


	*/
	State *string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post oauth token params
func (o *PostOauthTokenParams) WithTimeout(timeout time.Duration) *PostOauthTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post oauth token params
func (o *PostOauthTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post oauth token params
func (o *PostOauthTokenParams) WithContext(ctx context.Context) *PostOauthTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post oauth token params
func (o *PostOauthTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithClientID adds the clientID to the post oauth token params
func (o *PostOauthTokenParams) WithClientID(clientID string) *PostOauthTokenParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the post oauth token params
func (o *PostOauthTokenParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithClientSecret adds the clientSecret to the post oauth token params
func (o *PostOauthTokenParams) WithClientSecret(clientSecret string) *PostOauthTokenParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the post oauth token params
func (o *PostOauthTokenParams) SetClientSecret(clientSecret string) {
	o.ClientSecret = clientSecret
}

// WithCode adds the code to the post oauth token params
func (o *PostOauthTokenParams) WithCode(code *string) *PostOauthTokenParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the post oauth token params
func (o *PostOauthTokenParams) SetCode(code *string) {
	o.Code = code
}

// WithGrantType adds the grantType to the post oauth token params
func (o *PostOauthTokenParams) WithGrantType(grantType string) *PostOauthTokenParams {
	o.SetGrantType(grantType)
	return o
}

// SetGrantType adds the grantType to the post oauth token params
func (o *PostOauthTokenParams) SetGrantType(grantType string) {
	o.GrantType = grantType
}

// WithRedirectURI adds the redirectURI to the post oauth token params
func (o *PostOauthTokenParams) WithRedirectURI(redirectURI *string) *PostOauthTokenParams {
	o.SetRedirectURI(redirectURI)
	return o
}

// SetRedirectURI adds the redirectUri to the post oauth token params
func (o *PostOauthTokenParams) SetRedirectURI(redirectURI *string) {
	o.RedirectURI = redirectURI
}

// WithRefreshToken adds the refreshToken to the post oauth token params
func (o *PostOauthTokenParams) WithRefreshToken(refreshToken *string) *PostOauthTokenParams {
	o.SetRefreshToken(refreshToken)
	return o
}

// SetRefreshToken adds the refreshToken to the post oauth token params
func (o *PostOauthTokenParams) SetRefreshToken(refreshToken *string) {
	o.RefreshToken = refreshToken
}

// WithState adds the state to the post oauth token params
func (o *PostOauthTokenParams) WithState(state *string) *PostOauthTokenParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the post oauth token params
func (o *PostOauthTokenParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *PostOauthTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// form param client_id
	frClientID := o.ClientID
	fClientID := frClientID
	if fClientID != "" {
		if err := r.SetFormParam("client_id", fClientID); err != nil {
			return err
		}
	}

	// form param client_secret
	frClientSecret := o.ClientSecret
	fClientSecret := frClientSecret
	if fClientSecret != "" {
		if err := r.SetFormParam("client_secret", fClientSecret); err != nil {
			return err
		}
	}

	if o.Code != nil {

		// form param code
		var frCode string
		if o.Code != nil {
			frCode = *o.Code
		}
		fCode := frCode
		if fCode != "" {
			if err := r.SetFormParam("code", fCode); err != nil {
				return err
			}
		}

	}

	// form param grant_type
	frGrantType := o.GrantType
	fGrantType := frGrantType
	if fGrantType != "" {
		if err := r.SetFormParam("grant_type", fGrantType); err != nil {
			return err
		}
	}

	if o.RedirectURI != nil {

		// form param redirect_uri
		var frRedirectURI string
		if o.RedirectURI != nil {
			frRedirectURI = *o.RedirectURI
		}
		fRedirectURI := frRedirectURI
		if fRedirectURI != "" {
			if err := r.SetFormParam("redirect_uri", fRedirectURI); err != nil {
				return err
			}
		}

	}

	if o.RefreshToken != nil {

		// form param refresh_token
		var frRefreshToken string
		if o.RefreshToken != nil {
			frRefreshToken = *o.RefreshToken
		}
		fRefreshToken := frRefreshToken
		if fRefreshToken != "" {
			if err := r.SetFormParam("refresh_token", fRefreshToken); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// form param state
		var frState string
		if o.State != nil {
			frState = *o.State
		}
		fState := frState
		if fState != "" {
			if err := r.SetFormParam("state", fState); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
