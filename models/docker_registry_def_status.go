package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DockerRegistryDefStatus Docker Registry
//
// Docker registry
// swagger:model docker_registry_def_status
type DockerRegistryDefStatus struct {

	// Name of the docker registry
	// Required: true
	// Max Length: 140
	Name *string `json:"name"`

	// resources
	// Required: true
	Resources *DockerRegistryResources `json:"resources"`
}

// Validate validates this docker registry def status
func (m *DockerRegistryDefStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
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

func (m *DockerRegistryDefStatus) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *DockerRegistryDefStatus) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	if m.Resources != nil {

		if err := m.Resources.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}