package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// WorkflowListMetadata Metadata for workflow list
//
// All api calls that return a list will have this metadata block
//
// swagger:model workflow_list_metadata
type WorkflowListMetadata struct {

	// The filter used for the results
	Filter string `json:"filter,omitempty"`

	// The kind name
	// Required: true
	// Read Only: true
	Kind string `json:"kind"`

	// The number of records to retrieve relative to the offset
	Length int64 `json:"length,omitempty"`

	// Offset from the start of the entity list
	Offset int64 `json:"offset,omitempty"`

	// The attribute to perform sort on
	SortColumn string `json:"sort_column,omitempty"`

	// sort order
	SortOrder SortOrder `json:"sort_order,omitempty"`

	// Total matches found for the kind
	TotalMatches int64 `json:"total_matches,omitempty"`
}

// Validate validates this workflow list metadata
func (m *WorkflowListMetadata) Validate(formats strfmt.Registry) error {
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

var workflowListMetadataTypeKindPropEnum []interface{}

const (
	workflowListMetadataKindWorkflow string = "workflow"
)

// prop value enum
func (m *WorkflowListMetadata) validateKindEnum(path, location string, value string) error {
	if workflowListMetadataTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["workflow"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			workflowListMetadataTypeKindPropEnum = append(workflowListMetadataTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, workflowListMetadataTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowListMetadata) validateKind(formats strfmt.Registry) error {

	if err := validate.RequiredString("kind", "body", string(m.Kind)); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowListMetadata) validateSortOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.SortOrder) { // not required
		return nil
	}

	if err := m.SortOrder.Validate(formats); err != nil {
		return err
	}

	return nil
}
