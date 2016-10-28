package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VMBackupProfileResources VM backup profile creation/modification spec
//
// VM backup profile creation/modification spec
// swagger:model vm_backup_profile_resources
type VMBackupProfileResources struct {

	// A description for the VM backup profile
	Description string `json:"description,omitempty"`

	// VM backup profile name
	Name string `json:"name,omitempty"`

	// resources
	Resources *VMBackupProfileResourcesResources `json:"resources,omitempty"`
}

// Validate validates this vm backup profile resources
func (m *VMBackupProfileResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMBackupProfileResources) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {

		if err := m.Resources.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// VMBackupProfileResourcesResources VM backup profile resource
// swagger:model VMBackupProfileResourcesResources
type VMBackupProfileResourcesResources struct {

	// backup policy
	// Required: true
	BackupPolicy *VMBackupProfileResourcesResourcesBackupPolicy `json:"backup_policy"`
}

// Validate validates this VM backup profile resources resources
func (m *VMBackupProfileResourcesResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupPolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMBackupProfileResourcesResources) validateBackupPolicy(formats strfmt.Registry) error {

	if err := validate.Required("resources"+"."+"backup_policy", "body", m.BackupPolicy); err != nil {
		return err
	}

	if m.BackupPolicy != nil {

		if err := m.BackupPolicy.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// VMBackupProfileResourcesResourcesBackupPolicy Describes how and where to backup the VM
// swagger:model VMBackupProfileResourcesResourcesBackupPolicy
type VMBackupProfileResourcesResourcesBackupPolicy struct {

	// snapshot policy
	SnapshotPolicy *SnapshotPolicy `json:"snapshot_policy,omitempty"`
}

// Validate validates this VM backup profile resources resources backup policy
func (m *VMBackupProfileResourcesResourcesBackupPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshotPolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMBackupProfileResourcesResourcesBackupPolicy) validateSnapshotPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.SnapshotPolicy) { // not required
		return nil
	}

	if m.SnapshotPolicy != nil {

		if err := m.SnapshotPolicy.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
