package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// AvailableExtensionReference Reference to a available_extension
//
// The reference to a available_extension
// swagger:model available_extension_reference
type AvailableExtensionReference struct {

	// The kind name
	// Required: true
	// Read Only: true
	Kind string `json:"kind"`

	// name
	// Max Length: 140
	Name string `json:"name,omitempty"`

	// uuid
	// Pattern: \A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this available extension reference
func (m *AvailableExtensionReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var availableExtensionReferenceTypeKindPropEnum []interface{}

const (
	availableExtensionReferenceKindAvailableExtension string = "available_extension"
)

// prop value enum
func (m *AvailableExtensionReference) validateKindEnum(path, location string, value string) error {
	if availableExtensionReferenceTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["available_extension"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			availableExtensionReferenceTypeKindPropEnum = append(availableExtensionReferenceTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, availableExtensionReferenceTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AvailableExtensionReference) validateKind(formats strfmt.Registry) error {

	if err := validate.RequiredString("kind", "body", string(m.Kind)); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *AvailableExtensionReference) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *AvailableExtensionReference) validateUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.Pattern("uuid", "body", string(m.UUID), `\A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z`); err != nil {
		return err
	}

	return nil
}
