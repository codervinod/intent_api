package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// ConnectionResult Connection Result
//
// Connection Result
// swagger:model connection_result
type ConnectionResult struct {

	// api version
	APIVersion string `json:"api_version,omitempty"`

	// Connection status
	ConnectionStatus bool `json:"connection_status,omitempty"`

	// Connection status message
	Message string `json:"message,omitempty"`

	// metadata
	Metadata *DirectoryServiceMetadata `json:"metadata,omitempty"`
}

// Validate validates this connection result
func (m *ConnectionResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionResult) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {

		if err := m.Metadata.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
