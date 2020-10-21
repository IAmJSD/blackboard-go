// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AttemptFile attempt file
//
// swagger:model AttemptFile
type AttemptFile struct {

	// The primary key of the file attachment.
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// The name of the file which has been attached to an Attempt including the file extension.
	Name string `json:"name,omitempty"`

	// If Learn has a registered viewer for the file, this will be a URL for that viewer to render the file. This may not be populated depending on multiple factors including but not limited to Learn configuration, file type, and viewer provider.
	// Required: true
	// Read Only: true
	ViewURL string `json:"viewUrl"`
}

// Validate validates this attempt file
func (m *AttemptFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AttemptFile) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *AttemptFile) validateViewURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("viewUrl", "body", string(m.ViewURL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttemptFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttemptFile) UnmarshalBinary(b []byte) error {
	var res AttemptFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
