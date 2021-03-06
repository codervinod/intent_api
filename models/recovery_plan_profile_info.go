package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// RecoveryPlanProfileInfo Description of profile information for recovery plan
//
// Profile information for recovery plan
// swagger:model recovery_plan_profile_info
type RecoveryPlanProfileInfo struct {

	// Default profile to be used on recovery site
	DefaultProfile string `json:"default_profile,omitempty"`

	// profile mapping
	ProfileMapping []*ProfileMapping `json:"profile_mapping,omitempty"`
}

// Validate validates this recovery plan profile info
func (m *RecoveryPlanProfileInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProfileMapping(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanProfileInfo) validateProfileMapping(formats strfmt.Registry) error {

	if swag.IsZero(m.ProfileMapping) { // not required
		return nil
	}

	for i := 0; i < len(m.ProfileMapping); i++ {

		if swag.IsZero(m.ProfileMapping[i]) { // not required
			continue
		}

		if m.ProfileMapping[i] != nil {

			if err := m.ProfileMapping[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
