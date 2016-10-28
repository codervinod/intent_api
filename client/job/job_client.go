package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new job API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for job API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteJobsUUID deletes job

Delete a Job given its UUID
*/
func (a *Client) DeleteJobsUUID(params *DeleteJobsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteJobsUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteJobsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteJobsUUID",
		Method:             "DELETE",
		PathPattern:        "/jobs/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteJobsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteJobsUUIDOK), nil

}

/*
GetJobsUUID gets job

Given a UUID, returns a Job definition
*/
func (a *Client) GetJobsUUID(params *GetJobsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetJobsUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetJobsUUID",
		Method:             "GET",
		PathPattern:        "/jobs/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJobsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJobsUUIDOK), nil

}

/*
PatchJobsUUID updates job partial

Given an intentful spec and Job UUID, update Job using partial update semantics
*/
func (a *Client) PatchJobsUUID(params *PatchJobsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchJobsUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchJobsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchJobsUUID",
		Method:             "PATCH",
		PathPattern:        "/jobs/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchJobsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchJobsUUIDOK), nil

}

/*
PostJobs creates job

Create a Job for a workflow with associated metadata
*/
func (a *Client) PostJobs(params *PostJobsParams, authInfo runtime.ClientAuthInfoWriter) (*PostJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostJobsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostJobs",
		Method:             "POST",
		PathPattern:        "/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostJobsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostJobsOK), nil

}

/*
PostJobsList lists the jobs

List of all the jobs.
*/
func (a *Client) PostJobsList(params *PostJobsListParams, authInfo runtime.ClientAuthInfoWriter) (*PostJobsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostJobsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostJobsList",
		Method:             "POST",
		PathPattern:        "/jobs/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostJobsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostJobsListOK), nil

}

/*
PutJobsUUID updates job

Given an intenful spec and Job uuid, update Job
*/
func (a *Client) PutJobsUUID(params *PutJobsUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutJobsUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutJobsUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutJobsUUID",
		Method:             "PUT",
		PathPattern:        "/jobs/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutJobsUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutJobsUUIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}