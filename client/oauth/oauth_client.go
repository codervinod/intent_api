package oauth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new oauth API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for oauth API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteOauthClientClientID deletes an existing oauth client

Delete existing Oauth client information
*/
func (a *Client) DeleteOauthClientClientID(params *DeleteOauthClientClientIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOauthClientClientIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOauthClientClientIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOauthClientClientID",
		Method:             "DELETE",
		PathPattern:        "/oauth/client/{client_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOauthClientClientIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOauthClientClientIDOK), nil

}

/*
DeleteOauthTokenUnregisterClientID unregisters client access token and revoke the access token issued to the client

Unregister client access token and revoke the access token issued to
the client

*/
func (a *Client) DeleteOauthTokenUnregisterClientID(params *DeleteOauthTokenUnregisterClientIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOauthTokenUnregisterClientIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOauthTokenUnregisterClientIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOauthTokenUnregisterClientID",
		Method:             "DELETE",
		PathPattern:        "/oauth/token/unregister/{client_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOauthTokenUnregisterClientIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOauthTokenUnregisterClientIDOK), nil

}

/*
GetOauthClientClientID useds to fetch existing oauth client details

Get Oauth client information
*/
func (a *Client) GetOauthClientClientID(params *GetOauthClientClientIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetOauthClientClientIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOauthClientClientIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOauthClientClientID",
		Method:             "GET",
		PathPattern:        "/oauth/client/{client_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOauthClientClientIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOauthClientClientIDOK), nil

}

/*
PostOauthAuthorize useds to get authorization code from server

Authorization confirmation post url
*/
func (a *Client) PostOauthAuthorize(params *PostOauthAuthorizeParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOauthAuthorizeParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOauthAuthorize",
		Method:             "POST",
		PathPattern:        "/oauth/authorize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOauthAuthorizeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PostOauthClient this is used to register new oauth client

Add new Oauth Client
*/
func (a *Client) PostOauthClient(params *PostOauthClientParams, authInfo runtime.ClientAuthInfoWriter) (*PostOauthClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOauthClientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOauthClient",
		Method:             "POST",
		PathPattern:        "/oauth/client",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOauthClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOauthClientOK), nil

}

/*
PostOauthClientList retrieves all oauth client

Retrieves all oauth clients
*/
func (a *Client) PostOauthClientList(params *PostOauthClientListParams, authInfo runtime.ClientAuthInfoWriter) (*PostOauthClientListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOauthClientListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOauthClientList",
		Method:             "POST",
		PathPattern:        "/oauth/client/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOauthClientListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOauthClientListOK), nil

}

/*
PostOauthGenerateToken this is used when the user is already authenticated and would like to create bearer tokens for distribution

Generates an oauth token
*/
func (a *Client) PostOauthGenerateToken(params *PostOauthGenerateTokenParams, authInfo runtime.ClientAuthInfoWriter) (*PostOauthGenerateTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOauthGenerateTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOauthGenerateToken",
		Method:             "POST",
		PathPattern:        "/oauth/generate_token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOauthGenerateTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOauthGenerateTokenOK), nil

}

/*
PostOauthToken this callback is used to get token from oauth2 provider

Returns an access token
*/
func (a *Client) PostOauthToken(params *PostOauthTokenParams, authInfo runtime.ClientAuthInfoWriter) (*PostOauthTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOauthTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOauthToken",
		Method:             "POST",
		PathPattern:        "/oauth/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOauthTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOauthTokenOK), nil

}

/*
PostOauthTokenRegister this is used when a client wants to create a communication channel with the server by using oauth tokens

Register an oauth token

*/
func (a *Client) PostOauthTokenRegister(params *PostOauthTokenRegisterParams, authInfo runtime.ClientAuthInfoWriter) (*PostOauthTokenRegisterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOauthTokenRegisterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOauthTokenRegister",
		Method:             "POST",
		PathPattern:        "/oauth/token/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOauthTokenRegisterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOauthTokenRegisterOK), nil

}

/*
PutOauthClientClientID useds to update existing client details

Update Oauth client information
*/
func (a *Client) PutOauthClientClientID(params *PutOauthClientClientIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutOauthClientClientIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOauthClientClientIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutOauthClientClientID",
		Method:             "PUT",
		PathPattern:        "/oauth/client/{client_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOauthClientClientIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOauthClientClientIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
