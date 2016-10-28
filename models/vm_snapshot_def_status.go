package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VMSnapshotDefStatus vm snapshot output definitions
//
// The output object that defines a vm snapshot
// swagger:model vm_snapshot_def_status
type VMSnapshotDefStatus struct {

	// Reference to the cluster where this snapshot exists or needs to be
	// replicated to. This is to support the replication and retrieve of the
	// snapshot to a cluster.
	//
	ClusterReference *ClusterReference `json:"cluster_reference,omitempty"`

	// This field is same for all the entities (irrespective of kind) that
	// were snapshotted together.
	//
	// Pattern: \A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z
	ConsistencyGroupIdentifier string `json:"consistency_group_identifier,omitempty"`

	// Handle of the data snapshot. Cerebro PD snapshot handle can be used
	// here as data snapshot uuid. If a CG contains multiple entities,
	// then all the entity snapshots will share the same data snapshot uuid.
	//
	// Pattern: \A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z
	DataSnapshotIdentifier string `json:"data_snapshot_identifier,omitempty"`

	// entity backup reference
	EntityBackupReference *VMBackupReference `json:"entity_backup_reference,omitempty"`

	// The time when this snapshot expires and will be garbage collected.
	// If not set, then the snapshot never expires.
	//
	ExpirationTimeMsecs int64 `json:"expiration_time_msecs,omitempty"`

	// If a snapshot is replicated to a different clusters, then all the
	// instances of same snapshot will share this identifier.
	//
	// Pattern: \A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z
	LocationAgnosticIdentifier string `json:"location_agnostic_identifier,omitempty"`

	// Name of the snapshot
	// Max Length: 140
	Name string `json:"name,omitempty"`

	// resources
	Resources *VMSnapshotDefStatusResources `json:"resources,omitempty"`

	// Describes the files that are included in the snapshot.
	//
	// Required: true
	SnapshotFileList []*VMSnapshotDefStatusSnapshotFileListItems0 `json:"snapshot_file_list"`

	// Crash consistent or Application Consistent snapshot
	SnapshotType string `json:"snapshot_type,omitempty"`
}

// Validate validates this vm snapshot def status
func (m *VMSnapshotDefStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConsistencyGroupIdentifier(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDataSnapshotIdentifier(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEntityBackupReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocationAgnosticIdentifier(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSnapshotFileList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSnapshotType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMSnapshotDefStatus) validateClusterReference(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterReference) { // not required
		return nil
	}

	if m.ClusterReference != nil {

		if err := m.ClusterReference.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VMSnapshotDefStatus) validateConsistencyGroupIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.ConsistencyGroupIdentifier) { // not required
		return nil
	}

	if err := validate.Pattern("consistency_group_identifier", "body", string(m.ConsistencyGroupIdentifier), `\A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z`); err != nil {
		return err
	}

	return nil
}

func (m *VMSnapshotDefStatus) validateDataSnapshotIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.DataSnapshotIdentifier) { // not required
		return nil
	}

	if err := validate.Pattern("data_snapshot_identifier", "body", string(m.DataSnapshotIdentifier), `\A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z`); err != nil {
		return err
	}

	return nil
}

func (m *VMSnapshotDefStatus) validateEntityBackupReference(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityBackupReference) { // not required
		return nil
	}

	if m.EntityBackupReference != nil {

		if err := m.EntityBackupReference.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VMSnapshotDefStatus) validateLocationAgnosticIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.LocationAgnosticIdentifier) { // not required
		return nil
	}

	if err := validate.Pattern("location_agnostic_identifier", "body", string(m.LocationAgnosticIdentifier), `\A[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}\Z`); err != nil {
		return err
	}

	return nil
}

func (m *VMSnapshotDefStatus) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *VMSnapshotDefStatus) validateResources(formats strfmt.Registry) error {

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

func (m *VMSnapshotDefStatus) validateSnapshotFileList(formats strfmt.Registry) error {

	if err := validate.Required("snapshot_file_list", "body", m.SnapshotFileList); err != nil {
		return err
	}

	for i := 0; i < len(m.SnapshotFileList); i++ {

		if swag.IsZero(m.SnapshotFileList[i]) { // not required
			continue
		}

		if m.SnapshotFileList[i] != nil {

			if err := m.SnapshotFileList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var vmSnapshotDefStatusTypeSnapshotTypePropEnum []interface{}

const (
	vmSnapshotDefStatusSnapshotTypeCRASHCONSISTENT       string = "CRASH_CONSISTENT"
	vmSnapshotDefStatusSnapshotTypeAPPLICATIONCONSISTENT string = "APPLICATION_CONSISTENT"
)

// prop value enum
func (m *VMSnapshotDefStatus) validateSnapshotTypeEnum(path, location string, value string) error {
	if vmSnapshotDefStatusTypeSnapshotTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["CRASH_CONSISTENT","APPLICATION_CONSISTENT"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			vmSnapshotDefStatusTypeSnapshotTypePropEnum = append(vmSnapshotDefStatusTypeSnapshotTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, vmSnapshotDefStatusTypeSnapshotTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VMSnapshotDefStatus) validateSnapshotType(formats strfmt.Registry) error {

	if swag.IsZero(m.SnapshotType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSnapshotTypeEnum("snapshot_type", "body", m.SnapshotType); err != nil {
		return err
	}

	return nil
}

// VMSnapshotDefStatusResources Snapshot Resources
// swagger:model VMSnapshotDefStatusResources
type VMSnapshotDefStatusResources struct {

	// UUID of the base entity for which snapshot need to be taken
	//
	EntityUUID string `json:"entity_uuid,omitempty"`
}

// Validate validates this VM snapshot def status resources
func (m *VMSnapshotDefStatusResources) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// VMSnapshotDefStatusSnapshotFileListItems0 VM snapshot def status snapshot file list items0
// swagger:model VMSnapshotDefStatusSnapshotFileListItems0
type VMSnapshotDefStatusSnapshotFileListItems0 struct {

	// Pathname of the file that participated in the snapshot.
	//
	FilePath string `json:"file_path,omitempty"`

	// Pathname of the snapshot of the file.
	SnapshotFilePath string `json:"snapshot_file_path,omitempty"`
}

// Validate validates this VM snapshot def status snapshot file list items0
func (m *VMSnapshotDefStatusSnapshotFileListItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
