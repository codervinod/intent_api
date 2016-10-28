package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JobListIntentResponse Entity Intent List Response
//
// Response object for intentful operation of jobs
// swagger:model job_list_intent_response
type JobListIntentResponse struct {

	// api version
	// Required: true
	APIVersion *string `json:"api_version"`

	// entities
	Entities []*JobIntentResource `json:"entities,omitempty"`

	// metadata
	// Required: true
	Metadata *JobListMetadata `json:"metadata"`
}

// Validate validates this job list intent response
func (m *JobListIntentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEntities(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobListIntentResponse) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.Required("api_version", "body", m.APIVersion); err != nil {
		return err
	}

	return nil
}

func (m *JobListIntentResponse) validateEntities(formats strfmt.Registry) error {

	if swag.IsZero(m.Entities) { // not required
		return nil
	}

	for i := 0; i < len(m.Entities); i++ {

		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {

			if err := m.Entities[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *JobListIntentResponse) validateMetadata(formats strfmt.Registry) error {

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
