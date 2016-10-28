package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ResourceDomainDetails Resource Domain Status
//
// the intentful specification and status for resource domains
// swagger:model resource_domain_details
type ResourceDomainDetails struct {

	// the utilization/limit for resource types
	// Required: true
	Resources []*ResourceUtilizationStatus `json:"resources"`

	// the uuids of the domains charges should be rolled up to
	// Required: true
	RollupDomainReferenceList []*ResourceDomainReference `json:"rollup_domain_reference_list"`
}

// Validate validates this resource domain details
func (m *ResourceDomainDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRollupDomainReferenceList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceDomainDetails) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	for i := 0; i < len(m.Resources); i++ {

		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {

			if err := m.Resources[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ResourceDomainDetails) validateRollupDomainReferenceList(formats strfmt.Registry) error {

	if err := validate.Required("rollup_domain_reference_list", "body", m.RollupDomainReferenceList); err != nil {
		return err
	}

	for i := 0; i < len(m.RollupDomainReferenceList); i++ {

		if swag.IsZero(m.RollupDomainReferenceList[i]) { // not required
			continue
		}

		if m.RollupDomainReferenceList[i] != nil {

			if err := m.RollupDomainReferenceList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
