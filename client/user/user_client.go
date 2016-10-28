package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetLogout logs out the current user

Invalidates the session for current user
*/
func (a *Client) GetLogout(params *GetLogoutParams, authInfo runtime.ClientAuthInfoWriter) (*GetLogoutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogoutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLogout",
		Method:             "GET",
		PathPattern:        "/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLogoutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLogoutOK), nil

}

/*
GetUsersMe retrieves currently logged in user

Retrieves currently logged in user
*/
func (a *Client) GetUsersMe(params *GetUsersMeParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersMeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersMeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsersMe",
		Method:             "GET",
		PathPattern:        "/users/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersMeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsersMeOK), nil

}

/*
GetUsersUUID retrieves specified user

Retrieves specified user
*/
func (a *Client) GetUsersUUID(params *GetUsersUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsersUUID",
		Method:             "GET",
		PathPattern:        "/users/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsersUUIDOK), nil

}

/*
PutUsersUUID updates user

Update user
*/
func (a *Client) PutUsersUUID(params *PutUsersUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutUsersUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUsersUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutUsersUUID",
		Method:             "PUT",
		PathPattern:        "/users/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutUsersUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutUsersUUIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
