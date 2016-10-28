package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ImageType Image type
//
// The type of image
// swagger:model image_type
type ImageType string

const (
	ImageTypeDISKIMAGE ImageType = "DISK_IMAGE"
	ImageTypeISOIMAGE  ImageType = "ISO_IMAGE"
)

// for schema
var imageTypeEnum []interface{}

func (m ImageType) validateImageTypeEnum(path, location string, value ImageType) error {
	if imageTypeEnum == nil {
		var res []ImageType
		if err := json.Unmarshal([]byte(`["DISK_IMAGE","ISO_IMAGE"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			imageTypeEnum = append(imageTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, imageTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this image type
func (m ImageType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateImageTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}