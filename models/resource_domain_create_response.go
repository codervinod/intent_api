package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ResourceDomainCreateResponse Resource domain create response
//
// The response for the resource domain create request
// swagger:model resource_domain_create_response
type ResourceDomainCreateResponse struct {

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this resource domain create response
func (m *ResourceDomainCreateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}