package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// RecoveryPlanSiteInfo Description of sites involved in recovery plan
//
// Information about sites involved in recovery plan
// swagger:model recovery_plan_site_info
type RecoveryPlanSiteInfo struct {

	// Type of management entity. It can be another PC or Xi.
	ManagementEntityType string `json:"management_entity_type,omitempty"`

	// Identifier of management plane
	ManagementIdentifier string `json:"management_identifier,omitempty"`

	// Recovery site information
	// Max Length: 140
	Name string `json:"name,omitempty"`
}

// Validate validates this recovery plan site info
func (m *RecoveryPlanSiteInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanSiteInfo) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}
