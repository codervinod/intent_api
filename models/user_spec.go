package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// UserSpec User spec
//
// User spec
// swagger:model user_spec
type UserSpec struct {

	// api version
	// Required: true
	// Read Only: true
	APIVersion string `json:"api_version"`

	// metadata
	// Required: true
	Metadata *UserMetadata `json:"metadata"`

	// spec
	Spec *UserResourcesSpec `json:"spec,omitempty"`
}

// Validate validates this user spec
func (m *UserSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSpec) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.RequiredString("api_version", "body", string(m.APIVersion)); err != nil {
		return err
	}

	return nil
}

func (m *UserSpec) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {

		if err := m.Metadata.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *UserSpec) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {

		if err := m.Spec.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
