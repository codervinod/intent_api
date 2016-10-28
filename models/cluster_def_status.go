package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ClusterDefStatus Cluster Status
//
// Cluster status definition
// swagger:model cluster_def_status
type ClusterDefStatus struct {

	// Cluster Name
	// Max Length: 140
	Name string `json:"name,omitempty"`

	// policies
	Policies Policies `json:"policies,omitempty"`

	// providers
	Providers Providers `json:"providers,omitempty"`

	// resources
	// Required: true
	Resources *ClusterDefStatusResources `json:"resources"`

	// Cluster UUID
	// Required: true
	UUID *string `json:"uuid"`
}

// Validate validates this cluster def status
func (m *ClusterDefStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDefStatus) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *ClusterDefStatus) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	if m.Resources != nil {

		if err := m.Resources.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ClusterDefStatus) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// ClusterDefStatusResources Cluster resources
// swagger:model ClusterDefStatusResources
type ClusterDefStatusResources struct {

	// config
	// Required: true
	Config *ClusterConfig `json:"config"`

	// network
	// Required: true
	Network *ClusterNetwork `json:"network"`

	// nodes
	// Required: true
	Nodes *ClusterNodes `json:"nodes"`

	// Cluster current attributes and onging operations
	RuntimeStatus []string `json:"runtime_status,omitempty"`

	// Stats data map
	// Read Only: true
	Stats map[string]int32 `json:"stats,omitempty"`
}

// Validate validates this cluster def status resources
func (m *ClusterDefStatusResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRuntimeStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDefStatusResources) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("resources"+"."+"config", "body", m.Config); err != nil {
		return err
	}

	if m.Config != nil {

		if err := m.Config.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ClusterDefStatusResources) validateNetwork(formats strfmt.Registry) error {

	if err := validate.Required("resources"+"."+"network", "body", m.Network); err != nil {
		return err
	}

	if m.Network != nil {

		if err := m.Network.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ClusterDefStatusResources) validateNodes(formats strfmt.Registry) error {

	if err := validate.Required("resources"+"."+"nodes", "body", m.Nodes); err != nil {
		return err
	}

	if m.Nodes != nil {

		if err := m.Nodes.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var clusterDefStatusResourcesRuntimeStatusItemsEnum []interface{}

func (m *ClusterDefStatusResources) validateRuntimeStatusItemsEnum(path, location string, value string) error {
	if clusterDefStatusResourcesRuntimeStatusItemsEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["UPGRADE_IN_PROGRESS","HAVE_SELF_ENCRYPTING_DRIVE"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			clusterDefStatusResourcesRuntimeStatusItemsEnum = append(clusterDefStatusResourcesRuntimeStatusItemsEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, clusterDefStatusResourcesRuntimeStatusItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterDefStatusResources) validateRuntimeStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.RuntimeStatus) { // not required
		return nil
	}

	for i := 0; i < len(m.RuntimeStatus); i++ {

		// value enum
		if err := m.validateRuntimeStatusItemsEnum("resources"+"."+"runtime_status"+"."+strconv.Itoa(i), "body", m.RuntimeStatus[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ClusterDefStatusResources) validateStats(formats strfmt.Registry) error {

	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if err := validate.Required("resources"+"."+"stats", "body", m.Stats); err != nil {
		return err
	}

	return nil
}
