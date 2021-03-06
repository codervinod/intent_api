package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GroupsAttributeType Attribute Type
//
// The type of an attribute being used for grouping - may be continuous or discrete
// swagger:model groups_attribute_type
type GroupsAttributeType string

const (
	GroupsAttributeTypeCONTINUOUS GroupsAttributeType = "CONTINUOUS"
	GroupsAttributeTypeDISCRETE   GroupsAttributeType = "DISCRETE"
)

// for schema
var groupsAttributeTypeEnum []interface{}

func (m GroupsAttributeType) validateGroupsAttributeTypeEnum(path, location string, value GroupsAttributeType) error {
	if groupsAttributeTypeEnum == nil {
		var res []GroupsAttributeType
		if err := json.Unmarshal([]byte(`["CONTINUOUS","DISCRETE"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			groupsAttributeTypeEnum = append(groupsAttributeTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, groupsAttributeTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this groups attribute type
func (m GroupsAttributeType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGroupsAttributeTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
