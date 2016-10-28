package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JobInfo Information about a job associated workflow
//
// Information about a job associated workflow
// swagger:model job_info
type JobInfo struct {

	// UUID of a job associated with this workflow.
	// Required: true
	JobUUID *string `json:"job_uuid"`
}

// Validate validates this job info
func (m *JobInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobInfo) validateJobUUID(formats strfmt.Registry) error {

	if err := validate.Required("job_uuid", "body", m.JobUUID); err != nil {
		return err
	}

	return nil
}