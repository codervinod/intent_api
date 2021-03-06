package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ResourceDomainCreateRequest Resource Domain Create Request
//
// Request to create a resource domain
// swagger:model resource_domain_create_request
type ResourceDomainCreateRequest struct {

	// metadata
	// Required: true
	Metadata *ResourceDomainMetadata `json:"metadata"`

	// spec
	// Required: true
	Spec *ResourceDomainSpec `json:"spec"`
}

// Validate validates this resource domain create request
func (m *ResourceDomainCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *ResourceDomainCreateRequest) validateMetadata(formats strfmt.Registry) error {

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

func (m *ResourceDomainCreateRequest) validateSpec(formats strfmt.Registry) error {

	if err := validate.Required("spec", "body", m.Spec); err != nil {
		return err
	}

	if m.Spec != nil {

		if err := m.Spec.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
