package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DirectoryServiceSearchInput Directory Service search input
//
// Directory Service search input
// swagger:model directory_service_search_input
type DirectoryServiceSearchInput struct {

	// string to be searched
	// Required: true
	Query *string `json:"query"`

	// If set search will only return these attributes in the result else will send the defualt attribute list
	RequestedAttributeList []string `json:"requested_attribute_list,omitempty"`

	// If set search will be performed on these attributes else will be performed on cn (common name)
	SearchAttributeList []string `json:"search_attribute_list,omitempty"`
}

// Validate validates this directory service search input
func (m *DirectoryServiceSearchInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRequestedAttributeList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSearchAttributeList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectoryServiceSearchInput) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

func (m *DirectoryServiceSearchInput) validateRequestedAttributeList(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedAttributeList) { // not required
		return nil
	}

	return nil
}

func (m *DirectoryServiceSearchInput) validateSearchAttributeList(formats strfmt.Registry) error {

	if swag.IsZero(m.SearchAttributeList) { // not required
		return nil
	}

	return nil
}