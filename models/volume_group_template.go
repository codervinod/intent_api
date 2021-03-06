package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VolumeGroupTemplate Template volume_group
//
// Template for volume_group
// swagger:model volume_group_template
type VolumeGroupTemplate struct {

	// metadata
	Metadata *VolumeGroupTemplateMetadata `json:"metadata,omitempty"`

	// spec
	Spec *VolumeGroupTemplateSpec `json:"spec,omitempty"`
}

// Validate validates this volume group template
func (m *VolumeGroupTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeGroupTemplate) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {

		if err := m.Metadata.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VolumeGroupTemplate) validateSpec(formats strfmt.Registry) error {

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

// VolumeGroupTemplateMetadata volume group template metadata
// swagger:model VolumeGroupTemplateMetadata
type VolumeGroupTemplateMetadata struct {

	// kind
	// Read Only: true
	Kind string `json:"kind,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this volume group template metadata
func (m *VolumeGroupTemplateMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var volumeGroupTemplateMetadataTypeKindPropEnum []interface{}

const (
	volumeGroupTemplateMetadataKindVolumeGroupTemplate string = "volume_group_template"
)

// prop value enum
func (m *VolumeGroupTemplateMetadata) validateKindEnum(path, location string, value string) error {
	if volumeGroupTemplateMetadataTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["volume_group_template"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			volumeGroupTemplateMetadataTypeKindPropEnum = append(volumeGroupTemplateMetadataTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, volumeGroupTemplateMetadataTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VolumeGroupTemplateMetadata) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("metadata"+"."+"kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

// VolumeGroupTemplateSpec volume group template spec
// swagger:model VolumeGroupTemplateSpec
type VolumeGroupTemplateSpec struct {

	// name
	// Max Length: 140
	Name string `json:"name,omitempty"`

	// profile reference list
	ProfileReferenceList []*VolumeGroupProfileReference `json:"profile_reference_list,omitempty"`

	// volume group spec
	VolumeGroupSpec *VolumeGroup `json:"volume_group}_spec,omitempty"`
}

// Validate validates this volume group template spec
func (m *VolumeGroupTemplateSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProfileReferenceList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeGroupSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeGroupTemplateSpec) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("spec"+"."+"name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *VolumeGroupTemplateSpec) validateProfileReferenceList(formats strfmt.Registry) error {

	if swag.IsZero(m.ProfileReferenceList) { // not required
		return nil
	}

	for i := 0; i < len(m.ProfileReferenceList); i++ {

		if swag.IsZero(m.ProfileReferenceList[i]) { // not required
			continue
		}

		if m.ProfileReferenceList[i] != nil {

			if err := m.ProfileReferenceList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VolumeGroupTemplateSpec) validateVolumeGroupSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeGroupSpec) { // not required
		return nil
	}

	if m.VolumeGroupSpec != nil {

		if err := m.VolumeGroupSpec.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
