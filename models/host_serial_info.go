package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// HostSerialInfo Host serial numbers
//
// Host serial information
// swagger:model host_serial_info
type HostSerialInfo struct {

	// Rackable unit model name
	BlockModel string `json:"block_model,omitempty"`

	// Rackable unit serial number
	BlockSerial string `json:"block_serial,omitempty"`

	// Node serial number
	NodeSerial string `json:"node_serial,omitempty"`
}

// Validate validates this host serial info
func (m *HostSerialInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
