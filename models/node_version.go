package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// NodeVersion Node Version
//
// Cluster service node version
// swagger:model node_version
type NodeVersion struct {

	// Current NCC version for the cluster. Each of the nodes in the cluster
	// should be at this NCC version or upgrading to this version
	//
	NccVersion string `json:"ncc_version,omitempty"`

	// nos full version
	NosFullVersion string `json:"nos_full_version,omitempty"`

	// nos version
	NosVersion string `json:"nos_version,omitempty"`
}

// Validate validates this node version
func (m *NodeVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
