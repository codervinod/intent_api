package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// CatalogTemplateItem Catalog template item spec
//
// Template item
// swagger:model catalog_template_item
type CatalogTemplateItem struct {

	// image template
	ImageTemplate *CatalogImageTemplate `json:"image_template,omitempty"`

	// metadata store
	MetadataStore CatalogMetadataStore `json:"metadata_store,omitempty"`

	// network template
	NetworkTemplate *CatalogNetworkTemplate `json:"network_template,omitempty"`

	// storage template
	StorageTemplate *CatalogStorageTemplate `json:"storage_template,omitempty"`

	// vm template
	VMTemplate *CatalogVMTemplate `json:"vm_template,omitempty"`

	// volume template
	VolumeTemplate *CatalogVolumeTemplate `json:"volume_template,omitempty"`
}

// Validate validates this catalog template item
func (m *CatalogTemplateItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageTemplate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworkTemplate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStorageTemplate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVMTemplate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeTemplate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogTemplateItem) validateImageTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageTemplate) { // not required
		return nil
	}

	if m.ImageTemplate != nil {

		if err := m.ImageTemplate.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CatalogTemplateItem) validateNetworkTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkTemplate) { // not required
		return nil
	}

	if m.NetworkTemplate != nil {

		if err := m.NetworkTemplate.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CatalogTemplateItem) validateStorageTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.StorageTemplate) { // not required
		return nil
	}

	if m.StorageTemplate != nil {

		if err := m.StorageTemplate.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CatalogTemplateItem) validateVMTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.VMTemplate) { // not required
		return nil
	}

	if m.VMTemplate != nil {

		if err := m.VMTemplate.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CatalogTemplateItem) validateVolumeTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeTemplate) { // not required
		return nil
	}

	if m.VolumeTemplate != nil {

		if err := m.VolumeTemplate.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
