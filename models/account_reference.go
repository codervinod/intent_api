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

// AccountReference Reference to a account
//
// The reference to a account
// swagger:model account_reference
type AccountReference struct {

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

// Validate validates this account reference
func (m *AccountReference) Validate(formats strfmt.Registry) error {
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

var accountReferenceTypeKindPropEnum []interface{}

const (
	accountReferenceKindAccount string = "account"
)

// prop value enum
func (m *AccountReference) validateKindEnum(path, location string, value string) error {
	if accountReferenceTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["account"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			accountReferenceTypeKindPropEnum = append(accountReferenceTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, accountReferenceTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountReference) validateKind(formats strfmt.Registry) error {

	if err := validate.RequiredString("kind", "body", string(m.Kind)); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *AccountReference) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *AccountReference) validateUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.Pattern("uuid", "body", string(m.UUID), `\A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z`); err != nil {
		return err
	}

	return nil
}
