package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VMBackupProfileList VM backup profile entity list
//
// VM backup profile entity list
// swagger:model vm_backup_profile_list
type VMBackupProfileList struct {

	// api version
	// Required: true
	APIVersion *string `json:"api_version"`

	// entities
	// Required: true
	Entities []*VMBackupProfileResources `json:"entities"`

	// metadata
	// Required: true
	Metadata *VMBackupProfileListMetadata `json:"metadata"`
}

// Validate validates this vm backup profile list
func (m *VMBackupProfileList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEntities(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMBackupProfileList) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.Required("api_version", "body", m.APIVersion); err != nil {
		return err
	}

	return nil
}

func (m *VMBackupProfileList) validateEntities(formats strfmt.Registry) error {

	if err := validate.Required("entities", "body", m.Entities); err != nil {
		return err
	}

	for i := 0; i < len(m.Entities); i++ {

		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {

			if err := m.Entities[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VMBackupProfileList) validateMetadata(formats strfmt.Registry) error {

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
