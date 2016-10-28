package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// WebhookDefStatus Webhook status definition
//
// Webhook status definition
// swagger:model webhook_def_status
type WebhookDefStatus struct {

	// A description or user annotation for the webhook
	// Max Length: 1000
	Description string `json:"description,omitempty"`

	// Webhook name
	// Required: true
	// Max Length: 140
	Name *string `json:"name"`

	// resources
	// Required: true
	Resources *WebhookResources `json:"resources"`
}

// Validate validates this webhook def status
func (m *WebhookDefStatus) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *WebhookDefStatus) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *WebhookDefStatus) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *WebhookDefStatus) validateResources(formats strfmt.Registry) error {

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
