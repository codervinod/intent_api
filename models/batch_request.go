package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// BatchRequest Batch request
//
// Batch request
// swagger:model batch_request
type BatchRequest struct {

	// Intent api list
	// Required: true
	APIRequestList []*APIRequest `json:"api_request_list"`

	// api version
	// Required: true
	// Read Only: true
	APIVersion string `json:"api_version"`

	// If true - continue processing remaining APIs in the list
	// Required: true
	ContinueOnFailure bool `json:"continue_on_failure"`

	// Determines whether the API execution proceeds sequentially or not
	ExecuteOrder *string `json:"execute_order,omitempty"`
}

// Validate validates this batch request
func (m *BatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIRequestList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAPIVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContinueOnFailure(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExecuteOrder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchRequest) validateAPIRequestList(formats strfmt.Registry) error {

	if err := validate.Required("api_request_list", "body", m.APIRequestList); err != nil {
		return err
	}

	for i := 0; i < len(m.APIRequestList); i++ {

		if swag.IsZero(m.APIRequestList[i]) { // not required
			continue
		}

		if m.APIRequestList[i] != nil {

			if err := m.APIRequestList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *BatchRequest) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.RequiredString("api_version", "body", string(m.APIVersion)); err != nil {
		return err
	}

	return nil
}

func (m *BatchRequest) validateContinueOnFailure(formats strfmt.Registry) error {

	if err := validate.Required("continue_on_failure", "body", bool(m.ContinueOnFailure)); err != nil {
		return err
	}

	return nil
}

var batchRequestTypeExecuteOrderPropEnum []interface{}

const (
	batchRequestExecuteOrderSEQUENTIAL    string = "SEQUENTIAL"
	batchRequestExecuteOrderNONSEQUENTIAL string = "NON_SEQUENTIAL"
)

// prop value enum
func (m *BatchRequest) validateExecuteOrderEnum(path, location string, value string) error {
	if batchRequestTypeExecuteOrderPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["SEQUENTIAL","NON_SEQUENTIAL"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			batchRequestTypeExecuteOrderPropEnum = append(batchRequestTypeExecuteOrderPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, batchRequestTypeExecuteOrderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BatchRequest) validateExecuteOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecuteOrder) { // not required
		return nil
	}

	// value enum
	if err := m.validateExecuteOrderEnum("execute_order", "body", *m.ExecuteOrder); err != nil {
		return err
	}

	return nil
}