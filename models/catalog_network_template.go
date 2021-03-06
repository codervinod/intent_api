package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CatalogNetworkTemplate Catalog Network template item spec
//
// Network template
// swagger:model catalog_network_template
type CatalogNetworkTemplate struct {

	// network uuid
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	NetworkUUID string `json:"network_uuid,omitempty"`

	// vlan id
	VlanID int32 `json:"vlan_id,omitempty"`
}

// Validate validates this catalog network template
func (m *CatalogNetworkTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogNetworkTemplate) validateNetworkUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkUUID) { // not required
		return nil
	}

	if err := validate.Pattern("network_uuid", "body", string(m.NetworkUUID), `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
		return err
	}

	return nil
}
