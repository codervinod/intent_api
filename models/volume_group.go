package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VolumeGroup Volume group object
//
// Volume group object
// swagger:model volume_group
type VolumeGroup struct {

	// api version
	// Required: true
	// Read Only: true
	APIVersion string `json:"api_version"`

	// metadata
	// Required: true
	Metadata *VolumeGroupMetadata `json:"metadata"`

	// spec
	Spec *VolumeGroupSpec `json:"spec,omitempty"`

	// status
	Status *VolumeGroupResource `json:"status,omitempty"`
}

// Validate validates this volume group
func (m *VolumeGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeGroup) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.RequiredString("api_version", "body", string(m.APIVersion)); err != nil {
		return err
	}

	return nil
}

func (m *VolumeGroup) validateMetadata(formats strfmt.Registry) error {

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

func (m *VolumeGroup) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {

		if err := m.Spec.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VolumeGroup) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {

		if err := m.Status.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// VolumeGroupSpec volume group spec
// swagger:model VolumeGroupSpec
type VolumeGroupSpec struct {

	// resources
	Resources *VolumeGroupResource `json:"resources,omitempty"`
}

// Validate validates this volume group spec
func (m *VolumeGroupSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeGroupSpec) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {

		if err := m.Resources.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
