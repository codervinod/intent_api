package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CategoryResourceStatus Category object status
//
// Category object status
// swagger:model category_resource_status
type CategoryResourceStatus struct {

	// List of allowed values for category
	AllowedValues []*CategoryValue `json:"allowed_values,omitempty"`

	// Name of the category
	// Max Length: 140
	Name string `json:"name,omitempty"`
}

// Validate validates this category resource status
func (m *CategoryResourceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CategoryResourceStatus) validateAllowedValues(formats strfmt.Registry) error {

	if swag.IsZero(m.AllowedValues) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowedValues); i++ {

		if swag.IsZero(m.AllowedValues[i]) { // not required
			continue
		}

		if m.AllowedValues[i] != nil {

			if err := m.AllowedValues[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *CategoryResourceStatus) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}