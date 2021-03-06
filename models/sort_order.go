package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SortOrder Sort order
//
// The sort order in which results are returned
// swagger:model sort_order
type SortOrder string

const (
	SortOrderASCENDING  SortOrder = "ASCENDING"
	SortOrderDESCENDING SortOrder = "DESCENDING"
)

// for schema
var sortOrderEnum []interface{}

func (m SortOrder) validateSortOrderEnum(path, location string, value SortOrder) error {
	if sortOrderEnum == nil {
		var res []SortOrder
		if err := json.Unmarshal([]byte(`["ASCENDING","DESCENDING"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			sortOrderEnum = append(sortOrderEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, sortOrderEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this sort order
func (m SortOrder) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSortOrderEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
