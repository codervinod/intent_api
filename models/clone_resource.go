package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// CloneResource Creates new disk based on existing specification
//
// Creates new disk based on existing specification
// swagger:model clone_resource
type CloneResource struct {

	// Path of the image to clone from. If this is specified, then vmdisk_uuid and ndfs_filepath should not be set.
	//
	ImagePath string `json:"image_path,omitempty"`

	// The minimum size of the resulting clone in MiB. This should only be specified if vmdisk_uuid is set.
	//
	MinimumSizeMb int64 `json:"minimum_size_mb,omitempty"`

	// Path of the image to clone from. If this is specified, then vmdisk_uuid and image_path should not be set.
	//
	NdfsFilepath string `json:"ndfs_filepath,omitempty"`

	// UUID of the disk to clone from. If this is specified, then image_path and ndfs_filepath should not be set.
	//
	VmdiskUUID string `json:"vmdisk_uuid,omitempty"`
}

// Validate validates this clone resource
func (m *CloneResource) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
