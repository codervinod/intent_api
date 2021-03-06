package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ChangedRegions Changed regions
//
// Returns information about all the regions that have changed in the interval: [start_offset, next_offset).
//
// swagger:model changed_regions
type ChangedRegions struct {

	// Size of the file specified by snapshot_file_path.
	// Required: true
	FileSize *int64 `json:"file_size"`

	// The offset from where the client must continue the request. This field will not be set when there are no more changed regions to be returned. Note that the next_offset can be outside the endOffset specified by the client in the request. This helps clients reach the next changed offset faster.
	//
	NextOffset int64 `json:"next_offset,omitempty"`

	// List of regions describing the change for the interval [start_offset, next_offset).
	//
	Regions []*Region `json:"regions,omitempty"`
}

// Validate validates this changed regions
func (m *ChangedRegions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangedRegions) validateFileSize(formats strfmt.Registry) error {

	if err := validate.Required("file_size", "body", m.FileSize); err != nil {
		return err
	}

	return nil
}

func (m *ChangedRegions) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	for i := 0; i < len(m.Regions); i++ {

		if swag.IsZero(m.Regions[i]) { // not required
			continue
		}

		if m.Regions[i] != nil {

			if err := m.Regions[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
