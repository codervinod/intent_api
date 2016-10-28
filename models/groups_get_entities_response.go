package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GroupsGetEntitiesResponse Get Entities Response
//
// Get Entities Response
// swagger:model groups_get_entities_response
type GroupsGetEntitiesResponse struct {

	// entity type
	EntityType string `json:"entity_type,omitempty"`

	// filtered entity count
	FilteredEntityCount int64 `json:"filtered_entity_count,omitempty"`

	// group results
	GroupResults []*GroupsGroupResult `json:"group_results,omitempty"`

	// total entity count
	TotalEntityCount int64 `json:"total_entity_count,omitempty"`

	// total group count
	TotalGroupCount int64 `json:"total_group_count,omitempty"`
}

// Validate validates this groups get entities response
func (m *GroupsGetEntitiesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupResults(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupsGetEntitiesResponse) validateGroupResults(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupResults) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupResults); i++ {

		if swag.IsZero(m.GroupResults[i]) { // not required
			continue
		}

		if m.GroupResults[i] != nil {

			if err := m.GroupResults[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
