package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// PolicyNetworkEndpoint Network endpoint
//
// Endpoint for network
// swagger:model policy_network_endpoint
type PolicyNetworkEndpoint struct {

	// ip subnet
	IPSubnet *PolicyEndpointIPSubnet `json:"ip_subnet,omitempty"`

	// mac address
	MacAddress string `json:"mac_address,omitempty"`

	// port range
	PortRange *PolicyEndpointPortRange `json:"port_range,omitempty"`
}

// Validate validates this policy network endpoint
func (m *PolicyNetworkEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPSubnet(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePortRange(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyNetworkEndpoint) validateIPSubnet(formats strfmt.Registry) error {

	if swag.IsZero(m.IPSubnet) { // not required
		return nil
	}

	if m.IPSubnet != nil {

		if err := m.IPSubnet.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *PolicyNetworkEndpoint) validatePortRange(formats strfmt.Registry) error {

	if swag.IsZero(m.PortRange) { // not required
		return nil
	}

	if m.PortRange != nil {

		if err := m.PortRange.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
