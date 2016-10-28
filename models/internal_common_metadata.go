package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// InternalCommonMetadata Common Metadata
//
// The common set of metadata attributes for all entity kinds.
// swagger:model internal_common_metadata
type InternalCommonMetadata struct {

	// The date and time when the object was created.
	// Read Only: true
	CreationTime strfmt.DateTime `json:"creation_time,omitempty"`

	// The date and time when the object was deleted. This field
	// is only updated by the system as a side-effect of a delete operation.
	// Policies indicating "expiration_time" of objects should be
	// expressed in the "spec" section so that they can be updated by the
	// user.
	//
	// Read Only: true
	DeletionTime strfmt.DateTime `json:"deletion_time,omitempty"`

	// Monotonically increasing version of the "spec".
	// Read Only: true
	EntityVersion int64 `json:"entity_version,omitempty"`

	// The entitys kind
	// Read Only: true
	Kind string `json:"kind,omitempty"`

	// The date and time when the object was last updated.
	// Read Only: true
	LastUpdateTime strfmt.DateTime `json:"last_update_time,omitempty"`

	// The relative url that produced this result.
	// Read Only: true
	Rel string `json:"rel,omitempty"`

	// The UUID of the entity.
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this internal common metadata
func (m *InternalCommonMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}