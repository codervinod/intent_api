package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ClusterStatus Response Kind
//
// The status of a REST API call. Only used when there is a failure to
// report.
//
// swagger:model cluster_status
type ClusterStatus struct {

	// api version
	// Read Only: true
	APIVersion string `json:"api_version,omitempty"`

	// The HTTP error code.
	// Read Only: true
	Code int64 `json:"code,omitempty"`

	// Custom key-value details relevant to the status.
	// Read Only: true
	Details map[string]string `json:"details,omitempty"`

	// The kind name
	// Read Only: true
	Kind string `json:"kind,omitempty"`

	// A sentence explaining the reason for the status.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// One snake_case word.
	// Read Only: true
	Reason string `json:"reason,omitempty"`

	// Only value possible is "failure".
	// Read Only: true
	Status string `json:"status,omitempty"`
}

// Validate validates this cluster status
func (m *ClusterStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterStatus) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if err := validate.Required("details", "body", m.Details); err != nil {
		return err
	}

	return nil
}

var clusterStatusTypeKindPropEnum []interface{}

const (
	clusterStatusKindCluster string = "cluster"
)

// prop value enum
func (m *ClusterStatus) validateKindEnum(path, location string, value string) error {
	if clusterStatusTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["cluster"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			clusterStatusTypeKindPropEnum = append(clusterStatusTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, clusterStatusTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterStatus) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}
