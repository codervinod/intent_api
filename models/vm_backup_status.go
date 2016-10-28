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

// VMBackupStatus Response Kind
//
// The status of a REST API call. Only used when there is a failure to
// report.
//
// swagger:model vm_backup_status
type VMBackupStatus struct {

	// api version
	// Read Only: true
	APIVersion string `json:"api_version,omitempty"`

	// The HTTP error code.
	// Read Only: true
	Code int64 `json:"code,omitempty"`

	// Custom key-value details relevant to the status.
	// Read Only: true
	Details map[string]string `json:"details,omitempty"`

	// The kind backup name
	// Read Only: true
	Kind string `json:"kind,omitempty"`

	// A sentence explaining the reason for the status.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// One snake_case word.
	// Read Only: true
	Reason string `json:"reason,omitempty"`

	// Only value possible is "failure".
	// Read Only: true
	Status string `json:"status,omitempty"`
}

// Validate validates this vm backup status
func (m *VMBackupStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMBackupStatus) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if err := validate.Required("details", "body", m.Details); err != nil {
		return err
	}

	return nil
}

var vmBackupStatusTypeKindPropEnum []interface{}

const (
	vmBackupStatusKindVMBackup string = "vm_backup"
)

// prop value enum
func (m *VMBackupStatus) validateKindEnum(path, location string, value string) error {
	if vmBackupStatusTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["vm_backup"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			vmBackupStatusTypeKindPropEnum = append(vmBackupStatusTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, vmBackupStatusTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VMBackupStatus) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}
