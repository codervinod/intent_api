package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CaChainSpec CA chain spec
//
// CA chain spec
// swagger:model ca_chain_spec
type CaChainSpec struct {

	// Content of CA chain
	// Required: true
	CaChain *strfmt.Base64 `json:"ca_chain"`

	// The name of the CA Chain file
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this ca chain spec
func (m *CaChainSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaChain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CaChainSpec) validateCaChain(formats strfmt.Registry) error {

	if err := validate.Required("ca_chain", "body", m.CaChain); err != nil {
		return err
	}

	return nil
}

func (m *CaChainSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
