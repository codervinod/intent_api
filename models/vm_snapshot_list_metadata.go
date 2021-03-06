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

// VMSnapshotListMetadata Metadata for vm snapshot list
//
// All api calls that return a list will have this metadata block
//
// swagger:model vm_snapshot_list_metadata
type VMSnapshotListMetadata struct {

	// The filter used for the results
	Filter string `json:"filter,omitempty"`

	// The kind snapshot name
	// Read Only: true
	Kind string `json:"kind,omitempty"`

	// The number of records to retrieve relative to the offset
	Length int64 `json:"length,omitempty"`

	// Offset from the start of the entity list
	Offset int64 `json:"offset,omitempty"`

	// The attribute to perform sort on
	SortColumn string `json:"sort_column,omitempty"`

	// sort order
	SortOrder SortOrder `json:"sort_order,omitempty"`

	// Total matches found
	TotalMatches int64 `json:"total_matches,omitempty"`
}

// Validate validates this vm snapshot list metadata
func (m *VMSnapshotListMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSortOrder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vmSnapshotListMetadataTypeKindPropEnum []interface{}

const (
	vmSnapshotListMetadataKindVMSnapshot string = "vm_snapshot"
)

// prop value enum
func (m *VMSnapshotListMetadata) validateKindEnum(path, location string, value string) error {
	if vmSnapshotListMetadataTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["vm_snapshot"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			vmSnapshotListMetadataTypeKindPropEnum = append(vmSnapshotListMetadataTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, vmSnapshotListMetadataTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VMSnapshotListMetadata) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *VMSnapshotListMetadata) validateSortOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.SortOrder) { // not required
		return nil
	}

	if err := m.SortOrder.Validate(formats); err != nil {
		return err
	}

	return nil
}
