package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CatalogImageTemplate Catalog Image template item spec
//
// Image template
// swagger:model catalog_image_template
type CatalogImageTemplate struct {

	// image type
	ImageType CatalogImageType `json:"image_type,omitempty"`

	// image uuid
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	ImageUUID string `json:"image_uuid,omitempty"`
}

// Validate validates this catalog image template
func (m *CatalogImageTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImageUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogImageTemplate) validateImageType(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageType) { // not required
		return nil
	}

	if err := m.ImageType.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *CatalogImageTemplate) validateImageUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageUUID) { // not required
		return nil
	}

	if err := validate.Pattern("image_uuid", "body", string(m.ImageUUID), `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
		return err
	}

	return nil
}
