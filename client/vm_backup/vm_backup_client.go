package vm_backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vm backup API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vm backup API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetVMBackupsUUID gets kind backup

Given a UUID, returns a kind backup definition
*/
func (a *Client) GetVMBackupsUUID(params *GetVMBackupsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetVMBackupsUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMBackupsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVMBackupsUUID",
		Method:             "GET",
		PathPattern:        "/vm_backups/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVMBackupsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVMBackupsUUIDOK), nil

}

/*
PostVMBackupsList gets kind backups

Get kind backups
*/
func (a *Client) PostVMBackupsList(params *PostVMBackupsListParams, authInfo runtime.ClientAuthInfoWriter) (*PostVMBackupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVMBackupsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostVMBackupsList",
		Method:             "POST",
		PathPattern:        "/vm_backups/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostVMBackupsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostVMBackupsListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
