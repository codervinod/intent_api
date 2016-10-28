package vm_backup_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vm backup profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vm backup profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteVMBackupProfilesUUID deletes a VM backup profile

Delete a VM backup profile
*/
func (a *Client) DeleteVMBackupProfilesUUID(params *DeleteVMBackupProfilesUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteVMBackupProfilesUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVMBackupProfilesUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVMBackupProfilesUUID",
		Method:             "DELETE",
		PathPattern:        "/vm_backup_profiles/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVMBackupProfilesUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteVMBackupProfilesUUIDOK), nil

}

/*
GetVMBackupProfilesUUID gets details for a VM backup profile

Get details for a VM backup profile
*/
func (a *Client) GetVMBackupProfilesUUID(params *GetVMBackupProfilesUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetVMBackupProfilesUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMBackupProfilesUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVMBackupProfilesUUID",
		Method:             "GET",
		PathPattern:        "/vm_backup_profiles/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVMBackupProfilesUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVMBackupProfilesUUIDOK), nil

}

/*
PostVMBackupProfiles creates a VM backup profile

Create a VM backup profile
*/
func (a *Client) PostVMBackupProfiles(params *PostVMBackupProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*PostVMBackupProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVMBackupProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVMBackupProfiles",
		Method:             "POST",
		PathPattern:        "/vm_backup_profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVMBackupProfilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostVMBackupProfilesOK), nil

}

/*
PostVMBackupProfilesList gets VM backup profiles

Get VM backup profiles
*/
func (a *Client) PostVMBackupProfilesList(params *PostVMBackupProfilesListParams, authInfo runtime.ClientAuthInfoWriter) (*PostVMBackupProfilesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVMBackupProfilesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVMBackupProfilesList",
		Method:             "POST",
		PathPattern:        "/vm_backup_profiles/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVMBackupProfilesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostVMBackupProfilesListOK), nil

}

/*
PutVMBackupProfilesUUID modifies VM backup profile

Modify VM backup profile
*/
func (a *Client) PutVMBackupProfilesUUID(params *PutVMBackupProfilesUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutVMBackupProfilesUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutVMBackupProfilesUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutVMBackupProfilesUUID",
		Method:             "PUT",
		PathPattern:        "/vm_backup_profiles/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutVMBackupProfilesUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutVMBackupProfilesUUIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}