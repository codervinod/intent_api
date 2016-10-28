package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GroupsFieldData Group Summary Data
//
// Group Summary Data
// swagger:model groups_field_data
type GroupsFieldData struct {

	// buckets
	Buckets GroupsBucketSummaryMap `json:"buckets,omitempty"`

	// name
	// Max Length: 140
	Name string `json:"name,omitempty"`

	// values
	Values []*GroupsTimevaluePair `json:"values,omitempty"`
}

// Validate validates this groups field data
func (m *GroupsFieldData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupsFieldData) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *GroupsFieldData) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {

		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {

			if err := m.Values[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}