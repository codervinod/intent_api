package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DirectoryServiceDefStatus directory service
//
// directory service configuration
// swagger:model directory_service_def_status
type DirectoryServiceDefStatus struct {

	// directory service name
	// Required: true
	// Max Length: 140
	Name *string `json:"name"`

	// resources
	// Required: true
	Resources *DirectoryServiceDefStatusResources `json:"resources"`
}

// Validate validates this directory service def status
func (m *DirectoryServiceDefStatus) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *DirectoryServiceDefStatus) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *DirectoryServiceDefStatus) validateResources(formats strfmt.Registry) error {

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

// DirectoryServiceDefStatusResources Directory service resources
// swagger:model DirectoryServiceDefStatusResources
type DirectoryServiceDefStatusResources struct {

	// List of users assigned as an admin
	AdminUserReferenceList []*UserReference `json:"admin_user_reference_list,omitempty"`

	// domain name
	// Required: true
	Domain *string `json:"domain"`

	// Url of the directory
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this directory service def status resources
func (m *DirectoryServiceDefStatusResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminUserReferenceList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectoryServiceDefStatusResources) validateAdminUserReferenceList(formats strfmt.Registry) error {

	if swag.IsZero(m.AdminUserReferenceList) { // not required
		return nil
	}

	for i := 0; i < len(m.AdminUserReferenceList); i++ {

		if swag.IsZero(m.AdminUserReferenceList[i]) { // not required
			continue
		}

		if m.AdminUserReferenceList[i] != nil {

			if err := m.AdminUserReferenceList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *DirectoryServiceDefStatusResources) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("resources"+"."+"domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *DirectoryServiceDefStatusResources) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("resources"+"."+"url", "body", m.URL); err != nil {
		return err
	}

	return nil
}