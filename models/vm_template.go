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

// VMTemplate Template vm
//
// Template for vm
// swagger:model vm_template
type VMTemplate struct {

	// metadata
	Metadata *VMTemplateMetadata `json:"metadata,omitempty"`

	// spec
	Spec *VMTemplateSpec `json:"spec,omitempty"`
}

// Validate validates this vm template
func (m *VMTemplate) Validate(formats strfmt.Registry) error {
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

func (m *VMTemplate) validateMetadata(formats strfmt.Registry) error {

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

func (m *VMTemplate) validateSpec(formats strfmt.Registry) error {

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

// VMTemplateMetadata VM template metadata
// swagger:model VMTemplateMetadata
type VMTemplateMetadata struct {

	// kind
	// Read Only: true
	Kind string `json:"kind,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this VM template metadata
func (m *VMTemplateMetadata) Validate(formats strfmt.Registry) error {
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

var vmTemplateMetadataTypeKindPropEnum []interface{}

const (
	vmTemplateMetadataKindVMTemplate string = "vm_template"
)

// prop value enum
func (m *VMTemplateMetadata) validateKindEnum(path, location string, value string) error {
	if vmTemplateMetadataTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["vm_template"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			vmTemplateMetadataTypeKindPropEnum = append(vmTemplateMetadataTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, vmTemplateMetadataTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VMTemplateMetadata) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("metadata"+"."+"kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

// VMTemplateSpec VM template spec
// swagger:model VMTemplateSpec
type VMTemplateSpec struct {

	// name
	// Max Length: 140
	Name string `json:"name,omitempty"`

	// profile reference list
	ProfileReferenceList []*VMProfileReference `json:"profile_reference_list,omitempty"`

	// vm spec
	VMSpec *VM `json:"vm}_spec,omitempty"`
}

// Validate validates this VM template spec
func (m *VMTemplateSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProfileReferenceList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVMSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplateSpec) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("spec"+"."+"name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateSpec) validateProfileReferenceList(formats strfmt.Registry) error {

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

func (m *VMTemplateSpec) validateVMSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.VMSpec) { // not required
		return nil
	}

	if m.VMSpec != nil {

		if err := m.VMSpec.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
