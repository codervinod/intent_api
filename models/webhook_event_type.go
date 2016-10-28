package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// WebhookEventType Webhook event type
//
// Webhook event type
// swagger:model webhook_event_type
type WebhookEventType string

const (
	WebhookEventTypeVMON       WebhookEventType = "VM_ON"
	WebhookEventTypeVMOFF      WebhookEventType = "VM_OFF"
	WebhookEventTypeVMUPDATE   WebhookEventType = "VM_UPDATE"
	WebhookEventTypeVMMIGRATE  WebhookEventType = "VM_MIGRATE"
	WebhookEventTypeVNICPLUG   WebhookEventType = "VNIC_PLUG"
	WebhookEventTypeVNICUNPLUG WebhookEventType = "VNIC_UNPLUG"
)

// for schema
var webhookEventTypeEnum []interface{}

func (m WebhookEventType) validateWebhookEventTypeEnum(path, location string, value WebhookEventType) error {
	if webhookEventTypeEnum == nil {
		var res []WebhookEventType
		if err := json.Unmarshal([]byte(`["VM_ON","VM_OFF","VM_UPDATE","VM_MIGRATE","VNIC_PLUG","VNIC_UNPLUG"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			webhookEventTypeEnum = append(webhookEventTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, webhookEventTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this webhook event type
func (m WebhookEventType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateWebhookEventTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}