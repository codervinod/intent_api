package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CaCert CA certificate info
//
// CA certificate info
// swagger:model ca_cert
type CaCert struct {

	// Name of the certificate authority
	// Required: true
	CaName *string `json:"ca_name"`

	// Certificate content
	// Required: true
	Certificate *strfmt.Base64 `json:"certificate"`
}

// Validate validates this ca cert
func (m *CaCert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCertificate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CaCert) validateCaName(formats strfmt.Registry) error {

	if err := validate.Required("ca_name", "body", m.CaName); err != nil {
		return err
	}

	return nil
}

func (m *CaCert) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}
