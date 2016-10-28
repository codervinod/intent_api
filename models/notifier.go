package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// Notifier Notifier
//
// A notifier entity that can be used to monitor progress with respect to a
// set of requests issued to a system.
//
// swagger:model notifier
type Notifier struct {

	// api version
	// Read Only: true
	APIVersion string `json:"api_version,omitempty"`

	// metadata
	Metadata *InternalCommonMetadata `json:"metadata,omitempty"`
}

// Validate validates this notifier
func (m *Notifier) Validate(formats strfmt.Registry) error {
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

func (m *Notifier) validateMetadata(formats strfmt.Registry) error {

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
