package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// WorkflowDefStatus Workflow Status Definition
//
// Current status of workflow
// swagger:model workflow_def_status
type WorkflowDefStatus struct {

	// A description or user annotation for the Workflow
	// Max Length: 1000
	Description string `json:"description,omitempty"`

	// Workflow Name
	// Required: true
	// Max Length: 140
	Name *string `json:"name"`

	// resources
	// Required: true
	Resources *WorkflowDefStatusResources `json:"resources"`
}

// Validate validates this workflow def status
func (m *WorkflowDefStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowDefStatus) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowDefStatus) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowDefStatus) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	if m.Resources != nil {

		if err := m.Resources.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowDefStatusResources Workflow resources
// swagger:model WorkflowDefStatusResources
type WorkflowDefStatusResources struct {

	// Jobs associated with the workflow.
	Jobs []*JobInfo `json:"jobs,omitempty"`

	// workflow state
	WorkflowState *Workflow `json:"workflow_state,omitempty"`
}

// Validate validates this workflow def status resources
func (m *WorkflowDefStatusResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWorkflowState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowDefStatusResources) validateJobs(formats strfmt.Registry) error {

	if swag.IsZero(m.Jobs) { // not required
		return nil
	}

	for i := 0; i < len(m.Jobs); i++ {

		if swag.IsZero(m.Jobs[i]) { // not required
			continue
		}

		if m.Jobs[i] != nil {

			if err := m.Jobs[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowDefStatusResources) validateWorkflowState(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkflowState) { // not required
		return nil
	}

	if m.WorkflowState != nil {

		if err := m.WorkflowState.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
