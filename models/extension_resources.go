package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ExtensionResources Extension resources
//
// Extension resources
// swagger:model extension_resources
type ExtensionResources struct {

	// available extension ref
	// Required: true
	AvailableExtensionRef *AvailableExtensionResources `json:"available_extension_ref"`

	// Container host address
	ContainerHost string `json:"container_host,omitempty"`

	// State of extension container
	ContainerState string `json:"container_state,omitempty"`

	// State of extension
	// Required: true
	State *string `json:"state"`
}

// Validate validates this extension resources
func (m *ExtensionResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableExtensionRef(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainerState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionResources) validateAvailableExtensionRef(formats strfmt.Registry) error {

	if err := validate.Required("available_extension_ref", "body", m.AvailableExtensionRef); err != nil {
		return err
	}

	if m.AvailableExtensionRef != nil {

		if err := m.AvailableExtensionRef.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var extensionResourcesTypeContainerStatePropEnum []interface{}

const (
	extensionResourcesContainerStateQUEUED          string = "QUEUED"
	extensionResourcesContainerStateIMAGEDOWNLOADED string = "IMAGE_DOWNLOADED"
	extensionResourcesContainerStateCREATED         string = "CREATED"
	extensionResourcesContainerStateRUNNING         string = "RUNNING"
	extensionResourcesContainerStateSTOPPED         string = "STOPPED"
	extensionResourcesContainerStateKILLED          string = "KILLED"
	extensionResourcesContainerStateREMOVED         string = "REMOVED"
	extensionResourcesContainerStateDELETED         string = "DELETED"
	extensionResourcesContainerStatePAUSED          string = "PAUSED"
)

// prop value enum
func (m *ExtensionResources) validateContainerStateEnum(path, location string, value string) error {
	if extensionResourcesTypeContainerStatePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["QUEUED","IMAGE_DOWNLOADED","CREATED","RUNNING","STOPPED","KILLED","REMOVED","DELETED","PAUSED"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			extensionResourcesTypeContainerStatePropEnum = append(extensionResourcesTypeContainerStatePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, extensionResourcesTypeContainerStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ExtensionResources) validateContainerState(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainerState) { // not required
		return nil
	}

	// value enum
	if err := m.validateContainerStateEnum("container_state", "body", m.ContainerState); err != nil {
		return err
	}

	return nil
}

var extensionResourcesTypeStatePropEnum []interface{}

const (
	extensionResourcesStateINSTALLING  string = "INSTALLING"
	extensionResourcesStateRUNNING     string = "RUNNING"
	extensionResourcesStateSTOPPED     string = "STOPPED"
	extensionResourcesStateUNINSTALLED string = "UNINSTALLED"
)

// prop value enum
func (m *ExtensionResources) validateStateEnum(path, location string, value string) error {
	if extensionResourcesTypeStatePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["INSTALLING","RUNNING","STOPPED","UNINSTALLED"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			extensionResourcesTypeStatePropEnum = append(extensionResourcesTypeStatePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, extensionResourcesTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ExtensionResources) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}
