package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PermissionResourcesStatus Permission entity status
//
// Permission entity status
// swagger:model permission_resources_status
type PermissionResourcesStatus struct {

	// Human readable description of the permission
	// Max Length: 1000
	Description string `json:"description,omitempty"`

	// If state is kError, a message describing the error
	Message string `json:"message,omitempty"`

	// If state is kError, a machine_readable snake cased string
	Reason string `json:"reason,omitempty"`

	// resources
	Resources *PermissionResources `json:"resources,omitempty"`

	// The state of the permission resource
	State string `json:"state,omitempty"`
}

// Validate validates this permission resources status
func (m *PermissionResourcesStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PermissionResourcesStatus) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *PermissionResourcesStatus) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {

		if err := m.Resources.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
