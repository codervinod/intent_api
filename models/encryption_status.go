package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// EncryptionStatus Cluster encryption status
//
// Cluster encryption status
// swagger:model encryption_status
type EncryptionStatus struct {

	// List of disk encryption status
	Disks []*DiskEncryptionStatus `json:"disks,omitempty"`

	// Flag to indicate if the cluster has password protection for
	// encrption
	//
	Protected bool `json:"protected,omitempty"`

	// Indicates if the current cluster configuration is sufficient to
	// turn on password protection for encryption
	//
	ReadyForEncryption bool `json:"ready_for_encryption,omitempty"`
}

// Validate validates this encryption status
func (m *EncryptionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncryptionStatus) validateDisks(formats strfmt.Registry) error {

	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {

		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {

			if err := m.Disks[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
