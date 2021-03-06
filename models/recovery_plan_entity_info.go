package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// RecoveryPlanEntityInfo Information about entity to be recovered
//
// Information about entity to be recovered
// swagger:model recovery_plan_entity_info
type RecoveryPlanEntityInfo struct {

	// Expected state of entity after recovery
	EntityConfigUpdateInfo *RecoveryPlanEntityConfigUpdateInfo `json:"entity_config_update_info,omitempty"`

	// Name of entity
	// Max Length: 140
	EntityName string `json:"entity_name,omitempty"`

	// Type of entity
	EntityType string `json:"entity_type,omitempty"`

	// Uuid of entity
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	EntityUUID string `json:"entity_uuid,omitempty"`

	// Uuid of the cluster where entity is present
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	SourceClusterUUID string `json:"source_cluster_uuid,omitempty"`

	// Uuid of the cluster where entity has to be recovered
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	TargetClusterUUID string `json:"target_cluster_uuid,omitempty"`
}

// Validate validates this recovery plan entity info
func (m *RecoveryPlanEntityInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityConfigUpdateInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEntityName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEntityUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSourceClusterUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTargetClusterUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanEntityInfo) validateEntityConfigUpdateInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityConfigUpdateInfo) { // not required
		return nil
	}

	if m.EntityConfigUpdateInfo != nil {

		if err := m.EntityConfigUpdateInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *RecoveryPlanEntityInfo) validateEntityName(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityName) { // not required
		return nil
	}

	if err := validate.MaxLength("entity_name", "body", string(m.EntityName), 140); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryPlanEntityInfo) validateEntityUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityUUID) { // not required
		return nil
	}

	if err := validate.Pattern("entity_uuid", "body", string(m.EntityUUID), `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryPlanEntityInfo) validateSourceClusterUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceClusterUUID) { // not required
		return nil
	}

	if err := validate.Pattern("source_cluster_uuid", "body", string(m.SourceClusterUUID), `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryPlanEntityInfo) validateTargetClusterUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetClusterUUID) { // not required
		return nil
	}

	if err := validate.Pattern("target_cluster_uuid", "body", string(m.TargetClusterUUID), `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
		return err
	}

	return nil
}
