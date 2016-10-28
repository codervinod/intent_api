package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// WebhookResources Webhook resources
//
// Webhook resources
// swagger:model webhook_resources
type WebhookResources struct {

	// List of events subscribed to by the webhook
	EventsFilterList []WebhookEventType `json:"events_filter_list,omitempty"`

	// Target kind for the webhook
	// Required: true
	TargetKind *string `json:"target_kind"`

	// Url for the webhook
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this webhook resources
func (m *WebhookResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventsFilterList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTargetKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhookResources) validateEventsFilterList(formats strfmt.Registry) error {

	if swag.IsZero(m.EventsFilterList) { // not required
		return nil
	}

	return nil
}

func (m *WebhookResources) validateTargetKind(formats strfmt.Registry) error {

	if err := validate.Required("target_kind", "body", m.TargetKind); err != nil {
		return err
	}

	return nil
}

func (m *WebhookResources) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}
