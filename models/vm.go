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

// VM VM
//
// VM Input Definition
// swagger:model vm
type VM struct {

	// backup policy
	BackupPolicy *VMBackupPolicy `json:"backup_policy,omitempty"`

	// Reference to the cluster where this VM exists or needs to be migrated
	// to. This is to support migration of a VM from one cluster to another
	// cluster.
	//
	ClusterReference *ClusterReference `json:"cluster_reference,omitempty"`

	// A description or user annotation for the VM
	// Max Length: 1000
	Description string `json:"description,omitempty"`

	// VM Name
	// Required: true
	// Max Length: 140
	Name *string `json:"name"`

	// policies
	Policies Policies `json:"policies,omitempty"`

	// providers
	Providers Providers `json:"providers,omitempty"`

	// resources
	// Required: true
	Resources *VMResources `json:"resources"`
}

// Validate validates this vm
func (m *VM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupPolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClusterReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VM) validateBackupPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.BackupPolicy) { // not required
		return nil
	}

	if m.BackupPolicy != nil {

		if err := m.BackupPolicy.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VM) validateClusterReference(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterReference) { // not required
		return nil
	}

	if m.ClusterReference != nil {

		if err := m.ClusterReference.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VM) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *VM) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *VM) validateResources(formats strfmt.Registry) error {

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

// VMBackupPolicy Describes how and where to backup the VM
// swagger:model VMBackupPolicy
type VMBackupPolicy struct {

	// Consistency group to which this VM belongs to.
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	ConsistencyGroupIdentifier string `json:"consistency_group_identifier,omitempty"`

	// snapshot policy
	SnapshotPolicy *SnapshotPolicy `json:"snapshot_policy,omitempty"`
}

// Validate validates this VM backup policy
func (m *VMBackupPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsistencyGroupIdentifier(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSnapshotPolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMBackupPolicy) validateConsistencyGroupIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.ConsistencyGroupIdentifier) { // not required
		return nil
	}

	if err := validate.Pattern("backup_policy"+"."+"consistency_group_identifier", "body", string(m.ConsistencyGroupIdentifier), `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
		return err
	}

	return nil
}

func (m *VMBackupPolicy) validateSnapshotPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.SnapshotPolicy) { // not required
		return nil
	}

	if m.SnapshotPolicy != nil {

		if err := m.SnapshotPolicy.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// VMResources VM Resources
// swagger:model VMResources
type VMResources struct {

	// Reference to an entity that the VM should be cloned from.
	CloneReference *Reference `json:"clone_reference,omitempty"`

	// Memory size in MB
	// Minimum: 1
	MemorySizeMb int32 `json:"memory_size_mb,omitempty"`

	// Number of cores per Vcpu
	// Minimum: 1
	NumCoresPerVcpu int32 `json:"num_cores_per_vcpu,omitempty"`

	// Number of Vcpus
	// Minimum: 1
	NumVcpus int32 `json:"num_vcpus,omitempty"`

	// The current or desired power state of the VM
	PowerState string `json:"power_state,omitempty"`

	// The base VM name if this is a system VM.
	SystemVMBaseName string `json:"system_vm_base_name,omitempty"`

	// Virtual Disks attached to the VM
	VirtualDiskList []*VirtualDisk `json:"virtual_disk_list,omitempty"`

	// Referenced virtual disks attached to the VM
	VirtualDiskReferenceList []*VirtualDiskReference `json:"virtual_disk_reference_list,omitempty"`

	// Virtual NICs attached to the VM
	VirtualNicList []*VirtualNic `json:"virtual_nic_list,omitempty"`
}

// Validate validates this VM resources
func (m *VMResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloneReference(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMemorySizeMb(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNumCoresPerVcpu(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNumVcpus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePowerState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVirtualDiskList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVirtualDiskReferenceList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVirtualNicList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMResources) validateCloneReference(formats strfmt.Registry) error {

	if swag.IsZero(m.CloneReference) { // not required
		return nil
	}

	if m.CloneReference != nil {

		if err := m.CloneReference.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VMResources) validateMemorySizeMb(formats strfmt.Registry) error {

	if swag.IsZero(m.MemorySizeMb) { // not required
		return nil
	}

	if err := validate.MinimumInt("resources"+"."+"memory_size_mb", "body", int64(m.MemorySizeMb), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *VMResources) validateNumCoresPerVcpu(formats strfmt.Registry) error {

	if swag.IsZero(m.NumCoresPerVcpu) { // not required
		return nil
	}

	if err := validate.MinimumInt("resources"+"."+"num_cores_per_vcpu", "body", int64(m.NumCoresPerVcpu), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *VMResources) validateNumVcpus(formats strfmt.Registry) error {

	if swag.IsZero(m.NumVcpus) { // not required
		return nil
	}

	if err := validate.MinimumInt("resources"+"."+"num_vcpus", "body", int64(m.NumVcpus), 1, false); err != nil {
		return err
	}

	return nil
}

var vmResourcesTypePowerStatePropEnum []interface{}

const (
	vmResourcesPowerStatePOWEREDON  string = "POWERED_ON"
	vmResourcesPowerStatePOWEREDOFF string = "POWERED_OFF"
	vmResourcesPowerStatePAUSED     string = "PAUSED"
	vmResourcesPowerStateRESETTING  string = "RESETTING"
)

// prop value enum
func (m *VMResources) validatePowerStateEnum(path, location string, value string) error {
	if vmResourcesTypePowerStatePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["POWERED_ON","POWERED_OFF","PAUSED","RESETTING"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			vmResourcesTypePowerStatePropEnum = append(vmResourcesTypePowerStatePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, vmResourcesTypePowerStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VMResources) validatePowerState(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerState) { // not required
		return nil
	}

	// value enum
	if err := m.validatePowerStateEnum("resources"+"."+"power_state", "body", m.PowerState); err != nil {
		return err
	}

	return nil
}

func (m *VMResources) validateVirtualDiskList(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualDiskList) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualDiskList); i++ {

		if swag.IsZero(m.VirtualDiskList[i]) { // not required
			continue
		}

		if m.VirtualDiskList[i] != nil {

			if err := m.VirtualDiskList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VMResources) validateVirtualDiskReferenceList(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualDiskReferenceList) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualDiskReferenceList); i++ {

		if swag.IsZero(m.VirtualDiskReferenceList[i]) { // not required
			continue
		}

		if m.VirtualDiskReferenceList[i] != nil {

			if err := m.VirtualDiskReferenceList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VMResources) validateVirtualNicList(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualNicList) { // not required
		return nil
	}

	for i := 0; i < len(m.VirtualNicList); i++ {

		if swag.IsZero(m.VirtualNicList[i]) { // not required
			continue
		}

		if m.VirtualNicList[i] != nil {

			if err := m.VirtualNicList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
