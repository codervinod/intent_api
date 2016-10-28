package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new permission API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for permission API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeletePermissionsUUID deletes a permission object

Delete a permission object
*/
func (a *Client) DeletePermissionsUUID(params *DeletePermissionsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePermissionsUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePermissionsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePermissionsUUID",
		Method:             "DELETE",
		PathPattern:        "/permissions/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePermissionsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePermissionsUUIDOK), nil

}

/*
GetPermissionsUUID gets a permission object

Get a permission object
*/
func (a *Client) GetPermissionsUUID(params *GetPermissionsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPermissionsUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPermissionsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPermissionsUUID",
		Method:             "GET",
		PathPattern:        "/permissions/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPermissionsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPermissionsUUIDOK), nil

}

/*
PostPermissions creates a permission

Creates a permission
*/
func (a *Client) PostPermissions(params *PostPermissionsParams, authInfo runtime.ClientAuthInfoWriter) (*PostPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPermissionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPermissions",
		Method:             "POST",
		PathPattern:        "/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPermissionsOK), nil

}

/*
PostPermissionsList gets permissions

Get permissions
*/
func (a *Client) PostPermissionsList(params *PostPermissionsListParams, authInfo runtime.ClientAuthInfoWriter) (*PostPermissionsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPermissionsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPermissionsList",
		Method:             "POST",
		PathPattern:        "/permissions/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPermissionsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPermissionsListOK), nil

}

/*
PutPermissionsUUID updates permission

Update permission
*/
func (a *Client) PutPermissionsUUID(params *PutPermissionsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutPermissionsUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutPermissionsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutPermissionsUUID",
		Method:             "PUT",
		PathPattern:        "/permissions/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutPermissionsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutPermissionsUUIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
