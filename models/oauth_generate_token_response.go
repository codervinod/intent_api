package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// OauthGenerateTokenResponse Oauth token generate response
//
// Oauth token generate response
// swagger:model oauth_generate_token_response
type OauthGenerateTokenResponse struct {

	// client id of the Oauth Client
	// Required: true
	ClientID *string `json:"client_id"`

	// client secret of the Oauth Client
	// Required: true
	ClientSecret *string `json:"client_secret"`

	// oauth token
	// Required: true
	OauthToken *OauthToken `json:"oauth_token"`
}

// Validate validates this oauth generate token response
func (m *OauthGenerateTokenResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClientSecret(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOauthToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OauthGenerateTokenResponse) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *OauthGenerateTokenResponse) validateClientSecret(formats strfmt.Registry) error {

	if err := validate.Required("client_secret", "body", m.ClientSecret); err != nil {
		return err
	}

	return nil
}

func (m *OauthGenerateTokenResponse) validateOauthToken(formats strfmt.Registry) error {

	if err := validate.Required("oauth_token", "body", m.OauthToken); err != nil {
		return err
	}

	if m.OauthToken != nil {

		if err := m.OauthToken.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
