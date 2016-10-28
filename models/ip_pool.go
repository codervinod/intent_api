package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// IPPool IP pool
//
// IP pool
// swagger:model ip_pool
type IPPool struct {

	// range
	Range string `json:"range,omitempty"`
}

// Validate validates this ip pool
func (m *IPPool) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
