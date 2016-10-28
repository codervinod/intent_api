package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DiskEncryptionStatus Cluster disk encryption status
//
// Cluster disk encryption status
// swagger:model disk_encryption_status
type DiskEncryptionStatus struct {

	// Indicates if disk is encryption capable
	EncryptionCapable bool `json:"encryption_capable,omitempty"`

	// Request to set the new passwords for the encryption capable disk
	//
	Rekey bool `json:"rekey,omitempty"`

	// Disk UUID
	// Required: true
	UUID *string `json:"uuid"`
}

// Validate validates this disk encryption status
func (m *DiskEncryptionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiskEncryptionStatus) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}
