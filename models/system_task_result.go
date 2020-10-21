// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemTaskResult system task result
//
// swagger:model SystemTaskResult
type SystemTaskResult struct {

	// The identifier used to determine type of the system task result.
	ID string `json:"id,omitempty"`
}

// Validate validates this system task result
func (m *SystemTaskResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemTaskResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemTaskResult) UnmarshalBinary(b []byte) error {
	var res SystemTaskResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}