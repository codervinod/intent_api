package models

import "github.com/go-openapi/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Options Additional Options
//
// Behavioral options that need to be passed and are not a part of the spec
// go here
//
// swagger:model options
type Options map[string]interface{}

// Validate validates this options
func (m Options) Validate(formats strfmt.Registry) error {
	return nil
}
