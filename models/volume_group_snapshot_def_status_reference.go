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

// VolumeGroupSnapshotDefStatusReference Reference to a volume_group_snapshot_def_status
//
// The reference to a volume_group_snapshot_def_status
// swagger:model volume_group_snapshot_def_status_reference
type VolumeGroupSnapshotDefStatusReference struct {

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

// Validate validates this volume group snapshot def status reference
func (m *VolumeGroupSnapshotDefStatusReference) Validate(formats strfmt.Registry) error {
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

var volumeGroupSnapshotDefStatusReferenceTypeKindPropEnum []interface{}

const (
	volumeGroupSnapshotDefStatusReferenceKindVolumeGroupSnapshotDefStatus string = "volume_group_snapshot_def_status"
)

// prop value enum
func (m *VolumeGroupSnapshotDefStatusReference) validateKindEnum(path, location string, value string) error {
	if volumeGroupSnapshotDefStatusReferenceTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["volume_group_snapshot_def_status"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			volumeGroupSnapshotDefStatusReferenceTypeKindPropEnum = append(volumeGroupSnapshotDefStatusReferenceTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, volumeGroupSnapshotDefStatusReferenceTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VolumeGroupSnapshotDefStatusReference) validateKind(formats strfmt.Registry) error {

	if err := validate.RequiredString("kind", "body", string(m.Kind)); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *VolumeGroupSnapshotDefStatusReference) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *VolumeGroupSnapshotDefStatusReference) validateUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.Pattern("uuid", "body", string(m.UUID), `\A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z`); err != nil {
		return err
	}

	return nil
}
