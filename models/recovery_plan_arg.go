package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// RecoveryPlanArg Arguments for recovery plan
//
// Arguments for recovery plan
// swagger:model recovery_plan_arg
type RecoveryPlanArg struct {

	// Network information to be used for recovery plan
	NetworkInfo *RecoveryPlanNetworkInfo `json:"network_info,omitempty"`

	// Profile information to be used for recovery plan
	ProfileInfo *RecoveryPlanProfileInfo `json:"profile_info,omitempty"`

	// Information about protected site for recovery plan
	ProtectedSiteInfo *RecoveryPlanSiteInfo `json:"protected_site_info,omitempty"`

	// Information about recovery site for recovery plan
	RecoverySiteInfo *RecoveryPlanSiteInfo `json:"recovery_site_info,omitempty"`
}

// Validate validates this recovery plan arg
func (m *RecoveryPlanArg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProfileInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProtectedSiteInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecoverySiteInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanArg) validateNetworkInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkInfo) { // not required
		return nil
	}

	if m.NetworkInfo != nil {

		if err := m.NetworkInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *RecoveryPlanArg) validateProfileInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ProfileInfo) { // not required
		return nil
	}

	if m.ProfileInfo != nil {

		if err := m.ProfileInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *RecoveryPlanArg) validateProtectedSiteInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ProtectedSiteInfo) { // not required
		return nil
	}

	if m.ProtectedSiteInfo != nil {

		if err := m.ProtectedSiteInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *RecoveryPlanArg) validateRecoverySiteInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.RecoverySiteInfo) { // not required
		return nil
	}

	if m.RecoverySiteInfo != nil {

		if err := m.RecoverySiteInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}