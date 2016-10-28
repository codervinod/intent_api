package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// UserMetadata user metadata
//
// The user kind metadata
// swagger:model user_metadata
type UserMetadata struct {

	// account reference
	AccountReference *AccountReference `json:"account_reference,omitempty"`

	// Categories for the user
	Categories map[string]string `json:"categories,omitempty"`

	// Time when user was created
	CreationTime string `json:"creation_time,omitempty"`

	// Monotonically increasing number
	EntityVersion int64 `json:"entity_version,omitempty"`

	// The kind name
	// Required: true
	// Read Only: true
	Kind string `json:"kind"`

	// Time when user was last updated
	LastUpdateTime string `json:"last_update_time,omitempty"`

	// user name
	// Max Length: 140
	Name string `json:"name,omitempty"`

	// owner reference
	OwnerReference *UserReference `json:"owner_reference,omitempty"`

	// If present, spec was copied from the referenced entity. Could have
	// been a live entity, backup, snapshot or template entity. This field
	// is special as it is the only field in metadata that is modifiable by
	// the user in case of template-based upgrades and in-place reverts.
	//
	// Pattern: \A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z
	ParentReference string `json:"parent_reference,omitempty"`

	// user uuid
	// Pattern: \A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this user metadata
func (m *UserMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCategories(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOwnerReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParentReference(formats); err != nil {
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

func (m *UserMetadata) validateAccountReference(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountReference) { // not required
		return nil
	}

	if m.AccountReference != nil {

		if err := m.AccountReference.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *UserMetadata) validateCategories(formats strfmt.Registry) error {

	if swag.IsZero(m.Categories) { // not required
		return nil
	}

	if err := validate.Required("categories", "body", m.Categories); err != nil {
		return err
	}

	return nil
}

var userMetadataTypeKindPropEnum []interface{}

const (
	userMetadataKindUser string = "user"
)

// prop value enum
func (m *UserMetadata) validateKindEnum(path, location string, value string) error {
	if userMetadataTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["user"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			userMetadataTypeKindPropEnum = append(userMetadataTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, userMetadataTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UserMetadata) validateKind(formats strfmt.Registry) error {

	if err := validate.RequiredString("kind", "body", string(m.Kind)); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *UserMetadata) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *UserMetadata) validateOwnerReference(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnerReference) { // not required
		return nil
	}

	if m.OwnerReference != nil {

		if err := m.OwnerReference.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *UserMetadata) validateParentReference(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentReference) { // not required
		return nil
	}

	if err := validate.Pattern("parent_reference", "body", string(m.ParentReference), `\A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z`); err != nil {
		return err
	}

	return nil
}

func (m *UserMetadata) validateUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.Pattern("uuid", "body", string(m.UUID), `\A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z`); err != nil {
		return err
	}

	return nil
}