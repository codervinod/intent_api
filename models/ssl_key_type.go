package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SslKeyType SSL key type
//
// SSL key type
// swagger:model ssl_key_type
type SslKeyType string

const (
	SslKeyTypeRSA2048  SslKeyType = "RSA_2048"
	SslKeyTypeECDSA256 SslKeyType = "ECDSA_256"
	SslKeyTypeECDSA384 SslKeyType = "ECDSA_384"
)

// for schema
var sslKeyTypeEnum []interface{}

func (m SslKeyType) validateSslKeyTypeEnum(path, location string, value SslKeyType) error {
	if sslKeyTypeEnum == nil {
		var res []SslKeyType
		if err := json.Unmarshal([]byte(`["RSA_2048","ECDSA_256","ECDSA_384"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			sslKeyTypeEnum = append(sslKeyTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, sslKeyTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this ssl key type
func (m SslKeyType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSslKeyTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
