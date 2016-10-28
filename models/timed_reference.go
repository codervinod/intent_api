package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// TimedReference Timed Reference
//
// A reference that is invalid after the given expiration time.
// swagger:model timed_reference
type TimedReference struct {

	// The time after which the reference is not to be considered valid.
	//
	ExpirationTime strfmt.DateTime `json:"expiration_time,omitempty"`

	// reference
	Reference *Reference `json:"reference,omitempty"`
}

// Validate validates this timed reference
func (m *TimedReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimedReference) validateReference(formats strfmt.Registry) error {

	if swag.IsZero(m.Reference) { // not required
		return nil
	}

	if m.Reference != nil {

		if err := m.Reference.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}