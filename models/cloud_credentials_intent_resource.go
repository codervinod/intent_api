package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// CloudCredentialsIntentResource cloud_credentials Intent Response
//
// Response object for intentful operations on a cloud_credentials
// swagger:model cloud_credentials_intent_resource
type CloudCredentialsIntentResource struct {

	// api version
	APIVersion string `json:"api_version,omitempty"`

	// metadata
	// Required: true
	Metadata *CloudCredentialsMetadata `json:"metadata"`

	Spec CloudCredentials `json:"-"`

	// status
	Status *CloudCredentialsDefStatus `json:"status,omitempty"`
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CloudCredentialsIntentResource) UnmarshalJSON(raw []byte) error {
	var data struct {
		APIVersion string `json:"api_version,omitempty"`

		Metadata *CloudCredentialsMetadata `json:"metadata"`

		Spec json.RawMessage `json:"spec,omitempty"`

		Status *CloudCredentialsDefStatus `json:"status,omitempty"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	spec, err := UnmarshalCloudCredentials(bytes.NewBuffer(data.Spec), runtime.JSONConsumer())
	if err != nil {
		return err
	}

	var result CloudCredentialsIntentResource
	result.APIVersion = data.APIVersion
	result.Metadata = data.Metadata
	result.Spec = spec
	result.Status = data.Status
	*m = result
	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CloudCredentialsIntentResource) MarshalJSON() ([]byte, error) {
	var b1, b2 []byte
	var err error
	b1, err = json.Marshal(struct {
		APIVersion string `json:"api_version,omitempty"`

		Metadata *CloudCredentialsMetadata `json:"metadata"`

		Status *CloudCredentialsDefStatus `json:"status,omitempty"`
	}{
		APIVersion: m.APIVersion,
		Metadata:   m.Metadata,
		Status:     m.Status,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Spec CloudCredentials `json:"spec,omitempty"`
	}{
		Spec: m.Spec,
	})
	if err != nil {
		return nil, err
	}
	return swag.ConcatJSON(b1, b2), nil
}

// Validate validates this cloud credentials intent resource
func (m *CloudCredentialsIntentResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudCredentialsIntentResource) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {

		if err := m.Metadata.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CloudCredentialsIntentResource) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {

		if err := m.Status.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
