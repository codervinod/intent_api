package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JobIntentInput job Intent Entity
//
// An intentful representation of a job
// swagger:model job_intent_input
type JobIntentInput struct {

	// api version
	// Required: true
	APIVersion *string `json:"api_version"`

	// metadata
	// Required: true
	Metadata *JobMetadata `json:"metadata"`

	// options
	Options Options `json:"options,omitempty"`

	// spec
	Spec *Job `json:"spec,omitempty"`
}

// Validate validates this job intent input
func (m *JobIntentInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobIntentInput) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.Required("api_version", "body", m.APIVersion); err != nil {
		return err
	}

	return nil
}

func (m *JobIntentInput) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {

		if err := m.Metadata.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *JobIntentInput) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {

		if err := m.Spec.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
