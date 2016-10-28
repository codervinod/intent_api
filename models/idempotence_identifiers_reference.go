package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// IdempotenceIdentifiersReference Reference to a idempotence_identifiers
//
// The reference to a idempotence_identifiers
// swagger:model idempotence_identifiers_reference
type IdempotenceIdentifiersReference struct {

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

// Validate validates this idempotence identifiers reference
func (m *IdempotenceIdentifiersReference) Validate(formats strfmt.Registry) error {
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

var idempotenceIdentifiersReferenceTypeKindPropEnum []interface{}

const (
	idempotenceIdentifiersReferenceKindIdempotenceIdentifiers string = "idempotence_identifiers"
)

// prop value enum
func (m *IdempotenceIdentifiersReference) validateKindEnum(path, location string, value string) error {
	if idempotenceIdentifiersReferenceTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["idempotence_identifiers"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			idempotenceIdentifiersReferenceTypeKindPropEnum = append(idempotenceIdentifiersReferenceTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, idempotenceIdentifiersReferenceTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IdempotenceIdentifiersReference) validateKind(formats strfmt.Registry) error {

	if err := validate.RequiredString("kind", "body", string(m.Kind)); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *IdempotenceIdentifiersReference) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *IdempotenceIdentifiersReference) validateUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.Pattern("uuid", "body", string(m.UUID), `\A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z`); err != nil {
		return err
	}

	return nil
}
