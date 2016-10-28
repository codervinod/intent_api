package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// AvailableExtensionResources Available Extension resources
//
// Available Extension resources
// swagger:model available_extension_resources
type AvailableExtensionResources struct {

	// Company which developed the extension
	// Required: true
	Company *string `json:"company"`

	// Containter image name
	ContainerImageName string `json:"container_image_name,omitempty"`

	// Continater image tag
	ContainerImageTag string `json:"container_image_tag,omitempty"`

	// Port at which container service is running
	ContainerPort int64 `json:"container_port,omitempty"`

	// container spec
	ContainerSpec *ExtensionContainerOptions `json:"container_spec,omitempty"`

	// Description
	// Required: true
	// Max Length: 1000
	Description *string `json:"description"`

	// Document root for landing page of extension
	DocumentRoot string `json:"document_root,omitempty"`

	// Icon image uuid4
	// Required: true
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	IconImageUUID *string `json:"icon_image_uuid"`

	// Internal name of extension
	InternalName string `json:"internal_name,omitempty"`

	// Used only when same instance of app is used for all users
	MultiUserApp *bool `json:"multi_user_app,omitempty"`

	// Name of extension
	// Required: true
	// Max Length: 140
	Name *string `json:"name"`

	// list of oauth scopes
	OauthScopes []string `json:"oauth_scopes,omitempty"`

	// UUID for registry where container image is stored
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	RegistryUUID string `json:"registry_uuid,omitempty"`

	// Screen shot image uuid list
	ScreenShotUuids []string `json:"screen_shot_uuids,omitempty"`

	// Version string
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this available extension resources
func (m *AvailableExtensionResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompany(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainerSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIconImageUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOauthScopes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegistryUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScreenShotUuids(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailableExtensionResources) validateCompany(formats strfmt.Registry) error {

	if err := validate.Required("company", "body", m.Company); err != nil {
		return err
	}

	return nil
}

func (m *AvailableExtensionResources) validateContainerSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainerSpec) { // not required
		return nil
	}

	if m.ContainerSpec != nil {

		if err := m.ContainerSpec.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *AvailableExtensionResources) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *AvailableExtensionResources) validateIconImageUUID(formats strfmt.Registry) error {

	if err := validate.Required("icon_image_uuid", "body", m.IconImageUUID); err != nil {
		return err
	}

	if err := validate.Pattern("icon_image_uuid", "body", string(*m.IconImageUUID), `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableExtensionResources) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *AvailableExtensionResources) validateOauthScopes(formats strfmt.Registry) error {

	if swag.IsZero(m.OauthScopes) { // not required
		return nil
	}

	return nil
}

func (m *AvailableExtensionResources) validateRegistryUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.RegistryUUID) { // not required
		return nil
	}

	if err := validate.Pattern("registry_uuid", "body", string(m.RegistryUUID), `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
		return err
	}

	return nil
}

func (m *AvailableExtensionResources) validateScreenShotUuids(formats strfmt.Registry) error {

	if swag.IsZero(m.ScreenShotUuids) { // not required
		return nil
	}

	for i := 0; i < len(m.ScreenShotUuids); i++ {

		if err := validate.Pattern("screen_shot_uuids"+"."+strconv.Itoa(i), "body", string(m.ScreenShotUuids[i]), `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *AvailableExtensionResources) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}