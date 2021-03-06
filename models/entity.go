package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Entity Directory service Search response
//
// Directory service Search response
// swagger:model entity
type Entity struct {

	// attribute set
	AttributeSet []*Attribute `json:"attribute_set,omitempty"`

	// canonical name of the entity
	// Max Length: 140
	Name string `json:"name,omitempty"`

	// user or group in the directory service
	Type string `json:"type,omitempty"`
}

// Validate validates this entity
func (m *Entity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeSet(formats); err != nil {
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

func (m *Entity) validateAttributeSet(formats strfmt.Registry) error {

	if swag.IsZero(m.AttributeSet) { // not required
		return nil
	}

	for i := 0; i < len(m.AttributeSet); i++ {

		if swag.IsZero(m.AttributeSet[i]) { // not required
			continue
		}

		if m.AttributeSet[i] != nil {

			if err := m.AttributeSet[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Entity) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}
