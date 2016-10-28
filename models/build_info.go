package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// BuildInfo Cluter build details
//
// Cluter build details
// swagger:model build_info
type BuildInfo struct {

	// Build type, one of {dbg, opt, release}
	// Required: true
	BuildType *string `json:"build_type"`

	// Date/time of the last commit
	// Required: true
	CommitDate *string `json:"commit_date"`

	// Last Git commit id which the build is based on
	// Required: true
	CommitID *string `json:"commit_id"`

	// First 6 characters of the last Git commit id
	// Required: true
	ShortCommitID *string `json:"short_commit_id"`

	// Version string in format <code_name>-<version_numbers>-<branch_type>,
	// i.e master, danube-4.5.0.2-stable
	//
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this build info
func (m *BuildInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCommitDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCommitID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShortCommitID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildInfo) validateBuildType(formats strfmt.Registry) error {

	if err := validate.Required("build_type", "body", m.BuildType); err != nil {
		return err
	}

	return nil
}

func (m *BuildInfo) validateCommitDate(formats strfmt.Registry) error {

	if err := validate.Required("commit_date", "body", m.CommitDate); err != nil {
		return err
	}

	return nil
}

func (m *BuildInfo) validateCommitID(formats strfmt.Registry) error {

	if err := validate.Required("commit_id", "body", m.CommitID); err != nil {
		return err
	}

	return nil
}

func (m *BuildInfo) validateShortCommitID(formats strfmt.Registry) error {

	if err := validate.Required("short_commit_id", "body", m.ShortCommitID); err != nil {
		return err
	}

	return nil
}

func (m *BuildInfo) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}
