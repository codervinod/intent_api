package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VirtualDisk Virtual Disk (VM Disk)
//
// Virtual Disk (VM Disk)
// swagger:model virtual_disk
type VirtualDisk struct {

	// data source reference
	DataSourceReference *Reference `json:"data_source_reference,omitempty"`

	// device properties
	DeviceProperties *VirtualDiskDeviceProperties `json:"device_properties,omitempty"`

	// Size for the disk in MB.
	// Minimum: 1
	DiskSizeMb int32 `json:"disk_size_mb,omitempty"`

	// uuid
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this virtual disk
func (m *VirtualDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataSourceReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDeviceProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDiskSizeMb(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualDisk) validateDataSourceReference(formats strfmt.Registry) error {

	if swag.IsZero(m.DataSourceReference) { // not required
		return nil
	}

	if m.DataSourceReference != nil {

		if err := m.DataSourceReference.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VirtualDisk) validateDeviceProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceProperties) { // not required
		return nil
	}

	if m.DeviceProperties != nil {

		if err := m.DeviceProperties.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VirtualDisk) validateDiskSizeMb(formats strfmt.Registry) error {

	if swag.IsZero(m.DiskSizeMb) { // not required
		return nil
	}

	if err := validate.MinimumInt("disk_size_mb", "body", int64(m.DiskSizeMb), 1, false); err != nil {
		return err
	}

	return nil
}

// VirtualDiskDeviceProperties virtual disk device properties
// swagger:model VirtualDiskDeviceProperties
type VirtualDiskDeviceProperties struct {

	// device type
	DeviceType *string `json:"device_type,omitempty"`

	// disk address
	DiskAddress *DiskAddress `json:"disk_address,omitempty"`
}

// Validate validates this virtual disk device properties
func (m *VirtualDiskDeviceProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDiskAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var virtualDiskDevicePropertiesTypeDeviceTypePropEnum []interface{}

const (
	virtualDiskDevicePropertiesDeviceTypeDISK  string = "DISK"
	virtualDiskDevicePropertiesDeviceTypeCDROM string = "CDROM"
)

// prop value enum
func (m *VirtualDiskDeviceProperties) validateDeviceTypeEnum(path, location string, value string) error {
	if virtualDiskDevicePropertiesTypeDeviceTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["DISK","CDROM"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			virtualDiskDevicePropertiesTypeDeviceTypePropEnum = append(virtualDiskDevicePropertiesTypeDeviceTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, virtualDiskDevicePropertiesTypeDeviceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VirtualDiskDeviceProperties) validateDeviceType(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeviceTypeEnum("device_properties"+"."+"device_type", "body", *m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDiskDeviceProperties) validateDiskAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.DiskAddress) { // not required
		return nil
	}

	if m.DiskAddress != nil {

		if err := m.DiskAddress.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
