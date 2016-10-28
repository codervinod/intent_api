package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ContainerResources Container Resources Definition
//
// Container Resources Definition
// swagger:model container_resources
type ContainerResources struct {

	// container basic options
	ContainerBasicOptions *ContainerBasicOptions `json:"container_basic_options,omitempty"`

	// container options
	ContainerOptions *ContainerOptions `json:"container_options,omitempty"`

	// Desired state of the Container
	ContainerState string `json:"container_state,omitempty"`

	// Image name
	// Required: true
	// Max Length: 140
	ImageName *string `json:"image_name"`

	// Networks associated with this Container
	NetworkReferenceList []*NetworkReference `json:"network_reference_list,omitempty"`

	// registry reference
	RegistryReference *DockerRegistryReference `json:"registry_reference,omitempty"`

	// Volumes associated with this Container
	VolumeList []*ContainerVolume `json:"volume_list,omitempty"`

	// Referenced Volumes associated with this Container
	VolumeReferenceList []*ContainerVolumeReference `json:"volume_reference_list,omitempty"`
}

// Validate validates this container resources
func (m *ContainerResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerBasicOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainerOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainerState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImageName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworkReferenceList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegistryReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumeReferenceList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerResources) validateContainerBasicOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainerBasicOptions) { // not required
		return nil
	}

	if m.ContainerBasicOptions != nil {

		if err := m.ContainerBasicOptions.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ContainerResources) validateContainerOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainerOptions) { // not required
		return nil
	}

	if m.ContainerOptions != nil {

		if err := m.ContainerOptions.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var containerResourcesTypeContainerStatePropEnum []interface{}

const (
	containerResourcesContainerStateQUEUED          string = "QUEUED"
	containerResourcesContainerStateIMAGEDOWNLOADED string = "IMAGE_DOWNLOADED"
	containerResourcesContainerStateCREATED         string = "CREATED"
	containerResourcesContainerStateRUNNING         string = "RUNNING"
	containerResourcesContainerStateSTOPPED         string = "STOPPED"
	containerResourcesContainerStateKILLED          string = "KILLED"
	containerResourcesContainerStateREMOVED         string = "REMOVED"
	containerResourcesContainerStateDELETED         string = "DELETED"
	containerResourcesContainerStatePAUSED          string = "PAUSED"
)

// prop value enum
func (m *ContainerResources) validateContainerStateEnum(path, location string, value string) error {
	if containerResourcesTypeContainerStatePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["QUEUED","IMAGE_DOWNLOADED","CREATED","RUNNING","STOPPED","KILLED","REMOVED","DELETED","PAUSED"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			containerResourcesTypeContainerStatePropEnum = append(containerResourcesTypeContainerStatePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, containerResourcesTypeContainerStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ContainerResources) validateContainerState(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainerState) { // not required
		return nil
	}

	// value enum
	if err := m.validateContainerStateEnum("container_state", "body", m.ContainerState); err != nil {
		return err
	}

	return nil
}

func (m *ContainerResources) validateImageName(formats strfmt.Registry) error {

	if err := validate.Required("image_name", "body", m.ImageName); err != nil {
		return err
	}

	if err := validate.MaxLength("image_name", "body", string(*m.ImageName), 140); err != nil {
		return err
	}

	return nil
}

func (m *ContainerResources) validateNetworkReferenceList(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkReferenceList) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkReferenceList); i++ {

		if swag.IsZero(m.NetworkReferenceList[i]) { // not required
			continue
		}

		if m.NetworkReferenceList[i] != nil {

			if err := m.NetworkReferenceList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ContainerResources) validateRegistryReference(formats strfmt.Registry) error {

	if swag.IsZero(m.RegistryReference) { // not required
		return nil
	}

	if m.RegistryReference != nil {

		if err := m.RegistryReference.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ContainerResources) validateVolumeList(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeList) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeList); i++ {

		if swag.IsZero(m.VolumeList[i]) { // not required
			continue
		}

		if m.VolumeList[i] != nil {

			if err := m.VolumeList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ContainerResources) validateVolumeReferenceList(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeReferenceList) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeReferenceList); i++ {

		if swag.IsZero(m.VolumeReferenceList[i]) { // not required
			continue
		}

		if m.VolumeReferenceList[i] != nil {

			if err := m.VolumeReferenceList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}