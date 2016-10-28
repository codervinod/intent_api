package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ChangedRegionsQuery Changed Regions Query
//
// Instances of this type are used to specify the input for the changed regions query.
//
// swagger:model changed_regions_query
type ChangedRegionsQuery struct {

	// The absolute offset in bytes up to which to query for the changed regions. Note that the interval specified by the start_offset together with the end_offset is right half-open. If the end_offset is not specified, the portion from the start_offset till the end of the file will be included in the query.
	//
	// Minimum: 0
	EndOffset *int64 `json:"end_offset,omitempty"`

	// Absolute path of a file within a snapshot that must be used as the reference in the computation of the changed regions. If this path is not specified, then the changed regions will not be computed. Instead, the sparse and the non-sparse regions of the file specified in snapshot_file_path will be returned.
	//
	ReferenceSnapshotFilePath string `json:"reference_snapshot_file_path,omitempty"`

	// Absolute path of a file within a snapshot of an entity such as a virtual machine, a volume group or a protection domain.
	//
	// Required: true
	SnapshotFilePath *string `json:"snapshot_file_path"`

	// The absolute offset in bytes from where to query for the changed regions.
	//
	// Minimum: 0
	StartOffset *int64 `json:"start_offset,omitempty"`
}

// Validate validates this changed regions query
func (m *ChangedRegionsQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndOffset(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSnapshotFilePath(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartOffset(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangedRegionsQuery) validateEndOffset(formats strfmt.Registry) error {

	if swag.IsZero(m.EndOffset) { // not required
		return nil
	}

	if err := validate.MinimumInt("end_offset", "body", int64(*m.EndOffset), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChangedRegionsQuery) validateSnapshotFilePath(formats strfmt.Registry) error {

	if err := validate.Required("snapshot_file_path", "body", m.SnapshotFilePath); err != nil {
		return err
	}

	return nil
}

func (m *ChangedRegionsQuery) validateStartOffset(formats strfmt.Registry) error {

	if swag.IsZero(m.StartOffset) { // not required
		return nil
	}

	if err := validate.MinimumInt("start_offset", "body", int64(*m.StartOffset), 0, false); err != nil {
		return err
	}

	return nil
}