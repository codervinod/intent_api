package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VolumeGroupBackupDefStatus volume_group backup output definitions
//
// The output object that defines a volume_group backup
// swagger:model volume_group_backup_def_status
type VolumeGroupBackupDefStatus struct {

	// Name of the backup entity
	// Max Length: 140
	Name string `json:"name,omitempty"`

	// Latest read only copy of the entity policies
	// Read Only: true
	Policies interface{} `json:"policies,omitempty"`

	// List of snapshots
	SnapshotReferenceList []*VolumeGroupSnapshotReference `json:"snapshot_reference_list,omitempty"`

	// Backup is a synchronous replica
	SynchronousCopyEnabled bool `json:"synchronous_copy_enabled,omitempty"`
}

// Validate validates this volume group backup def status
func (m *VolumeGroupBackupDefStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSnapshotReferenceList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeGroupBackupDefStatus) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *VolumeGroupBackupDefStatus) validateSnapshotReferenceList(formats strfmt.Registry) error {

	if swag.IsZero(m.SnapshotReferenceList) { // not required
		return nil
	}

	for i := 0; i < len(m.SnapshotReferenceList); i++ {

		if swag.IsZero(m.SnapshotReferenceList[i]) { // not required
			continue
		}

		if m.SnapshotReferenceList[i] != nil {

			if err := m.SnapshotReferenceList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
