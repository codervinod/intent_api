package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// StageProgress Progress of a stage of a workflow
//
// Progress of a stage of a workflow
// swagger:model stage_progress
type StageProgress struct {

	// execution status
	// Required: true
	ExecutionStatus *ExecutionStatus `json:"execution_status"`

	// Tasks to be done for this Stage
	Tasks []*StepProgress `json:"tasks,omitempty"`
}

// Validate validates this stage progress
func (m *StageProgress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutionStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StageProgress) validateExecutionStatus(formats strfmt.Registry) error {

	if err := validate.Required("execution_status", "body", m.ExecutionStatus); err != nil {
		return err
	}

	if m.ExecutionStatus != nil {

		if err := m.ExecutionStatus.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *StageProgress) validateTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {

		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {

			if err := m.Tasks[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
