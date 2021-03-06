package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// ExecuteWorkflows Workflows to be executed for this stage
//
// Workflows to be executed for this stage
// swagger:model execute_workflows
type ExecuteWorkflows struct {

	// List of workflows
	WorkflowsList []string `json:"workflows_list,omitempty"`
}

// Validate validates this execute workflows
func (m *ExecuteWorkflows) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkflowsList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecuteWorkflows) validateWorkflowsList(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkflowsList) { // not required
		return nil
	}

	return nil
}
