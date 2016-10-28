package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ExecutionPlan Execution plan for multiple requests
//
// Execution plan for multiple requests
// swagger:model execution_plan
type ExecutionPlan string

const (
	ExecutionPlanSEQUENTIAL ExecutionPlan = "SEQUENTIAL"
	ExecutionPlanPARALLEL   ExecutionPlan = "PARALLEL"
)

// for schema
var executionPlanEnum []interface{}

func (m ExecutionPlan) validateExecutionPlanEnum(path, location string, value ExecutionPlan) error {
	if executionPlanEnum == nil {
		var res []ExecutionPlan
		if err := json.Unmarshal([]byte(`["SEQUENTIAL","PARALLEL"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			executionPlanEnum = append(executionPlanEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, executionPlanEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this execution plan
func (m ExecutionPlan) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateExecutionPlanEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}