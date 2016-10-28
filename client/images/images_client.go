package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new images API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for images API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteImagesUUID deletes image

Delete an image
*/
func (a *Client) DeleteImagesUUID(params *DeleteImagesUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteImagesUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImagesUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteImagesUUID",
		Method:             "DELETE",
		PathPattern:        "/images/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteImagesUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteImagesUUIDOK), nil

}

/*
GetImagesUUID gets image

Get details for an image
*/
func (a *Client) GetImagesUUID(params *GetImagesUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetImagesUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImagesUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetImagesUUID",
		Method:             "GET",
		PathPattern:        "/images/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetImagesUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetImagesUUIDOK), nil

}

/*
PostImages creates image

Create an image
*/
func (a *Client) PostImages(params *PostImagesParams, authInfo runtime.ClientAuthInfoWriter) (*PostImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostImages",
		Method:             "POST",
		PathPattern:        "/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostImagesOK), nil

}

/*
PostImagesList lists images

List images
*/
func (a *Client) PostImagesList(params *PostImagesListParams, authInfo runtime.ClientAuthInfoWriter) (*PostImagesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostImagesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostImagesList",
		Method:             "POST",
		PathPattern:        "/images/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostImagesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostImagesListOK), nil

}

/*
PutImagesUUID modifies image

Modify image
*/
func (a *Client) PutImagesUUID(params *PutImagesUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutImagesUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutImagesUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutImagesUUID",
		Method:             "PUT",
		PathPattern:        "/images/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutImagesUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutImagesUUIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
